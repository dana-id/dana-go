# QueryDivisionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MerchantId** | Pointer to **string** | Merchant identifier. Required when divisionIdType is EXTERNAL_ID | [optional] 
**DivisionId** | **string** | Division identifier or external identifier. Length depends on divisionIdType - INNER_ID (21 max), EXTERNAL_ID (64 max) | 
**DivisionIdType** | **string** | Division identifier type | 

## Methods

### NewQueryDivisionRequest

`func NewQueryDivisionRequest(divisionId string, divisionIdType string, ) *QueryDivisionRequest`

NewQueryDivisionRequest instantiates a new QueryDivisionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryDivisionRequestWithDefaults

`func NewQueryDivisionRequestWithDefaults() *QueryDivisionRequest`

NewQueryDivisionRequestWithDefaults instantiates a new QueryDivisionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMerchantId

`func (o *QueryDivisionRequest) GetMerchantId() string`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *QueryDivisionRequest) GetMerchantIdOk() (*string, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *QueryDivisionRequest) SetMerchantId(v string)`

SetMerchantId sets MerchantId field to given value.

### HasMerchantId

`func (o *QueryDivisionRequest) HasMerchantId() bool`

HasMerchantId returns a boolean if a field has been set.

### GetDivisionId

`func (o *QueryDivisionRequest) GetDivisionId() string`

GetDivisionId returns the DivisionId field if non-nil, zero value otherwise.

### GetDivisionIdOk

`func (o *QueryDivisionRequest) GetDivisionIdOk() (*string, bool)`

GetDivisionIdOk returns a tuple with the DivisionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDivisionId

`func (o *QueryDivisionRequest) SetDivisionId(v string)`

SetDivisionId sets DivisionId field to given value.


### GetDivisionIdType

`func (o *QueryDivisionRequest) GetDivisionIdType() string`

GetDivisionIdType returns the DivisionIdType field if non-nil, zero value otherwise.

### GetDivisionIdTypeOk

`func (o *QueryDivisionRequest) GetDivisionIdTypeOk() (*string, bool)`

GetDivisionIdTypeOk returns a tuple with the DivisionIdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDivisionIdType

`func (o *QueryDivisionRequest) SetDivisionIdType(v string)`

SetDivisionIdType sets DivisionIdType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


