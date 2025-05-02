# ConsultPayRequestAdditionalInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Buyer** | [**Buyer**](Buyer.md) |  | 
**EnvInfo** | [**EnvInfo**](EnvInfo.md) |  | 
**MerchantTransType** | Pointer to **string** |  | [optional] 

## Methods

### NewConsultPayRequestAdditionalInfo

`func NewConsultPayRequestAdditionalInfo(buyer Buyer, envInfo EnvInfo, ) *ConsultPayRequestAdditionalInfo`

NewConsultPayRequestAdditionalInfo instantiates a new ConsultPayRequestAdditionalInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsultPayRequestAdditionalInfoWithDefaults

`func NewConsultPayRequestAdditionalInfoWithDefaults() *ConsultPayRequestAdditionalInfo`

NewConsultPayRequestAdditionalInfoWithDefaults instantiates a new ConsultPayRequestAdditionalInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuyer

`func (o *ConsultPayRequestAdditionalInfo) GetBuyer() Buyer`

GetBuyer returns the Buyer field if non-nil, zero value otherwise.

### GetBuyerOk

`func (o *ConsultPayRequestAdditionalInfo) GetBuyerOk() (*Buyer, bool)`

GetBuyerOk returns a tuple with the Buyer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyer

`func (o *ConsultPayRequestAdditionalInfo) SetBuyer(v Buyer)`

SetBuyer sets Buyer field to given value.


### GetEnvInfo

`func (o *ConsultPayRequestAdditionalInfo) GetEnvInfo() EnvInfo`

GetEnvInfo returns the EnvInfo field if non-nil, zero value otherwise.

### GetEnvInfoOk

`func (o *ConsultPayRequestAdditionalInfo) GetEnvInfoOk() (*EnvInfo, bool)`

GetEnvInfoOk returns a tuple with the EnvInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvInfo

`func (o *ConsultPayRequestAdditionalInfo) SetEnvInfo(v EnvInfo)`

SetEnvInfo sets EnvInfo field to given value.


### GetMerchantTransType

`func (o *ConsultPayRequestAdditionalInfo) GetMerchantTransType() string`

GetMerchantTransType returns the MerchantTransType field if non-nil, zero value otherwise.

### GetMerchantTransTypeOk

`func (o *ConsultPayRequestAdditionalInfo) GetMerchantTransTypeOk() (*string, bool)`

GetMerchantTransTypeOk returns a tuple with the MerchantTransType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantTransType

`func (o *ConsultPayRequestAdditionalInfo) SetMerchantTransType(v string)`

SetMerchantTransType sets MerchantTransType field to given value.

### HasMerchantTransType

`func (o *ConsultPayRequestAdditionalInfo) HasMerchantTransType() bool`

HasMerchantTransType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


