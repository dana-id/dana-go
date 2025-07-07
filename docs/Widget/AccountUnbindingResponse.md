# AccountUnbindingResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponseCode** | **string** | Response code. Refer to https://dashboard.dana.id/api-docs/read/108#HTML-AccountUnbinding-ResponseCodeandMessage | 
**ResponseMessage** | **string** | Response message. Refer to https://dashboard.dana.id/api-docs/read/108#HTML-AccountUnbinding-ResponseCodeandMessage | 
**ReferenceNo** | Pointer to **string** | Transaction identifier on DANA system | [optional] 
**PartnerReferenceNo** | Pointer to **string** | Unique transaction identifier on partner system which assigned to each transaction | [optional] 
**MerchantId** | Pointer to **string** | Merchant identifier that is unique per each merchant | [optional] 
**SubMerchantId** | Pointer to **string** | Information of sub merchant identifier | [optional] 
**LinkId** | Pointer to **string** | Information of link identifier | [optional] 
**UnlinkResult** | Pointer to **string** | Result of unlinking process | [optional] 
**AdditionalInfo** | Pointer to **map[string]interface{}** | Additional information | [optional] 

## Methods

### NewAccountUnbindingResponse

`func NewAccountUnbindingResponse(responseCode string, responseMessage string, ) *AccountUnbindingResponse`

NewAccountUnbindingResponse instantiates a new AccountUnbindingResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountUnbindingResponseWithDefaults

`func NewAccountUnbindingResponseWithDefaults() *AccountUnbindingResponse`

NewAccountUnbindingResponseWithDefaults instantiates a new AccountUnbindingResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponseCode

`func (o *AccountUnbindingResponse) GetResponseCode() string`

GetResponseCode returns the ResponseCode field if non-nil, zero value otherwise.

### GetResponseCodeOk

`func (o *AccountUnbindingResponse) GetResponseCodeOk() (*string, bool)`

GetResponseCodeOk returns a tuple with the ResponseCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseCode

`func (o *AccountUnbindingResponse) SetResponseCode(v string)`

SetResponseCode sets ResponseCode field to given value.


### GetResponseMessage

`func (o *AccountUnbindingResponse) GetResponseMessage() string`

GetResponseMessage returns the ResponseMessage field if non-nil, zero value otherwise.

### GetResponseMessageOk

`func (o *AccountUnbindingResponse) GetResponseMessageOk() (*string, bool)`

GetResponseMessageOk returns a tuple with the ResponseMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseMessage

`func (o *AccountUnbindingResponse) SetResponseMessage(v string)`

SetResponseMessage sets ResponseMessage field to given value.


### GetReferenceNo

`func (o *AccountUnbindingResponse) GetReferenceNo() string`

GetReferenceNo returns the ReferenceNo field if non-nil, zero value otherwise.

### GetReferenceNoOk

`func (o *AccountUnbindingResponse) GetReferenceNoOk() (*string, bool)`

GetReferenceNoOk returns a tuple with the ReferenceNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceNo

`func (o *AccountUnbindingResponse) SetReferenceNo(v string)`

SetReferenceNo sets ReferenceNo field to given value.

### HasReferenceNo

`func (o *AccountUnbindingResponse) HasReferenceNo() bool`

HasReferenceNo returns a boolean if a field has been set.

### GetPartnerReferenceNo

`func (o *AccountUnbindingResponse) GetPartnerReferenceNo() string`

GetPartnerReferenceNo returns the PartnerReferenceNo field if non-nil, zero value otherwise.

### GetPartnerReferenceNoOk

`func (o *AccountUnbindingResponse) GetPartnerReferenceNoOk() (*string, bool)`

GetPartnerReferenceNoOk returns a tuple with the PartnerReferenceNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerReferenceNo

`func (o *AccountUnbindingResponse) SetPartnerReferenceNo(v string)`

SetPartnerReferenceNo sets PartnerReferenceNo field to given value.

### HasPartnerReferenceNo

`func (o *AccountUnbindingResponse) HasPartnerReferenceNo() bool`

HasPartnerReferenceNo returns a boolean if a field has been set.

### GetMerchantId

`func (o *AccountUnbindingResponse) GetMerchantId() string`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *AccountUnbindingResponse) GetMerchantIdOk() (*string, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *AccountUnbindingResponse) SetMerchantId(v string)`

SetMerchantId sets MerchantId field to given value.

### HasMerchantId

`func (o *AccountUnbindingResponse) HasMerchantId() bool`

HasMerchantId returns a boolean if a field has been set.

### GetSubMerchantId

`func (o *AccountUnbindingResponse) GetSubMerchantId() string`

GetSubMerchantId returns the SubMerchantId field if non-nil, zero value otherwise.

### GetSubMerchantIdOk

`func (o *AccountUnbindingResponse) GetSubMerchantIdOk() (*string, bool)`

GetSubMerchantIdOk returns a tuple with the SubMerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubMerchantId

`func (o *AccountUnbindingResponse) SetSubMerchantId(v string)`

SetSubMerchantId sets SubMerchantId field to given value.

### HasSubMerchantId

`func (o *AccountUnbindingResponse) HasSubMerchantId() bool`

HasSubMerchantId returns a boolean if a field has been set.

### GetLinkId

`func (o *AccountUnbindingResponse) GetLinkId() string`

GetLinkId returns the LinkId field if non-nil, zero value otherwise.

### GetLinkIdOk

`func (o *AccountUnbindingResponse) GetLinkIdOk() (*string, bool)`

GetLinkIdOk returns a tuple with the LinkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkId

`func (o *AccountUnbindingResponse) SetLinkId(v string)`

SetLinkId sets LinkId field to given value.

### HasLinkId

`func (o *AccountUnbindingResponse) HasLinkId() bool`

HasLinkId returns a boolean if a field has been set.

### GetUnlinkResult

`func (o *AccountUnbindingResponse) GetUnlinkResult() string`

GetUnlinkResult returns the UnlinkResult field if non-nil, zero value otherwise.

### GetUnlinkResultOk

`func (o *AccountUnbindingResponse) GetUnlinkResultOk() (*string, bool)`

GetUnlinkResultOk returns a tuple with the UnlinkResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnlinkResult

`func (o *AccountUnbindingResponse) SetUnlinkResult(v string)`

SetUnlinkResult sets UnlinkResult field to given value.

### HasUnlinkResult

`func (o *AccountUnbindingResponse) HasUnlinkResult() bool`

HasUnlinkResult returns a boolean if a field has been set.

### GetAdditionalInfo

`func (o *AccountUnbindingResponse) GetAdditionalInfo() map[string]interface{}`

GetAdditionalInfo returns the AdditionalInfo field if non-nil, zero value otherwise.

### GetAdditionalInfoOk

`func (o *AccountUnbindingResponse) GetAdditionalInfoOk() (*map[string]interface{}, bool)`

GetAdditionalInfoOk returns a tuple with the AdditionalInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalInfo

`func (o *AccountUnbindingResponse) SetAdditionalInfo(v map[string]interface{})`

SetAdditionalInfo sets AdditionalInfo field to given value.

### HasAdditionalInfo

`func (o *AccountUnbindingResponse) HasAdditionalInfo() bool`

HasAdditionalInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


