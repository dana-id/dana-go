# IPGPaymentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponseCode** | **string** | Refer to response code list:<br /> * 2005400 - Successful<br /> * 4005400 - Bad Request - Retry request with proper parameter<br /> * 4005401 - Invalid Field Format - Retry request with proper parameter<br /> * 4005402 - Invalid Mandatory Field - Retry request with proper parameter<br /> * 4015400 - Unauthorized. [reason] - Retry request with proper parameter<br /> * 4035402 - Exceeds Transaction Amount Limit - Try to adjust the order amount<br /> * 4035405 - Do Not Honor - Retry request with proper parameter or can contact DANA to check the user/account status<br /> * 4035415 - Transaction Not Permitted - Retry request periodically or consult to DANA<br /> * 4045408 - Invalid Merchant - Retry request with proper parameter<br /> * 4045418 - Inconsistent Request - Retry with proper parameter<br /> * 4295400 - Too Many Requests - Retry request periodically by sending same request payload<br /> * 5005400 - General Error - Retry request periodically<br /> * 5005401 - Internal Server Error - Retry request periodically by sending same request payload<br />  | 
**ResponseMessage** | **string** | Human readable response message | 
**ReferenceNo** | Pointer to **string** | Transaction identifier on DANA system, returned when transaction is successfully processed | [optional] 
**PartnerReferenceNo** | **string** | Transaction identifier on partner system which assigned to each transaction | 
**WebRedirectUrl** | Pointer to **string** | DANA checkout URL, returned when transaction is successfully processed | [optional] 
**AdditionalInfo** | Pointer to **map[string]interface{}** | Additional information | [optional] 

## Methods

### NewIPGPaymentResponse

`func NewIPGPaymentResponse(responseCode string, responseMessage string, partnerReferenceNo string, ) *IPGPaymentResponse`

NewIPGPaymentResponse instantiates a new IPGPaymentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIPGPaymentResponseWithDefaults

`func NewIPGPaymentResponseWithDefaults() *IPGPaymentResponse`

NewIPGPaymentResponseWithDefaults instantiates a new IPGPaymentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponseCode

`func (o *IPGPaymentResponse) GetResponseCode() string`

GetResponseCode returns the ResponseCode field if non-nil, zero value otherwise.

### GetResponseCodeOk

`func (o *IPGPaymentResponse) GetResponseCodeOk() (*string, bool)`

GetResponseCodeOk returns a tuple with the ResponseCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseCode

`func (o *IPGPaymentResponse) SetResponseCode(v string)`

SetResponseCode sets ResponseCode field to given value.


### GetResponseMessage

`func (o *IPGPaymentResponse) GetResponseMessage() string`

GetResponseMessage returns the ResponseMessage field if non-nil, zero value otherwise.

### GetResponseMessageOk

`func (o *IPGPaymentResponse) GetResponseMessageOk() (*string, bool)`

GetResponseMessageOk returns a tuple with the ResponseMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseMessage

`func (o *IPGPaymentResponse) SetResponseMessage(v string)`

SetResponseMessage sets ResponseMessage field to given value.


### GetReferenceNo

`func (o *IPGPaymentResponse) GetReferenceNo() string`

GetReferenceNo returns the ReferenceNo field if non-nil, zero value otherwise.

### GetReferenceNoOk

`func (o *IPGPaymentResponse) GetReferenceNoOk() (*string, bool)`

GetReferenceNoOk returns a tuple with the ReferenceNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceNo

`func (o *IPGPaymentResponse) SetReferenceNo(v string)`

SetReferenceNo sets ReferenceNo field to given value.

### HasReferenceNo

`func (o *IPGPaymentResponse) HasReferenceNo() bool`

HasReferenceNo returns a boolean if a field has been set.

### GetPartnerReferenceNo

`func (o *IPGPaymentResponse) GetPartnerReferenceNo() string`

GetPartnerReferenceNo returns the PartnerReferenceNo field if non-nil, zero value otherwise.

### GetPartnerReferenceNoOk

`func (o *IPGPaymentResponse) GetPartnerReferenceNoOk() (*string, bool)`

GetPartnerReferenceNoOk returns a tuple with the PartnerReferenceNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerReferenceNo

`func (o *IPGPaymentResponse) SetPartnerReferenceNo(v string)`

SetPartnerReferenceNo sets PartnerReferenceNo field to given value.


### GetWebRedirectUrl

`func (o *IPGPaymentResponse) GetWebRedirectUrl() string`

GetWebRedirectUrl returns the WebRedirectUrl field if non-nil, zero value otherwise.

### GetWebRedirectUrlOk

`func (o *IPGPaymentResponse) GetWebRedirectUrlOk() (*string, bool)`

GetWebRedirectUrlOk returns a tuple with the WebRedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebRedirectUrl

`func (o *IPGPaymentResponse) SetWebRedirectUrl(v string)`

SetWebRedirectUrl sets WebRedirectUrl field to given value.

### HasWebRedirectUrl

`func (o *IPGPaymentResponse) HasWebRedirectUrl() bool`

HasWebRedirectUrl returns a boolean if a field has been set.

### GetAdditionalInfo

`func (o *IPGPaymentResponse) GetAdditionalInfo() map[string]interface{}`

GetAdditionalInfo returns the AdditionalInfo field if non-nil, zero value otherwise.

### GetAdditionalInfoOk

`func (o *IPGPaymentResponse) GetAdditionalInfoOk() (*map[string]interface{}, bool)`

GetAdditionalInfoOk returns a tuple with the AdditionalInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalInfo

`func (o *IPGPaymentResponse) SetAdditionalInfo(v map[string]interface{})`

SetAdditionalInfo sets AdditionalInfo field to given value.

### HasAdditionalInfo

`func (o *IPGPaymentResponse) HasAdditionalInfo() bool`

HasAdditionalInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


