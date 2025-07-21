# CreateDivisionResponseResponseBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResultInfo** | [**ResultInfo**](ResultInfo.md) |  | 
**DivisionId** | Pointer to **string** | Division identifier. Present when successfully processed | [optional] 

## Methods

### NewCreateDivisionResponseResponseBody

`func NewCreateDivisionResponseResponseBody(resultInfo ResultInfo, ) *CreateDivisionResponseResponseBody`

NewCreateDivisionResponseResponseBody instantiates a new CreateDivisionResponseResponseBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDivisionResponseResponseBodyWithDefaults

`func NewCreateDivisionResponseResponseBodyWithDefaults() *CreateDivisionResponseResponseBody`

NewCreateDivisionResponseResponseBodyWithDefaults instantiates a new CreateDivisionResponseResponseBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResultInfo

`func (o *CreateDivisionResponseResponseBody) GetResultInfo() ResultInfo`

GetResultInfo returns the ResultInfo field if non-nil, zero value otherwise.

### GetResultInfoOk

`func (o *CreateDivisionResponseResponseBody) GetResultInfoOk() (*ResultInfo, bool)`

GetResultInfoOk returns a tuple with the ResultInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultInfo

`func (o *CreateDivisionResponseResponseBody) SetResultInfo(v ResultInfo)`

SetResultInfo sets ResultInfo field to given value.


### GetDivisionId

`func (o *CreateDivisionResponseResponseBody) GetDivisionId() string`

GetDivisionId returns the DivisionId field if non-nil, zero value otherwise.

### GetDivisionIdOk

`func (o *CreateDivisionResponseResponseBody) GetDivisionIdOk() (*string, bool)`

GetDivisionIdOk returns a tuple with the DivisionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDivisionId

`func (o *CreateDivisionResponseResponseBody) SetDivisionId(v string)`

SetDivisionId sets DivisionId field to given value.

### HasDivisionId

`func (o *CreateDivisionResponseResponseBody) HasDivisionId() bool`

HasDivisionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


