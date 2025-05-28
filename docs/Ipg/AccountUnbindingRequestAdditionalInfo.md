# AccountUnbindingRequestAdditionalInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessToken** | **string** | Contains customer token, which has been obtained from binding process | 
**EndUserIpAddress** | Pointer to **string** | IP address of the end user (customer) using IPv4 format | [optional] 
**DeviceId** | **string** | Device identification on which the API services is currently being accessed by the end user (customer) | 
**Latitude** | Pointer to **string** | Location on which the API services is currently being accessed by the end user (customer), refer to ISO 6709 standard representation of geographic point location by coordinates | [optional] 
**Longitude** | Pointer to **string** | Location on which the API services is currently being accessed by the end user (customer), refer to ISO 6709 Standard representation of geographic point location by coordinates | [optional] 

## Methods

### NewAccountUnbindingRequestAdditionalInfo

`func NewAccountUnbindingRequestAdditionalInfo(accessToken string, deviceId string, ) *AccountUnbindingRequestAdditionalInfo`

NewAccountUnbindingRequestAdditionalInfo instantiates a new AccountUnbindingRequestAdditionalInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountUnbindingRequestAdditionalInfoWithDefaults

`func NewAccountUnbindingRequestAdditionalInfoWithDefaults() *AccountUnbindingRequestAdditionalInfo`

NewAccountUnbindingRequestAdditionalInfoWithDefaults instantiates a new AccountUnbindingRequestAdditionalInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessToken

`func (o *AccountUnbindingRequestAdditionalInfo) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *AccountUnbindingRequestAdditionalInfo) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *AccountUnbindingRequestAdditionalInfo) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.


### GetEndUserIpAddress

`func (o *AccountUnbindingRequestAdditionalInfo) GetEndUserIpAddress() string`

GetEndUserIpAddress returns the EndUserIpAddress field if non-nil, zero value otherwise.

### GetEndUserIpAddressOk

`func (o *AccountUnbindingRequestAdditionalInfo) GetEndUserIpAddressOk() (*string, bool)`

GetEndUserIpAddressOk returns a tuple with the EndUserIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndUserIpAddress

`func (o *AccountUnbindingRequestAdditionalInfo) SetEndUserIpAddress(v string)`

SetEndUserIpAddress sets EndUserIpAddress field to given value.

### HasEndUserIpAddress

`func (o *AccountUnbindingRequestAdditionalInfo) HasEndUserIpAddress() bool`

HasEndUserIpAddress returns a boolean if a field has been set.

### GetDeviceId

`func (o *AccountUnbindingRequestAdditionalInfo) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *AccountUnbindingRequestAdditionalInfo) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *AccountUnbindingRequestAdditionalInfo) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.


### GetLatitude

`func (o *AccountUnbindingRequestAdditionalInfo) GetLatitude() string`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *AccountUnbindingRequestAdditionalInfo) GetLatitudeOk() (*string, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *AccountUnbindingRequestAdditionalInfo) SetLatitude(v string)`

SetLatitude sets Latitude field to given value.

### HasLatitude

`func (o *AccountUnbindingRequestAdditionalInfo) HasLatitude() bool`

HasLatitude returns a boolean if a field has been set.

### GetLongitude

`func (o *AccountUnbindingRequestAdditionalInfo) GetLongitude() string`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *AccountUnbindingRequestAdditionalInfo) GetLongitudeOk() (*string, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *AccountUnbindingRequestAdditionalInfo) SetLongitude(v string)`

SetLongitude sets Longitude field to given value.

### HasLongitude

`func (o *AccountUnbindingRequestAdditionalInfo) HasLongitude() bool`

HasLongitude returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


