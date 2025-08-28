# DanaAccountInquiryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponseCode** | **string** | Refer to response code list | 
**ResponseMessage** | **string** | Refer to response code list | 
**ReferenceNo** | Pointer to **string** | Transaction identifier on DANA system | [optional] 
**PartnerReferenceNo** | Pointer to **string** | Unique transaction identifier on partner system which assigned to each transaction<br /> Notes:<br /> If the partner receives a timeout or an unexpected response from DANA and partner expects to perform retry request to DANA, please use the partnerReferenceNo that is the same as the one used in the transaction request process before  | [optional] 
**SessionId** | Pointer to **string** | Session identifier | [optional] 
**CustomerNumber** | Pointer to **string** | Customer account number, in format 628xxx | [optional] 
**CustomerName** | **string** | Customer account name | 
**CustomerMonthlyInLimit** | Pointer to **string** | Limitation of transfer to DANA balance for customer per month | [optional] 
**MinAmount** | [**Money**](Money.md) | Minimal amount. Contains two sub-fields:<br /> 1. Value: Amount, including the cents<br /> 2. Currency: Currency code based on ISO  | 
**MaxAmount** | [**Money**](Money.md) | Maximal amount. Contains two sub-fields:<br /> 1. Value: Amount, including the cents<br /> 2. Currency: Currency code based on ISO  | 
**Amount** | [**Money**](Money.md) | Amount. Contains two sub-fields:<br /> 1. Value: Transaction amount, including the cents<br /> 2. Currency: Currency code based on ISO  | 
**FeeAmount** | [**Money**](Money.md) | Fee amount. Contains two sub-fields:<br /> 1. Value: Amount, including the cents<br /> 2. Currency: Currency code based on ISO  | 
**FeeType** | Pointer to **string** | Type of fee for each transfer to DANA transaction. Such as admin fee | [optional] 
**AdditionalInfo** | Pointer to **map[string]interface{}** | Additional information | [optional] 

## Methods

### NewDanaAccountInquiryResponse

`func NewDanaAccountInquiryResponse(responseCode string, responseMessage string, customerName string, minAmount Money, maxAmount Money, amount Money, feeAmount Money, ) *DanaAccountInquiryResponse`

NewDanaAccountInquiryResponse instantiates a new DanaAccountInquiryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDanaAccountInquiryResponseWithDefaults

`func NewDanaAccountInquiryResponseWithDefaults() *DanaAccountInquiryResponse`

NewDanaAccountInquiryResponseWithDefaults instantiates a new DanaAccountInquiryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponseCode

`func (o *DanaAccountInquiryResponse) GetResponseCode() string`

GetResponseCode returns the ResponseCode field if non-nil, zero value otherwise.

### GetResponseCodeOk

`func (o *DanaAccountInquiryResponse) GetResponseCodeOk() (*string, bool)`

GetResponseCodeOk returns a tuple with the ResponseCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseCode

`func (o *DanaAccountInquiryResponse) SetResponseCode(v string)`

SetResponseCode sets ResponseCode field to given value.


### GetResponseMessage

`func (o *DanaAccountInquiryResponse) GetResponseMessage() string`

GetResponseMessage returns the ResponseMessage field if non-nil, zero value otherwise.

### GetResponseMessageOk

`func (o *DanaAccountInquiryResponse) GetResponseMessageOk() (*string, bool)`

GetResponseMessageOk returns a tuple with the ResponseMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseMessage

`func (o *DanaAccountInquiryResponse) SetResponseMessage(v string)`

SetResponseMessage sets ResponseMessage field to given value.


### GetReferenceNo

`func (o *DanaAccountInquiryResponse) GetReferenceNo() string`

GetReferenceNo returns the ReferenceNo field if non-nil, zero value otherwise.

### GetReferenceNoOk

`func (o *DanaAccountInquiryResponse) GetReferenceNoOk() (*string, bool)`

GetReferenceNoOk returns a tuple with the ReferenceNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceNo

`func (o *DanaAccountInquiryResponse) SetReferenceNo(v string)`

SetReferenceNo sets ReferenceNo field to given value.

### HasReferenceNo

`func (o *DanaAccountInquiryResponse) HasReferenceNo() bool`

HasReferenceNo returns a boolean if a field has been set.

### GetPartnerReferenceNo

`func (o *DanaAccountInquiryResponse) GetPartnerReferenceNo() string`

GetPartnerReferenceNo returns the PartnerReferenceNo field if non-nil, zero value otherwise.

### GetPartnerReferenceNoOk

`func (o *DanaAccountInquiryResponse) GetPartnerReferenceNoOk() (*string, bool)`

GetPartnerReferenceNoOk returns a tuple with the PartnerReferenceNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerReferenceNo

`func (o *DanaAccountInquiryResponse) SetPartnerReferenceNo(v string)`

SetPartnerReferenceNo sets PartnerReferenceNo field to given value.

### HasPartnerReferenceNo

`func (o *DanaAccountInquiryResponse) HasPartnerReferenceNo() bool`

HasPartnerReferenceNo returns a boolean if a field has been set.

### GetSessionId

`func (o *DanaAccountInquiryResponse) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *DanaAccountInquiryResponse) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *DanaAccountInquiryResponse) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *DanaAccountInquiryResponse) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### GetCustomerNumber

`func (o *DanaAccountInquiryResponse) GetCustomerNumber() string`

GetCustomerNumber returns the CustomerNumber field if non-nil, zero value otherwise.

### GetCustomerNumberOk

`func (o *DanaAccountInquiryResponse) GetCustomerNumberOk() (*string, bool)`

GetCustomerNumberOk returns a tuple with the CustomerNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerNumber

`func (o *DanaAccountInquiryResponse) SetCustomerNumber(v string)`

SetCustomerNumber sets CustomerNumber field to given value.

### HasCustomerNumber

`func (o *DanaAccountInquiryResponse) HasCustomerNumber() bool`

HasCustomerNumber returns a boolean if a field has been set.

### GetCustomerName

`func (o *DanaAccountInquiryResponse) GetCustomerName() string`

GetCustomerName returns the CustomerName field if non-nil, zero value otherwise.

### GetCustomerNameOk

`func (o *DanaAccountInquiryResponse) GetCustomerNameOk() (*string, bool)`

GetCustomerNameOk returns a tuple with the CustomerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerName

`func (o *DanaAccountInquiryResponse) SetCustomerName(v string)`

SetCustomerName sets CustomerName field to given value.


### GetCustomerMonthlyInLimit

`func (o *DanaAccountInquiryResponse) GetCustomerMonthlyInLimit() string`

GetCustomerMonthlyInLimit returns the CustomerMonthlyInLimit field if non-nil, zero value otherwise.

### GetCustomerMonthlyInLimitOk

`func (o *DanaAccountInquiryResponse) GetCustomerMonthlyInLimitOk() (*string, bool)`

GetCustomerMonthlyInLimitOk returns a tuple with the CustomerMonthlyInLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerMonthlyInLimit

`func (o *DanaAccountInquiryResponse) SetCustomerMonthlyInLimit(v string)`

SetCustomerMonthlyInLimit sets CustomerMonthlyInLimit field to given value.

### HasCustomerMonthlyInLimit

`func (o *DanaAccountInquiryResponse) HasCustomerMonthlyInLimit() bool`

HasCustomerMonthlyInLimit returns a boolean if a field has been set.

### GetMinAmount

`func (o *DanaAccountInquiryResponse) GetMinAmount() Money`

GetMinAmount returns the MinAmount field if non-nil, zero value otherwise.

### GetMinAmountOk

`func (o *DanaAccountInquiryResponse) GetMinAmountOk() (*Money, bool)`

GetMinAmountOk returns a tuple with the MinAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinAmount

`func (o *DanaAccountInquiryResponse) SetMinAmount(v Money)`

SetMinAmount sets MinAmount field to given value.


### GetMaxAmount

`func (o *DanaAccountInquiryResponse) GetMaxAmount() Money`

GetMaxAmount returns the MaxAmount field if non-nil, zero value otherwise.

### GetMaxAmountOk

`func (o *DanaAccountInquiryResponse) GetMaxAmountOk() (*Money, bool)`

GetMaxAmountOk returns a tuple with the MaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAmount

`func (o *DanaAccountInquiryResponse) SetMaxAmount(v Money)`

SetMaxAmount sets MaxAmount field to given value.


### GetAmount

`func (o *DanaAccountInquiryResponse) GetAmount() Money`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *DanaAccountInquiryResponse) GetAmountOk() (*Money, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *DanaAccountInquiryResponse) SetAmount(v Money)`

SetAmount sets Amount field to given value.


### GetFeeAmount

`func (o *DanaAccountInquiryResponse) GetFeeAmount() Money`

GetFeeAmount returns the FeeAmount field if non-nil, zero value otherwise.

### GetFeeAmountOk

`func (o *DanaAccountInquiryResponse) GetFeeAmountOk() (*Money, bool)`

GetFeeAmountOk returns a tuple with the FeeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeAmount

`func (o *DanaAccountInquiryResponse) SetFeeAmount(v Money)`

SetFeeAmount sets FeeAmount field to given value.


### GetFeeType

`func (o *DanaAccountInquiryResponse) GetFeeType() string`

GetFeeType returns the FeeType field if non-nil, zero value otherwise.

### GetFeeTypeOk

`func (o *DanaAccountInquiryResponse) GetFeeTypeOk() (*string, bool)`

GetFeeTypeOk returns a tuple with the FeeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeType

`func (o *DanaAccountInquiryResponse) SetFeeType(v string)`

SetFeeType sets FeeType field to given value.

### HasFeeType

`func (o *DanaAccountInquiryResponse) HasFeeType() bool`

HasFeeType returns a boolean if a field has been set.

### GetAdditionalInfo

`func (o *DanaAccountInquiryResponse) GetAdditionalInfo() map[string]interface{}`

GetAdditionalInfo returns the AdditionalInfo field if non-nil, zero value otherwise.

### GetAdditionalInfoOk

`func (o *DanaAccountInquiryResponse) GetAdditionalInfoOk() (*map[string]interface{}, bool)`

GetAdditionalInfoOk returns a tuple with the AdditionalInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalInfo

`func (o *DanaAccountInquiryResponse) SetAdditionalInfo(v map[string]interface{})`

SetAdditionalInfo sets AdditionalInfo field to given value.

### HasAdditionalInfo

`func (o *DanaAccountInquiryResponse) HasAdditionalInfo() bool`

HasAdditionalInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


