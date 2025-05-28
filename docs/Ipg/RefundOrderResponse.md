# RefundOrderResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponseCode** | **string** | Refer to response code list | 
**ResponseMessage** | **string** | Refer to response code list | 
**OriginalReferenceNo** | Pointer to **string** | Original transaction identifier on DANA system | [optional] 
**OriginalPartnerReferenceNo** | **string** | Original transaction identifier on partner system | 
**OriginalExternalId** | Pointer to **string** | Original external identifier on header message | [optional] 
**OriginalCaptureNo** | Pointer to **string** | DANA&#39;s capture identifier. Use to refund the corresponding capture order | [optional] 
**RefundNo** | Pointer to **string** | Refund number identifier on DANA system | [optional] 
**PartnerRefundNo** | **string** | Reference number from merchant for the refund | 
**RefundAmount** | [**Money**](Money.md) | Refund amount. Contains two sub-fields - 1. Value (Amount, including the cents) and 2. Currency (Currency code based on ISO) | 
**RefundTime** | Pointer to **string** | Refund time, in format YYYY-MM-DDTHH:mm:ss+07:00. Time must be in GMT+7 (Jakarta time) | [optional] 
**AdditionalInfo** | Pointer to **map[string]interface{}** | Additional information | [optional] 

## Methods

### NewRefundOrderResponse

`func NewRefundOrderResponse(responseCode string, responseMessage string, originalPartnerReferenceNo string, partnerRefundNo string, refundAmount Money, ) *RefundOrderResponse`

NewRefundOrderResponse instantiates a new RefundOrderResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRefundOrderResponseWithDefaults

`func NewRefundOrderResponseWithDefaults() *RefundOrderResponse`

NewRefundOrderResponseWithDefaults instantiates a new RefundOrderResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponseCode

`func (o *RefundOrderResponse) GetResponseCode() string`

GetResponseCode returns the ResponseCode field if non-nil, zero value otherwise.

### GetResponseCodeOk

`func (o *RefundOrderResponse) GetResponseCodeOk() (*string, bool)`

GetResponseCodeOk returns a tuple with the ResponseCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseCode

`func (o *RefundOrderResponse) SetResponseCode(v string)`

SetResponseCode sets ResponseCode field to given value.


### GetResponseMessage

`func (o *RefundOrderResponse) GetResponseMessage() string`

GetResponseMessage returns the ResponseMessage field if non-nil, zero value otherwise.

### GetResponseMessageOk

`func (o *RefundOrderResponse) GetResponseMessageOk() (*string, bool)`

GetResponseMessageOk returns a tuple with the ResponseMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseMessage

`func (o *RefundOrderResponse) SetResponseMessage(v string)`

SetResponseMessage sets ResponseMessage field to given value.


### GetOriginalReferenceNo

`func (o *RefundOrderResponse) GetOriginalReferenceNo() string`

GetOriginalReferenceNo returns the OriginalReferenceNo field if non-nil, zero value otherwise.

### GetOriginalReferenceNoOk

`func (o *RefundOrderResponse) GetOriginalReferenceNoOk() (*string, bool)`

GetOriginalReferenceNoOk returns a tuple with the OriginalReferenceNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalReferenceNo

`func (o *RefundOrderResponse) SetOriginalReferenceNo(v string)`

SetOriginalReferenceNo sets OriginalReferenceNo field to given value.

### HasOriginalReferenceNo

`func (o *RefundOrderResponse) HasOriginalReferenceNo() bool`

HasOriginalReferenceNo returns a boolean if a field has been set.

### GetOriginalPartnerReferenceNo

`func (o *RefundOrderResponse) GetOriginalPartnerReferenceNo() string`

GetOriginalPartnerReferenceNo returns the OriginalPartnerReferenceNo field if non-nil, zero value otherwise.

### GetOriginalPartnerReferenceNoOk

`func (o *RefundOrderResponse) GetOriginalPartnerReferenceNoOk() (*string, bool)`

GetOriginalPartnerReferenceNoOk returns a tuple with the OriginalPartnerReferenceNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalPartnerReferenceNo

`func (o *RefundOrderResponse) SetOriginalPartnerReferenceNo(v string)`

SetOriginalPartnerReferenceNo sets OriginalPartnerReferenceNo field to given value.


### GetOriginalExternalId

`func (o *RefundOrderResponse) GetOriginalExternalId() string`

GetOriginalExternalId returns the OriginalExternalId field if non-nil, zero value otherwise.

### GetOriginalExternalIdOk

`func (o *RefundOrderResponse) GetOriginalExternalIdOk() (*string, bool)`

GetOriginalExternalIdOk returns a tuple with the OriginalExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalExternalId

`func (o *RefundOrderResponse) SetOriginalExternalId(v string)`

SetOriginalExternalId sets OriginalExternalId field to given value.

### HasOriginalExternalId

`func (o *RefundOrderResponse) HasOriginalExternalId() bool`

HasOriginalExternalId returns a boolean if a field has been set.

### GetOriginalCaptureNo

`func (o *RefundOrderResponse) GetOriginalCaptureNo() string`

GetOriginalCaptureNo returns the OriginalCaptureNo field if non-nil, zero value otherwise.

### GetOriginalCaptureNoOk

`func (o *RefundOrderResponse) GetOriginalCaptureNoOk() (*string, bool)`

GetOriginalCaptureNoOk returns a tuple with the OriginalCaptureNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalCaptureNo

`func (o *RefundOrderResponse) SetOriginalCaptureNo(v string)`

SetOriginalCaptureNo sets OriginalCaptureNo field to given value.

### HasOriginalCaptureNo

`func (o *RefundOrderResponse) HasOriginalCaptureNo() bool`

HasOriginalCaptureNo returns a boolean if a field has been set.

### GetRefundNo

`func (o *RefundOrderResponse) GetRefundNo() string`

GetRefundNo returns the RefundNo field if non-nil, zero value otherwise.

### GetRefundNoOk

`func (o *RefundOrderResponse) GetRefundNoOk() (*string, bool)`

GetRefundNoOk returns a tuple with the RefundNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundNo

`func (o *RefundOrderResponse) SetRefundNo(v string)`

SetRefundNo sets RefundNo field to given value.

### HasRefundNo

`func (o *RefundOrderResponse) HasRefundNo() bool`

HasRefundNo returns a boolean if a field has been set.

### GetPartnerRefundNo

`func (o *RefundOrderResponse) GetPartnerRefundNo() string`

GetPartnerRefundNo returns the PartnerRefundNo field if non-nil, zero value otherwise.

### GetPartnerRefundNoOk

`func (o *RefundOrderResponse) GetPartnerRefundNoOk() (*string, bool)`

GetPartnerRefundNoOk returns a tuple with the PartnerRefundNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerRefundNo

`func (o *RefundOrderResponse) SetPartnerRefundNo(v string)`

SetPartnerRefundNo sets PartnerRefundNo field to given value.


### GetRefundAmount

`func (o *RefundOrderResponse) GetRefundAmount() Money`

GetRefundAmount returns the RefundAmount field if non-nil, zero value otherwise.

### GetRefundAmountOk

`func (o *RefundOrderResponse) GetRefundAmountOk() (*Money, bool)`

GetRefundAmountOk returns a tuple with the RefundAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundAmount

`func (o *RefundOrderResponse) SetRefundAmount(v Money)`

SetRefundAmount sets RefundAmount field to given value.


### GetRefundTime

`func (o *RefundOrderResponse) GetRefundTime() string`

GetRefundTime returns the RefundTime field if non-nil, zero value otherwise.

### GetRefundTimeOk

`func (o *RefundOrderResponse) GetRefundTimeOk() (*string, bool)`

GetRefundTimeOk returns a tuple with the RefundTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundTime

`func (o *RefundOrderResponse) SetRefundTime(v string)`

SetRefundTime sets RefundTime field to given value.

### HasRefundTime

`func (o *RefundOrderResponse) HasRefundTime() bool`

HasRefundTime returns a boolean if a field has been set.

### GetAdditionalInfo

`func (o *RefundOrderResponse) GetAdditionalInfo() map[string]interface{}`

GetAdditionalInfo returns the AdditionalInfo field if non-nil, zero value otherwise.

### GetAdditionalInfoOk

`func (o *RefundOrderResponse) GetAdditionalInfoOk() (*map[string]interface{}, bool)`

GetAdditionalInfoOk returns a tuple with the AdditionalInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalInfo

`func (o *RefundOrderResponse) SetAdditionalInfo(v map[string]interface{})`

SetAdditionalInfo sets AdditionalInfo field to given value.

### HasAdditionalInfo

`func (o *RefundOrderResponse) HasAdditionalInfo() bool`

HasAdditionalInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


