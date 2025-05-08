# StatusDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcquirementStatus** | **string** | The status of acquirement | 
**Frozen** | Pointer to **string** | Whether the frozen is true or not | [optional] 
**Cancelled** | Pointer to **string** | Whether the cancelled is true or not | [optional] 

## Methods

### NewStatusDetail

`func NewStatusDetail(acquirementStatus string, ) *StatusDetail`

NewStatusDetail instantiates a new StatusDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatusDetailWithDefaults

`func NewStatusDetailWithDefaults() *StatusDetail`

NewStatusDetailWithDefaults instantiates a new StatusDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcquirementStatus

`func (o *StatusDetail) GetAcquirementStatus() string`

GetAcquirementStatus returns the AcquirementStatus field if non-nil, zero value otherwise.

### GetAcquirementStatusOk

`func (o *StatusDetail) GetAcquirementStatusOk() (*string, bool)`

GetAcquirementStatusOk returns a tuple with the AcquirementStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcquirementStatus

`func (o *StatusDetail) SetAcquirementStatus(v string)`

SetAcquirementStatus sets AcquirementStatus field to given value.


### GetFrozen

`func (o *StatusDetail) GetFrozen() string`

GetFrozen returns the Frozen field if non-nil, zero value otherwise.

### GetFrozenOk

`func (o *StatusDetail) GetFrozenOk() (*string, bool)`

GetFrozenOk returns a tuple with the Frozen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrozen

`func (o *StatusDetail) SetFrozen(v string)`

SetFrozen sets Frozen field to given value.

### HasFrozen

`func (o *StatusDetail) HasFrozen() bool`

HasFrozen returns a boolean if a field has been set.

### GetCancelled

`func (o *StatusDetail) GetCancelled() string`

GetCancelled returns the Cancelled field if non-nil, zero value otherwise.

### GetCancelledOk

`func (o *StatusDetail) GetCancelledOk() (*string, bool)`

GetCancelledOk returns a tuple with the Cancelled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelled

`func (o *StatusDetail) SetCancelled(v string)`

SetCancelled sets Cancelled field to given value.

### HasCancelled

`func (o *StatusDetail) HasCancelled() bool`

HasCancelled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


