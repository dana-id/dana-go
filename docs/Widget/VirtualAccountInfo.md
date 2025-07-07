# VirtualAccountInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VirtualAccountCode** | **string** | Virtual account code (required if payMethod is VIRTUAL_ACCOUNT) | 
**VirtualAccountExpiryTime** | **string** | Expiry time of virtual account in format YYYY-MM-DDTHH:mm:ss+07:00 (Jakarta time) | 
**Signature** | **string** | Signature of virtual account | 

## Methods

### NewVirtualAccountInfo

`func NewVirtualAccountInfo(virtualAccountCode string, virtualAccountExpiryTime string, signature string, ) *VirtualAccountInfo`

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


