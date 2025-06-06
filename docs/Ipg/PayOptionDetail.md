# PayOptionDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PayMethod** | **string** | Payment Method, e.g. CREDIT_CARD | 
**PayOption** | **string** | Payment option which shows the provider of this payment e.g. CREDIT_CARD_VISA | 
**TransAmount** | Pointer to [**Money**](Money.md) | Trans amount. Contains value and currency | [optional] 
**FeeAmount** | Pointer to [**Money**](Money.md) | Fee amount. Contains value and currency | [optional] 
**CardToken** | Pointer to **string** | Card token used for this payment | [optional] 
**MerchantToken** | Pointer to **string** | Merchant token used for this payment | [optional] 
**AdditionalInfo** | Pointer to [**PayOptionDetailAdditionalInfo**](PayOptionDetailAdditionalInfo.md) |  | [optional] 

## Methods

### NewPayOptionDetail

`func NewPayOptionDetail(payMethod string, payOption string, ) *PayOptionDetail`

NewPayOptionDetail instantiates a new PayOptionDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPayOptionDetailWithDefaults

`func NewPayOptionDetailWithDefaults() *PayOptionDetail`

NewPayOptionDetailWithDefaults instantiates a new PayOptionDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPayMethod

`func (o *PayOptionDetail) GetPayMethod() string`

GetPayMethod returns the PayMethod field if non-nil, zero value otherwise.

### GetPayMethodOk

`func (o *PayOptionDetail) GetPayMethodOk() (*string, bool)`

GetPayMethodOk returns a tuple with the PayMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayMethod

`func (o *PayOptionDetail) SetPayMethod(v string)`

SetPayMethod sets PayMethod field to given value.


### GetPayOption

`func (o *PayOptionDetail) GetPayOption() string`

GetPayOption returns the PayOption field if non-nil, zero value otherwise.

### GetPayOptionOk

`func (o *PayOptionDetail) GetPayOptionOk() (*string, bool)`

GetPayOptionOk returns a tuple with the PayOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayOption

`func (o *PayOptionDetail) SetPayOption(v string)`

SetPayOption sets PayOption field to given value.


### GetTransAmount

`func (o *PayOptionDetail) GetTransAmount() Money`

GetTransAmount returns the TransAmount field if non-nil, zero value otherwise.

### GetTransAmountOk

`func (o *PayOptionDetail) GetTransAmountOk() (*Money, bool)`

GetTransAmountOk returns a tuple with the TransAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransAmount

`func (o *PayOptionDetail) SetTransAmount(v Money)`

SetTransAmount sets TransAmount field to given value.

### HasTransAmount

`func (o *PayOptionDetail) HasTransAmount() bool`

HasTransAmount returns a boolean if a field has been set.

### GetFeeAmount

`func (o *PayOptionDetail) GetFeeAmount() Money`

GetFeeAmount returns the FeeAmount field if non-nil, zero value otherwise.

### GetFeeAmountOk

`func (o *PayOptionDetail) GetFeeAmountOk() (*Money, bool)`

GetFeeAmountOk returns a tuple with the FeeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeAmount

`func (o *PayOptionDetail) SetFeeAmount(v Money)`

SetFeeAmount sets FeeAmount field to given value.

### HasFeeAmount

`func (o *PayOptionDetail) HasFeeAmount() bool`

HasFeeAmount returns a boolean if a field has been set.

### GetCardToken

`func (o *PayOptionDetail) GetCardToken() string`

GetCardToken returns the CardToken field if non-nil, zero value otherwise.

### GetCardTokenOk

`func (o *PayOptionDetail) GetCardTokenOk() (*string, bool)`

GetCardTokenOk returns a tuple with the CardToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardToken

`func (o *PayOptionDetail) SetCardToken(v string)`

SetCardToken sets CardToken field to given value.

### HasCardToken

`func (o *PayOptionDetail) HasCardToken() bool`

HasCardToken returns a boolean if a field has been set.

### GetMerchantToken

`func (o *PayOptionDetail) GetMerchantToken() string`

GetMerchantToken returns the MerchantToken field if non-nil, zero value otherwise.

### GetMerchantTokenOk

`func (o *PayOptionDetail) GetMerchantTokenOk() (*string, bool)`

GetMerchantTokenOk returns a tuple with the MerchantToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantToken

`func (o *PayOptionDetail) SetMerchantToken(v string)`

SetMerchantToken sets MerchantToken field to given value.

### HasMerchantToken

`func (o *PayOptionDetail) HasMerchantToken() bool`

HasMerchantToken returns a boolean if a field has been set.

### GetAdditionalInfo

`func (o *PayOptionDetail) GetAdditionalInfo() PayOptionDetailAdditionalInfo`

GetAdditionalInfo returns the AdditionalInfo field if non-nil, zero value otherwise.

### GetAdditionalInfoOk

`func (o *PayOptionDetail) GetAdditionalInfoOk() (*PayOptionDetailAdditionalInfo, bool)`

GetAdditionalInfoOk returns a tuple with the AdditionalInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalInfo

`func (o *PayOptionDetail) SetAdditionalInfo(v PayOptionDetailAdditionalInfo)`

SetAdditionalInfo sets AdditionalInfo field to given value.

### HasAdditionalInfo

`func (o *PayOptionDetail) HasAdditionalInfo() bool`

HasAdditionalInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


