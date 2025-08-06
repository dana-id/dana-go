# BalanceInquiryRequestAdditionalInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessToken** | **string** | Contains customer token, which has been obtained from binding process | 
**EndUserIpAddress** | Pointer to **string** | IP address of the end user (customer) using IPv4 format | [optional] 
**DeviceId** | **string** | Device identification on which the API services is currently being accessed by the end user (customer) | 
**Latitude** | Pointer to **string** | Location on which the API services is currently being accessed by the end user (customer), refer to ISO 6709 standard representation of geographic point location by coordinates | [optional] 
**Longitude** | Pointer to **string** | Location on which the API services is currently being accessed by the end user (customer), refer to ISO 6709 Standard representation of geographic point location by coordinates | [optional] 

## Methods

### NewBalanceInquiryRequestAdditionalInfo

`func NewBalanceInquiryRequestAdditionalInfo(accessToken string, deviceId string, ) *BalanceInquiryRequestAdditionalInfo`

NewBalanceInquiryRequestAdditionalInfo instantiates a new BalanceInquiryRequestAdditionalInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBalanceInquiryRequestAdditionalInfoWithDefaults

`func NewBalanceInquiryRequestAdditionalInfoWithDefaults() *BalanceInquiryRequestAdditionalInfo`

NewBalanceInquiryRequestAdditionalInfoWithDefaults instantiates a new BalanceInquiryRequestAdditionalInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessToken

`func (o *BalanceInquiryRequestAdditionalInfo) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *BalanceInquiryRequestAdditionalInfo) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *BalanceInquiryRequestAdditionalInfo) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.


### GetEndUserIpAddress

`func (o *BalanceInquiryRequestAdditionalInfo) GetEndUserIpAddress() string`

GetEndUserIpAddress returns the EndUserIpAddress field if non-nil, zero value otherwise.

### GetEndUserIpAddressOk

`func (o *BalanceInquiryRequestAdditionalInfo) GetEndUserIpAddressOk() (*string, bool)`

GetEndUserIpAddressOk returns a tuple with the EndUserIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndUserIpAddress

`func (o *BalanceInquiryRequestAdditionalInfo) SetEndUserIpAddress(v string)`

SetEndUserIpAddress sets EndUserIpAddress field to given value.

### HasEndUserIpAddress

`func (o *BalanceInquiryRequestAdditionalInfo) HasEndUserIpAddress() bool`

HasEndUserIpAddress returns a boolean if a field has been set.

### GetDeviceId

`func (o *BalanceInquiryRequestAdditionalInfo) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *BalanceInquiryRequestAdditionalInfo) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *BalanceInquiryRequestAdditionalInfo) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.


### GetLatitude

`func (o *BalanceInquiryRequestAdditionalInfo) GetLatitude() string`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *BalanceInquiryRequestAdditionalInfo) GetLatitudeOk() (*string, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *BalanceInquiryRequestAdditionalInfo) SetLatitude(v string)`

SetLatitude sets Latitude field to given value.

### HasLatitude

`func (o *BalanceInquiryRequestAdditionalInfo) HasLatitude() bool`

HasLatitude returns a boolean if a field has been set.

### GetLongitude

`func (o *BalanceInquiryRequestAdditionalInfo) GetLongitude() string`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *BalanceInquiryRequestAdditionalInfo) GetLongitudeOk() (*string, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *BalanceInquiryRequestAdditionalInfo) SetLongitude(v string)`

SetLongitude sets Longitude field to given value.

### HasLongitude

`func (o *BalanceInquiryRequestAdditionalInfo) HasLongitude() bool`

HasLongitude returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


