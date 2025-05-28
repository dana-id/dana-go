# RefundOrderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MerchantId** | **string** | Merchant identifier that is unique per each merchant | 
**SubMerchantId** | Pointer to **string** | Information of sub merchant identifier | [optional] 
**OriginalReferenceNo** | Pointer to **string** | Original transaction identifier on DANA system | [optional] 
**OriginalPartnerReferenceNo** | **string** | Original transaction identifier on partner system | 
**OriginalExternalId** | Pointer to **string** | Original external identifier on header message | [optional] 
**OriginalCaptureNo** | Pointer to **string** | DANA&#39;s capture identifier. Use to refund the corresponding capture order | [optional] 
**PartnerRefundNo** | **string** | Reference number from merchant for the refund | 
**RefundAmount** | [**Money**](Money.md) | Refund amount. Contains two sub-fields - 1. Value (Transaction amount, including the cents) and 2. Currency (Currency code based on ISO) | 
**ExternalStoreId** | Pointer to **string** | Store identifier to indicate to which store this payment belongs to | [optional] 
**Reason** | Pointer to **string** | Refund reason | [optional] 
**AdditionalInfo** | Pointer to [**RefundOrderRequestAdditionalInfo**](RefundOrderRequestAdditionalInfo.md) |  | [optional] 

## Methods

### NewRefundOrderRequest

`func NewRefundOrderRequest(merchantId string, originalPartnerReferenceNo string, partnerRefundNo string, refundAmount Money, ) *RefundOrderRequest`

NewRefundOrderRequest instantiates a new RefundOrderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRefundOrderRequestWithDefaults

`func NewRefundOrderRequestWithDefaults() *RefundOrderRequest`

NewRefundOrderRequestWithDefaults instantiates a new RefundOrderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMerchantId

`func (o *RefundOrderRequest) GetMerchantId() string`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *RefundOrderRequest) GetMerchantIdOk() (*string, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *RefundOrderRequest) SetMerchantId(v string)`

SetMerchantId sets MerchantId field to given value.


### GetSubMerchantId

`func (o *RefundOrderRequest) GetSubMerchantId() string`

GetSubMerchantId returns the SubMerchantId field if non-nil, zero value otherwise.

### GetSubMerchantIdOk

`func (o *RefundOrderRequest) GetSubMerchantIdOk() (*string, bool)`

GetSubMerchantIdOk returns a tuple with the SubMerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubMerchantId

`func (o *RefundOrderRequest) SetSubMerchantId(v string)`

SetSubMerchantId sets SubMerchantId field to given value.

### HasSubMerchantId

`func (o *RefundOrderRequest) HasSubMerchantId() bool`

HasSubMerchantId returns a boolean if a field has been set.

### GetOriginalReferenceNo

`func (o *RefundOrderRequest) GetOriginalReferenceNo() string`

GetOriginalReferenceNo returns the OriginalReferenceNo field if non-nil, zero value otherwise.

### GetOriginalReferenceNoOk

`func (o *RefundOrderRequest) GetOriginalReferenceNoOk() (*string, bool)`

GetOriginalReferenceNoOk returns a tuple with the OriginalReferenceNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalReferenceNo

`func (o *RefundOrderRequest) SetOriginalReferenceNo(v string)`

SetOriginalReferenceNo sets OriginalReferenceNo field to given value.

### HasOriginalReferenceNo

`func (o *RefundOrderRequest) HasOriginalReferenceNo() bool`

HasOriginalReferenceNo returns a boolean if a field has been set.

### GetOriginalPartnerReferenceNo

`func (o *RefundOrderRequest) GetOriginalPartnerReferenceNo() string`

GetOriginalPartnerReferenceNo returns the OriginalPartnerReferenceNo field if non-nil, zero value otherwise.

### GetOriginalPartnerReferenceNoOk

`func (o *RefundOrderRequest) GetOriginalPartnerReferenceNoOk() (*string, bool)`

GetOriginalPartnerReferenceNoOk returns a tuple with the OriginalPartnerReferenceNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalPartnerReferenceNo

`func (o *RefundOrderRequest) SetOriginalPartnerReferenceNo(v string)`

SetOriginalPartnerReferenceNo sets OriginalPartnerReferenceNo field to given value.


### GetOriginalExternalId

`func (o *RefundOrderRequest) GetOriginalExternalId() string`

GetOriginalExternalId returns the OriginalExternalId field if non-nil, zero value otherwise.

### GetOriginalExternalIdOk

`func (o *RefundOrderRequest) GetOriginalExternalIdOk() (*string, bool)`

GetOriginalExternalIdOk returns a tuple with the OriginalExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalExternalId

`func (o *RefundOrderRequest) SetOriginalExternalId(v string)`

SetOriginalExternalId sets OriginalExternalId field to given value.

### HasOriginalExternalId

`func (o *RefundOrderRequest) HasOriginalExternalId() bool`

HasOriginalExternalId returns a boolean if a field has been set.

### GetOriginalCaptureNo

`func (o *RefundOrderRequest) GetOriginalCaptureNo() string`

GetOriginalCaptureNo returns the OriginalCaptureNo field if non-nil, zero value otherwise.

### GetOriginalCaptureNoOk

`func (o *RefundOrderRequest) GetOriginalCaptureNoOk() (*string, bool)`

GetOriginalCaptureNoOk returns a tuple with the OriginalCaptureNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalCaptureNo

`func (o *RefundOrderRequest) SetOriginalCaptureNo(v string)`

SetOriginalCaptureNo sets OriginalCaptureNo field to given value.

### HasOriginalCaptureNo

`func (o *RefundOrderRequest) HasOriginalCaptureNo() bool`

HasOriginalCaptureNo returns a boolean if a field has been set.

### GetPartnerRefundNo

`func (o *RefundOrderRequest) GetPartnerRefundNo() string`

GetPartnerRefundNo returns the PartnerRefundNo field if non-nil, zero value otherwise.

### GetPartnerRefundNoOk

`func (o *RefundOrderRequest) GetPartnerRefundNoOk() (*string, bool)`

GetPartnerRefundNoOk returns a tuple with the PartnerRefundNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerRefundNo

`func (o *RefundOrderRequest) SetPartnerRefundNo(v string)`

SetPartnerRefundNo sets PartnerRefundNo field to given value.


### GetRefundAmount

`func (o *RefundOrderRequest) GetRefundAmount() Money`

GetRefundAmount returns the RefundAmount field if non-nil, zero value otherwise.

### GetRefundAmountOk

`func (o *RefundOrderRequest) GetRefundAmountOk() (*Money, bool)`

GetRefundAmountOk returns a tuple with the RefundAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundAmount

`func (o *RefundOrderRequest) SetRefundAmount(v Money)`

SetRefundAmount sets RefundAmount field to given value.


### GetExternalStoreId

`func (o *RefundOrderRequest) GetExternalStoreId() string`

GetExternalStoreId returns the ExternalStoreId field if non-nil, zero value otherwise.

### GetExternalStoreIdOk

`func (o *RefundOrderRequest) GetExternalStoreIdOk() (*string, bool)`

GetExternalStoreIdOk returns a tuple with the ExternalStoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalStoreId

`func (o *RefundOrderRequest) SetExternalStoreId(v string)`

SetExternalStoreId sets ExternalStoreId field to given value.

### HasExternalStoreId

`func (o *RefundOrderRequest) HasExternalStoreId() bool`

HasExternalStoreId returns a boolean if a field has been set.

### GetReason

`func (o *RefundOrderRequest) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *RefundOrderRequest) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *RefundOrderRequest) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *RefundOrderRequest) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetAdditionalInfo

`func (o *RefundOrderRequest) GetAdditionalInfo() RefundOrderRequestAdditionalInfo`

GetAdditionalInfo returns the AdditionalInfo field if non-nil, zero value otherwise.

### GetAdditionalInfoOk

`func (o *RefundOrderRequest) GetAdditionalInfoOk() (*RefundOrderRequestAdditionalInfo, bool)`

GetAdditionalInfoOk returns a tuple with the AdditionalInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalInfo

`func (o *RefundOrderRequest) SetAdditionalInfo(v RefundOrderRequestAdditionalInfo)`

SetAdditionalInfo sets AdditionalInfo field to given value.

### HasAdditionalInfo

`func (o *RefundOrderRequest) HasAdditionalInfo() bool`

HasAdditionalInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


