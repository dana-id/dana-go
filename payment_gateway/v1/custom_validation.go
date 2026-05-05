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
	"unicode/utf8"

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
		validateConditionalPayOptionAdditionalInfoCreateOrderRequest,
		validateSandboxPayMethodAndPayOption,
		validateOptionalFieldsWithRequiredNestedCreateOrderRequest,
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
			var validationErrors []error
			for _, validator := range validators {
				if err := validator(req); err != nil {
					validationErrors = append(validationErrors, err)
				}
			}
			if len(validationErrors) > 0 {
				messages := make([]string, 0, len(validationErrors))
				for _, err := range validationErrors {
					if err == nil {
						continue
					}
					messages = append(messages, err.Error())
				}

				return &exceptions.GenericOpenAPIError{
					ErrorMsg: fmt.Sprintf("validation failed: %s", strings.Join(messages, "; ")),
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

func runeLen(s string) int {
	return utf8.RuneCountInString(s)
}

func isCardPayment(payMethod, payOption string) bool {
	payMethod = strings.TrimSpace(payMethod)
	payOption = strings.TrimSpace(payOption)

	if payMethod == string(PAYMETHOD_CARD_) ||
		payMethod == string(PAYMETHOD_CREDIT_CARD_) {
		return true
	}
	return payOption == string(PAYOPTION_NETWORK_PAY_PG_CARD_)
}

func isEwalletPayment(payOption string) bool {
	payOption = strings.TrimSpace(payOption)
	switch payOption {
	case string(PAYOPTION_NETWORK_PAY_PG_SPAY_),
		string(PAYOPTION_NETWORK_PAY_PG_OVO_),
		string(PAYOPTION_NETWORK_PAY_PG_GOPAY_),
		string(PAYOPTION_NETWORK_PAY_PG_LINKAJA_):
		return true
	default:
		return false
	}
}

func isVirtualAccountPayMethod(payMethod string) bool {
	return strings.TrimSpace(payMethod) == string(PAYMETHOD_VIRTUAL_ACCOUNT_)
}

// validateConditionalPayOptionAdditionalInfoCreateOrderRequest enforces conditional fields inside:
// createOrderByApiRequest.payOptionDetails[].additionalInfo
//
// Spec (from OpenAPI):
//   - phoneNumber is required for Card and e-wallet payments, and must be 1-15 characters.
//     Card: payMethod CARD (and credit/debit-card methods) OR payOption NETWORK_PAY_PG_CARD
//     E-wallet: payOption NETWORK_PAY_PG_{SPAY,OVO,GOPAY,LINKAJA}
func validateConditionalPayOptionAdditionalInfoCreateOrderRequest(request interface{}) error {
	req, ok := request.(*CreateOrderRequest)
	if !ok || req == nil {
		return nil
	}
	if req.CreateOrderByApiRequest == nil || req.CreateOrderByApiRequest.PayOptionDetails == nil {
		return nil
	}

	for i, detail := range req.CreateOrderByApiRequest.PayOptionDetails {
		payMethod := strings.TrimSpace(detail.PayMethod)
		payOption := strings.TrimSpace(detail.PayOption)

		additional := detail.AdditionalInfo
		phoneNumber := ""
		if additional != nil && additional.PhoneNumber != nil {
			phoneNumber = strings.TrimSpace(*additional.PhoneNumber)
		}

		// Card and e-wallet payments require phoneNumber.
		if isCardPayment(payMethod, payOption) || isEwalletPayment(payOption) {
			if phoneNumber == "" {
				return &exceptions.GenericOpenAPIError{
					ErrorMsg: fmt.Sprintf("phoneNumber is required for card/e-wallet payment (payOptionDetails[%d])", i),
				}
			}
			if runeLen(phoneNumber) < 1 || runeLen(phoneNumber) > 15 {
				return &exceptions.GenericOpenAPIError{
					ErrorMsg: fmt.Sprintf("phoneNumber must be between 1 and 15 characters (payOptionDetails[%d])", i),
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
	"CARD": true, "QRIS": true, "BRI": true, "PANI": true,
	"CIMB": true, "MANDIRI": true, "BTPN": true, "BSI_PAYMENT": true,
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

// validateOptionalFieldsWithRequiredNestedCreateOrderRequest enforces: when an optional field is present,
// its required nested fields must be present (per payment_gateway spec). E.g. goods is optional but when
// filled, each good must have name and other required Goods fields; shippingInfo when filled must have
// required ShippingInfo fields in each item.
func validateOptionalFieldsWithRequiredNestedCreateOrderRequest(request interface{}) error {
	req, ok := request.(*CreateOrderRequest)
	if !ok || req == nil {
		return nil
	}
	if req.CreateOrderByApiRequest != nil && req.CreateOrderByApiRequest.AdditionalInfo != nil {
		if err := validateOrderOptionalNestedFields(req.CreateOrderByApiRequest.AdditionalInfo.Order); err != nil {
			return err
		}
	}
	if req.CreateOrderByRedirectRequest != nil && req.CreateOrderByRedirectRequest.AdditionalInfo != nil {
		if req.CreateOrderByRedirectRequest.AdditionalInfo.Order != nil {
			if err := validateOrderOptionalNestedFieldsRedirect(*req.CreateOrderByRedirectRequest.AdditionalInfo.Order); err != nil {
				return err
			}
		}
	}
	return nil
}

// validateOrderOptionalNestedFields validates optional fields on OrderApiObject that have required nested content when present:
// buyer (externalUserType/externalUserId pair), goods, and shippingInfo.
func validateOrderOptionalNestedFields(order *OrderApiObject) error {
	if order == nil {
		return nil
	}
	if err := validateBuyerWhenPresent(order.Buyer, "additionalInfo.order.buyer"); err != nil {
		return err
	}
	if len(order.Goods) > 0 {
		for i, g := range order.Goods {
			if err := validateGoodsRequiredFieldsWhenPresent(&g, "additionalInfo.order.goods", i); err != nil {
				return err
			}
		}
	}
	if len(order.ShippingInfo) > 0 {
		for i, s := range order.ShippingInfo {
			if err := validateShippingInfoRequiredFieldsWhenPresent(&s, "additionalInfo.order.shippingInfo", i); err != nil {
				return err
			}
		}
	}
	return nil
}

// validateOrderOptionalNestedFieldsRedirect validates optional fields on OrderRedirectObject that have required nested content when present:
// buyer (externalUserType/externalUserId pair), goods, and shippingInfo.
func validateOrderOptionalNestedFieldsRedirect(order OrderRedirectObject) error {
	if err := validateBuyerWhenPresent(order.Buyer, "additionalInfo.order.buyer"); err != nil {
		return err
	}
	if len(order.Goods) > 0 {
		for i, g := range order.Goods {
			if err := validateGoodsRequiredFieldsWhenPresent(&g, "additionalInfo.order.goods", i); err != nil {
				return err
			}
		}
	}
	if len(order.ShippingInfo) > 0 {
		for i, s := range order.ShippingInfo {
			if err := validateShippingInfoRequiredFieldsWhenPresent(&s, "additionalInfo.order.shippingInfo", i); err != nil {
				return err
			}
		}
	}
	return nil
}

// validateBuyerWhenPresent enforces spec: when buyer is filled, externalUserType and externalUserId
// are mutually dependent (Required if externalUserId is filled / Required if externalUserType is filled).
// So when buyer is present, if one is set the other must be set.
func validateBuyerWhenPresent(buyer *Buyer, fieldPath string) error {
	if buyer == nil {
		return nil
	}
	hasType := buyer.ExternalUserType != nil && strings.TrimSpace(*buyer.ExternalUserType) != ""
	hasID := buyer.ExternalUserId != nil && strings.TrimSpace(*buyer.ExternalUserId) != ""
	if hasType && !hasID {
		return &exceptions.GenericOpenAPIError{
			ErrorMsg: fmt.Sprintf("%s.externalUserId is required when %s.externalUserType is filled", fieldPath, fieldPath),
		}
	}
	if hasID && !hasType {
		return &exceptions.GenericOpenAPIError{
			ErrorMsg: fmt.Sprintf("%s.externalUserType is required when %s.externalUserId is filled", fieldPath, fieldPath),
		}
	}
	return nil
}

// validateGoodsRequiredFieldsWhenPresent ensures that when goods is filled, each item has required fields per spec
// (name, merchantGoodsId, description, category, price, quantity).
func validateGoodsRequiredFieldsWhenPresent(g *Goods, fieldPath string, index int) error {
	if g == nil {
		return nil
	}
	if strings.TrimSpace(g.Name) == "" {
		return &exceptions.GenericOpenAPIError{
			ErrorMsg: fmt.Sprintf("%s[%d].name is required when goods is filled", fieldPath, index),
		}
	}
	if strings.TrimSpace(g.MerchantGoodsId) == "" {
		return &exceptions.GenericOpenAPIError{
			ErrorMsg: fmt.Sprintf("%s[%d].merchantGoodsId is required when goods is filled", fieldPath, index),
		}
	}
	if strings.TrimSpace(g.Description) == "" {
		return &exceptions.GenericOpenAPIError{
			ErrorMsg: fmt.Sprintf("%s[%d].description is required when goods is filled", fieldPath, index),
		}
	}
	if strings.TrimSpace(g.Category) == "" {
		return &exceptions.GenericOpenAPIError{
			ErrorMsg: fmt.Sprintf("%s[%d].category is required when goods is filled", fieldPath, index),
		}
	}
	if strings.TrimSpace(g.Quantity) == "" {
		return &exceptions.GenericOpenAPIError{
			ErrorMsg: fmt.Sprintf("%s[%d].quantity is required when goods is filled", fieldPath, index),
		}
	}
	if strings.TrimSpace(g.Price.Value) == "" {
		return &exceptions.GenericOpenAPIError{
			ErrorMsg: fmt.Sprintf("%s[%d].price.value is required when goods is filled", fieldPath, index),
		}
	}
	if strings.TrimSpace(g.Price.Currency) == "" {
		return &exceptions.GenericOpenAPIError{
			ErrorMsg: fmt.Sprintf("%s[%d].price.currency is required when goods is filled", fieldPath, index),
		}
	}
	return nil
}

// validateShippingInfoRequiredFieldsWhenPresent ensures that when shippingInfo is filled, each item has required fields per spec.
func validateShippingInfoRequiredFieldsWhenPresent(s *ShippingInfo, fieldPath string, index int) error {
	if s == nil {
		return nil
	}
	if strings.TrimSpace(s.MerchantShippingId) == "" {
		return &exceptions.GenericOpenAPIError{
			ErrorMsg: fmt.Sprintf("%s[%d].merchantShippingId is required when shippingInfo is filled", fieldPath, index),
		}
	}
	if strings.TrimSpace(s.CountryName) == "" {
		return &exceptions.GenericOpenAPIError{
			ErrorMsg: fmt.Sprintf("%s[%d].countryName is required when shippingInfo is filled", fieldPath, index),
		}
	}
	if strings.TrimSpace(s.StateName) == "" {
		return &exceptions.GenericOpenAPIError{
			ErrorMsg: fmt.Sprintf("%s[%d].stateName is required when shippingInfo is filled", fieldPath, index),
		}
	}
	if strings.TrimSpace(s.CityName) == "" {
		return &exceptions.GenericOpenAPIError{
			ErrorMsg: fmt.Sprintf("%s[%d].cityName is required when shippingInfo is filled", fieldPath, index),
		}
	}
	if strings.TrimSpace(s.Address1) == "" {
		return &exceptions.GenericOpenAPIError{
			ErrorMsg: fmt.Sprintf("%s[%d].address1 is required when shippingInfo is filled", fieldPath, index),
		}
	}
	if strings.TrimSpace(s.FirstName) == "" {
		return &exceptions.GenericOpenAPIError{
			ErrorMsg: fmt.Sprintf("%s[%d].firstName is required when shippingInfo is filled", fieldPath, index),
		}
	}
	if strings.TrimSpace(s.LastName) == "" {
		return &exceptions.GenericOpenAPIError{
			ErrorMsg: fmt.Sprintf("%s[%d].lastName is required when shippingInfo is filled", fieldPath, index),
		}
	}
	if strings.TrimSpace(s.ZipCode) == "" {
		return &exceptions.GenericOpenAPIError{
			ErrorMsg: fmt.Sprintf("%s[%d].zipCode is required when shippingInfo is filled", fieldPath, index),
		}
	}
	return nil
}
