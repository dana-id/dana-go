# TransferToBankInquiryStatusRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OriginalPartnerReferenceNo** | Pointer to **string** | Original transaction identifier on partner system<br /> Notes:<br /> Required to be filled using the partnerReferenceNo that is the same as the one used in Transfer to Bank  | [optional] 
**OriginalReferenceNo** | Pointer to **string** | Original transaction identifier on DANA system | [optional] 
**OriginalExternalId** | Pointer to **string** | Original external identifier on header message | [optional] 
**ServiceCode** | **string** | Transaction type indicator is based on the service code of the original transaction request, value always 00 | [default to "00"]
**AdditionalInfo** | Pointer to **map[string]interface{}** | Additional information | [optional] 

## Methods

### NewTransferToBankInquiryStatusRequest

`func NewTransferToBankInquiryStatusRequest(serviceCode string, ) *TransferToBankInquiryStatusRequest`

NewTransferToBankInquiryStatusRequest instantiates a new TransferToBankInquiryStatusRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransferToBankInquiryStatusRequestWithDefaults

`func NewTransferToBankInquiryStatusRequestWithDefaults() *TransferToBankInquiryStatusRequest`

NewTransferToBankInquiryStatusRequestWithDefaults instantiates a new TransferToBankInquiryStatusRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOriginalPartnerReferenceNo

`func (o *TransferToBankInquiryStatusRequest) GetOriginalPartnerReferenceNo() string`

GetOriginalPartnerReferenceNo returns the OriginalPartnerReferenceNo field if non-nil, zero value otherwise.

### GetOriginalPartnerReferenceNoOk

`func (o *TransferToBankInquiryStatusRequest) GetOriginalPartnerReferenceNoOk() (*string, bool)`

GetOriginalPartnerReferenceNoOk returns a tuple with the OriginalPartnerReferenceNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalPartnerReferenceNo

`func (o *TransferToBankInquiryStatusRequest) SetOriginalPartnerReferenceNo(v string)`

SetOriginalPartnerReferenceNo sets OriginalPartnerReferenceNo field to given value.

### HasOriginalPartnerReferenceNo

`func (o *TransferToBankInquiryStatusRequest) HasOriginalPartnerReferenceNo() bool`

HasOriginalPartnerReferenceNo returns a boolean if a field has been set.

### GetOriginalReferenceNo

`func (o *TransferToBankInquiryStatusRequest) GetOriginalReferenceNo() string`

GetOriginalReferenceNo returns the OriginalReferenceNo field if non-nil, zero value otherwise.

### GetOriginalReferenceNoOk

`func (o *TransferToBankInquiryStatusRequest) GetOriginalReferenceNoOk() (*string, bool)`

GetOriginalReferenceNoOk returns a tuple with the OriginalReferenceNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalReferenceNo

`func (o *TransferToBankInquiryStatusRequest) SetOriginalReferenceNo(v string)`

SetOriginalReferenceNo sets OriginalReferenceNo field to given value.

### HasOriginalReferenceNo

`func (o *TransferToBankInquiryStatusRequest) HasOriginalReferenceNo() bool`

HasOriginalReferenceNo returns a boolean if a field has been set.

### GetOriginalExternalId

`func (o *TransferToBankInquiryStatusRequest) GetOriginalExternalId() string`

GetOriginalExternalId returns the OriginalExternalId field if non-nil, zero value otherwise.

### GetOriginalExternalIdOk

`func (o *TransferToBankInquiryStatusRequest) GetOriginalExternalIdOk() (*string, bool)`

GetOriginalExternalIdOk returns a tuple with the OriginalExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalExternalId

`func (o *TransferToBankInquiryStatusRequest) SetOriginalExternalId(v string)`

SetOriginalExternalId sets OriginalExternalId field to given value.

### HasOriginalExternalId

`func (o *TransferToBankInquiryStatusRequest) HasOriginalExternalId() bool`

HasOriginalExternalId returns a boolean if a field has been set.

### GetServiceCode

`func (o *TransferToBankInquiryStatusRequest) GetServiceCode() string`

GetServiceCode returns the ServiceCode field if non-nil, zero value otherwise.

### GetServiceCodeOk

`func (o *TransferToBankInquiryStatusRequest) GetServiceCodeOk() (*string, bool)`

GetServiceCodeOk returns a tuple with the ServiceCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceCode

`func (o *TransferToBankInquiryStatusRequest) SetServiceCode(v string)`

SetServiceCode sets ServiceCode field to given value.


### GetAdditionalInfo

`func (o *TransferToBankInquiryStatusRequest) GetAdditionalInfo() map[string]interface{}`

GetAdditionalInfo returns the AdditionalInfo field if non-nil, zero value otherwise.

### GetAdditionalInfoOk

`func (o *TransferToBankInquiryStatusRequest) GetAdditionalInfoOk() (*map[string]interface{}, bool)`

GetAdditionalInfoOk returns a tuple with the AdditionalInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalInfo

`func (o *TransferToBankInquiryStatusRequest) SetAdditionalInfo(v map[string]interface{})`

SetAdditionalInfo sets AdditionalInfo field to given value.

### HasAdditionalInfo

`func (o *TransferToBankInquiryStatusRequest) HasAdditionalInfo() bool`

HasAdditionalInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


