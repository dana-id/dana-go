# BankAccountInquiryResponseAdditionalInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FeeAmount** | [**Money**](Money.md) | Additional information of fee amount. Contains two sub-fields:<br /> 1. Value: Amount, including the cents<br /> 2. Currency: Currency code based on ISO  | 

## Methods

### NewBankAccountInquiryResponseAdditionalInfo

`func NewBankAccountInquiryResponseAdditionalInfo(feeAmount Money, ) *BankAccountInquiryResponseAdditionalInfo`

NewBankAccountInquiryResponseAdditionalInfo instantiates a new BankAccountInquiryResponseAdditionalInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBankAccountInquiryResponseAdditionalInfoWithDefaults

`func NewBankAccountInquiryResponseAdditionalInfoWithDefaults() *BankAccountInquiryResponseAdditionalInfo`

NewBankAccountInquiryResponseAdditionalInfoWithDefaults instantiates a new BankAccountInquiryResponseAdditionalInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeeAmount

`func (o *BankAccountInquiryResponseAdditionalInfo) GetFeeAmount() Money`

GetFeeAmount returns the FeeAmount field if non-nil, zero value otherwise.

### GetFeeAmountOk

`func (o *BankAccountInquiryResponseAdditionalInfo) GetFeeAmountOk() (*Money, bool)`

GetFeeAmountOk returns a tuple with the FeeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeAmount

`func (o *BankAccountInquiryResponseAdditionalInfo) SetFeeAmount(v Money)`

SetFeeAmount sets FeeAmount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


