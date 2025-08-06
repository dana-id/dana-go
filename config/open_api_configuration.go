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

package config

import (
	"context"
	"fmt"
	"strings"
)

// OpenApiAuthSettings stores OPEN_API specific authentication settings
type OpenApiAuthSettings struct {
	ClientSecret   string
	ClientID       string
	Env            string
	AccessToken    string
	PrivateKey     string
	PrivateKeyPath string
}

// OpenApiConfiguration stores the configuration for OPEN_API endpoints
// This configuration handles CLIENT_SECRET authentication instead of SNAP signature authentication
type OpenApiConfiguration struct {
	*Configuration
	OpenApiAuth *OpenApiAuthSettings
}

// NewOpenApiConfiguration creates a new OpenApiConfiguration object
func NewOpenApiConfiguration() *OpenApiConfiguration {
	cfg := NewConfiguration()

	return &OpenApiConfiguration{
		Configuration: cfg,
		OpenApiAuth:   &OpenApiAuthSettings{},
	}
}

// NewOpenApiConfigurationWithAuth creates a new OpenApiConfiguration with authentication settings
func NewOpenApiConfigurationWithAuth(auth *OpenApiAuthSettings) *OpenApiConfiguration {
	cfg := NewConfiguration()

	// Set API key fields from OPEN_API auth settings
	if cfg.APIKey == nil {
		cfg.APIKey = &APIKey{}
	}

	cfg.APIKey.CLIENT_SECRET = auth.ClientSecret
	cfg.APIKey.CLIENT_ID = auth.ClientID
	cfg.APIKey.ENV = auth.Env
	cfg.APIKey.DANA_ENV = auth.Env
	cfg.APIKey.ACCESS_TOKEN = auth.AccessToken
	cfg.APIKey.PRIVATE_KEY = auth.PrivateKey
	cfg.APIKey.PRIVATE_KEY_PATH = auth.PrivateKeyPath

	// Set appropriate server based on environment
	if auth.Env == ENV_PRODUCTION {
		cfg.Servers = ServerConfigurations{
			{
				URL:         "https://api.dana.id",
				Description: "DANA API Gateway Production",
			},
		}
	} else {
		cfg.Servers = ServerConfigurations{
			{
				URL:         "https://api-sandbox.dana.id",
				Description: "DANA API Gateway Sandbox",
			},
		}
	}

	return &OpenApiConfiguration{
		Configuration: cfg,
		OpenApiAuth:   auth,
	}
}

// GetOpenApiAuthSettings returns the authentication settings for OPEN_API endpoints
func (c *OpenApiConfiguration) GetOpenApiAuthSettings() map[string]interface{} {
	auth := make(map[string]interface{})

	if c.OpenApiAuth != nil {
		if c.OpenApiAuth.ClientSecret != "" {
			auth["CLIENT_SECRET"] = map[string]interface{}{
				"type":  "api_key",
				"in":    "header",
				"key":   "clientSecret",
				"value": c.OpenApiAuth.ClientSecret,
			}
		}

		if c.OpenApiAuth.ClientID != "" {
			auth["CLIENT_ID"] = map[string]interface{}{
				"type":  "api_key",
				"in":    "header",
				"key":   "clientId",
				"value": c.OpenApiAuth.ClientID,
			}
		}

		if c.OpenApiAuth.AccessToken != "" {
			auth["ACCESS_TOKEN"] = map[string]interface{}{
				"type":  "api_key",
				"in":    "header",
				"key":   "accessToken",
				"value": c.OpenApiAuth.AccessToken,
			}
		}
	}

	return auth
}

// GetHostSettings returns the appropriate host settings for OPEN_API endpoints
func (c *OpenApiConfiguration) GetHostSettings() ServerConfigurations {
	if c.OpenApiAuth != nil && c.OpenApiAuth.Env == ENV_PRODUCTION {
		return ServerConfigurations{
			{
				URL:         "https://api.dana.id",
				Description: "DANA API Gateway Production",
			},
		}
	}

	return ServerConfigurations{
		{
			URL:         "https://api-sandbox.dana.id",
			Description: "DANA API Gateway Sandbox",
		},
	}
}

// ServerURLWithContext returns a new server URL given an endpoint for OPEN_API
func (c *OpenApiConfiguration) ServerURLWithContext(ctx context.Context, endpoint string) (string, error) {
	// Get the appropriate server configuration
	servers := c.GetHostSettings()

	if len(servers) == 0 {
		return "", fmt.Errorf("no server configuration available for endpoint: %s", endpoint)
	}

	// For OPEN_API, we typically use the first (and usually only) server
	serverURL, err := servers.URL(0, nil)
	if err != nil {
		return "", fmt.Errorf("failed to get server URL: %v", err)
	}

	return serverURL, nil
}

// ToDebugReport returns debugging information for OPEN_API configuration
func (c *OpenApiConfiguration) ToDebugReport() string {
	report := "DANA Go SDK Debug Report:\n"
	report += fmt.Sprintf("Host: %s\n", c.Host)
	report += fmt.Sprintf("User Agent: %s\n", c.UserAgent)
	report += fmt.Sprintf("Debug: %v\n", c.Debug)

	if c.OpenApiAuth != nil {
		report += "\nOPEN_API Configuration:\n"
		report += fmt.Sprintf("Environment: %s\n", c.OpenApiAuth.Env)
		report += fmt.Sprintf("Client ID: %s\n", maskSensitive(c.OpenApiAuth.ClientID))
		report += fmt.Sprintf("Client Secret: %s\n", maskSensitive(c.OpenApiAuth.ClientSecret))

		if c.OpenApiAuth.AccessToken != "" {
			report += fmt.Sprintf("Access Token: %s\n", maskSensitive(c.OpenApiAuth.AccessToken))
		}

		if c.OpenApiAuth.PrivateKey != "" {
			report += "Private Key: [SET]\n"
		} else if c.OpenApiAuth.PrivateKeyPath != "" {
			report += fmt.Sprintf("Private Key Path: %s\n", c.OpenApiAuth.PrivateKeyPath)
		}
	}

	return report
}

// maskSensitive masks sensitive information for logging
func maskSensitive(value string) string {
	if len(value) <= 8 {
		return strings.Repeat("*", len(value))
	}
	return value[:4] + strings.Repeat("*", len(value)-8) + value[len(value)-4:]
}

// Validate checks if the OPEN_API configuration is valid
func (c *OpenApiConfiguration) Validate() error {
	if c.OpenApiAuth == nil {
		return fmt.Errorf("OPEN_API auth settings are required")
	}

	if c.OpenApiAuth.ClientSecret == "" {
		return fmt.Errorf("CLIENT_SECRET is required for OPEN_API authentication")
	}

	if c.OpenApiAuth.Env == "" {
		return fmt.Errorf("ENV or DANA_ENV is required for OPEN_API configuration")
	}

	if c.OpenApiAuth.Env != ENV_SANDBOX && c.OpenApiAuth.Env != ENV_PRODUCTION {
		return fmt.Errorf("ENV or DANA_ENV must be either %s or %s", ENV_SANDBOX, ENV_PRODUCTION)
	}

	return nil
}

// SetDefaultOpenApiConfiguration sets a default OPEN_API configuration
func SetDefaultOpenApiConfiguration(config *OpenApiConfiguration) {
	// This could be used to set a global default configuration
	// Implementation depends on your specific needs
}

// GetDefaultOpenApiConfiguration returns a default OPEN_API configuration
func GetDefaultOpenApiConfiguration() *OpenApiConfiguration {
	return NewOpenApiConfiguration()
}
