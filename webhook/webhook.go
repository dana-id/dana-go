package webhook

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"encoding/pem"
	"fmt"

	payment_gateway "github.com/dana-id/go_client/payment_gateway/v1"
)

type WebhookParser struct {
	publicKey *rsa.PublicKey
}

func NewWebhookParser(gatewayPublicKeyPEM string) (*WebhookParser, error) {
	block, _ := pem.Decode([]byte(gatewayPublicKeyPEM))
	if block == nil {
		return nil, fmt.Errorf("failed to parse PEM block containing the public key")
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

// MinifyJSON compacts a JSON string by removing unnecessary whitespace.
func MinifyJSON(jsonStr string) (string, error) {
	var obj interface{}
	if err := json.Unmarshal([]byte(jsonStr), &obj); err != nil {
		return "", fmt.Errorf("MinifyJSON: failed to unmarshal JSON: %w", err)
	}

	minifiedBytes, err := json.Marshal(obj)
	if err != nil {
		return "", fmt.Errorf("MinifyJSON: failed to marshal JSON for minification: %w", err)
	}
	return string(minifiedBytes), nil
}

func sha256LowerHex(data string) string {
	hasher := sha256.New()
	hasher.Write([]byte(data))
	return hex.EncodeToString(hasher.Sum(nil))
}

func (p *WebhookParser) constructStringToVerify(httpMethod, relativePathURL, body, xTimestamp string) (string, error) {
	minifiedBody, err := MinifyJSON(body)
	if err != nil {
		return "", fmt.Errorf("constructStringToVerify: failed to minify JSON body: %w", err)
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
) (*payment_gateway.FinishNotifyRequest, error) {
	xSignature := headers.Get("X-SIGNATURE")
	xTimestamp := headers.Get("X-TIMESTAMP")

	if xSignature == "" || xTimestamp == "" {
		return nil, fmt.Errorf("missing X-SIGNATURE or X-TIMESTAMP header")
	}

	stringToVerify, err := p.constructStringToVerify(httpMethod, relativePathURL, body, xTimestamp)
	if err != nil {
		return nil, fmt.Errorf("failed to construct string for verification: %w", err)
	}

	signatureBytes, err := base64.StdEncoding.DecodeString(xSignature)
	if err != nil {
		return nil, fmt.Errorf("failed to base64 decode X-SIGNATURE: %w", err)
	}

	hashedStringToVerify := sha256.Sum256([]byte(stringToVerify))

	err = rsa.VerifyPKCS1v15(p.publicKey, crypto.SHA256, hashedStringToVerify[:], signatureBytes)
	if err != nil {
		if err == rsa.ErrVerification {
			return nil, fmt.Errorf("signature verification failed: %w. StringToVerify: '%s'", err, stringToVerify)
		}
		return nil, fmt.Errorf("error during signature verification: %w. StringToVerify: '%s'", err, stringToVerify)
	}

	var payload payment_gateway.FinishNotifyRequest
	if err := json.Unmarshal([]byte(body), &payload); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON body into FinishNotifyRequest: %w", err)
	}

	return &payload, nil
}
