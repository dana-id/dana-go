# ConsultPayRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MerchantId** | **string** | Merchant identifier that is unique per each merchant | 
**Amount** | [**Money**](Money.md) | Amount. Contains two sub-fields:<br /> 1. Value: Transaction amount, including the cents<br /> 2. Currency: Currency code based on ISO<br />  | 
**AdditionalInfo** | [**ConsultPayRequestAdditionalInfo**](ConsultPayRequestAdditionalInfo.md) | Additional information | 
**ExternalStoreId** | Pointer to **string** | Store identifier to indicate to which store this payment belongs to. Need to be provided to show QRIS payment method. | [optional] 

## Methods

### NewConsultPayRequest

`func NewConsultPayRequest(merchantId string, amount Money, additionalInfo ConsultPayRequestAdditionalInfo, ) *ConsultPayRequest`

NewConsultPayRequest instantiates a new ConsultPayRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsultPayRequestWithDefaults

`func NewConsultPayRequestWithDefaults() *ConsultPayRequest`

NewConsultPayRequestWithDefaults instantiates a new ConsultPayRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMerchantId

`func (o *ConsultPayRequest) GetMerchantId() string`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *ConsultPayRequest) GetMerchantIdOk() (*string, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *ConsultPayRequest) SetMerchantId(v string)`

SetMerchantId sets MerchantId field to given value.


### GetAmount

`func (o *ConsultPayRequest) GetAmount() Money`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ConsultPayRequest) GetAmountOk() (*Money, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ConsultPayRequest) SetAmount(v Money)`

SetAmount sets Amount field to given value.


### GetAdditionalInfo

`func (o *ConsultPayRequest) GetAdditionalInfo() ConsultPayRequestAdditionalInfo`

GetAdditionalInfo returns the AdditionalInfo field if non-nil, zero value otherwise.

### GetAdditionalInfoOk

`func (o *ConsultPayRequest) GetAdditionalInfoOk() (*ConsultPayRequestAdditionalInfo, bool)`

GetAdditionalInfoOk returns a tuple with the AdditionalInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalInfo

`func (o *ConsultPayRequest) SetAdditionalInfo(v ConsultPayRequestAdditionalInfo)`

SetAdditionalInfo sets AdditionalInfo field to given value.


### GetExternalStoreId

`func (o *ConsultPayRequest) GetExternalStoreId() string`

GetExternalStoreId returns the ExternalStoreId field if non-nil, zero value otherwise.

### GetExternalStoreIdOk

`func (o *ConsultPayRequest) GetExternalStoreIdOk() (*string, bool)`

GetExternalStoreIdOk returns a tuple with the ExternalStoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalStoreId

`func (o *ConsultPayRequest) SetExternalStoreId(v string)`

SetExternalStoreId sets ExternalStoreId field to given value.

### HasExternalStoreId

`func (o *ConsultPayRequest) HasExternalStoreId() bool`

HasExternalStoreId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


