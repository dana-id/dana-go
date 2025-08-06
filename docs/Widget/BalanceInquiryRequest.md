# BalanceInquiryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PartnerReferenceNo** | Pointer to **string** | Unique transaction identifier on partner system which assigned to each transaction<br /> Notes:<br /> If the partner receives a timeout or an unexpected response from DANA and partner expects to perform retry request to DANA, please use the partnerReferenceNo that is the same as the one used in the transaction request process before  | [optional] 
**BalanceTypes** | Pointer to **[]string** | Information of balance types to specify which balance type expected to be returned. Will return all available balance type if this parameter empty | [optional] 
**AdditionalInfo** | [**BalanceInquiryRequestAdditionalInfo**](BalanceInquiryRequestAdditionalInfo.md) | Additional information | 

## Methods

### NewBalanceInquiryRequest

`func NewBalanceInquiryRequest(additionalInfo BalanceInquiryRequestAdditionalInfo, ) *BalanceInquiryRequest`

NewBalanceInquiryRequest instantiates a new BalanceInquiryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBalanceInquiryRequestWithDefaults

`func NewBalanceInquiryRequestWithDefaults() *BalanceInquiryRequest`

NewBalanceInquiryRequestWithDefaults instantiates a new BalanceInquiryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPartnerReferenceNo

`func (o *BalanceInquiryRequest) GetPartnerReferenceNo() string`

GetPartnerReferenceNo returns the PartnerReferenceNo field if non-nil, zero value otherwise.

### GetPartnerReferenceNoOk

`func (o *BalanceInquiryRequest) GetPartnerReferenceNoOk() (*string, bool)`

GetPartnerReferenceNoOk returns a tuple with the PartnerReferenceNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerReferenceNo

`func (o *BalanceInquiryRequest) SetPartnerReferenceNo(v string)`

SetPartnerReferenceNo sets PartnerReferenceNo field to given value.

### HasPartnerReferenceNo

`func (o *BalanceInquiryRequest) HasPartnerReferenceNo() bool`

HasPartnerReferenceNo returns a boolean if a field has been set.

### GetBalanceTypes

`func (o *BalanceInquiryRequest) GetBalanceTypes() []string`

GetBalanceTypes returns the BalanceTypes field if non-nil, zero value otherwise.

### GetBalanceTypesOk

`func (o *BalanceInquiryRequest) GetBalanceTypesOk() (*[]string, bool)`

GetBalanceTypesOk returns a tuple with the BalanceTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceTypes

`func (o *BalanceInquiryRequest) SetBalanceTypes(v []string)`

SetBalanceTypes sets BalanceTypes field to given value.

### HasBalanceTypes

`func (o *BalanceInquiryRequest) HasBalanceTypes() bool`

HasBalanceTypes returns a boolean if a field has been set.

### GetAdditionalInfo

`func (o *BalanceInquiryRequest) GetAdditionalInfo() BalanceInquiryRequestAdditionalInfo`

GetAdditionalInfo returns the AdditionalInfo field if non-nil, zero value otherwise.

### GetAdditionalInfoOk

`func (o *BalanceInquiryRequest) GetAdditionalInfoOk() (*BalanceInquiryRequestAdditionalInfo, bool)`

GetAdditionalInfoOk returns a tuple with the AdditionalInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalInfo

`func (o *BalanceInquiryRequest) SetAdditionalInfo(v BalanceInquiryRequestAdditionalInfo)`

SetAdditionalInfo sets AdditionalInfo field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


