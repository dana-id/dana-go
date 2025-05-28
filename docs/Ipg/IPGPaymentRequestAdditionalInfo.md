# IPGPaymentRequestAdditionalInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SupportDeepLinkCheckoutUrl** | Pointer to **string** | Additional information of deeplink checkout URL. For Mini Program, DANA will treat as false | [optional] 
**PhoneNumber** | Pointer to **string** | Additional information of user&#39;s phone number | [optional] 
**PublicUserId** | Pointer to **string** | Additional information of public user&#39;s identifier | [optional] 
**ProductCode** | **string** | Additional information of product code | 
**ServiceInfo** | Pointer to [**ServiceInfo**](ServiceInfo.md) |  | [optional] 
**Order** | [**Order**](Order.md) |  | 
**Mcc** | **string** | Additional information of merchant category code. This parameter is used to identify the type of business in which a merchant is engaged. | 
**EnvInfo** | [**EnvInfo**](EnvInfo.md) |  | 
**ExtendInfo** | Pointer to **string** | Additional information of extend | [optional] 

## Methods

### NewIPGPaymentRequestAdditionalInfo

`func NewIPGPaymentRequestAdditionalInfo(productCode string, order Order, mcc string, envInfo EnvInfo, ) *IPGPaymentRequestAdditionalInfo`

NewIPGPaymentRequestAdditionalInfo instantiates a new IPGPaymentRequestAdditionalInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIPGPaymentRequestAdditionalInfoWithDefaults

`func NewIPGPaymentRequestAdditionalInfoWithDefaults() *IPGPaymentRequestAdditionalInfo`

NewIPGPaymentRequestAdditionalInfoWithDefaults instantiates a new IPGPaymentRequestAdditionalInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSupportDeepLinkCheckoutUrl

`func (o *IPGPaymentRequestAdditionalInfo) GetSupportDeepLinkCheckoutUrl() string`

GetSupportDeepLinkCheckoutUrl returns the SupportDeepLinkCheckoutUrl field if non-nil, zero value otherwise.

### GetSupportDeepLinkCheckoutUrlOk

`func (o *IPGPaymentRequestAdditionalInfo) GetSupportDeepLinkCheckoutUrlOk() (*string, bool)`

GetSupportDeepLinkCheckoutUrlOk returns a tuple with the SupportDeepLinkCheckoutUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportDeepLinkCheckoutUrl

`func (o *IPGPaymentRequestAdditionalInfo) SetSupportDeepLinkCheckoutUrl(v string)`

SetSupportDeepLinkCheckoutUrl sets SupportDeepLinkCheckoutUrl field to given value.

### HasSupportDeepLinkCheckoutUrl

`func (o *IPGPaymentRequestAdditionalInfo) HasSupportDeepLinkCheckoutUrl() bool`

HasSupportDeepLinkCheckoutUrl returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *IPGPaymentRequestAdditionalInfo) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *IPGPaymentRequestAdditionalInfo) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *IPGPaymentRequestAdditionalInfo) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *IPGPaymentRequestAdditionalInfo) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetPublicUserId

`func (o *IPGPaymentRequestAdditionalInfo) GetPublicUserId() string`

GetPublicUserId returns the PublicUserId field if non-nil, zero value otherwise.

### GetPublicUserIdOk

`func (o *IPGPaymentRequestAdditionalInfo) GetPublicUserIdOk() (*string, bool)`

GetPublicUserIdOk returns a tuple with the PublicUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicUserId

`func (o *IPGPaymentRequestAdditionalInfo) SetPublicUserId(v string)`

SetPublicUserId sets PublicUserId field to given value.

### HasPublicUserId

`func (o *IPGPaymentRequestAdditionalInfo) HasPublicUserId() bool`

HasPublicUserId returns a boolean if a field has been set.

### GetProductCode

`func (o *IPGPaymentRequestAdditionalInfo) GetProductCode() string`

GetProductCode returns the ProductCode field if non-nil, zero value otherwise.

### GetProductCodeOk

`func (o *IPGPaymentRequestAdditionalInfo) GetProductCodeOk() (*string, bool)`

GetProductCodeOk returns a tuple with the ProductCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductCode

`func (o *IPGPaymentRequestAdditionalInfo) SetProductCode(v string)`

SetProductCode sets ProductCode field to given value.


### GetServiceInfo

`func (o *IPGPaymentRequestAdditionalInfo) GetServiceInfo() ServiceInfo`

GetServiceInfo returns the ServiceInfo field if non-nil, zero value otherwise.

### GetServiceInfoOk

`func (o *IPGPaymentRequestAdditionalInfo) GetServiceInfoOk() (*ServiceInfo, bool)`

GetServiceInfoOk returns a tuple with the ServiceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceInfo

`func (o *IPGPaymentRequestAdditionalInfo) SetServiceInfo(v ServiceInfo)`

SetServiceInfo sets ServiceInfo field to given value.

### HasServiceInfo

`func (o *IPGPaymentRequestAdditionalInfo) HasServiceInfo() bool`

HasServiceInfo returns a boolean if a field has been set.

### GetOrder

`func (o *IPGPaymentRequestAdditionalInfo) GetOrder() Order`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *IPGPaymentRequestAdditionalInfo) GetOrderOk() (*Order, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *IPGPaymentRequestAdditionalInfo) SetOrder(v Order)`

SetOrder sets Order field to given value.


### GetMcc

`func (o *IPGPaymentRequestAdditionalInfo) GetMcc() string`

GetMcc returns the Mcc field if non-nil, zero value otherwise.

### GetMccOk

`func (o *IPGPaymentRequestAdditionalInfo) GetMccOk() (*string, bool)`

GetMccOk returns a tuple with the Mcc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcc

`func (o *IPGPaymentRequestAdditionalInfo) SetMcc(v string)`

SetMcc sets Mcc field to given value.


### GetEnvInfo

`func (o *IPGPaymentRequestAdditionalInfo) GetEnvInfo() EnvInfo`

GetEnvInfo returns the EnvInfo field if non-nil, zero value otherwise.

### GetEnvInfoOk

`func (o *IPGPaymentRequestAdditionalInfo) GetEnvInfoOk() (*EnvInfo, bool)`

GetEnvInfoOk returns a tuple with the EnvInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvInfo

`func (o *IPGPaymentRequestAdditionalInfo) SetEnvInfo(v EnvInfo)`

SetEnvInfo sets EnvInfo field to given value.


### GetExtendInfo

`func (o *IPGPaymentRequestAdditionalInfo) GetExtendInfo() string`

GetExtendInfo returns the ExtendInfo field if non-nil, zero value otherwise.

### GetExtendInfoOk

`func (o *IPGPaymentRequestAdditionalInfo) GetExtendInfoOk() (*string, bool)`

GetExtendInfoOk returns a tuple with the ExtendInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendInfo

`func (o *IPGPaymentRequestAdditionalInfo) SetExtendInfo(v string)`

SetExtendInfo sets ExtendInfo field to given value.

### HasExtendInfo

`func (o *IPGPaymentRequestAdditionalInfo) HasExtendInfo() bool`

HasExtendInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


