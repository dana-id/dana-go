# MobileNoInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MobileNo** | Pointer to **string** | Mobile number. Accepted formats: 0xxxxxxxx, 62xxxxxxx, 62-xxxxxxx, +62xxxxxx | [optional] 
**MobileId** | Pointer to **string** | Mobile ID | [optional] 
**Verified** | Pointer to **string** | Verification status | [optional] 

## Methods

### NewMobileNoInfo

`func NewMobileNoInfo() *MobileNoInfo`

NewMobileNoInfo instantiates a new MobileNoInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMobileNoInfoWithDefaults

`func NewMobileNoInfoWithDefaults() *MobileNoInfo`

NewMobileNoInfoWithDefaults instantiates a new MobileNoInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMobileNo

`func (o *MobileNoInfo) GetMobileNo() string`

GetMobileNo returns the MobileNo field if non-nil, zero value otherwise.

### GetMobileNoOk

`func (o *MobileNoInfo) GetMobileNoOk() (*string, bool)`

GetMobileNoOk returns a tuple with the MobileNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobileNo

`func (o *MobileNoInfo) SetMobileNo(v string)`

SetMobileNo sets MobileNo field to given value.

### HasMobileNo

`func (o *MobileNoInfo) HasMobileNo() bool`

HasMobileNo returns a boolean if a field has been set.

### GetMobileId

`func (o *MobileNoInfo) GetMobileId() string`

GetMobileId returns the MobileId field if non-nil, zero value otherwise.

### GetMobileIdOk

`func (o *MobileNoInfo) GetMobileIdOk() (*string, bool)`

GetMobileIdOk returns a tuple with the MobileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobileId

`func (o *MobileNoInfo) SetMobileId(v string)`

SetMobileId sets MobileId field to given value.

### HasMobileId

`func (o *MobileNoInfo) HasMobileId() bool`

HasMobileId returns a boolean if a field has been set.

### GetVerified

`func (o *MobileNoInfo) GetVerified() string`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *MobileNoInfo) GetVerifiedOk() (*string, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *MobileNoInfo) SetVerified(v string)`

SetVerified sets Verified field to given value.

### HasVerified

`func (o *MobileNoInfo) HasVerified() bool`

HasVerified returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


