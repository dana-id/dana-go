// Copyright 2025 PT Espay Debit Indonesia Koe
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package test

import (
	gocontext "context"
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"os"
	"time"

	"github.com/playwright-community/playwright-go"
)

// OAuthConfig holds the configuration for the OAuth flow
type OAuthConfig struct {
	PIN         string
	OAuthURL    string
	PhoneNumber string
}

// MobileData represents the seamlessData JSON structure in the OAuth URL
type MobileData struct {
	Mobile string `json:"mobile"`
}

// ExtractMobileFromURL extracts the mobile number from the OAuth URL's seamlessData parameter
func ExtractMobileFromURL(oauthURL string) string {
	parsedURL, err := url.Parse(oauthURL)
	if err != nil {
		log.Printf("Error parsing URL: %v", err)
		return "0811742234" // Default fallback number
	}

	queryParams := parsedURL.Query()
	seamlessData := queryParams.Get("seamlessData")
	if seamlessData == "" {
		return "0811742234" // Default fallback number
	}

	// URL decode and parse JSON
	unquoted, err := url.QueryUnescape(seamlessData)
	if err != nil {
		log.Printf("Error unescaping seamlessData: %v", err)
		return "0811742234"
	}

	var mobileData MobileData
	if err := json.Unmarshal([]byte(unquoted), &mobileData); err != nil {
		log.Printf("Error parsing seamlessData JSON: %v", err)
		return "0811742234"
	}

	if mobileData.Mobile == "" {
		return "0811742234"
	}

	return mobileData.Mobile
}

// AutomateOAuth performs the OAuth login flow and returns the authorization code
func AutomateOAuth(config *OAuthConfig) (string, error) {
	// If phone number is not provided, try to extract it from the OAuth URL
	if config.PhoneNumber == "" {
		config.PhoneNumber = ExtractMobileFromURL(config.OAuthURL)
	}

	// Install playwright if it's not already installed
	err := playwright.Install()
	if err != nil {
		return "", fmt.Errorf("could not install playwright: %w", err)
	}

	// Launch playwright
	pw, err := playwright.Run()
	if err != nil {
		return "", fmt.Errorf("could not launch playwright: %w", err)
	}
	defer pw.Stop()

	// Launch browser
	browserType := pw.Chromium
	browser, err := browserType.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(true),
	})
	if err != nil {
		return "", fmt.Errorf("could not launch browser: %w", err)
	}
	defer browser.Close()

	// Create a new context with mobile device emulation
	deviceName := "iPhone 13"
	device := pw.Devices[deviceName]
	if device == nil {
		return "", fmt.Errorf("device %s not found", deviceName)
	}

	// Create a viewport with the device dimensions
	viewport := &playwright.Size{
		Width:  device.Viewport.Width,
		Height: device.Viewport.Height,
	}

	context, err := browser.NewContext(playwright.BrowserNewContextOptions{
		UserAgent:         playwright.String(device.UserAgent),
		Viewport:          viewport,
		DeviceScaleFactor: playwright.Float(device.DeviceScaleFactor),
		IsMobile:          playwright.Bool(device.IsMobile),
		HasTouch:          playwright.Bool(device.HasTouch),
	})
	if err != nil {
		return "", fmt.Errorf("could not create browser context: %w", err)
	}
	defer context.Close()

	// Enable request/response logging if needed
	context.SetDefaultTimeout(60 * 1000) // 60 seconds timeout

	// Create a new page
	page, err := context.NewPage()
	if err != nil {
		return "", fmt.Errorf("could not create page: %w", err)
	}

	// Navigate to the OAuth URL
	_, err = page.Goto(config.OAuthURL, playwright.PageGotoOptions{
		// Use 'domcontentloaded' instead of 'networkidle' for better compatibility
		WaitUntil: playwright.WaitUntilStateDomcontentloaded,
		Timeout:   playwright.Float(60000),
	})
	if err != nil {
		return "", fmt.Errorf("could not navigate to OAuth URL: %w", err)
	}

	// Wait for the login page to load
	time.Sleep(3 * time.Second) // Give it a moment to stabilize

	// Check if we're already on the PIN entry page
	isPinEntryVisible, err := page.Locator("input[type=tel]").IsVisible()

	if err != nil || !isPinEntryVisible {
		// If not, we need to enter the phone number first
		page.Evaluate(`
        (mobile) => {
            const inputs = document.querySelectorAll('input');
            for (const input of inputs) {
                if (input.type === 'tel' || 
                    input.placeholder === '12312345678' || 
                    input.maxLength === 13 ||
                    input.className.includes('phone-number')) {
                    input.value = mobile;
                    input.dispatchEvent(new Event('input', { bubbles: true }));
                    return true;
                }
            }
            return false;
        }
        `, config.PhoneNumber)

		// Submit the phone number using JavaScript (similar to Python version)
		submitResult, err := page.Evaluate(`
	() => {
		const buttons = document.querySelectorAll('button');
		for (const button of buttons) {
			if (button.type === 'submit' || 
				button.innerText.includes('Next') || 
				button.innerText.includes('Lanjutkan') || 
				button.innerText.includes('Continue')) {
				console.log('Found submit button:', button);
				button.click();
				return true;
			}
		}
		return false;
	}
	`)
		if err != nil {
			return "", fmt.Errorf("could not find submit button: %w", err)
		}

		// If we didn't find the submit button through JS, try the traditional method
		if submitResult == nil || submitResult.(bool) == false {
			submitButton, err := page.WaitForSelector("button[type=submit]", playwright.PageWaitForSelectorOptions{
				State:   playwright.WaitForSelectorStateVisible,
				Timeout: playwright.Float(5000),
			})
			if err != nil {
				return "", fmt.Errorf("could not find submit button: %w", err)
			}

			// Click the submit button
			err = submitButton.Click()
			if err != nil {
				return "", fmt.Errorf("could not click submit button: %w", err)
			}
		}

		// Wait (up to 3 s) for any of the PIN-input selectors to appear.
		_, err = page.WaitForSelector(".txt-input-pin-field, input[maxlength=\"6\"], input[type=\"password\"]",
			playwright.PageWaitForSelectorOptions{Timeout: playwright.Float(5000)})
		if err != nil {
			log.Println("Timeout waiting for PIN field")
		}

		// Evaluate a small script that clicks the “Continue” button if found.
		needToContinue, err := page.Evaluate(`() => {
		const btn = document.querySelector('button.btn-continue.fs-unmask.btn.btn-primary');
		if (btn) { btn.click(); return true; }
		return false;
	}`)
		if err != nil {
			log.Fatalf("evaluate failed: %v", err)
		}

		if needToContinue.(bool) {
			time.Sleep(500 * time.Millisecond)
		}

		// Enter PIN using JavaScript (similar to phone input approach)
		pinResult, err := page.Evaluate(`
	(pinCode) => {
            const specificPinInput = document.querySelector('.txt-input-pin-field');
            if (specificPinInput) {
                specificPinInput.value = pinCode;
                specificPinInput.dispatchEvent(new Event('input', { bubbles: true }));
                specificPinInput.dispatchEvent(new Event('change', { bubbles: true }));
                return true;
            }
            const inputs = document.querySelectorAll('input');
            const singlePinInput = Array.from(inputs).find(input => 
                input.maxLength === 6 && 
                (input.type === 'text' || input.type === 'tel' || input.type === 'number' || input.inputMode === 'numeric')
            );
            if (singlePinInput) {
                singlePinInput.value = pinCode;
                singlePinInput.dispatchEvent(new Event('input', { bubbles: true }));
                singlePinInput.dispatchEvent(new Event('change', { bubbles: true }));
                return true;
            }
            const pinInputs = Array.from(inputs).filter(input => 
                input.maxLength === 1 || 
                input.type === 'password' || 
                input.className.includes('pin')
            );
            if (pinInputs.length >= pinCode.length) {
                for (let i = 0; i < pinCode.length; i++) {
                    pinInputs[i].value = pinCode.charAt(i);
                    pinInputs[i].dispatchEvent(new Event('input', { bubbles: true }));
                    pinInputs[i].dispatchEvent(new Event('change', { bubbles: true }));
                }
                return true;
            }
            return false;
        }
	`, config.PIN)

		if pinResult != true {
			// Check if pinResult is a map and contains success field
			return "", fmt.Errorf("could not enter PIN")
		}

		// Create a context with a timeout using Go's standard context package
		ctx, cancel := gocontext.WithTimeout(gocontext.Background(), 20*time.Second)
		defer cancel()

		var authCode string

		// Listen for URL changes to capture the authorization code
		page.On("framenavigated", func(frame playwright.Frame) {
			currentURL := frame.URL()

			// Check if this is the redirect URL with auth code
			if parsedURL, err := url.Parse(currentURL); err == nil {
				if parsedURL.Host == "google.com" || parsedURL.Query().Has("auth_code") {
					authCode = parsedURL.Query().Get("auth_code")
					cancel() // Stop waiting once we have the code
				}
			}
		})

		// Wait until context is done (either timeout or auth code captured)
		<-ctx.Done()

		if authCode == "" {
			// If we didn't get the auth code, check URL one more time
			currentURL := page.URL()
			if parsedURL, err := url.Parse(currentURL); err == nil {
				authCode = parsedURL.Query().Get("auth_code")
			}
		}

		if authCode == "" {
			return "", fmt.Errorf("could not capture authorization code")
		}

		log.Printf("OAuth flow completed successfully, auth code: %s", authCode)
		return authCode, nil
	}
	return "", fmt.Errorf("OAuth automation failed")
}

// GetDefaultOAuthConfig returns a default configuration for OAuth automation
func GetDefaultOAuthConfig() *OAuthConfig {
	// Check for PIN in environment variables first
	pin := os.Getenv("DANA_PIN")
	if pin == "" {
		pin = "123321" // Default test PIN
	}

	// OAuth URL can be provided via environment or use default test URL
	oauthURL := os.Getenv("DANA_OAUTH_URL")
	if oauthURL == "" {
		oauthURL = "https://m.sandbox.dana.id/n/ipg/oauth?partnerId=2025021714245533502768&scopes=CASHIER,AGREEMENT_PAY,QUERY_BALANCE,DEFAULT_BASIC_PROFILE,MINI_DANA&externalId=b3b7e164-a295-4461-8475-8db546f0d509&state=02c92610-aa7c-42b0-bf26-23bb06e4d475&channelId=2025021714245533502768&redirectUrl=https://google.com&timestamp=2025-05-20T22:27:01+00:00&seamlessData=%7B%22mobile%22%3A%220811742234%22%7D&seamlessSign=IN8MCZZMge6C0SOQG4otmP9WE5yTll0%2F6OwgmUV0cfFi1e9Hj8PAWuD1ZanQ9MZcrKH1nwJEKnYTtqtQ3AdpLa24E%2B2W3%2BNJD8nh7FLicjPFQuDEIAdE%2ByEqTfpU5Z8%2B1tdB%2BW3HN4p6ko%2BiSXu28XHZOnxxXbfMZzQ0qwpwhTp76xSi2tH5eU7ksp37G9sjCB3eyFXBR8bWr7NCjDzFL5cxVlTuEZCmLieDLYh%2FiGClPfWj%2F7tnzzppyiPJsG7PjWkuM25%2B%2BwHBcb7DUA1yVllq30lxUpKeogZ3AuY%2Be9%2FeRHrhz6d%2BBFnzowI3Fk2ZA64BR9E8TSpNHyzWKCNc1A%3D%3D&isSnapBI=true"
	}

	// Phone number is extracted from OAuth URL if not provided
	return &OAuthConfig{
		PIN:      pin,
		OAuthURL: oauthURL,
	}
}
