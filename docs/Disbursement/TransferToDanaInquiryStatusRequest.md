# TransferToDanaInquiryStatusRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OriginalPartnerReferenceNo** | **string** | Original transaction identifier on partner system | 
**OriginalReferenceNo** | Pointer to **string** | Original transaction identifier on DANA system | [optional] 
**OriginalExternalId** | Pointer to **string** | Original external identifier on header message | [optional] 
**ServiceCode** | **string** | Transaction type indicator is based on the service code of the original transaction request, value always 38 | [default to "38"]
**AdditionalInfo** | Pointer to **map[string]interface{}** | Additional information | [optional] 

## Methods

### NewTransferToDanaInquiryStatusRequest

`func NewTransferToDanaInquiryStatusRequest(originalPartnerReferenceNo string, serviceCode string, ) *TransferToDanaInquiryStatusRequest`

NewTransferToDanaInquiryStatusRequest instantiates a new TransferToDanaInquiryStatusRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransferToDanaInquiryStatusRequestWithDefaults

`func NewTransferToDanaInquiryStatusRequestWithDefaults() *TransferToDanaInquiryStatusRequest`

NewTransferToDanaInquiryStatusRequestWithDefaults instantiates a new TransferToDanaInquiryStatusRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOriginalPartnerReferenceNo

`func (o *TransferToDanaInquiryStatusRequest) GetOriginalPartnerReferenceNo() string`

GetOriginalPartnerReferenceNo returns the OriginalPartnerReferenceNo field if non-nil, zero value otherwise.

### GetOriginalPartnerReferenceNoOk

`func (o *TransferToDanaInquiryStatusRequest) GetOriginalPartnerReferenceNoOk() (*string, bool)`

GetOriginalPartnerReferenceNoOk returns a tuple with the OriginalPartnerReferenceNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalPartnerReferenceNo

`func (o *TransferToDanaInquiryStatusRequest) SetOriginalPartnerReferenceNo(v string)`

SetOriginalPartnerReferenceNo sets OriginalPartnerReferenceNo field to given value.


### GetOriginalReferenceNo

`func (o *TransferToDanaInquiryStatusRequest) GetOriginalReferenceNo() string`

GetOriginalReferenceNo returns the OriginalReferenceNo field if non-nil, zero value otherwise.

### GetOriginalReferenceNoOk

`func (o *TransferToDanaInquiryStatusRequest) GetOriginalReferenceNoOk() (*string, bool)`

GetOriginalReferenceNoOk returns a tuple with the OriginalReferenceNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalReferenceNo

`func (o *TransferToDanaInquiryStatusRequest) SetOriginalReferenceNo(v string)`

SetOriginalReferenceNo sets OriginalReferenceNo field to given value.

### HasOriginalReferenceNo

`func (o *TransferToDanaInquiryStatusRequest) HasOriginalReferenceNo() bool`

HasOriginalReferenceNo returns a boolean if a field has been set.

### GetOriginalExternalId

`func (o *TransferToDanaInquiryStatusRequest) GetOriginalExternalId() string`

GetOriginalExternalId returns the OriginalExternalId field if non-nil, zero value otherwise.

### GetOriginalExternalIdOk

`func (o *TransferToDanaInquiryStatusRequest) GetOriginalExternalIdOk() (*string, bool)`

GetOriginalExternalIdOk returns a tuple with the OriginalExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalExternalId

`func (o *TransferToDanaInquiryStatusRequest) SetOriginalExternalId(v string)`

SetOriginalExternalId sets OriginalExternalId field to given value.

### HasOriginalExternalId

`func (o *TransferToDanaInquiryStatusRequest) HasOriginalExternalId() bool`

HasOriginalExternalId returns a boolean if a field has been set.

### GetServiceCode

`func (o *TransferToDanaInquiryStatusRequest) GetServiceCode() string`

GetServiceCode returns the ServiceCode field if non-nil, zero value otherwise.

### GetServiceCodeOk

`func (o *TransferToDanaInquiryStatusRequest) GetServiceCodeOk() (*string, bool)`

GetServiceCodeOk returns a tuple with the ServiceCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceCode

`func (o *TransferToDanaInquiryStatusRequest) SetServiceCode(v string)`

SetServiceCode sets ServiceCode field to given value.


### GetAdditionalInfo

`func (o *TransferToDanaInquiryStatusRequest) GetAdditionalInfo() map[string]interface{}`

GetAdditionalInfo returns the AdditionalInfo field if non-nil, zero value otherwise.

### GetAdditionalInfoOk

`func (o *TransferToDanaInquiryStatusRequest) GetAdditionalInfoOk() (*map[string]interface{}, bool)`

GetAdditionalInfoOk returns a tuple with the AdditionalInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalInfo

`func (o *TransferToDanaInquiryStatusRequest) SetAdditionalInfo(v map[string]interface{})`

SetAdditionalInfo sets AdditionalInfo field to given value.

### HasAdditionalInfo

`func (o *TransferToDanaInquiryStatusRequest) HasAdditionalInfo() bool`

HasAdditionalInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


