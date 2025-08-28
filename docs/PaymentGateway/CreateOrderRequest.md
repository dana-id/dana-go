# CreateOrderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalInfo** | Pointer to [**CreateOrderByApiAdditionalInfo**](CreateOrderByApiAdditionalInfo.md) |  | [optional] 
**PartnerReferenceNo** | **string** | Transaction identifier on partner system | 
**MerchantId** | **string** | Merchant identifier that is unique per each merchant | 
**SubMerchantId** | Pointer to **string** | Information of sub merchant identifier | [optional] 
**Amount** | [**Money**](Money.md) | Amount. Contains two sub-fields:<br /> 1. Value: Transaction amount, including the cents<br /> 2. Currency: Currency code based on ISO<br />  | 
**ExternalStoreId** | Pointer to **string** | Store identifier to indicate to which store this payment belongs to | [optional] 
**ValidUpTo** | Pointer to **string** | The time when the payment will be automatically expired, in format YYYY-MM-DDTHH:mm:ss+07:00. Time must be in GMT+7 (Jakarta time) and cannot be more than one week in the future. | [optional] 
**DisabledPayMethods** | Pointer to **string** | Payment method(s) that cannot be used for this | [optional] 
**UrlParams** | [**[]UrlParam**](UrlParam.md) | Notify URL that DANA must send the payment notification to | 
**PayOptionDetails** | [**[]PayOptionDetail**](PayOptionDetail.md) |  | 

## Methods

### NewCreateOrderRequest

`func NewCreateOrderRequest(partnerReferenceNo string, merchantId string, amount Money, urlParams []UrlParam, payOptionDetails []PayOptionDetail, ) *CreateOrderRequest`

NewCreateOrderRequest instantiates a new CreateOrderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrderRequestWithDefaults

`func NewCreateOrderRequestWithDefaults() *CreateOrderRequest`

NewCreateOrderRequestWithDefaults instantiates a new CreateOrderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalInfo

`func (o *CreateOrderRequest) GetAdditionalInfo() CreateOrderByApiAdditionalInfo`

GetAdditionalInfo returns the AdditionalInfo field if non-nil, zero value otherwise.

### GetAdditionalInfoOk

`func (o *CreateOrderRequest) GetAdditionalInfoOk() (*CreateOrderByApiAdditionalInfo, bool)`

GetAdditionalInfoOk returns a tuple with the AdditionalInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalInfo

`func (o *CreateOrderRequest) SetAdditionalInfo(v CreateOrderByApiAdditionalInfo)`

SetAdditionalInfo sets AdditionalInfo field to given value.

### HasAdditionalInfo

`func (o *CreateOrderRequest) HasAdditionalInfo() bool`

HasAdditionalInfo returns a boolean if a field has been set.

### GetPartnerReferenceNo

`func (o *CreateOrderRequest) GetPartnerReferenceNo() string`

GetPartnerReferenceNo returns the PartnerReferenceNo field if non-nil, zero value otherwise.

### GetPartnerReferenceNoOk

`func (o *CreateOrderRequest) GetPartnerReferenceNoOk() (*string, bool)`

GetPartnerReferenceNoOk returns a tuple with the PartnerReferenceNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerReferenceNo

`func (o *CreateOrderRequest) SetPartnerReferenceNo(v string)`

SetPartnerReferenceNo sets PartnerReferenceNo field to given value.


### GetMerchantId

`func (o *CreateOrderRequest) GetMerchantId() string`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *CreateOrderRequest) GetMerchantIdOk() (*string, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *CreateOrderRequest) SetMerchantId(v string)`

SetMerchantId sets MerchantId field to given value.


### GetSubMerchantId

`func (o *CreateOrderRequest) GetSubMerchantId() string`

GetSubMerchantId returns the SubMerchantId field if non-nil, zero value otherwise.

### GetSubMerchantIdOk

`func (o *CreateOrderRequest) GetSubMerchantIdOk() (*string, bool)`

GetSubMerchantIdOk returns a tuple with the SubMerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubMerchantId

`func (o *CreateOrderRequest) SetSubMerchantId(v string)`

SetSubMerchantId sets SubMerchantId field to given value.

### HasSubMerchantId

`func (o *CreateOrderRequest) HasSubMerchantId() bool`

HasSubMerchantId returns a boolean if a field has been set.

### GetAmount

`func (o *CreateOrderRequest) GetAmount() Money`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CreateOrderRequest) GetAmountOk() (*Money, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CreateOrderRequest) SetAmount(v Money)`

SetAmount sets Amount field to given value.


### GetExternalStoreId

`func (o *CreateOrderRequest) GetExternalStoreId() string`

GetExternalStoreId returns the ExternalStoreId field if non-nil, zero value otherwise.

### GetExternalStoreIdOk

`func (o *CreateOrderRequest) GetExternalStoreIdOk() (*string, bool)`

GetExternalStoreIdOk returns a tuple with the ExternalStoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalStoreId

`func (o *CreateOrderRequest) SetExternalStoreId(v string)`

SetExternalStoreId sets ExternalStoreId field to given value.

### HasExternalStoreId

`func (o *CreateOrderRequest) HasExternalStoreId() bool`

HasExternalStoreId returns a boolean if a field has been set.

### GetValidUpTo

`func (o *CreateOrderRequest) GetValidUpTo() string`

GetValidUpTo returns the ValidUpTo field if non-nil, zero value otherwise.

### GetValidUpToOk

`func (o *CreateOrderRequest) GetValidUpToOk() (*string, bool)`

GetValidUpToOk returns a tuple with the ValidUpTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidUpTo

`func (o *CreateOrderRequest) SetValidUpTo(v string)`

SetValidUpTo sets ValidUpTo field to given value.

### HasValidUpTo

`func (o *CreateOrderRequest) HasValidUpTo() bool`

HasValidUpTo returns a boolean if a field has been set.

### GetDisabledPayMethods

`func (o *CreateOrderRequest) GetDisabledPayMethods() string`

GetDisabledPayMethods returns the DisabledPayMethods field if non-nil, zero value otherwise.

### GetDisabledPayMethodsOk

`func (o *CreateOrderRequest) GetDisabledPayMethodsOk() (*string, bool)`

GetDisabledPayMethodsOk returns a tuple with the DisabledPayMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledPayMethods

`func (o *CreateOrderRequest) SetDisabledPayMethods(v string)`

SetDisabledPayMethods sets DisabledPayMethods field to given value.

### HasDisabledPayMethods

`func (o *CreateOrderRequest) HasDisabledPayMethods() bool`

HasDisabledPayMethods returns a boolean if a field has been set.

### GetUrlParams

`func (o *CreateOrderRequest) GetUrlParams() []UrlParam`

GetUrlParams returns the UrlParams field if non-nil, zero value otherwise.

### GetUrlParamsOk

`func (o *CreateOrderRequest) GetUrlParamsOk() (*[]UrlParam, bool)`

GetUrlParamsOk returns a tuple with the UrlParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlParams

`func (o *CreateOrderRequest) SetUrlParams(v []UrlParam)`

SetUrlParams sets UrlParams field to given value.


### GetPayOptionDetails

`func (o *CreateOrderRequest) GetPayOptionDetails() []PayOptionDetail`

GetPayOptionDetails returns the PayOptionDetails field if non-nil, zero value otherwise.

### GetPayOptionDetailsOk

`func (o *CreateOrderRequest) GetPayOptionDetailsOk() (*[]PayOptionDetail, bool)`

GetPayOptionDetailsOk returns a tuple with the PayOptionDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayOptionDetails

`func (o *CreateOrderRequest) SetPayOptionDetails(v []PayOptionDetail)`

SetPayOptionDetails sets PayOptionDetails field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


