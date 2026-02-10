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

package widget

import (
	"fmt"

	exceptions "github.com/dana-id/dana-go/v2/exceptions"
	utils "github.com/dana-id/dana-go/v2/utils"
)

// ValidationFunc is a function type for custom validations
type ValidationFunc func(interface{}) error

// validationRegistry maps request type names to their validation functions
var validationRegistry = map[string][]ValidationFunc{
	"WidgetPaymentRequest": {
		validateValidUpToWidgetPaymentRequest,
		// Add more validation functions here as needed
		// Example: validateAmountWidgetPaymentRequest,
		// Example: validateMerchantIdWidgetPaymentRequest,
	},
	// Add more request types and their validations here as needed
}

// CustomValidation performs custom validations on the request based on its type
func CustomValidation(request interface{}, returnValue interface{}) error {
	if request == nil {
		return nil
	}

	switch req := request.(type) {
	case *WidgetPaymentRequest:
		if validators, exists := validationRegistry["WidgetPaymentRequest"]; exists {
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

// validateValidUpToWidgetPaymentRequest validates the validUpTo field in WidgetPaymentRequest
func validateValidUpToWidgetPaymentRequest(request interface{}) error {
	req, ok := request.(*WidgetPaymentRequest)
	if !ok {
		return nil // Type assertion failed, skip this validation
	}

	if req == nil {
		return nil
	}

	if err := utils.ValidateValidUpToDate(req.ValidUpTo); err != nil {
		return &exceptions.GenericOpenAPIError{
			ErrorMsg: fmt.Sprintf("validUpTo validation failed: %v", err),
		}
	}

	return nil
}
