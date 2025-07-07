# WidgetPaymentRequestAdditionalInfo

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

### NewWidgetPaymentRequestAdditionalInfo

`func NewWidgetPaymentRequestAdditionalInfo(productCode string, order Order, mcc string, envInfo EnvInfo, ) *WidgetPaymentRequestAdditionalInfo`

NewWidgetPaymentRequestAdditionalInfo instantiates a new WidgetPaymentRequestAdditionalInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWidgetPaymentRequestAdditionalInfoWithDefaults

`func NewWidgetPaymentRequestAdditionalInfoWithDefaults() *WidgetPaymentRequestAdditionalInfo`

NewWidgetPaymentRequestAdditionalInfoWithDefaults instantiates a new WidgetPaymentRequestAdditionalInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSupportDeepLinkCheckoutUrl

`func (o *WidgetPaymentRequestAdditionalInfo) GetSupportDeepLinkCheckoutUrl() string`

GetSupportDeepLinkCheckoutUrl returns the SupportDeepLinkCheckoutUrl field if non-nil, zero value otherwise.

### GetSupportDeepLinkCheckoutUrlOk

`func (o *WidgetPaymentRequestAdditionalInfo) GetSupportDeepLinkCheckoutUrlOk() (*string, bool)`

GetSupportDeepLinkCheckoutUrlOk returns a tuple with the SupportDeepLinkCheckoutUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportDeepLinkCheckoutUrl

`func (o *WidgetPaymentRequestAdditionalInfo) SetSupportDeepLinkCheckoutUrl(v string)`

SetSupportDeepLinkCheckoutUrl sets SupportDeepLinkCheckoutUrl field to given value.

### HasSupportDeepLinkCheckoutUrl

`func (o *WidgetPaymentRequestAdditionalInfo) HasSupportDeepLinkCheckoutUrl() bool`

HasSupportDeepLinkCheckoutUrl returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *WidgetPaymentRequestAdditionalInfo) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *WidgetPaymentRequestAdditionalInfo) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *WidgetPaymentRequestAdditionalInfo) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *WidgetPaymentRequestAdditionalInfo) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetPublicUserId

`func (o *WidgetPaymentRequestAdditionalInfo) GetPublicUserId() string`

GetPublicUserId returns the PublicUserId field if non-nil, zero value otherwise.

### GetPublicUserIdOk

`func (o *WidgetPaymentRequestAdditionalInfo) GetPublicUserIdOk() (*string, bool)`

GetPublicUserIdOk returns a tuple with the PublicUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicUserId

`func (o *WidgetPaymentRequestAdditionalInfo) SetPublicUserId(v string)`

SetPublicUserId sets PublicUserId field to given value.

### HasPublicUserId

`func (o *WidgetPaymentRequestAdditionalInfo) HasPublicUserId() bool`

HasPublicUserId returns a boolean if a field has been set.

### GetProductCode

`func (o *WidgetPaymentRequestAdditionalInfo) GetProductCode() string`

GetProductCode returns the ProductCode field if non-nil, zero value otherwise.

### GetProductCodeOk

`func (o *WidgetPaymentRequestAdditionalInfo) GetProductCodeOk() (*string, bool)`

GetProductCodeOk returns a tuple with the ProductCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductCode

`func (o *WidgetPaymentRequestAdditionalInfo) SetProductCode(v string)`

SetProductCode sets ProductCode field to given value.


### GetServiceInfo

`func (o *WidgetPaymentRequestAdditionalInfo) GetServiceInfo() ServiceInfo`

GetServiceInfo returns the ServiceInfo field if non-nil, zero value otherwise.

### GetServiceInfoOk

`func (o *WidgetPaymentRequestAdditionalInfo) GetServiceInfoOk() (*ServiceInfo, bool)`

GetServiceInfoOk returns a tuple with the ServiceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceInfo

`func (o *WidgetPaymentRequestAdditionalInfo) SetServiceInfo(v ServiceInfo)`

SetServiceInfo sets ServiceInfo field to given value.

### HasServiceInfo

`func (o *WidgetPaymentRequestAdditionalInfo) HasServiceInfo() bool`

HasServiceInfo returns a boolean if a field has been set.

### GetOrder

`func (o *WidgetPaymentRequestAdditionalInfo) GetOrder() Order`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *WidgetPaymentRequestAdditionalInfo) GetOrderOk() (*Order, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *WidgetPaymentRequestAdditionalInfo) SetOrder(v Order)`

SetOrder sets Order field to given value.


### GetMcc

`func (o *WidgetPaymentRequestAdditionalInfo) GetMcc() string`

GetMcc returns the Mcc field if non-nil, zero value otherwise.

### GetMccOk

`func (o *WidgetPaymentRequestAdditionalInfo) GetMccOk() (*string, bool)`

GetMccOk returns a tuple with the Mcc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcc

`func (o *WidgetPaymentRequestAdditionalInfo) SetMcc(v string)`

SetMcc sets Mcc field to given value.


### GetEnvInfo

`func (o *WidgetPaymentRequestAdditionalInfo) GetEnvInfo() EnvInfo`

GetEnvInfo returns the EnvInfo field if non-nil, zero value otherwise.

### GetEnvInfoOk

`func (o *WidgetPaymentRequestAdditionalInfo) GetEnvInfoOk() (*EnvInfo, bool)`

GetEnvInfoOk returns a tuple with the EnvInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvInfo

`func (o *WidgetPaymentRequestAdditionalInfo) SetEnvInfo(v EnvInfo)`

SetEnvInfo sets EnvInfo field to given value.


### GetExtendInfo

`func (o *WidgetPaymentRequestAdditionalInfo) GetExtendInfo() string`

GetExtendInfo returns the ExtendInfo field if non-nil, zero value otherwise.

### GetExtendInfoOk

`func (o *WidgetPaymentRequestAdditionalInfo) GetExtendInfoOk() (*string, bool)`

GetExtendInfoOk returns a tuple with the ExtendInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendInfo

`func (o *WidgetPaymentRequestAdditionalInfo) SetExtendInfo(v string)`

SetExtendInfo sets ExtendInfo field to given value.

### HasExtendInfo

`func (o *WidgetPaymentRequestAdditionalInfo) HasExtendInfo() bool`

HasExtendInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


