# QueryAssetCardListRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MemberId** | **string** | User identifier or merchant identifier | 
**BindingId** | Pointer to **string** | Asset card bind identifier | [optional] 
**EnableOnly** | Pointer to **string** | Return only active cards when &#x60;\&quot;true\&quot;&#x60; | [optional] 
**ContactBizTypeList** | Pointer to **[]string** | Contact biz type list (&#x60;ContactBizTypeEnum&#x60;) | [optional] 
**AssetTypeList** | Pointer to **[]string** | Asset type list (&#x60;AssetCardTypeEnum&#x60;) | [optional] 

## Methods

### NewQueryAssetCardListRequest

`func NewQueryAssetCardListRequest(memberId string, ) *QueryAssetCardListRequest`

NewQueryAssetCardListRequest instantiates a new QueryAssetCardListRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryAssetCardListRequestWithDefaults

`func NewQueryAssetCardListRequestWithDefaults() *QueryAssetCardListRequest`

NewQueryAssetCardListRequestWithDefaults instantiates a new QueryAssetCardListRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMemberId

`func (o *QueryAssetCardListRequest) GetMemberId() string`

GetMemberId returns the MemberId field if non-nil, zero value otherwise.

### GetMemberIdOk

`func (o *QueryAssetCardListRequest) GetMemberIdOk() (*string, bool)`

GetMemberIdOk returns a tuple with the MemberId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberId

`func (o *QueryAssetCardListRequest) SetMemberId(v string)`

SetMemberId sets MemberId field to given value.


### GetBindingId

`func (o *QueryAssetCardListRequest) GetBindingId() string`

GetBindingId returns the BindingId field if non-nil, zero value otherwise.

### GetBindingIdOk

`func (o *QueryAssetCardListRequest) GetBindingIdOk() (*string, bool)`

GetBindingIdOk returns a tuple with the BindingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindingId

`func (o *QueryAssetCardListRequest) SetBindingId(v string)`

SetBindingId sets BindingId field to given value.

### HasBindingId

`func (o *QueryAssetCardListRequest) HasBindingId() bool`

HasBindingId returns a boolean if a field has been set.

### GetEnableOnly

`func (o *QueryAssetCardListRequest) GetEnableOnly() string`

GetEnableOnly returns the EnableOnly field if non-nil, zero value otherwise.

### GetEnableOnlyOk

`func (o *QueryAssetCardListRequest) GetEnableOnlyOk() (*string, bool)`

GetEnableOnlyOk returns a tuple with the EnableOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableOnly

`func (o *QueryAssetCardListRequest) SetEnableOnly(v string)`

SetEnableOnly sets EnableOnly field to given value.

### HasEnableOnly

`func (o *QueryAssetCardListRequest) HasEnableOnly() bool`

HasEnableOnly returns a boolean if a field has been set.

### GetContactBizTypeList

`func (o *QueryAssetCardListRequest) GetContactBizTypeList() []string`

GetContactBizTypeList returns the ContactBizTypeList field if non-nil, zero value otherwise.

### GetContactBizTypeListOk

`func (o *QueryAssetCardListRequest) GetContactBizTypeListOk() (*[]string, bool)`

GetContactBizTypeListOk returns a tuple with the ContactBizTypeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactBizTypeList

`func (o *QueryAssetCardListRequest) SetContactBizTypeList(v []string)`

SetContactBizTypeList sets ContactBizTypeList field to given value.

### HasContactBizTypeList

`func (o *QueryAssetCardListRequest) HasContactBizTypeList() bool`

HasContactBizTypeList returns a boolean if a field has been set.

### GetAssetTypeList

`func (o *QueryAssetCardListRequest) GetAssetTypeList() []string`

GetAssetTypeList returns the AssetTypeList field if non-nil, zero value otherwise.

### GetAssetTypeListOk

`func (o *QueryAssetCardListRequest) GetAssetTypeListOk() (*[]string, bool)`

GetAssetTypeListOk returns a tuple with the AssetTypeList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetTypeList

`func (o *QueryAssetCardListRequest) SetAssetTypeList(v []string)`

SetAssetTypeList sets AssetTypeList field to given value.

### HasAssetTypeList

`func (o *QueryAssetCardListRequest) HasAssetTypeList() bool`

HasAssetTypeList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


