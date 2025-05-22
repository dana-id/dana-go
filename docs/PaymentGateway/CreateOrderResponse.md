# CreateOrderResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponseCode** | **string** | Response code. Refer to https://dashboard.dana.id/api-docs/read/243#paymentgatewayprod-paymentRedirect-ResponseCodeandMessage | 
**ResponseMessage** | **string** | Response message. Refer to https://dashboard.dana.id/api-docs/read/243#paymentgatewayprod-paymentRedirect-ResponseCodeandMessage | 
**ReferenceNo** | Pointer to **string** | Transaction identifier on DANA system. Present if successfully processed | [optional] 
**PartnerReferenceNo** | **string** | Transaction identifier on partner system | 
**WebRedirectUrl** | Pointer to **string** | Checkout URLs. Present if successfully processed and payment method is not OVO/Virtual Account/QRIS | [optional] 
**AdditionalInfo** | Pointer to [**CreateOrderResponseAdditionalInfo**](CreateOrderResponseAdditionalInfo.md) | Additional information | [optional] 
**ExternalOrderId** | Pointer to **string** | External order identifier | [optional] 

## Methods

### NewCreateOrderResponse

`func NewCreateOrderResponse(responseCode string, responseMessage string, partnerReferenceNo string, ) *CreateOrderResponse`

NewCreateOrderResponse instantiates a new CreateOrderResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrderResponseWithDefaults

`func NewCreateOrderResponseWithDefaults() *CreateOrderResponse`

NewCreateOrderResponseWithDefaults instantiates a new CreateOrderResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponseCode

`func (o *CreateOrderResponse) GetResponseCode() string`

GetResponseCode returns the ResponseCode field if non-nil, zero value otherwise.

### GetResponseCodeOk

`func (o *CreateOrderResponse) GetResponseCodeOk() (*string, bool)`

GetResponseCodeOk returns a tuple with the ResponseCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseCode

`func (o *CreateOrderResponse) SetResponseCode(v string)`

SetResponseCode sets ResponseCode field to given value.


### GetResponseMessage

`func (o *CreateOrderResponse) GetResponseMessage() string`

GetResponseMessage returns the ResponseMessage field if non-nil, zero value otherwise.

### GetResponseMessageOk

`func (o *CreateOrderResponse) GetResponseMessageOk() (*string, bool)`

GetResponseMessageOk returns a tuple with the ResponseMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseMessage

`func (o *CreateOrderResponse) SetResponseMessage(v string)`

SetResponseMessage sets ResponseMessage field to given value.


### GetReferenceNo

`func (o *CreateOrderResponse) GetReferenceNo() string`

GetReferenceNo returns the ReferenceNo field if non-nil, zero value otherwise.

### GetReferenceNoOk

`func (o *CreateOrderResponse) GetReferenceNoOk() (*string, bool)`

GetReferenceNoOk returns a tuple with the ReferenceNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceNo

`func (o *CreateOrderResponse) SetReferenceNo(v string)`

SetReferenceNo sets ReferenceNo field to given value.

### HasReferenceNo

`func (o *CreateOrderResponse) HasReferenceNo() bool`

HasReferenceNo returns a boolean if a field has been set.

### GetPartnerReferenceNo

`func (o *CreateOrderResponse) GetPartnerReferenceNo() string`

GetPartnerReferenceNo returns the PartnerReferenceNo field if non-nil, zero value otherwise.

### GetPartnerReferenceNoOk

`func (o *CreateOrderResponse) GetPartnerReferenceNoOk() (*string, bool)`

GetPartnerReferenceNoOk returns a tuple with the PartnerReferenceNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerReferenceNo

`func (o *CreateOrderResponse) SetPartnerReferenceNo(v string)`

SetPartnerReferenceNo sets PartnerReferenceNo field to given value.


### GetWebRedirectUrl

`func (o *CreateOrderResponse) GetWebRedirectUrl() string`

GetWebRedirectUrl returns the WebRedirectUrl field if non-nil, zero value otherwise.

### GetWebRedirectUrlOk

`func (o *CreateOrderResponse) GetWebRedirectUrlOk() (*string, bool)`

GetWebRedirectUrlOk returns a tuple with the WebRedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebRedirectUrl

`func (o *CreateOrderResponse) SetWebRedirectUrl(v string)`

SetWebRedirectUrl sets WebRedirectUrl field to given value.

### HasWebRedirectUrl

`func (o *CreateOrderResponse) HasWebRedirectUrl() bool`

HasWebRedirectUrl returns a boolean if a field has been set.

### GetAdditionalInfo

`func (o *CreateOrderResponse) GetAdditionalInfo() CreateOrderResponseAdditionalInfo`

GetAdditionalInfo returns the AdditionalInfo field if non-nil, zero value otherwise.

### GetAdditionalInfoOk

`func (o *CreateOrderResponse) GetAdditionalInfoOk() (*CreateOrderResponseAdditionalInfo, bool)`

GetAdditionalInfoOk returns a tuple with the AdditionalInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalInfo

`func (o *CreateOrderResponse) SetAdditionalInfo(v CreateOrderResponseAdditionalInfo)`

SetAdditionalInfo sets AdditionalInfo field to given value.

### HasAdditionalInfo

`func (o *CreateOrderResponse) HasAdditionalInfo() bool`

HasAdditionalInfo returns a boolean if a field has been set.

### GetExternalOrderId

`func (o *CreateOrderResponse) GetExternalOrderId() string`

GetExternalOrderId returns the ExternalOrderId field if non-nil, zero value otherwise.

### GetExternalOrderIdOk

`func (o *CreateOrderResponse) GetExternalOrderIdOk() (*string, bool)`

GetExternalOrderIdOk returns a tuple with the ExternalOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalOrderId

`func (o *CreateOrderResponse) SetExternalOrderId(v string)`

SetExternalOrderId sets ExternalOrderId field to given value.

### HasExternalOrderId

`func (o *CreateOrderResponse) HasExternalOrderId() bool`

HasExternalOrderId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


