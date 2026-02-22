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
	"regexp"

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
