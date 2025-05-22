# ConsultPayResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponseCode** | Pointer to **string** | Response code. Refer to https://dashboard.dana.id/api-docs/read/237#paymentgatewayprod-ConsultPay-ResponseCodeandMessage | [optional] 
**ResponseMessage** | Pointer to **string** | Response message. Refer to https://dashboard.dana.id/api-docs/read/237#paymentgatewayprod-ConsultPay-ResponseCodeandMessage | [optional] 
**PaymentInfos** | Pointer to [**[]ConsultPayPaymentInfo**](ConsultPayPaymentInfo.md) | Define list of payment information that includes payment method and payment option for transaction | [optional] 

## Methods

### NewConsultPayResponse

`func NewConsultPayResponse() *ConsultPayResponse`

NewConsultPayResponse instantiates a new ConsultPayResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsultPayResponseWithDefaults

`func NewConsultPayResponseWithDefaults() *ConsultPayResponse`

NewConsultPayResponseWithDefaults instantiates a new ConsultPayResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponseCode

`func (o *ConsultPayResponse) GetResponseCode() string`

GetResponseCode returns the ResponseCode field if non-nil, zero value otherwise.

### GetResponseCodeOk

`func (o *ConsultPayResponse) GetResponseCodeOk() (*string, bool)`

GetResponseCodeOk returns a tuple with the ResponseCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseCode

`func (o *ConsultPayResponse) SetResponseCode(v string)`

SetResponseCode sets ResponseCode field to given value.

### HasResponseCode

`func (o *ConsultPayResponse) HasResponseCode() bool`

HasResponseCode returns a boolean if a field has been set.

### GetResponseMessage

`func (o *ConsultPayResponse) GetResponseMessage() string`

GetResponseMessage returns the ResponseMessage field if non-nil, zero value otherwise.

### GetResponseMessageOk

`func (o *ConsultPayResponse) GetResponseMessageOk() (*string, bool)`

GetResponseMessageOk returns a tuple with the ResponseMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseMessage

`func (o *ConsultPayResponse) SetResponseMessage(v string)`

SetResponseMessage sets ResponseMessage field to given value.

### HasResponseMessage

`func (o *ConsultPayResponse) HasResponseMessage() bool`

HasResponseMessage returns a boolean if a field has been set.

### GetPaymentInfos

`func (o *ConsultPayResponse) GetPaymentInfos() []ConsultPayPaymentInfo`

GetPaymentInfos returns the PaymentInfos field if non-nil, zero value otherwise.

### GetPaymentInfosOk

`func (o *ConsultPayResponse) GetPaymentInfosOk() (*[]ConsultPayPaymentInfo, bool)`

GetPaymentInfosOk returns a tuple with the PaymentInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentInfos

`func (o *ConsultPayResponse) SetPaymentInfos(v []ConsultPayPaymentInfo)`

SetPaymentInfos sets PaymentInfos field to given value.

### HasPaymentInfos

`func (o *ConsultPayResponse) HasPaymentInfos() bool`

HasPaymentInfos returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


