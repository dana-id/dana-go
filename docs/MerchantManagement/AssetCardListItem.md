# AssetCardListItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContactBizType** | **string** | Contact business type (&#x60;ContactBizTypeEnum&#x60;) | 
**CardIndexNo** | **string** | Card index number | 
**CardNoLength** | **string** | Card number length based on card index number | 
**MaskedCardNo** | **string** | Card number (masked) | 
**AssetType** | **string** | Asset / card type (&#x60;AssetCardTypeEnum&#x60;) | 
**HolderName** | **map[string]interface{}** | Holder name (JSON object per DANA spec) | 
**InstLogoUrl** | Pointer to **string** | Institution logo URL | [optional] 
**InstId** | **string** | Institution identifier | 
**InstOfficialName** | **string** | Institution official name based on &#x60;instId&#x60; | 
**ExpiryYear** | **string** | Expiry year (e.g. debit, credit, virtual account) | 
**ExpiryMonth** | **string** | Expiry month | 
**Verified** | **string** | Whether the user&#39;s card is verified | 
**BindingId** | Pointer to **string** | Asset card bind identifier | [optional] 
**DefaultAsset** | Pointer to **string** | Whether this asset is the user&#39;s default payment | [optional] 
**EnableStatus** | Pointer to **string** | Whether the card status is enabled; &#x60;\&quot;true\&quot;&#x60; when &#x60;enableOnly&#x60; in request was true | [optional] 
**DirectDebit** | Pointer to **string** | Whether payment uses direct debit | [optional] 

## Methods

### NewAssetCardListItem

`func NewAssetCardListItem(contactBizType string, cardIndexNo string, cardNoLength string, maskedCardNo string, assetType string, holderName map[string]interface{}, instId string, instOfficialName string, expiryYear string, expiryMonth string, verified string, ) *AssetCardListItem`

NewAssetCardListItem instantiates a new AssetCardListItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetCardListItemWithDefaults

`func NewAssetCardListItemWithDefaults() *AssetCardListItem`

NewAssetCardListItemWithDefaults instantiates a new AssetCardListItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContactBizType

`func (o *AssetCardListItem) GetContactBizType() string`

GetContactBizType returns the ContactBizType field if non-nil, zero value otherwise.

### GetContactBizTypeOk

`func (o *AssetCardListItem) GetContactBizTypeOk() (*string, bool)`

GetContactBizTypeOk returns a tuple with the ContactBizType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactBizType

`func (o *AssetCardListItem) SetContactBizType(v string)`

SetContactBizType sets ContactBizType field to given value.


### GetCardIndexNo

`func (o *AssetCardListItem) GetCardIndexNo() string`

GetCardIndexNo returns the CardIndexNo field if non-nil, zero value otherwise.

### GetCardIndexNoOk

`func (o *AssetCardListItem) GetCardIndexNoOk() (*string, bool)`

GetCardIndexNoOk returns a tuple with the CardIndexNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardIndexNo

`func (o *AssetCardListItem) SetCardIndexNo(v string)`

SetCardIndexNo sets CardIndexNo field to given value.


### GetCardNoLength

`func (o *AssetCardListItem) GetCardNoLength() string`

GetCardNoLength returns the CardNoLength field if non-nil, zero value otherwise.

### GetCardNoLengthOk

`func (o *AssetCardListItem) GetCardNoLengthOk() (*string, bool)`

GetCardNoLengthOk returns a tuple with the CardNoLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardNoLength

`func (o *AssetCardListItem) SetCardNoLength(v string)`

SetCardNoLength sets CardNoLength field to given value.


### GetMaskedCardNo

`func (o *AssetCardListItem) GetMaskedCardNo() string`

GetMaskedCardNo returns the MaskedCardNo field if non-nil, zero value otherwise.

### GetMaskedCardNoOk

`func (o *AssetCardListItem) GetMaskedCardNoOk() (*string, bool)`

GetMaskedCardNoOk returns a tuple with the MaskedCardNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaskedCardNo

`func (o *AssetCardListItem) SetMaskedCardNo(v string)`

SetMaskedCardNo sets MaskedCardNo field to given value.


### GetAssetType

`func (o *AssetCardListItem) GetAssetType() string`

GetAssetType returns the AssetType field if non-nil, zero value otherwise.

### GetAssetTypeOk

`func (o *AssetCardListItem) GetAssetTypeOk() (*string, bool)`

GetAssetTypeOk returns a tuple with the AssetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetType

`func (o *AssetCardListItem) SetAssetType(v string)`

SetAssetType sets AssetType field to given value.


### GetHolderName

`func (o *AssetCardListItem) GetHolderName() map[string]interface{}`

GetHolderName returns the HolderName field if non-nil, zero value otherwise.

### GetHolderNameOk

`func (o *AssetCardListItem) GetHolderNameOk() (*map[string]interface{}, bool)`

GetHolderNameOk returns a tuple with the HolderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHolderName

`func (o *AssetCardListItem) SetHolderName(v map[string]interface{})`

SetHolderName sets HolderName field to given value.


### GetInstLogoUrl

`func (o *AssetCardListItem) GetInstLogoUrl() string`

GetInstLogoUrl returns the InstLogoUrl field if non-nil, zero value otherwise.

### GetInstLogoUrlOk

`func (o *AssetCardListItem) GetInstLogoUrlOk() (*string, bool)`

GetInstLogoUrlOk returns a tuple with the InstLogoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstLogoUrl

`func (o *AssetCardListItem) SetInstLogoUrl(v string)`

SetInstLogoUrl sets InstLogoUrl field to given value.

### HasInstLogoUrl

`func (o *AssetCardListItem) HasInstLogoUrl() bool`

HasInstLogoUrl returns a boolean if a field has been set.

### GetInstId

`func (o *AssetCardListItem) GetInstId() string`

GetInstId returns the InstId field if non-nil, zero value otherwise.

### GetInstIdOk

`func (o *AssetCardListItem) GetInstIdOk() (*string, bool)`

GetInstIdOk returns a tuple with the InstId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstId

`func (o *AssetCardListItem) SetInstId(v string)`

SetInstId sets InstId field to given value.


### GetInstOfficialName

`func (o *AssetCardListItem) GetInstOfficialName() string`

GetInstOfficialName returns the InstOfficialName field if non-nil, zero value otherwise.

### GetInstOfficialNameOk

`func (o *AssetCardListItem) GetInstOfficialNameOk() (*string, bool)`

GetInstOfficialNameOk returns a tuple with the InstOfficialName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstOfficialName

`func (o *AssetCardListItem) SetInstOfficialName(v string)`

SetInstOfficialName sets InstOfficialName field to given value.


### GetExpiryYear

`func (o *AssetCardListItem) GetExpiryYear() string`

GetExpiryYear returns the ExpiryYear field if non-nil, zero value otherwise.

### GetExpiryYearOk

`func (o *AssetCardListItem) GetExpiryYearOk() (*string, bool)`

GetExpiryYearOk returns a tuple with the ExpiryYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryYear

`func (o *AssetCardListItem) SetExpiryYear(v string)`

SetExpiryYear sets ExpiryYear field to given value.


### GetExpiryMonth

`func (o *AssetCardListItem) GetExpiryMonth() string`

GetExpiryMonth returns the ExpiryMonth field if non-nil, zero value otherwise.

### GetExpiryMonthOk

`func (o *AssetCardListItem) GetExpiryMonthOk() (*string, bool)`

GetExpiryMonthOk returns a tuple with the ExpiryMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryMonth

`func (o *AssetCardListItem) SetExpiryMonth(v string)`

SetExpiryMonth sets ExpiryMonth field to given value.


### GetVerified

`func (o *AssetCardListItem) GetVerified() string`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *AssetCardListItem) GetVerifiedOk() (*string, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *AssetCardListItem) SetVerified(v string)`

SetVerified sets Verified field to given value.


### GetBindingId

`func (o *AssetCardListItem) GetBindingId() string`

GetBindingId returns the BindingId field if non-nil, zero value otherwise.

### GetBindingIdOk

`func (o *AssetCardListItem) GetBindingIdOk() (*string, bool)`

GetBindingIdOk returns a tuple with the BindingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindingId

`func (o *AssetCardListItem) SetBindingId(v string)`

SetBindingId sets BindingId field to given value.

### HasBindingId

`func (o *AssetCardListItem) HasBindingId() bool`

HasBindingId returns a boolean if a field has been set.

### GetDefaultAsset

`func (o *AssetCardListItem) GetDefaultAsset() string`

GetDefaultAsset returns the DefaultAsset field if non-nil, zero value otherwise.

### GetDefaultAssetOk

`func (o *AssetCardListItem) GetDefaultAssetOk() (*string, bool)`

GetDefaultAssetOk returns a tuple with the DefaultAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultAsset

`func (o *AssetCardListItem) SetDefaultAsset(v string)`

SetDefaultAsset sets DefaultAsset field to given value.

### HasDefaultAsset

`func (o *AssetCardListItem) HasDefaultAsset() bool`

HasDefaultAsset returns a boolean if a field has been set.

### GetEnableStatus

`func (o *AssetCardListItem) GetEnableStatus() string`

GetEnableStatus returns the EnableStatus field if non-nil, zero value otherwise.

### GetEnableStatusOk

`func (o *AssetCardListItem) GetEnableStatusOk() (*string, bool)`

GetEnableStatusOk returns a tuple with the EnableStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableStatus

`func (o *AssetCardListItem) SetEnableStatus(v string)`

SetEnableStatus sets EnableStatus field to given value.

### HasEnableStatus

`func (o *AssetCardListItem) HasEnableStatus() bool`

HasEnableStatus returns a boolean if a field has been set.

### GetDirectDebit

`func (o *AssetCardListItem) GetDirectDebit() string`

GetDirectDebit returns the DirectDebit field if non-nil, zero value otherwise.

### GetDirectDebitOk

`func (o *AssetCardListItem) GetDirectDebitOk() (*string, bool)`

GetDirectDebitOk returns a tuple with the DirectDebit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectDebit

`func (o *AssetCardListItem) SetDirectDebit(v string)`

SetDirectDebit sets DirectDebit field to given value.

### HasDirectDebit

`func (o *AssetCardListItem) HasDirectDebit() bool`

HasDirectDebit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


