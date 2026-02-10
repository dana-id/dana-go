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

package webhook

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"os"
	"strings"

	"github.com/dana-id/dana-go/utils"
)

// Sandbox gateway public key used for webhook signature verification when ENV is sandbox.
const sandboxWebhookPublicKey = `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAnaKVGRbin4Wh4KN35OPh
ytJBjYTz7QZKSZjmHfiHxFmulfT87rta+IvGJ0rCBgg+1EtKk1hX8G5gPGJs1htJ
5jHa3/jCk9l+luzjnuT9UVlwJahvzmFw+IoDoM7hIPjsLtnIe04SgYo0tZBpEmkQ
vUGhmHPqYnUGSSMIpDLJDvbyr8gtwluja1SbRphgDCoYVXq+uUJ5HzPS049aaxTS
nfXh/qXuDoB9EzCrgppLDS2ubmk21+dr7WaO/3RFjnwx5ouv6w+iC1XOJKar3CTk
X6JV1OSST1C9sbPGzMHZ8AGB51BM0mok7davD/5irUk+f0C25OgzkwtxAt80dkDo
/QIDAQAB
-----END PUBLIC KEY-----`

type WebhookParser struct {
	publicKey *rsa.PublicKey
}

func NewWebhookParser(publicKey *string, publicKeyPath *string) (*WebhookParser, error) {
	var keyInputContent string

	env := os.Getenv("DANA_ENV")
	if env == "" {
		env = os.Getenv("ENV")
	}
	if env == "" || strings.ToLower(env) == "sandbox" {
		keyInputContent = sandboxWebhookPublicKey
	} else if publicKeyPath != nil && *publicKeyPath != "" {
		fileBytes, errReadFile := os.ReadFile(*publicKeyPath)
		if errReadFile != nil {
			return nil, fmt.Errorf("failed to read key from file path '%s': %w", *publicKeyPath, errReadFile)
		}
		keyInputContent = string(fileBytes)
	} else if publicKey != nil && *publicKey != "" {
		keyInputContent = *publicKey
	} else {
		return nil, fmt.Errorf("either 'publicKey' or 'publicKeyPath' must be provided")
	}

	keyInputContent = strings.TrimSpace(keyInputContent)
	if keyInputContent == "" {
		return nil, fmt.Errorf("key content is empty after processing input")
	}

	normalizedPEM, err := utils.NormalizePEMKeyWithType(keyInputContent, "PUBLIC KEY")
	if err != nil {
		return nil, fmt.Errorf("failed to normalize gateway public key: %w\nOriginal input string (or file content): %s", err, keyInputContent)
	}

	block, _ := pem.Decode(normalizedPEM)
	if block == nil {
		return nil, fmt.Errorf("failed to parse PEM block containing the public key. Processed key:\n%s", string(normalizedPEM))
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse DER encoded public key: %w", err)
	}

	switch pub := pub.(type) {
	case *rsa.PublicKey:
		return &WebhookParser{publicKey: pub}, nil
	default:
		return nil, fmt.Errorf("public key is not of type RSA")
	}
}

func sha256LowerHex(data string) string {
	hasher := sha256.New()
	hasher.Write([]byte(data))
	return hex.EncodeToString(hasher.Sum(nil))
}

func (p *WebhookParser) constructStringToVerify(httpMethod, relativePathURL, body, xTimestamp string) (string, error) {
	minifiedBody, err := utils.EnsureMinifiedJSON(body)
	if err != nil {
		return "", fmt.Errorf("constructStringToVerify: failed to ensure JSON is minified: %w", err)
	}

	bodyHash := sha256LowerHex(minifiedBody)
	return fmt.Sprintf("%s:%s:%s:%s", httpMethod, relativePathURL, bodyHash, xTimestamp), nil
}

type HeaderGetter interface {
	Get(key string) string
}

type MapHeaderGetter map[string]string

func (m MapHeaderGetter) Get(key string) string {
	return m[key]
}

func (p *WebhookParser) ParseWebhook(
	httpMethod string,
	relativePathURL string,
	headers HeaderGetter,
	body string,
) (*FinishNotifyRequest, error) {
	normalizedHeaders := make(map[string]string)

	headerKeys := []string{"x-signature", "x-timestamp"}

	for _, key := range headerKeys {
		if value := headers.Get(key); value != "" {
			normalizedHeaders[strings.ToUpper(key)] = value
		}
		if value := headers.Get(strings.ToUpper(key)); value != "" {
			normalizedHeaders[strings.ToUpper(key)] = value
		}
		if value := headers.Get(strings.Title(key)); value != "" {
			normalizedHeaders[strings.ToUpper(key)] = value
		}
	}

	xSignature := normalizedHeaders["X-SIGNATURE"]
	xTimestamp := normalizedHeaders["X-TIMESTAMP"]

	if xSignature == "" || xTimestamp == "" {
		return nil, fmt.Errorf("missing X-SIGNATURE or X-TIMESTAMP header")
	}

	processedBody, err := utils.EnsureMinifiedJSON(body)
	if err != nil {
		return nil, fmt.Errorf("failed to process JSON body: %w", err)
	}

	stringToVerify, err := p.constructStringToVerify(httpMethod, relativePathURL, processedBody, xTimestamp)
	if err != nil {
		return nil, fmt.Errorf("failed to construct string for verification: %w", err)
	}

	signatureBytes, err := base64.StdEncoding.DecodeString(xSignature)
	if err != nil {
		return nil, fmt.Errorf("failed to base64 decode X-SIGNATURE: %w", err)
	}

	hasher := sha256.New()
	hasher.Write([]byte(stringToVerify))
	hashedStringToVerify := hasher.Sum(nil)

	err = rsa.VerifyPKCS1v15(p.publicKey, crypto.SHA256, hashedStringToVerify[:], signatureBytes)
	if err != nil {
		return nil, fmt.Errorf("signature verification failed: %w", err)
	}
	var result FinishNotifyRequest
	if err := utils.FlexibleUnmarshal([]byte(processedBody), &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal webhook body: %w", err)
	}
	return &result, nil
}
