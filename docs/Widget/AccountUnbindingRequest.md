# AccountUnbindingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MerchantId** | **string** | Merchant identifier that is unique per each merchant | 
**SubMerchantId** | Pointer to **string** | Information of sub merchant identifier | [optional] 
**PartnerReferenceNo** | Pointer to **string** | Unique transaction identifier on partner system which assigned to each transaction | [optional] 
**LinkId** | Pointer to **string** | Information of link identifier | [optional] 
**TokenId** | Pointer to **string** | Information of token identifier | [optional] 
**AdditionalInfo** | [**AccountUnbindingRequestAdditionalInfo**](AccountUnbindingRequestAdditionalInfo.md) | Additional information | 

## Methods

### NewAccountUnbindingRequest

`func NewAccountUnbindingRequest(merchantId string, additionalInfo AccountUnbindingRequestAdditionalInfo, ) *AccountUnbindingRequest`

NewAccountUnbindingRequest instantiates a new AccountUnbindingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountUnbindingRequestWithDefaults

`func NewAccountUnbindingRequestWithDefaults() *AccountUnbindingRequest`

NewAccountUnbindingRequestWithDefaults instantiates a new AccountUnbindingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMerchantId

`func (o *AccountUnbindingRequest) GetMerchantId() string`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *AccountUnbindingRequest) GetMerchantIdOk() (*string, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *AccountUnbindingRequest) SetMerchantId(v string)`

SetMerchantId sets MerchantId field to given value.


### GetSubMerchantId

`func (o *AccountUnbindingRequest) GetSubMerchantId() string`

GetSubMerchantId returns the SubMerchantId field if non-nil, zero value otherwise.

### GetSubMerchantIdOk

`func (o *AccountUnbindingRequest) GetSubMerchantIdOk() (*string, bool)`

GetSubMerchantIdOk returns a tuple with the SubMerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubMerchantId

`func (o *AccountUnbindingRequest) SetSubMerchantId(v string)`

SetSubMerchantId sets SubMerchantId field to given value.

### HasSubMerchantId

`func (o *AccountUnbindingRequest) HasSubMerchantId() bool`

HasSubMerchantId returns a boolean if a field has been set.

### GetPartnerReferenceNo

`func (o *AccountUnbindingRequest) GetPartnerReferenceNo() string`

GetPartnerReferenceNo returns the PartnerReferenceNo field if non-nil, zero value otherwise.

### GetPartnerReferenceNoOk

`func (o *AccountUnbindingRequest) GetPartnerReferenceNoOk() (*string, bool)`

GetPartnerReferenceNoOk returns a tuple with the PartnerReferenceNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerReferenceNo

`func (o *AccountUnbindingRequest) SetPartnerReferenceNo(v string)`

SetPartnerReferenceNo sets PartnerReferenceNo field to given value.

### HasPartnerReferenceNo

`func (o *AccountUnbindingRequest) HasPartnerReferenceNo() bool`

HasPartnerReferenceNo returns a boolean if a field has been set.

### GetLinkId

`func (o *AccountUnbindingRequest) GetLinkId() string`

GetLinkId returns the LinkId field if non-nil, zero value otherwise.

### GetLinkIdOk

`func (o *AccountUnbindingRequest) GetLinkIdOk() (*string, bool)`

GetLinkIdOk returns a tuple with the LinkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkId

`func (o *AccountUnbindingRequest) SetLinkId(v string)`

SetLinkId sets LinkId field to given value.

### HasLinkId

`func (o *AccountUnbindingRequest) HasLinkId() bool`

HasLinkId returns a boolean if a field has been set.

### GetTokenId

`func (o *AccountUnbindingRequest) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *AccountUnbindingRequest) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *AccountUnbindingRequest) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.

### HasTokenId

`func (o *AccountUnbindingRequest) HasTokenId() bool`

HasTokenId returns a boolean if a field has been set.

### GetAdditionalInfo

`func (o *AccountUnbindingRequest) GetAdditionalInfo() AccountUnbindingRequestAdditionalInfo`

GetAdditionalInfo returns the AdditionalInfo field if non-nil, zero value otherwise.

### GetAdditionalInfoOk

`func (o *AccountUnbindingRequest) GetAdditionalInfoOk() (*AccountUnbindingRequestAdditionalInfo, bool)`

GetAdditionalInfoOk returns a tuple with the AdditionalInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalInfo

`func (o *AccountUnbindingRequest) SetAdditionalInfo(v AccountUnbindingRequestAdditionalInfo)`

SetAdditionalInfo sets AdditionalInfo field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


