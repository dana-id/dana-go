# ConsultPayPaymentInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PayMethod** | **string** |  | 
**PayOption** | Pointer to **string** |  | [optional] 
**PromoInfos** | Pointer to [**[]PromoInfo**](PromoInfo.md) |  | [optional] 

## Methods

### NewConsultPayPaymentInfo

`func NewConsultPayPaymentInfo(payMethod string, ) *ConsultPayPaymentInfo`

NewConsultPayPaymentInfo instantiates a new ConsultPayPaymentInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsultPayPaymentInfoWithDefaults

`func NewConsultPayPaymentInfoWithDefaults() *ConsultPayPaymentInfo`

NewConsultPayPaymentInfoWithDefaults instantiates a new ConsultPayPaymentInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPayMethod

`func (o *ConsultPayPaymentInfo) GetPayMethod() string`

GetPayMethod returns the PayMethod field if non-nil, zero value otherwise.

### GetPayMethodOk

`func (o *ConsultPayPaymentInfo) GetPayMethodOk() (*string, bool)`

GetPayMethodOk returns a tuple with the PayMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayMethod

`func (o *ConsultPayPaymentInfo) SetPayMethod(v string)`

SetPayMethod sets PayMethod field to given value.


### GetPayOption

`func (o *ConsultPayPaymentInfo) GetPayOption() string`

GetPayOption returns the PayOption field if non-nil, zero value otherwise.

### GetPayOptionOk

`func (o *ConsultPayPaymentInfo) GetPayOptionOk() (*string, bool)`

GetPayOptionOk returns a tuple with the PayOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayOption

`func (o *ConsultPayPaymentInfo) SetPayOption(v string)`

SetPayOption sets PayOption field to given value.

### HasPayOption

`func (o *ConsultPayPaymentInfo) HasPayOption() bool`

HasPayOption returns a boolean if a field has been set.

### GetPromoInfos

`func (o *ConsultPayPaymentInfo) GetPromoInfos() []PromoInfo`

GetPromoInfos returns the PromoInfos field if non-nil, zero value otherwise.

### GetPromoInfosOk

`func (o *ConsultPayPaymentInfo) GetPromoInfosOk() (*[]PromoInfo, bool)`

GetPromoInfosOk returns a tuple with the PromoInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromoInfos

`func (o *ConsultPayPaymentInfo) SetPromoInfos(v []PromoInfo)`

SetPromoInfos sets PromoInfos field to given value.

### HasPromoInfos

`func (o *ConsultPayPaymentInfo) HasPromoInfos() bool`

HasPromoInfos returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


