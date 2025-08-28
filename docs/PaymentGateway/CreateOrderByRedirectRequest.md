# CreateOrderByRedirectRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalInfo** | Pointer to [**CreateOrderByRedirectAdditionalInfo**](CreateOrderByRedirectAdditionalInfo.md) |  | [optional] 
**PartnerReferenceNo** | **string** | Transaction identifier on partner system | 
**MerchantId** | **string** | Merchant identifier that is unique per each merchant | 
**SubMerchantId** | Pointer to **string** | Information of sub merchant identifier | [optional] 
**Amount** | [**Money**](Money.md) | Amount. Contains two sub-fields:<br /> 1. Value: Transaction amount, including the cents<br /> 2. Currency: Currency code based on ISO<br />  | 
**ExternalStoreId** | Pointer to **string** | Store identifier to indicate to which store this payment belongs to | [optional] 
**ValidUpTo** | Pointer to **string** | The time when the payment will be automatically expired, in format YYYY-MM-DDTHH:mm:ss+07:00. Time must be in GMT+7 (Jakarta time) and cannot be more than one week in the future. | [optional] 
**DisabledPayMethods** | Pointer to **string** | Payment method(s) that cannot be used for this | [optional] 
**UrlParams** | [**[]UrlParam**](UrlParam.md) | Notify URL that DANA must send the payment notification to | 

## Methods

### NewCreateOrderByRedirectRequest

`func NewCreateOrderByRedirectRequest(partnerReferenceNo string, merchantId string, amount Money, urlParams []UrlParam, ) *CreateOrderByRedirectRequest`

NewCreateOrderByRedirectRequest instantiates a new CreateOrderByRedirectRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrderByRedirectRequestWithDefaults

`func NewCreateOrderByRedirectRequestWithDefaults() *CreateOrderByRedirectRequest`

NewCreateOrderByRedirectRequestWithDefaults instantiates a new CreateOrderByRedirectRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalInfo

`func (o *CreateOrderByRedirectRequest) GetAdditionalInfo() CreateOrderByRedirectAdditionalInfo`

GetAdditionalInfo returns the AdditionalInfo field if non-nil, zero value otherwise.

### GetAdditionalInfoOk

`func (o *CreateOrderByRedirectRequest) GetAdditionalInfoOk() (*CreateOrderByRedirectAdditionalInfo, bool)`

GetAdditionalInfoOk returns a tuple with the AdditionalInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalInfo

`func (o *CreateOrderByRedirectRequest) SetAdditionalInfo(v CreateOrderByRedirectAdditionalInfo)`

SetAdditionalInfo sets AdditionalInfo field to given value.

### HasAdditionalInfo

`func (o *CreateOrderByRedirectRequest) HasAdditionalInfo() bool`

HasAdditionalInfo returns a boolean if a field has been set.

### GetPartnerReferenceNo

`func (o *CreateOrderByRedirectRequest) GetPartnerReferenceNo() string`

GetPartnerReferenceNo returns the PartnerReferenceNo field if non-nil, zero value otherwise.

### GetPartnerReferenceNoOk

`func (o *CreateOrderByRedirectRequest) GetPartnerReferenceNoOk() (*string, bool)`

GetPartnerReferenceNoOk returns a tuple with the PartnerReferenceNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerReferenceNo

`func (o *CreateOrderByRedirectRequest) SetPartnerReferenceNo(v string)`

SetPartnerReferenceNo sets PartnerReferenceNo field to given value.


### GetMerchantId

`func (o *CreateOrderByRedirectRequest) GetMerchantId() string`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *CreateOrderByRedirectRequest) GetMerchantIdOk() (*string, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *CreateOrderByRedirectRequest) SetMerchantId(v string)`

SetMerchantId sets MerchantId field to given value.


### GetSubMerchantId

`func (o *CreateOrderByRedirectRequest) GetSubMerchantId() string`

GetSubMerchantId returns the SubMerchantId field if non-nil, zero value otherwise.

### GetSubMerchantIdOk

`func (o *CreateOrderByRedirectRequest) GetSubMerchantIdOk() (*string, bool)`

GetSubMerchantIdOk returns a tuple with the SubMerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubMerchantId

`func (o *CreateOrderByRedirectRequest) SetSubMerchantId(v string)`

SetSubMerchantId sets SubMerchantId field to given value.

### HasSubMerchantId

`func (o *CreateOrderByRedirectRequest) HasSubMerchantId() bool`

HasSubMerchantId returns a boolean if a field has been set.

### GetAmount

`func (o *CreateOrderByRedirectRequest) GetAmount() Money`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CreateOrderByRedirectRequest) GetAmountOk() (*Money, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CreateOrderByRedirectRequest) SetAmount(v Money)`

SetAmount sets Amount field to given value.


### GetExternalStoreId

`func (o *CreateOrderByRedirectRequest) GetExternalStoreId() string`

GetExternalStoreId returns the ExternalStoreId field if non-nil, zero value otherwise.

### GetExternalStoreIdOk

`func (o *CreateOrderByRedirectRequest) GetExternalStoreIdOk() (*string, bool)`

GetExternalStoreIdOk returns a tuple with the ExternalStoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalStoreId

`func (o *CreateOrderByRedirectRequest) SetExternalStoreId(v string)`

SetExternalStoreId sets ExternalStoreId field to given value.

### HasExternalStoreId

`func (o *CreateOrderByRedirectRequest) HasExternalStoreId() bool`

HasExternalStoreId returns a boolean if a field has been set.

### GetValidUpTo

`func (o *CreateOrderByRedirectRequest) GetValidUpTo() string`

GetValidUpTo returns the ValidUpTo field if non-nil, zero value otherwise.

### GetValidUpToOk

`func (o *CreateOrderByRedirectRequest) GetValidUpToOk() (*string, bool)`

GetValidUpToOk returns a tuple with the ValidUpTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidUpTo

`func (o *CreateOrderByRedirectRequest) SetValidUpTo(v string)`

SetValidUpTo sets ValidUpTo field to given value.

### HasValidUpTo

`func (o *CreateOrderByRedirectRequest) HasValidUpTo() bool`

HasValidUpTo returns a boolean if a field has been set.

### GetDisabledPayMethods

`func (o *CreateOrderByRedirectRequest) GetDisabledPayMethods() string`

GetDisabledPayMethods returns the DisabledPayMethods field if non-nil, zero value otherwise.

### GetDisabledPayMethodsOk

`func (o *CreateOrderByRedirectRequest) GetDisabledPayMethodsOk() (*string, bool)`

GetDisabledPayMethodsOk returns a tuple with the DisabledPayMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledPayMethods

`func (o *CreateOrderByRedirectRequest) SetDisabledPayMethods(v string)`

SetDisabledPayMethods sets DisabledPayMethods field to given value.

### HasDisabledPayMethods

`func (o *CreateOrderByRedirectRequest) HasDisabledPayMethods() bool`

HasDisabledPayMethods returns a boolean if a field has been set.

### GetUrlParams

`func (o *CreateOrderByRedirectRequest) GetUrlParams() []UrlParam`

GetUrlParams returns the UrlParams field if non-nil, zero value otherwise.

### GetUrlParamsOk

`func (o *CreateOrderByRedirectRequest) GetUrlParamsOk() (*[]UrlParam, bool)`

GetUrlParamsOk returns a tuple with the UrlParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlParams

`func (o *CreateOrderByRedirectRequest) SetUrlParams(v []UrlParam)`

SetUrlParams sets UrlParams field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


