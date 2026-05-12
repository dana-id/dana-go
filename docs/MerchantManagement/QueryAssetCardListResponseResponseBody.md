# QueryAssetCardListResponseResponseBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResultInfo** | [**MemberAssetResultInfo**](MemberAssetResultInfo.md) |  | 
**AssetCardList** | Pointer to [**[]AssetCardListItem**](AssetCardListItem.md) | Asset card list detail (present when applicable) | [optional] 

## Methods

### NewQueryAssetCardListResponseResponseBody

`func NewQueryAssetCardListResponseResponseBody(resultInfo MemberAssetResultInfo, ) *QueryAssetCardListResponseResponseBody`

NewQueryAssetCardListResponseResponseBody instantiates a new QueryAssetCardListResponseResponseBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryAssetCardListResponseResponseBodyWithDefaults

`func NewQueryAssetCardListResponseResponseBodyWithDefaults() *QueryAssetCardListResponseResponseBody`

NewQueryAssetCardListResponseResponseBodyWithDefaults instantiates a new QueryAssetCardListResponseResponseBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResultInfo

`func (o *QueryAssetCardListResponseResponseBody) GetResultInfo() MemberAssetResultInfo`

GetResultInfo returns the ResultInfo field if non-nil, zero value otherwise.

### GetResultInfoOk

`func (o *QueryAssetCardListResponseResponseBody) GetResultInfoOk() (*MemberAssetResultInfo, bool)`

GetResultInfoOk returns a tuple with the ResultInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultInfo

`func (o *QueryAssetCardListResponseResponseBody) SetResultInfo(v MemberAssetResultInfo)`

SetResultInfo sets ResultInfo field to given value.


### GetAssetCardList

`func (o *QueryAssetCardListResponseResponseBody) GetAssetCardList() []AssetCardListItem`

GetAssetCardList returns the AssetCardList field if non-nil, zero value otherwise.

### GetAssetCardListOk

`func (o *QueryAssetCardListResponseResponseBody) GetAssetCardListOk() (*[]AssetCardListItem, bool)`

GetAssetCardListOk returns a tuple with the AssetCardList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetCardList

`func (o *QueryAssetCardListResponseResponseBody) SetAssetCardList(v []AssetCardListItem)`

SetAssetCardList sets AssetCardList field to given value.

### HasAssetCardList

`func (o *QueryAssetCardListResponseResponseBody) HasAssetCardList() bool`

HasAssetCardList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


