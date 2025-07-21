# BalanceInquiryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponseCode** | **string** | Refer to response code list | 
**ResponseMessage** | **string** | Refer to response code list | 
**ReferenceNo** | Pointer to **string** | Transaction identifier on DANA system | [optional] 
**PartnerReferenceNo** | Pointer to **string** | Unique transaction identifier on partner system which assigned to each transaction<br /> Notes:<br /> If the partner receives a timeout or an unexpected response from DANA and partner expects to perform retry request to DANA, please use the partnerReferenceNo that is the same as the one used in the transaction request process before  | [optional] 
**Name** | Pointer to **string** | Customer account name | [optional] 
**AccountInfos** | Pointer to [**[]AccountInfo**](AccountInfo.md) | Account information | [optional] 
**AdditionalInfo** | Pointer to **map[string]interface{}** | Additional information | [optional] 

## Methods

### NewBalanceInquiryResponse

`func NewBalanceInquiryResponse(responseCode string, responseMessage string, ) *BalanceInquiryResponse`

NewBalanceInquiryResponse instantiates a new BalanceInquiryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBalanceInquiryResponseWithDefaults

`func NewBalanceInquiryResponseWithDefaults() *BalanceInquiryResponse`

NewBalanceInquiryResponseWithDefaults instantiates a new BalanceInquiryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponseCode

`func (o *BalanceInquiryResponse) GetResponseCode() string`

GetResponseCode returns the ResponseCode field if non-nil, zero value otherwise.

### GetResponseCodeOk

`func (o *BalanceInquiryResponse) GetResponseCodeOk() (*string, bool)`

GetResponseCodeOk returns a tuple with the ResponseCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseCode

`func (o *BalanceInquiryResponse) SetResponseCode(v string)`

SetResponseCode sets ResponseCode field to given value.


### GetResponseMessage

`func (o *BalanceInquiryResponse) GetResponseMessage() string`

GetResponseMessage returns the ResponseMessage field if non-nil, zero value otherwise.

### GetResponseMessageOk

`func (o *BalanceInquiryResponse) GetResponseMessageOk() (*string, bool)`

GetResponseMessageOk returns a tuple with the ResponseMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseMessage

`func (o *BalanceInquiryResponse) SetResponseMessage(v string)`

SetResponseMessage sets ResponseMessage field to given value.


### GetReferenceNo

`func (o *BalanceInquiryResponse) GetReferenceNo() string`

GetReferenceNo returns the ReferenceNo field if non-nil, zero value otherwise.

### GetReferenceNoOk

`func (o *BalanceInquiryResponse) GetReferenceNoOk() (*string, bool)`

GetReferenceNoOk returns a tuple with the ReferenceNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceNo

`func (o *BalanceInquiryResponse) SetReferenceNo(v string)`

SetReferenceNo sets ReferenceNo field to given value.

### HasReferenceNo

`func (o *BalanceInquiryResponse) HasReferenceNo() bool`

HasReferenceNo returns a boolean if a field has been set.

### GetPartnerReferenceNo

`func (o *BalanceInquiryResponse) GetPartnerReferenceNo() string`

GetPartnerReferenceNo returns the PartnerReferenceNo field if non-nil, zero value otherwise.

### GetPartnerReferenceNoOk

`func (o *BalanceInquiryResponse) GetPartnerReferenceNoOk() (*string, bool)`

GetPartnerReferenceNoOk returns a tuple with the PartnerReferenceNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerReferenceNo

`func (o *BalanceInquiryResponse) SetPartnerReferenceNo(v string)`

SetPartnerReferenceNo sets PartnerReferenceNo field to given value.

### HasPartnerReferenceNo

`func (o *BalanceInquiryResponse) HasPartnerReferenceNo() bool`

HasPartnerReferenceNo returns a boolean if a field has been set.

### GetName

`func (o *BalanceInquiryResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BalanceInquiryResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BalanceInquiryResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BalanceInquiryResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAccountInfos

`func (o *BalanceInquiryResponse) GetAccountInfos() []AccountInfo`

GetAccountInfos returns the AccountInfos field if non-nil, zero value otherwise.

### GetAccountInfosOk

`func (o *BalanceInquiryResponse) GetAccountInfosOk() (*[]AccountInfo, bool)`

GetAccountInfosOk returns a tuple with the AccountInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountInfos

`func (o *BalanceInquiryResponse) SetAccountInfos(v []AccountInfo)`

SetAccountInfos sets AccountInfos field to given value.

### HasAccountInfos

`func (o *BalanceInquiryResponse) HasAccountInfos() bool`

HasAccountInfos returns a boolean if a field has been set.

### GetAdditionalInfo

`func (o *BalanceInquiryResponse) GetAdditionalInfo() map[string]interface{}`

GetAdditionalInfo returns the AdditionalInfo field if non-nil, zero value otherwise.

### GetAdditionalInfoOk

`func (o *BalanceInquiryResponse) GetAdditionalInfoOk() (*map[string]interface{}, bool)`

GetAdditionalInfoOk returns a tuple with the AdditionalInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalInfo

`func (o *BalanceInquiryResponse) SetAdditionalInfo(v map[string]interface{})`

SetAdditionalInfo sets AdditionalInfo field to given value.

### HasAdditionalInfo

`func (o *BalanceInquiryResponse) HasAdditionalInfo() bool`

HasAdditionalInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


