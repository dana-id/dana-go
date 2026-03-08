// Copyright 2026 PT Espay Debit Indonesia Koe
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

package payment_gateway

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	exceptions "github.com/dana-id/dana-go/v2/exceptions"
	utils "github.com/dana-id/dana-go/v2/utils"
)

// moneyValuePattern: value must be digits (1-16) + "." + exactly 2 digits (e.g. 10000.00)
var moneyValuePattern = regexp.MustCompile(`^\d{1,16}\.\d{2}$`)

// ValidationFunc is a function type for custom validations
type ValidationFunc func(interface{}) error

// validationRegistry maps request type names to their validation functions
var validationRegistry = map[string][]ValidationFunc{
	"CreateOrderRequest": {
		validateAdditionalInfoRequiredCreateOrderRequest,
		validateMoneyValueCreateOrderRequest,
		validateValidUpToCreateOrderRequest,
		validateExternalStoreIdForQris,
		validateSandboxPayMethodAndPayOption,
	},
	// Add more request types and their validations here as needed
}

// CustomValidation performs custom validations on the request based on its type
func CustomValidation(request interface{}, returnValue interface{}) error {
	if request == nil {
		return nil
	}

	switch req := request.(type) {
	case *CreateOrderRequest:
		if validators, exists := validationRegistry["CreateOrderRequest"]; exists {
			for _, validator := range validators {
				if err := validator(req); err != nil {
					return err
				}
			}
		}
	default:
		return nil
	}

	return nil
}

// validateAdditionalInfoRequiredCreateOrderRequest validates that AdditionalInfo must exist for CreateOrderByApiRequest and CreateOrderByRedirectRequest
func validateAdditionalInfoRequiredCreateOrderRequest(request interface{}) error {
	req, ok := request.(*CreateOrderRequest)
	if !ok || req == nil {
		return nil
	}
	if req.CreateOrderByApiRequest != nil {
		if req.CreateOrderByApiRequest.AdditionalInfo == nil {
			return &exceptions.GenericOpenAPIError{
				ErrorMsg: "additionalInfo is required",
			}
		}
	}
	if req.CreateOrderByRedirectRequest != nil {
		if req.CreateOrderByRedirectRequest.AdditionalInfo == nil {
			return &exceptions.GenericOpenAPIError{
				ErrorMsg: "additionalInfo is required",
			}
		}
	}
	return nil
}

// validateMoneyValueCreateOrderRequest validates that Money value fields match pattern ^\d{1,16}\.\d{2}$
func validateMoneyValueCreateOrderRequest(request interface{}) error {
	req, ok := request.(*CreateOrderRequest)
	if !ok || req == nil {
		return nil
	}
	validateMoney := func(m Money, fieldName string) error {
		v := m.Value
		if v == "" {
			return &exceptions.GenericOpenAPIError{
				ErrorMsg: fmt.Sprintf("%s.value is required", fieldName),
			}
		}
		if !moneyValuePattern.MatchString(v) {
			return &exceptions.GenericOpenAPIError{
				ErrorMsg: fmt.Sprintf("%s.value must match pattern (e.g. 10000.00): got %q", fieldName, v),
			}
		}
		return nil
	}
	if req.CreateOrderByApiRequest != nil {
		if err := validateMoney(req.CreateOrderByApiRequest.Amount, "amount"); err != nil {
			return err
		}
	}
	if req.CreateOrderByRedirectRequest != nil {
		if err := validateMoney(req.CreateOrderByRedirectRequest.Amount, "amount"); err != nil {
			return err
		}
	}
	return nil
}

// validateValidUpToCreateOrderRequest validates the validUpTo field in CreateOrderRequest
func validateValidUpToCreateOrderRequest(request interface{}) error {
	req, ok := request.(*CreateOrderRequest)
	if !ok {
		return nil // Type assertion failed, skip this validation
	}

	if req == nil {
		return nil
	}

	// Validate validUpTo for CreateOrderByApiRequest variant
	if req.CreateOrderByApiRequest != nil {
		if err := utils.ValidateValidUpToDate(req.CreateOrderByApiRequest.ValidUpTo); err != nil {
			return &exceptions.GenericOpenAPIError{
				ErrorMsg: fmt.Sprintf("validUpTo validation failed: %v", err),
			}
		}
	}

	// Validate validUpTo for CreateOrderByRedirectRequest variant
	if req.CreateOrderByRedirectRequest != nil {
		if err := utils.ValidateValidUpToDate(req.CreateOrderByRedirectRequest.ValidUpTo); err != nil {
			return &exceptions.GenericOpenAPIError{
				ErrorMsg: fmt.Sprintf("validUpTo validation failed: %v", err),
			}
		}
	}

	return nil
}

// validateExternalStoreIdForQris validates that externalStoreId is required when payOption is NETWORK_PAY_PG_QRIS
func validateExternalStoreIdForQris(request interface{}) error {
	req, ok := request.(*CreateOrderRequest)
	if !ok {
		return nil // Type assertion failed, skip this validation
	}

	if req == nil {
		return nil
	}

	// Validate for CreateOrderByApiRequest variant
	if req.CreateOrderByApiRequest != nil {
		apiReq := req.CreateOrderByApiRequest
		hasQris := false

		// Check if any payOption is NETWORK_PAY_PG_QRIS
		if apiReq.PayOptionDetails != nil {
			for _, payOptionDetail := range apiReq.PayOptionDetails {
				if payOptionDetail.PayOption == string(PAYOPTION_NETWORK_PAY_PG_QRIS_) {
					hasQris = true
					break
				}
			}
		}

		// If QRIS is found, externalStoreId must be provided
		if hasQris {
			if apiReq.ExternalStoreId == nil || *apiReq.ExternalStoreId == "" {
				return &exceptions.GenericOpenAPIError{
					ErrorMsg: "externalStoreId is required when payOption is NETWORK_PAY_PG_QRIS",
				}
			}
		}
	}

	return nil
}

// In sandbox, only these payMethods are available (Payment Gateway).
var sandboxAllowedPayMethods = map[string]bool{
	"BALANCE": true, "CREDIT_CARD": true, "DEBIT_CARD": true,
	"VIRTUAL_ACCOUNT": true, "NETWORK_PAY": true,
}

// In sandbox, only these payOptions are available (exact or suffix match e.g. VIRTUAL_ACCOUNT_BRI).
var sandboxAllowedPayOptions = map[string]bool{
	"CARD": true, "QRIS": true, "BRI": true, "PANIN": true,
	"CIMB": true, "MANDIRI": true, "BTPN": true,
}

func isSandbox() bool {
	env := os.Getenv("DANA_ENV")
	if env == "" {
		env = os.Getenv("ENV")
	}
	if env == "" {
		env = "sandbox"
	}
	return strings.ToLower(env) == "sandbox"
}

func payOptionAllowedInSandbox(value string) bool {
	value = strings.TrimSpace(value)
	if value == "" {
		return false
	}
	if sandboxAllowedPayOptions[value] {
		return true
	}
	for opt := range sandboxAllowedPayOptions {
		if strings.HasSuffix(value, "_"+opt) {
			return true
		}
	}
	return false
}

// validateSandboxPayMethodAndPayOption validates that in sandbox only allowed payMethod and payOption values are used.
func validateSandboxPayMethodAndPayOption(request interface{}) error {
	if !isSandbox() {
		return nil
	}
	req, ok := request.(*CreateOrderRequest)
	if !ok || req == nil || req.CreateOrderByApiRequest == nil {
		return nil
	}
	apiReq := req.CreateOrderByApiRequest
	if apiReq.PayOptionDetails == nil {
		return nil
	}
	for i, detail := range apiReq.PayOptionDetails {
		pm := strings.TrimSpace(detail.PayMethod)
		if pm != "" && !sandboxAllowedPayMethods[pm] {
			return &exceptions.GenericOpenAPIError{
				ErrorMsg: fmt.Sprintf("in sandbox, payMethod must be one of [BALANCE, CREDIT_CARD, DEBIT_CARD, VIRTUAL_ACCOUNT, NETWORK_PAY]; got %q in payOptionDetails[%d]", pm, i),
			}
		}
		po := strings.TrimSpace(detail.PayOption)
		if po != "" && !payOptionAllowedInSandbox(po) {
			return &exceptions.GenericOpenAPIError{
				ErrorMsg: fmt.Sprintf("in sandbox, payOption must be one of [CARD, QRIS, BRI, PANIN, CIMB, MANDIRI, BTPN] (or suffix like VIRTUAL_ACCOUNT_BRI); got %q in payOptionDetails[%d]", po, i),
			}
		}
	}
	return nil
}
