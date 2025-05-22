# VirtualAccountInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VirtualAccountCode** | Pointer to **string** | Virtual account code. Present if successfully processed and payment method is VIRTUAL_ACCOUNT | [optional] 
**VirtualAccountExpiryTime** | Pointer to **string** | Expiry time of virtual account, in format YYYY-MM-DDTHH:mm:ss+07:00. Time must be in GMT+7 (Jakarta time). Present if successfully processed and payment method is VIRTUAL_ACCOUNT | [optional] 
**Signature** | Pointer to **string** | Signature of virtual account. Present if successfully processed and payment method is VIRTUAL_ACCOUNT | [optional] 

## Methods

### NewVirtualAccountInfo

`func NewVirtualAccountInfo() *VirtualAccountInfo`

NewVirtualAccountInfo instantiates a new VirtualAccountInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualAccountInfoWithDefaults

`func NewVirtualAccountInfoWithDefaults() *VirtualAccountInfo`

NewVirtualAccountInfoWithDefaults instantiates a new VirtualAccountInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVirtualAccountCode

`func (o *VirtualAccountInfo) GetVirtualAccountCode() string`

GetVirtualAccountCode returns the VirtualAccountCode field if non-nil, zero value otherwise.

### GetVirtualAccountCodeOk

`func (o *VirtualAccountInfo) GetVirtualAccountCodeOk() (*string, bool)`

GetVirtualAccountCodeOk returns a tuple with the VirtualAccountCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualAccountCode

`func (o *VirtualAccountInfo) SetVirtualAccountCode(v string)`

SetVirtualAccountCode sets VirtualAccountCode field to given value.

### HasVirtualAccountCode

`func (o *VirtualAccountInfo) HasVirtualAccountCode() bool`

HasVirtualAccountCode returns a boolean if a field has been set.

### GetVirtualAccountExpiryTime

`func (o *VirtualAccountInfo) GetVirtualAccountExpiryTime() string`

GetVirtualAccountExpiryTime returns the VirtualAccountExpiryTime field if non-nil, zero value otherwise.

### GetVirtualAccountExpiryTimeOk

`func (o *VirtualAccountInfo) GetVirtualAccountExpiryTimeOk() (*string, bool)`

GetVirtualAccountExpiryTimeOk returns a tuple with the VirtualAccountExpiryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualAccountExpiryTime

`func (o *VirtualAccountInfo) SetVirtualAccountExpiryTime(v string)`

SetVirtualAccountExpiryTime sets VirtualAccountExpiryTime field to given value.

### HasVirtualAccountExpiryTime

`func (o *VirtualAccountInfo) HasVirtualAccountExpiryTime() bool`

HasVirtualAccountExpiryTime returns a boolean if a field has been set.

### GetSignature

`func (o *VirtualAccountInfo) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *VirtualAccountInfo) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *VirtualAccountInfo) SetSignature(v string)`

SetSignature sets Signature field to given value.

### HasSignature

`func (o *VirtualAccountInfo) HasSignature() bool`

HasSignature returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


