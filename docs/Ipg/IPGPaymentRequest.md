# IPGPaymentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PartnerReferenceNo** | **string** | Unique transaction identifier on partner system which assigned to each transaction | 
**MerchantId** | **string** | Merchant identifier that is unique per each merchant | 
**SubMerchantId** | Pointer to **string** |  | [optional] 
**Amount** | [**Money**](Money.md) |  | 
**ExternalStoreId** | Pointer to **string** | Store identifier to indicate to which store this payment belongs to | [optional] 
**ValidUpTo** | Pointer to **string** | The time when the payment will be automatically expired, in format YYYY-MM-DDTHH:mm:ss+07:00. Time must be in GMT+7 (Jakarta time) | [optional] 
**PointOfInitiation** | Pointer to **string** | Used for getting more info regarding source of request of the user | [optional] 
**DisabledPayMethods** | Pointer to **string** | Payment method(s) that cannot be used for this payment | [optional] 
**PayOptionDetails** | Pointer to [**[]PayOptionDetail**](PayOptionDetail.md) | Payment option that will be used for this payment | [optional] 
**AdditionalInfo** | [**IPGPaymentRequestAdditionalInfo**](IPGPaymentRequestAdditionalInfo.md) |  | 
**UrlParams** | Pointer to [**[]UrlParam**](UrlParam.md) | Notify URL that DANA must send the payment notification to | [optional] 

## Methods

### NewIPGPaymentRequest

`func NewIPGPaymentRequest(partnerReferenceNo string, merchantId string, amount Money, additionalInfo IPGPaymentRequestAdditionalInfo, ) *IPGPaymentRequest`

NewIPGPaymentRequest instantiates a new IPGPaymentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIPGPaymentRequestWithDefaults

`func NewIPGPaymentRequestWithDefaults() *IPGPaymentRequest`

NewIPGPaymentRequestWithDefaults instantiates a new IPGPaymentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPartnerReferenceNo

`func (o *IPGPaymentRequest) GetPartnerReferenceNo() string`

GetPartnerReferenceNo returns the PartnerReferenceNo field if non-nil, zero value otherwise.

### GetPartnerReferenceNoOk

`func (o *IPGPaymentRequest) GetPartnerReferenceNoOk() (*string, bool)`

GetPartnerReferenceNoOk returns a tuple with the PartnerReferenceNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerReferenceNo

`func (o *IPGPaymentRequest) SetPartnerReferenceNo(v string)`

SetPartnerReferenceNo sets PartnerReferenceNo field to given value.


### GetMerchantId

`func (o *IPGPaymentRequest) GetMerchantId() string`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *IPGPaymentRequest) GetMerchantIdOk() (*string, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *IPGPaymentRequest) SetMerchantId(v string)`

SetMerchantId sets MerchantId field to given value.


### GetSubMerchantId

`func (o *IPGPaymentRequest) GetSubMerchantId() string`

GetSubMerchantId returns the SubMerchantId field if non-nil, zero value otherwise.

### GetSubMerchantIdOk

`func (o *IPGPaymentRequest) GetSubMerchantIdOk() (*string, bool)`

GetSubMerchantIdOk returns a tuple with the SubMerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubMerchantId

`func (o *IPGPaymentRequest) SetSubMerchantId(v string)`

SetSubMerchantId sets SubMerchantId field to given value.

### HasSubMerchantId

`func (o *IPGPaymentRequest) HasSubMerchantId() bool`

HasSubMerchantId returns a boolean if a field has been set.

### GetAmount

`func (o *IPGPaymentRequest) GetAmount() Money`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *IPGPaymentRequest) GetAmountOk() (*Money, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *IPGPaymentRequest) SetAmount(v Money)`

SetAmount sets Amount field to given value.


### GetExternalStoreId

`func (o *IPGPaymentRequest) GetExternalStoreId() string`

GetExternalStoreId returns the ExternalStoreId field if non-nil, zero value otherwise.

### GetExternalStoreIdOk

`func (o *IPGPaymentRequest) GetExternalStoreIdOk() (*string, bool)`

GetExternalStoreIdOk returns a tuple with the ExternalStoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalStoreId

`func (o *IPGPaymentRequest) SetExternalStoreId(v string)`

SetExternalStoreId sets ExternalStoreId field to given value.

### HasExternalStoreId

`func (o *IPGPaymentRequest) HasExternalStoreId() bool`

HasExternalStoreId returns a boolean if a field has been set.

### GetValidUpTo

`func (o *IPGPaymentRequest) GetValidUpTo() string`

GetValidUpTo returns the ValidUpTo field if non-nil, zero value otherwise.

### GetValidUpToOk

`func (o *IPGPaymentRequest) GetValidUpToOk() (*string, bool)`

GetValidUpToOk returns a tuple with the ValidUpTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidUpTo

`func (o *IPGPaymentRequest) SetValidUpTo(v string)`

SetValidUpTo sets ValidUpTo field to given value.

### HasValidUpTo

`func (o *IPGPaymentRequest) HasValidUpTo() bool`

HasValidUpTo returns a boolean if a field has been set.

### GetPointOfInitiation

`func (o *IPGPaymentRequest) GetPointOfInitiation() string`

GetPointOfInitiation returns the PointOfInitiation field if non-nil, zero value otherwise.

### GetPointOfInitiationOk

`func (o *IPGPaymentRequest) GetPointOfInitiationOk() (*string, bool)`

GetPointOfInitiationOk returns a tuple with the PointOfInitiation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPointOfInitiation

`func (o *IPGPaymentRequest) SetPointOfInitiation(v string)`

SetPointOfInitiation sets PointOfInitiation field to given value.

### HasPointOfInitiation

`func (o *IPGPaymentRequest) HasPointOfInitiation() bool`

HasPointOfInitiation returns a boolean if a field has been set.

### GetDisabledPayMethods

`func (o *IPGPaymentRequest) GetDisabledPayMethods() string`

GetDisabledPayMethods returns the DisabledPayMethods field if non-nil, zero value otherwise.

### GetDisabledPayMethodsOk

`func (o *IPGPaymentRequest) GetDisabledPayMethodsOk() (*string, bool)`

GetDisabledPayMethodsOk returns a tuple with the DisabledPayMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledPayMethods

`func (o *IPGPaymentRequest) SetDisabledPayMethods(v string)`

SetDisabledPayMethods sets DisabledPayMethods field to given value.

### HasDisabledPayMethods

`func (o *IPGPaymentRequest) HasDisabledPayMethods() bool`

HasDisabledPayMethods returns a boolean if a field has been set.

### GetPayOptionDetails

`func (o *IPGPaymentRequest) GetPayOptionDetails() []PayOptionDetail`

GetPayOptionDetails returns the PayOptionDetails field if non-nil, zero value otherwise.

### GetPayOptionDetailsOk

`func (o *IPGPaymentRequest) GetPayOptionDetailsOk() (*[]PayOptionDetail, bool)`

GetPayOptionDetailsOk returns a tuple with the PayOptionDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayOptionDetails

`func (o *IPGPaymentRequest) SetPayOptionDetails(v []PayOptionDetail)`

SetPayOptionDetails sets PayOptionDetails field to given value.

### HasPayOptionDetails

`func (o *IPGPaymentRequest) HasPayOptionDetails() bool`

HasPayOptionDetails returns a boolean if a field has been set.

### GetAdditionalInfo

`func (o *IPGPaymentRequest) GetAdditionalInfo() IPGPaymentRequestAdditionalInfo`

GetAdditionalInfo returns the AdditionalInfo field if non-nil, zero value otherwise.

### GetAdditionalInfoOk

`func (o *IPGPaymentRequest) GetAdditionalInfoOk() (*IPGPaymentRequestAdditionalInfo, bool)`

GetAdditionalInfoOk returns a tuple with the AdditionalInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalInfo

`func (o *IPGPaymentRequest) SetAdditionalInfo(v IPGPaymentRequestAdditionalInfo)`

SetAdditionalInfo sets AdditionalInfo field to given value.


### GetUrlParams

`func (o *IPGPaymentRequest) GetUrlParams() []UrlParam`

GetUrlParams returns the UrlParams field if non-nil, zero value otherwise.

### GetUrlParamsOk

`func (o *IPGPaymentRequest) GetUrlParamsOk() (*[]UrlParam, bool)`

GetUrlParamsOk returns a tuple with the UrlParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlParams

`func (o *IPGPaymentRequest) SetUrlParams(v []UrlParam)`

SetUrlParams sets UrlParams field to given value.

### HasUrlParams

`func (o *IPGPaymentRequest) HasUrlParams() bool`

HasUrlParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


