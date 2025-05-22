# QueryPaymentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OriginalPartnerReferenceNo** | Pointer to **string** | Original transaction identifier on partner system. Required if originalReferenceNo is not filled | [optional] 
**OriginalReferenceNo** | Pointer to **string** | Original transaction identifier on DANA system. Required if originalPartnerReferenceNo is not filled | [optional] 
**OriginalExternalId** | Pointer to **string** | Original external identifier on header message | [optional] 
**ServiceCode** | **string** | Transaction type indicator is based on the service code of the original transaction request:<br /> - IPG Cashier Pay - SNAP: 54<br /> - QRIS CPM (Acquirer) - SNAP: 60<br /> - QRIS MPM (Acquirer) - SNAP: 47<br /> - Payment Gateway: 54<br />  | [default to "54"]
**TransactionDate** | Pointer to **string** | Transaction date, in format YYYY-MM-DDTHH:mm:ss+07:00. Time must be in GMT+7 (Jakarta time) | [optional] 
**Amount** | Pointer to [**Money**](Money.md) | Amount. Contains two sub-fields:<br /> 1. Value: Transaction amount, including the cents<br /> 2. Currency: Currency code based on ISO<br />  | [optional] 
**MerchantId** | **string** | Merchant identifier that is unique per each merchant | 
**SubMerchantId** | Pointer to **string** | Information of sub merchant identifier | [optional] 
**ExternalStoreId** | Pointer to **string** | Store identifier to indicate to which store this payment belongs to | [optional] 
**AdditionalInfo** | Pointer to **map[string]interface{}** | Additional information | [optional] 

## Methods

### NewQueryPaymentRequest

`func NewQueryPaymentRequest(serviceCode string, merchantId string, ) *QueryPaymentRequest`

NewQueryPaymentRequest instantiates a new QueryPaymentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryPaymentRequestWithDefaults

`func NewQueryPaymentRequestWithDefaults() *QueryPaymentRequest`

NewQueryPaymentRequestWithDefaults instantiates a new QueryPaymentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOriginalPartnerReferenceNo

`func (o *QueryPaymentRequest) GetOriginalPartnerReferenceNo() string`

GetOriginalPartnerReferenceNo returns the OriginalPartnerReferenceNo field if non-nil, zero value otherwise.

### GetOriginalPartnerReferenceNoOk

`func (o *QueryPaymentRequest) GetOriginalPartnerReferenceNoOk() (*string, bool)`

GetOriginalPartnerReferenceNoOk returns a tuple with the OriginalPartnerReferenceNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalPartnerReferenceNo

`func (o *QueryPaymentRequest) SetOriginalPartnerReferenceNo(v string)`

SetOriginalPartnerReferenceNo sets OriginalPartnerReferenceNo field to given value.

### HasOriginalPartnerReferenceNo

`func (o *QueryPaymentRequest) HasOriginalPartnerReferenceNo() bool`

HasOriginalPartnerReferenceNo returns a boolean if a field has been set.

### GetOriginalReferenceNo

`func (o *QueryPaymentRequest) GetOriginalReferenceNo() string`

GetOriginalReferenceNo returns the OriginalReferenceNo field if non-nil, zero value otherwise.

### GetOriginalReferenceNoOk

`func (o *QueryPaymentRequest) GetOriginalReferenceNoOk() (*string, bool)`

GetOriginalReferenceNoOk returns a tuple with the OriginalReferenceNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalReferenceNo

`func (o *QueryPaymentRequest) SetOriginalReferenceNo(v string)`

SetOriginalReferenceNo sets OriginalReferenceNo field to given value.

### HasOriginalReferenceNo

`func (o *QueryPaymentRequest) HasOriginalReferenceNo() bool`

HasOriginalReferenceNo returns a boolean if a field has been set.

### GetOriginalExternalId

`func (o *QueryPaymentRequest) GetOriginalExternalId() string`

GetOriginalExternalId returns the OriginalExternalId field if non-nil, zero value otherwise.

### GetOriginalExternalIdOk

`func (o *QueryPaymentRequest) GetOriginalExternalIdOk() (*string, bool)`

GetOriginalExternalIdOk returns a tuple with the OriginalExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalExternalId

`func (o *QueryPaymentRequest) SetOriginalExternalId(v string)`

SetOriginalExternalId sets OriginalExternalId field to given value.

### HasOriginalExternalId

`func (o *QueryPaymentRequest) HasOriginalExternalId() bool`

HasOriginalExternalId returns a boolean if a field has been set.

### GetServiceCode

`func (o *QueryPaymentRequest) GetServiceCode() string`

GetServiceCode returns the ServiceCode field if non-nil, zero value otherwise.

### GetServiceCodeOk

`func (o *QueryPaymentRequest) GetServiceCodeOk() (*string, bool)`

GetServiceCodeOk returns a tuple with the ServiceCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceCode

`func (o *QueryPaymentRequest) SetServiceCode(v string)`

SetServiceCode sets ServiceCode field to given value.


### GetTransactionDate

`func (o *QueryPaymentRequest) GetTransactionDate() string`

GetTransactionDate returns the TransactionDate field if non-nil, zero value otherwise.

### GetTransactionDateOk

`func (o *QueryPaymentRequest) GetTransactionDateOk() (*string, bool)`

GetTransactionDateOk returns a tuple with the TransactionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionDate

`func (o *QueryPaymentRequest) SetTransactionDate(v string)`

SetTransactionDate sets TransactionDate field to given value.

### HasTransactionDate

`func (o *QueryPaymentRequest) HasTransactionDate() bool`

HasTransactionDate returns a boolean if a field has been set.

### GetAmount

`func (o *QueryPaymentRequest) GetAmount() Money`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *QueryPaymentRequest) GetAmountOk() (*Money, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *QueryPaymentRequest) SetAmount(v Money)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *QueryPaymentRequest) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetMerchantId

`func (o *QueryPaymentRequest) GetMerchantId() string`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *QueryPaymentRequest) GetMerchantIdOk() (*string, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *QueryPaymentRequest) SetMerchantId(v string)`

SetMerchantId sets MerchantId field to given value.


### GetSubMerchantId

`func (o *QueryPaymentRequest) GetSubMerchantId() string`

GetSubMerchantId returns the SubMerchantId field if non-nil, zero value otherwise.

### GetSubMerchantIdOk

`func (o *QueryPaymentRequest) GetSubMerchantIdOk() (*string, bool)`

GetSubMerchantIdOk returns a tuple with the SubMerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubMerchantId

`func (o *QueryPaymentRequest) SetSubMerchantId(v string)`

SetSubMerchantId sets SubMerchantId field to given value.

### HasSubMerchantId

`func (o *QueryPaymentRequest) HasSubMerchantId() bool`

HasSubMerchantId returns a boolean if a field has been set.

### GetExternalStoreId

`func (o *QueryPaymentRequest) GetExternalStoreId() string`

GetExternalStoreId returns the ExternalStoreId field if non-nil, zero value otherwise.

### GetExternalStoreIdOk

`func (o *QueryPaymentRequest) GetExternalStoreIdOk() (*string, bool)`

GetExternalStoreIdOk returns a tuple with the ExternalStoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalStoreId

`func (o *QueryPaymentRequest) SetExternalStoreId(v string)`

SetExternalStoreId sets ExternalStoreId field to given value.

### HasExternalStoreId

`func (o *QueryPaymentRequest) HasExternalStoreId() bool`

HasExternalStoreId returns a boolean if a field has been set.

### GetAdditionalInfo

`func (o *QueryPaymentRequest) GetAdditionalInfo() map[string]interface{}`

GetAdditionalInfo returns the AdditionalInfo field if non-nil, zero value otherwise.

### GetAdditionalInfoOk

`func (o *QueryPaymentRequest) GetAdditionalInfoOk() (*map[string]interface{}, bool)`

GetAdditionalInfoOk returns a tuple with the AdditionalInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalInfo

`func (o *QueryPaymentRequest) SetAdditionalInfo(v map[string]interface{})`

SetAdditionalInfo sets AdditionalInfo field to given value.

### HasAdditionalInfo

`func (o *QueryPaymentRequest) HasAdditionalInfo() bool`

HasAdditionalInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


