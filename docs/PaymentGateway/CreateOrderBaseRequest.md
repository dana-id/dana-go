# CreateOrderBaseRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PartnerReferenceNo** | **string** | Transaction identifier on partner system | 
**MerchantId** | **string** | Merchant identifier that is unique per each merchant | 
**SubMerchantId** | Pointer to **string** | Information of sub merchant identifier | [optional] 
**Amount** | [**Money**](Money.md) | Amount. Contains two sub-fields:<br /> 1. Value: Transaction amount, including the cents<br /> 2. Currency: Currency code based on ISO<br />  | 
**ExternalStoreId** | Pointer to **string** | Store identifier to indicate to which store this payment belongs to | [optional] 
**ValidUpTo** | Pointer to **string** | The time when the payment will be automatically expired, in format YYYY-MM-DDTHH:mm:ss+07:00. Time must be in GMT+7 (Jakarta time) | [optional] 
**DisabledPayMethods** | Pointer to **string** | Payment method(s) that cannot be used for this | [optional] 
**UrlParams** | [**[]UrlParam**](UrlParam.md) | Notify URL that DANA must send the payment notification to | 

## Methods

### NewCreateOrderBaseRequest

`func NewCreateOrderBaseRequest(partnerReferenceNo string, merchantId string, amount Money, urlParams []UrlParam, ) *CreateOrderBaseRequest`

NewCreateOrderBaseRequest instantiates a new CreateOrderBaseRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrderBaseRequestWithDefaults

`func NewCreateOrderBaseRequestWithDefaults() *CreateOrderBaseRequest`

NewCreateOrderBaseRequestWithDefaults instantiates a new CreateOrderBaseRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPartnerReferenceNo

`func (o *CreateOrderBaseRequest) GetPartnerReferenceNo() string`

GetPartnerReferenceNo returns the PartnerReferenceNo field if non-nil, zero value otherwise.

### GetPartnerReferenceNoOk

`func (o *CreateOrderBaseRequest) GetPartnerReferenceNoOk() (*string, bool)`

GetPartnerReferenceNoOk returns a tuple with the PartnerReferenceNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerReferenceNo

`func (o *CreateOrderBaseRequest) SetPartnerReferenceNo(v string)`

SetPartnerReferenceNo sets PartnerReferenceNo field to given value.


### GetMerchantId

`func (o *CreateOrderBaseRequest) GetMerchantId() string`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *CreateOrderBaseRequest) GetMerchantIdOk() (*string, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *CreateOrderBaseRequest) SetMerchantId(v string)`

SetMerchantId sets MerchantId field to given value.


### GetSubMerchantId

`func (o *CreateOrderBaseRequest) GetSubMerchantId() string`

GetSubMerchantId returns the SubMerchantId field if non-nil, zero value otherwise.

### GetSubMerchantIdOk

`func (o *CreateOrderBaseRequest) GetSubMerchantIdOk() (*string, bool)`

GetSubMerchantIdOk returns a tuple with the SubMerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubMerchantId

`func (o *CreateOrderBaseRequest) SetSubMerchantId(v string)`

SetSubMerchantId sets SubMerchantId field to given value.

### HasSubMerchantId

`func (o *CreateOrderBaseRequest) HasSubMerchantId() bool`

HasSubMerchantId returns a boolean if a field has been set.

### GetAmount

`func (o *CreateOrderBaseRequest) GetAmount() Money`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CreateOrderBaseRequest) GetAmountOk() (*Money, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CreateOrderBaseRequest) SetAmount(v Money)`

SetAmount sets Amount field to given value.


### GetExternalStoreId

`func (o *CreateOrderBaseRequest) GetExternalStoreId() string`

GetExternalStoreId returns the ExternalStoreId field if non-nil, zero value otherwise.

### GetExternalStoreIdOk

`func (o *CreateOrderBaseRequest) GetExternalStoreIdOk() (*string, bool)`

GetExternalStoreIdOk returns a tuple with the ExternalStoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalStoreId

`func (o *CreateOrderBaseRequest) SetExternalStoreId(v string)`

SetExternalStoreId sets ExternalStoreId field to given value.

### HasExternalStoreId

`func (o *CreateOrderBaseRequest) HasExternalStoreId() bool`

HasExternalStoreId returns a boolean if a field has been set.

### GetValidUpTo

`func (o *CreateOrderBaseRequest) GetValidUpTo() string`

GetValidUpTo returns the ValidUpTo field if non-nil, zero value otherwise.

### GetValidUpToOk

`func (o *CreateOrderBaseRequest) GetValidUpToOk() (*string, bool)`

GetValidUpToOk returns a tuple with the ValidUpTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidUpTo

`func (o *CreateOrderBaseRequest) SetValidUpTo(v string)`

SetValidUpTo sets ValidUpTo field to given value.

### HasValidUpTo

`func (o *CreateOrderBaseRequest) HasValidUpTo() bool`

HasValidUpTo returns a boolean if a field has been set.

### GetDisabledPayMethods

`func (o *CreateOrderBaseRequest) GetDisabledPayMethods() string`

GetDisabledPayMethods returns the DisabledPayMethods field if non-nil, zero value otherwise.

### GetDisabledPayMethodsOk

`func (o *CreateOrderBaseRequest) GetDisabledPayMethodsOk() (*string, bool)`

GetDisabledPayMethodsOk returns a tuple with the DisabledPayMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledPayMethods

`func (o *CreateOrderBaseRequest) SetDisabledPayMethods(v string)`

SetDisabledPayMethods sets DisabledPayMethods field to given value.

### HasDisabledPayMethods

`func (o *CreateOrderBaseRequest) HasDisabledPayMethods() bool`

HasDisabledPayMethods returns a boolean if a field has been set.

### GetUrlParams

`func (o *CreateOrderBaseRequest) GetUrlParams() []UrlParam`

GetUrlParams returns the UrlParams field if non-nil, zero value otherwise.

### GetUrlParamsOk

`func (o *CreateOrderBaseRequest) GetUrlParamsOk() (*[]UrlParam, bool)`

GetUrlParamsOk returns a tuple with the UrlParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlParams

`func (o *CreateOrderBaseRequest) SetUrlParams(v []UrlParam)`

SetUrlParams sets UrlParams field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


