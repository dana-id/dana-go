# CreateOrderResponseAdditionalInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentCode** | Pointer to **string** | Payment code (only for Virtual Account / QRIS) | [optional] 

## Methods

### NewCreateOrderResponseAdditionalInfo

`func NewCreateOrderResponseAdditionalInfo() *CreateOrderResponseAdditionalInfo`

NewCreateOrderResponseAdditionalInfo instantiates a new CreateOrderResponseAdditionalInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrderResponseAdditionalInfoWithDefaults

`func NewCreateOrderResponseAdditionalInfoWithDefaults() *CreateOrderResponseAdditionalInfo`

NewCreateOrderResponseAdditionalInfoWithDefaults instantiates a new CreateOrderResponseAdditionalInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaymentCode

`func (o *CreateOrderResponseAdditionalInfo) GetPaymentCode() string`

GetPaymentCode returns the PaymentCode field if non-nil, zero value otherwise.

### GetPaymentCodeOk

`func (o *CreateOrderResponseAdditionalInfo) GetPaymentCodeOk() (*string, bool)`

GetPaymentCodeOk returns a tuple with the PaymentCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentCode

`func (o *CreateOrderResponseAdditionalInfo) SetPaymentCode(v string)`

SetPaymentCode sets PaymentCode field to given value.

### HasPaymentCode

`func (o *CreateOrderResponseAdditionalInfo) HasPaymentCode() bool`

HasPaymentCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


