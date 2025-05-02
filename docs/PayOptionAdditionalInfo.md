# PayOptionAdditionalInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PhoneNumber** | Pointer to **string** |  | [optional] 
**PaymentCode** | Pointer to **string** |  | [optional] 
**PromoInfos** | Pointer to [**[]PromoInfo**](PromoInfo.md) |  | [optional] 

## Methods

### NewPayOptionAdditionalInfo

`func NewPayOptionAdditionalInfo() *PayOptionAdditionalInfo`

NewPayOptionAdditionalInfo instantiates a new PayOptionAdditionalInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPayOptionAdditionalInfoWithDefaults

`func NewPayOptionAdditionalInfoWithDefaults() *PayOptionAdditionalInfo`

NewPayOptionAdditionalInfoWithDefaults instantiates a new PayOptionAdditionalInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhoneNumber

`func (o *PayOptionAdditionalInfo) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *PayOptionAdditionalInfo) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *PayOptionAdditionalInfo) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *PayOptionAdditionalInfo) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetPaymentCode

`func (o *PayOptionAdditionalInfo) GetPaymentCode() string`

GetPaymentCode returns the PaymentCode field if non-nil, zero value otherwise.

### GetPaymentCodeOk

`func (o *PayOptionAdditionalInfo) GetPaymentCodeOk() (*string, bool)`

GetPaymentCodeOk returns a tuple with the PaymentCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentCode

`func (o *PayOptionAdditionalInfo) SetPaymentCode(v string)`

SetPaymentCode sets PaymentCode field to given value.

### HasPaymentCode

`func (o *PayOptionAdditionalInfo) HasPaymentCode() bool`

HasPaymentCode returns a boolean if a field has been set.

### GetPromoInfos

`func (o *PayOptionAdditionalInfo) GetPromoInfos() []PromoInfo`

GetPromoInfos returns the PromoInfos field if non-nil, zero value otherwise.

### GetPromoInfosOk

`func (o *PayOptionAdditionalInfo) GetPromoInfosOk() (*[]PromoInfo, bool)`

GetPromoInfosOk returns a tuple with the PromoInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromoInfos

`func (o *PayOptionAdditionalInfo) SetPromoInfos(v []PromoInfo)`

SetPromoInfos sets PromoInfos field to given value.

### HasPromoInfos

`func (o *PayOptionAdditionalInfo) HasPromoInfos() bool`

HasPromoInfos returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


