# TimeDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedTime** | **string** | Time of created order, format YYYY-MM-DDTHH:mm:ss+07:00 (Jakarta time) | 
**ExpiryTime** | **string** | Time of expiry order, format YYYY-MM-DDTHH:mm:ss+07:00 (Jakarta time) | 
**PaidTimes** | Pointer to **[]string** | Array of paid order times in format YYYY-MM-DDTHH:mm:ss+07:00 (Jakarta time) | [optional] 
**ConfirmedTimes** | Pointer to **[]string** | Array of confirmed order times in format YYYY-MM-DDTHH:mm:ss+07:00 (Jakarta time) | [optional] 
**CancelledTime** | Pointer to **string** | Time of cancelled order in format YYYY-MM-DDTHH:mm:ss+07:00 (Jakarta time) | [optional] 

## Methods

### NewTimeDetail

`func NewTimeDetail(createdTime string, expiryTime string, ) *TimeDetail`

NewTimeDetail instantiates a new TimeDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeDetailWithDefaults

`func NewTimeDetailWithDefaults() *TimeDetail`

NewTimeDetailWithDefaults instantiates a new TimeDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedTime

`func (o *TimeDetail) GetCreatedTime() string`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *TimeDetail) GetCreatedTimeOk() (*string, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *TimeDetail) SetCreatedTime(v string)`

SetCreatedTime sets CreatedTime field to given value.


### GetExpiryTime

`func (o *TimeDetail) GetExpiryTime() string`

GetExpiryTime returns the ExpiryTime field if non-nil, zero value otherwise.

### GetExpiryTimeOk

`func (o *TimeDetail) GetExpiryTimeOk() (*string, bool)`

GetExpiryTimeOk returns a tuple with the ExpiryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryTime

`func (o *TimeDetail) SetExpiryTime(v string)`

SetExpiryTime sets ExpiryTime field to given value.


### GetPaidTimes

`func (o *TimeDetail) GetPaidTimes() []string`

GetPaidTimes returns the PaidTimes field if non-nil, zero value otherwise.

### GetPaidTimesOk

`func (o *TimeDetail) GetPaidTimesOk() (*[]string, bool)`

GetPaidTimesOk returns a tuple with the PaidTimes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaidTimes

`func (o *TimeDetail) SetPaidTimes(v []string)`

SetPaidTimes sets PaidTimes field to given value.

### HasPaidTimes

`func (o *TimeDetail) HasPaidTimes() bool`

HasPaidTimes returns a boolean if a field has been set.

### GetConfirmedTimes

`func (o *TimeDetail) GetConfirmedTimes() []string`

GetConfirmedTimes returns the ConfirmedTimes field if non-nil, zero value otherwise.

### GetConfirmedTimesOk

`func (o *TimeDetail) GetConfirmedTimesOk() (*[]string, bool)`

GetConfirmedTimesOk returns a tuple with the ConfirmedTimes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmedTimes

`func (o *TimeDetail) SetConfirmedTimes(v []string)`

SetConfirmedTimes sets ConfirmedTimes field to given value.

### HasConfirmedTimes

`func (o *TimeDetail) HasConfirmedTimes() bool`

HasConfirmedTimes returns a boolean if a field has been set.

### GetCancelledTime

`func (o *TimeDetail) GetCancelledTime() string`

GetCancelledTime returns the CancelledTime field if non-nil, zero value otherwise.

### GetCancelledTimeOk

`func (o *TimeDetail) GetCancelledTimeOk() (*string, bool)`

GetCancelledTimeOk returns a tuple with the CancelledTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelledTime

`func (o *TimeDetail) SetCancelledTime(v string)`

SetCancelledTime sets CancelledTime field to given value.

### HasCancelledTime

`func (o *TimeDetail) HasCancelledTime() bool`

HasCancelledTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


