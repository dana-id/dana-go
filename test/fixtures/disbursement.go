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

package fixtures

import (
	"fmt"
	"time"

	disbursement "github.com/dana-id/dana-go/disbursement/v1"
)

// Helper function to create string pointer
func strPtr(s string) *string {
	return &s
}

// Helper function to create bool pointer
func boolPtr(b bool) *bool {
	return &b
}

// Helper function to format current time for DANA API
func getCurrentTransactionDate() string {
	return time.Now().Format("2006-01-02T15:04:05+07:00")
}

// Helper function to generate unique reference number with timestamp
func generateUniqueRef(prefix string) string {
	return fmt.Sprintf("%s-%s", prefix, time.Now().Format("20060102150405000"))
}

// GetDanaAccountInquiryRequest returns a fixture for DANA account inquiry request
func GetDanaAccountInquiryRequest() disbursement.DanaAccountInquiryRequest {
	additionalInfo := disbursement.DanaAccountInquiryRequestAdditionalInfo{
		FundType: "AGENT_TOPUP_FOR_USER_SETTLE",
	}

	amount := disbursement.Money{
		Value:    "1.00",
		Currency: "IDR",
	}

	request := disbursement.DanaAccountInquiryRequest{
		PartnerReferenceNo: strPtr(generateUniqueRef("ACCINQ")),
		CustomerNumber:     strPtr("62811742234"),
		Amount:             amount,
		TransactionDate:    strPtr(getCurrentTransactionDate()),
		AdditionalInfo:     additionalInfo,
	}

	return request
}

// GetDanaAccountInquiryRequestWithAccessToken returns a fixture for DANA account inquiry request with access token
func GetDanaAccountInquiryRequestWithAccessToken() disbursement.DanaAccountInquiryRequest {
	additionalInfo := disbursement.DanaAccountInquiryRequestAdditionalInfo{
		FundType:    "AGENT_TOPUP_FOR_USER_SETTLE",
		AccessToken: strPtr("test_access_token_123456"),
	}

	amount := disbursement.Money{
		Value:    "1.00",
		Currency: "IDR",
	}

	request := disbursement.DanaAccountInquiryRequest{
		PartnerReferenceNo: strPtr(generateUniqueRef("ACCINQ-TOKEN")),
		CustomerNumber:     strPtr("62811742234"),
		Amount:             amount,
		TransactionDate:    strPtr(getCurrentTransactionDate()),
		AdditionalInfo:     additionalInfo,
	}

	return request
}

// GetDanaAccountInquiryRequestWithCustomerId returns a fixture for DANA account inquiry request with customer ID
func GetDanaAccountInquiryRequestWithCustomerId() disbursement.DanaAccountInquiryRequest {
	additionalInfo := disbursement.DanaAccountInquiryRequestAdditionalInfo{
		FundType:   "AGENT_TOPUP_FOR_USER_SETTLE",
		CustomerId: strPtr("test_customer_123"),
	}

	amount := disbursement.Money{
		Value:    "1.00",
		Currency: "IDR",
	}

	request := disbursement.DanaAccountInquiryRequest{
		PartnerReferenceNo: strPtr(generateUniqueRef("ACCINQ-CUSTID")),
		CustomerNumber:     strPtr("620000000000"), // Default phone number format when using customer_id
		Amount:             amount,
		TransactionDate:    strPtr(getCurrentTransactionDate()),
		AdditionalInfo:     additionalInfo,
	}

	return request
}

// GetBankAccountInquiryRequest returns a fixture for bank account inquiry request
func GetBankAccountInquiryRequest() disbursement.BankAccountInquiryRequest {
	additionalInfo := disbursement.BankAccountInquiryRequestAdditionalInfo{
		FundType:            "MERCHANT_WITHDRAW_FOR_CORPORATE",
		BeneficiaryBankCode: "014",
	}

	amount := disbursement.Money{
		Value:    "50000.00",
		Currency: "IDR",
	}

	request := disbursement.BankAccountInquiryRequest{
		PartnerReferenceNo:       strPtr(generateUniqueRef("BANKINQ")),
		CustomerNumber:           "62811742234",
		BeneficiaryAccountNumber: "2460888509",
		Amount:                   amount,
		AdditionalInfo:           additionalInfo,
	}

	return request
}

// GetBankAccountInquiryRequestWithAccountName returns a fixture for bank account inquiry request with beneficiary account name
func GetBankAccountInquiryRequestWithAccountName() disbursement.BankAccountInquiryRequest {
	additionalInfo := disbursement.BankAccountInquiryRequestAdditionalInfo{
		FundType:               "MERCHANT_WITHDRAW_FOR_CORPORATE",
		BeneficiaryBankCode:    "014",
		BeneficiaryAccountName: strPtr("Test Account Holder"),
	}

	amount := disbursement.Money{
		Value:    "1.00",
		Currency: "IDR",
	}

	request := disbursement.BankAccountInquiryRequest{
		PartnerReferenceNo:       strPtr(generateUniqueRef("BANKINQ-NAME")),
		CustomerNumber:           "62811742234",
		BeneficiaryAccountNumber: "2460888509",
		Amount:                   amount,
		AdditionalInfo:           additionalInfo,
	}

	return request
}

// GetTransferToBankRequest returns a fixture for transfer to bank request
func GetTransferToBankRequest() disbursement.TransferToBankRequest {
	additionalInfo := disbursement.TransferToBankRequestAdditionalInfo{
		FundType: "MERCHANT_WITHDRAW_FOR_CORPORATE",
	}

	amount := disbursement.Money{
		Value:    "50000.00",
		Currency: "IDR",
	}

	request := disbursement.TransferToBankRequest{
		PartnerReferenceNo:       strPtr(generateUniqueRef("TRANSFER")),
		CustomerNumber:           "62811742234",
		BeneficiaryAccountNumber: "2460888509",
		BeneficiaryBankCode:      "014",
		Amount:                   amount,
		AdditionalInfo:           additionalInfo,
	}

	return request
}

// GetTransferToBankRequestWithNotification returns a fixture for transfer to bank request with notification enabled
func GetTransferToBankRequestWithNotification() disbursement.TransferToBankRequest {
	additionalInfo := disbursement.TransferToBankRequestAdditionalInfo{
		FundType:               "MERCHANT_WITHDRAW_FOR_CORPORATE",
		NeedNotify:             boolPtr(true),
		BeneficiaryAccountName: strPtr("Test Account Holder"),
	}

	amount := disbursement.Money{
		Value:    "1.00",
		Currency: "IDR",
	}

	request := disbursement.TransferToBankRequest{
		PartnerReferenceNo:       strPtr(generateUniqueRef("TRANSFER-NOTIFY")),
		CustomerNumber:           "62811742234",
		BeneficiaryAccountNumber: "2460888509",
		BeneficiaryBankCode:      "014",
		Amount:                   amount,
		AdditionalInfo:           additionalInfo,
	}

	return request
}

// GetTransferToBankInquiryStatusRequest returns a fixture for transfer to bank inquiry status request
func GetTransferToBankInquiryStatusRequest() disbursement.TransferToBankInquiryStatusRequest {
	request := disbursement.TransferToBankInquiryStatusRequest{
		OriginalPartnerReferenceNo: strPtr(generateUniqueRef("TRANSFER")),
		ServiceCode:                "00", // Service code for transfer to bank inquiry status
	}

	return request
}

// GetTransferToBankInquiryStatusRequestWithReferences returns a fixture for transfer to bank inquiry status request with all reference numbers
func GetTransferToBankInquiryStatusRequestWithReferences() disbursement.TransferToBankInquiryStatusRequest {
	request := disbursement.TransferToBankInquiryStatusRequest{
		OriginalPartnerReferenceNo: strPtr(generateUniqueRef("TRANSFER-REF")),
		OriginalReferenceNo:        strPtr("DANA123456789"),
		OriginalExternalId:         strPtr("EXT123456789"),
		ServiceCode:                "00", // Service code for transfer to bank inquiry status
	}

	return request
}

// GetTransferToDanaRequest returns a fixture for transfer to DANA request
func GetTransferToDanaRequest() disbursement.TransferToDanaRequest {
	additionalInfo := disbursement.TransferToDanaRequestAdditionalInfo{
		FundType: "AGENT_TOPUP_FOR_USER_SETTLE",
	}

	amount := disbursement.Money{
		Value:    "1.00",
		Currency: "IDR",
	}

	feeAmount := disbursement.Money{
		Value:    "1.00",
		Currency: "IDR",
	}

	request := disbursement.TransferToDanaRequest{
		PartnerReferenceNo: generateUniqueRef("TOPUP"),
		CustomerNumber:     strPtr("62811742234"),
		Amount:             amount,
		FeeAmount:          feeAmount,
		TransactionDate:    strPtr(getCurrentTransactionDate()),
		AdditionalInfo:     additionalInfo,
	}

	return request
}

// GetTransferToDanaRequestWithDivision returns a fixture for transfer to DANA request with division charge target
func GetTransferToDanaRequestWithDivision() disbursement.TransferToDanaRequest {
	additionalInfo := disbursement.TransferToDanaRequestAdditionalInfo{
		FundType:           "AGENT_TOPUP_FOR_USER_SETTLE",
		ExternalDivisionId: strPtr("DIV123456"),
		ChargeTarget:       strPtr("DIVISION"),
		AccountType:        strPtr("DANA_WALLET"),
	}

	amount := disbursement.Money{
		Value:    "1.00",
		Currency: "IDR",
	}

	feeAmount := disbursement.Money{
		Value:    "1.00",
		Currency: "IDR",
	}

	request := disbursement.TransferToDanaRequest{
		PartnerReferenceNo: generateUniqueRef("TOPUP-DIV"),
		CustomerNumber:     strPtr("62811742234"),
		Amount:             amount,
		FeeAmount:          feeAmount,
		TransactionDate:    strPtr(getCurrentTransactionDate()),
		Notes:              strPtr("Test transfer to DANA with division"),
		AdditionalInfo:     additionalInfo,
	}

	return request
}

// GetTransferToDanaInquiryStatusRequest returns a fixture for transfer to DANA inquiry status request
func GetTransferToDanaInquiryStatusRequest() disbursement.TransferToDanaInquiryStatusRequest {
	request := disbursement.TransferToDanaInquiryStatusRequest{
		OriginalPartnerReferenceNo: generateUniqueRef("TOPUP"),
		ServiceCode:                "38", // Service code for transfer to DANA
	}

	return request
}

// GetTransferToDanaInquiryStatusRequestWithReference returns a fixture for transfer to DANA inquiry status request with DANA reference
func GetTransferToDanaInquiryStatusRequestWithReference() disbursement.TransferToDanaInquiryStatusRequest {
	request := disbursement.TransferToDanaInquiryStatusRequest{
		OriginalPartnerReferenceNo: generateUniqueRef("TOPUP-REF"),
		OriginalReferenceNo:        strPtr("DANA123456789"),
		ServiceCode:                "38", // Service code for transfer to DANA
	}

	return request
}

// GetDynamicDanaAccountInquiryRequest returns a unique DANA account inquiry request for each call
func GetDynamicDanaAccountInquiryRequest() disbursement.DanaAccountInquiryRequest {
	additionalInfo := disbursement.DanaAccountInquiryRequestAdditionalInfo{
		FundType: "AGENT_TOPUP_FOR_USER_SETTLE",
	}

	amount := disbursement.Money{
		Value:    "1.00",
		Currency: "IDR",
	}

	request := disbursement.DanaAccountInquiryRequest{
		PartnerReferenceNo: strPtr(generateUniqueRef("ACCINQ")),
		CustomerNumber:     strPtr("62811742234"),
		Amount:             amount,
		TransactionDate:    strPtr(getCurrentTransactionDate()),
		AdditionalInfo:     additionalInfo,
	}

	return request
}

// GetDynamicTransferToDanaRequest returns a unique transfer to DANA request for each call
func GetDynamicTransferToDanaRequest() disbursement.TransferToDanaRequest {
	additionalInfo := disbursement.TransferToDanaRequestAdditionalInfo{
		FundType:    "AGENT_TOPUP_FOR_USER_SETTLE",
		AccountType: strPtr("DANA_WALLET"),
	}

	amount := disbursement.Money{
		Value:    "1.00",
		Currency: "IDR",
	}

	feeAmount := disbursement.Money{
		Value:    "1.00",
		Currency: "IDR",
	}

	request := disbursement.TransferToDanaRequest{
		PartnerReferenceNo: generateUniqueRef("TOPUP"),
		CustomerNumber:     strPtr("62811742234"),
		Amount:             amount,
		FeeAmount:          feeAmount,
		TransactionDate:    strPtr(getCurrentTransactionDate()),
		Notes:              strPtr("Test transfer to DANA"),
		AdditionalInfo:     additionalInfo,
	}

	return request
}

// GetDynamicBankAccountInquiryRequest returns a unique bank account inquiry request for each call
func GetDynamicBankAccountInquiryRequest() disbursement.BankAccountInquiryRequest {
	additionalInfo := disbursement.BankAccountInquiryRequestAdditionalInfo{
		FundType:            "MERCHANT_WITHDRAW_FOR_CORPORATE",
		BeneficiaryBankCode: "014",
	}

	amount := disbursement.Money{
		Value:    "1.00",
		Currency: "IDR",
	}

	request := disbursement.BankAccountInquiryRequest{
		PartnerReferenceNo:       strPtr(generateUniqueRef("BANKINQ")),
		CustomerNumber:           "62811742234",
		BeneficiaryAccountNumber: "2460888509",
		Amount:                   amount,
		AdditionalInfo:           additionalInfo,
	}

	return request
}

// GetDynamicTransferToBankRequest returns a unique transfer to bank request for each call
func GetDynamicTransferToBankRequest() disbursement.TransferToBankRequest {
	additionalInfo := disbursement.TransferToBankRequestAdditionalInfo{
		FundType: "MERCHANT_WITHDRAW_FOR_CORPORATE",
	}

	amount := disbursement.Money{
		Value:    "1.00",
		Currency: "IDR",
	}

	request := disbursement.TransferToBankRequest{
		PartnerReferenceNo:       strPtr(generateUniqueRef("TRANSFER")),
		CustomerNumber:           "62811742234",
		BeneficiaryAccountNumber: "2460888509",
		BeneficiaryBankCode:      "014",
		Amount:                   amount,
		AdditionalInfo:           additionalInfo,
	}

	return request
}
