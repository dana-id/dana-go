# InternationalOrderInfoExchangeRate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rate** | Pointer to **string** | Rate of exchange value represents the relation between two currencies. For example, 15917.2690 indicates that one USD is equivalent to 15917.2690 IDR | [optional] 
**ExchangeRelation** | Pointer to **string** | Exchange rate between two currencies. For example USD/IDR, refers to how much of one currency (in this case, Indonesian Rupiah or IDR) can be exchanged for one unit of another currency (in this case, US Dollar or USD) | [optional] 

## Methods

### NewInternationalOrderInfoExchangeRate

`func NewInternationalOrderInfoExchangeRate() *InternationalOrderInfoExchangeRate`

NewInternationalOrderInfoExchangeRate instantiates a new InternationalOrderInfoExchangeRate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternationalOrderInfoExchangeRateWithDefaults

`func NewInternationalOrderInfoExchangeRateWithDefaults() *InternationalOrderInfoExchangeRate`

NewInternationalOrderInfoExchangeRateWithDefaults instantiates a new InternationalOrderInfoExchangeRate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRate

`func (o *InternationalOrderInfoExchangeRate) GetRate() string`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *InternationalOrderInfoExchangeRate) GetRateOk() (*string, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *InternationalOrderInfoExchangeRate) SetRate(v string)`

SetRate sets Rate field to given value.

### HasRate

`func (o *InternationalOrderInfoExchangeRate) HasRate() bool`

HasRate returns a boolean if a field has been set.

### GetExchangeRelation

`func (o *InternationalOrderInfoExchangeRate) GetExchangeRelation() string`

GetExchangeRelation returns the ExchangeRelation field if non-nil, zero value otherwise.

### GetExchangeRelationOk

`func (o *InternationalOrderInfoExchangeRate) GetExchangeRelationOk() (*string, bool)`

GetExchangeRelationOk returns a tuple with the ExchangeRelation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeRelation

`func (o *InternationalOrderInfoExchangeRate) SetExchangeRelation(v string)`

SetExchangeRelation sets ExchangeRelation field to given value.

### HasExchangeRelation

`func (o *InternationalOrderInfoExchangeRate) HasExchangeRelation() bool`

HasExchangeRelation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


