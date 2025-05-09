# RefundOrderRequestAdditionalInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PayoutAccountNo** | Pointer to **string** | Additional information of payout account number. This param need to be filled if want to refund to specific payout account not that specified by DANA | [optional] 
**RefundAppliedTime** | Pointer to **string** | Additional information of refund applied time, in format YYYY-MM-DDTHH:mm:ss+07:00. Time must be in GMT+7 (Jakarta time) | [optional] 
**ActorType** | Pointer to **string** | Additional information of actor type, refer to ActorTypeEnum | [optional] 
**ReturnChargeToPayer** | Pointer to **string** | Additional information of return charge to payer | [optional] 
**Destination** | Pointer to **string** | Additional information of destination | [optional] 
**EnvInfo** | Pointer to **map[string]interface{}** | Additional information of environment | [optional] 
**AuditInfo** | Pointer to **map[string]interface{}** | Additional information of audit | [optional] 
**ActorContext** | Pointer to **map[string]interface{}** | Additional information of actor context | [optional] 
**RefundOptionBill** | Pointer to **[]map[string]interface{}** | Additional information of refund option bill | [optional] 
**ExtendInfo** | Pointer to **string** | Additional information of extend | [optional] 
**AsyncRefund** | Pointer to **string** | Additional information of async refund to determine the process of refund whether sync or async. The values is true/false | [optional] 

## Methods

### NewRefundOrderRequestAdditionalInfo

`func NewRefundOrderRequestAdditionalInfo() *RefundOrderRequestAdditionalInfo`

NewRefundOrderRequestAdditionalInfo instantiates a new RefundOrderRequestAdditionalInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRefundOrderRequestAdditionalInfoWithDefaults

`func NewRefundOrderRequestAdditionalInfoWithDefaults() *RefundOrderRequestAdditionalInfo`

NewRefundOrderRequestAdditionalInfoWithDefaults instantiates a new RefundOrderRequestAdditionalInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPayoutAccountNo

`func (o *RefundOrderRequestAdditionalInfo) GetPayoutAccountNo() string`

GetPayoutAccountNo returns the PayoutAccountNo field if non-nil, zero value otherwise.

### GetPayoutAccountNoOk

`func (o *RefundOrderRequestAdditionalInfo) GetPayoutAccountNoOk() (*string, bool)`

GetPayoutAccountNoOk returns a tuple with the PayoutAccountNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayoutAccountNo

`func (o *RefundOrderRequestAdditionalInfo) SetPayoutAccountNo(v string)`

SetPayoutAccountNo sets PayoutAccountNo field to given value.

### HasPayoutAccountNo

`func (o *RefundOrderRequestAdditionalInfo) HasPayoutAccountNo() bool`

HasPayoutAccountNo returns a boolean if a field has been set.

### GetRefundAppliedTime

`func (o *RefundOrderRequestAdditionalInfo) GetRefundAppliedTime() string`

GetRefundAppliedTime returns the RefundAppliedTime field if non-nil, zero value otherwise.

### GetRefundAppliedTimeOk

`func (o *RefundOrderRequestAdditionalInfo) GetRefundAppliedTimeOk() (*string, bool)`

GetRefundAppliedTimeOk returns a tuple with the RefundAppliedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundAppliedTime

`func (o *RefundOrderRequestAdditionalInfo) SetRefundAppliedTime(v string)`

SetRefundAppliedTime sets RefundAppliedTime field to given value.

### HasRefundAppliedTime

`func (o *RefundOrderRequestAdditionalInfo) HasRefundAppliedTime() bool`

HasRefundAppliedTime returns a boolean if a field has been set.

### GetActorType

`func (o *RefundOrderRequestAdditionalInfo) GetActorType() string`

GetActorType returns the ActorType field if non-nil, zero value otherwise.

### GetActorTypeOk

`func (o *RefundOrderRequestAdditionalInfo) GetActorTypeOk() (*string, bool)`

GetActorTypeOk returns a tuple with the ActorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorType

`func (o *RefundOrderRequestAdditionalInfo) SetActorType(v string)`

SetActorType sets ActorType field to given value.

### HasActorType

`func (o *RefundOrderRequestAdditionalInfo) HasActorType() bool`

HasActorType returns a boolean if a field has been set.

### GetReturnChargeToPayer

`func (o *RefundOrderRequestAdditionalInfo) GetReturnChargeToPayer() string`

GetReturnChargeToPayer returns the ReturnChargeToPayer field if non-nil, zero value otherwise.

### GetReturnChargeToPayerOk

`func (o *RefundOrderRequestAdditionalInfo) GetReturnChargeToPayerOk() (*string, bool)`

GetReturnChargeToPayerOk returns a tuple with the ReturnChargeToPayer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnChargeToPayer

`func (o *RefundOrderRequestAdditionalInfo) SetReturnChargeToPayer(v string)`

SetReturnChargeToPayer sets ReturnChargeToPayer field to given value.

### HasReturnChargeToPayer

`func (o *RefundOrderRequestAdditionalInfo) HasReturnChargeToPayer() bool`

HasReturnChargeToPayer returns a boolean if a field has been set.

### GetDestination

`func (o *RefundOrderRequestAdditionalInfo) GetDestination() string`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *RefundOrderRequestAdditionalInfo) GetDestinationOk() (*string, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *RefundOrderRequestAdditionalInfo) SetDestination(v string)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *RefundOrderRequestAdditionalInfo) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### GetEnvInfo

`func (o *RefundOrderRequestAdditionalInfo) GetEnvInfo() map[string]interface{}`

GetEnvInfo returns the EnvInfo field if non-nil, zero value otherwise.

### GetEnvInfoOk

`func (o *RefundOrderRequestAdditionalInfo) GetEnvInfoOk() (*map[string]interface{}, bool)`

GetEnvInfoOk returns a tuple with the EnvInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvInfo

`func (o *RefundOrderRequestAdditionalInfo) SetEnvInfo(v map[string]interface{})`

SetEnvInfo sets EnvInfo field to given value.

### HasEnvInfo

`func (o *RefundOrderRequestAdditionalInfo) HasEnvInfo() bool`

HasEnvInfo returns a boolean if a field has been set.

### GetAuditInfo

`func (o *RefundOrderRequestAdditionalInfo) GetAuditInfo() map[string]interface{}`

GetAuditInfo returns the AuditInfo field if non-nil, zero value otherwise.

### GetAuditInfoOk

`func (o *RefundOrderRequestAdditionalInfo) GetAuditInfoOk() (*map[string]interface{}, bool)`

GetAuditInfoOk returns a tuple with the AuditInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditInfo

`func (o *RefundOrderRequestAdditionalInfo) SetAuditInfo(v map[string]interface{})`

SetAuditInfo sets AuditInfo field to given value.

### HasAuditInfo

`func (o *RefundOrderRequestAdditionalInfo) HasAuditInfo() bool`

HasAuditInfo returns a boolean if a field has been set.

### GetActorContext

`func (o *RefundOrderRequestAdditionalInfo) GetActorContext() map[string]interface{}`

GetActorContext returns the ActorContext field if non-nil, zero value otherwise.

### GetActorContextOk

`func (o *RefundOrderRequestAdditionalInfo) GetActorContextOk() (*map[string]interface{}, bool)`

GetActorContextOk returns a tuple with the ActorContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorContext

`func (o *RefundOrderRequestAdditionalInfo) SetActorContext(v map[string]interface{})`

SetActorContext sets ActorContext field to given value.

### HasActorContext

`func (o *RefundOrderRequestAdditionalInfo) HasActorContext() bool`

HasActorContext returns a boolean if a field has been set.

### GetRefundOptionBill

`func (o *RefundOrderRequestAdditionalInfo) GetRefundOptionBill() []map[string]interface{}`

GetRefundOptionBill returns the RefundOptionBill field if non-nil, zero value otherwise.

### GetRefundOptionBillOk

`func (o *RefundOrderRequestAdditionalInfo) GetRefundOptionBillOk() (*[]map[string]interface{}, bool)`

GetRefundOptionBillOk returns a tuple with the RefundOptionBill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundOptionBill

`func (o *RefundOrderRequestAdditionalInfo) SetRefundOptionBill(v []map[string]interface{})`

SetRefundOptionBill sets RefundOptionBill field to given value.

### HasRefundOptionBill

`func (o *RefundOrderRequestAdditionalInfo) HasRefundOptionBill() bool`

HasRefundOptionBill returns a boolean if a field has been set.

### GetExtendInfo

`func (o *RefundOrderRequestAdditionalInfo) GetExtendInfo() string`

GetExtendInfo returns the ExtendInfo field if non-nil, zero value otherwise.

### GetExtendInfoOk

`func (o *RefundOrderRequestAdditionalInfo) GetExtendInfoOk() (*string, bool)`

GetExtendInfoOk returns a tuple with the ExtendInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendInfo

`func (o *RefundOrderRequestAdditionalInfo) SetExtendInfo(v string)`

SetExtendInfo sets ExtendInfo field to given value.

### HasExtendInfo

`func (o *RefundOrderRequestAdditionalInfo) HasExtendInfo() bool`

HasExtendInfo returns a boolean if a field has been set.

### GetAsyncRefund

`func (o *RefundOrderRequestAdditionalInfo) GetAsyncRefund() string`

GetAsyncRefund returns the AsyncRefund field if non-nil, zero value otherwise.

### GetAsyncRefundOk

`func (o *RefundOrderRequestAdditionalInfo) GetAsyncRefundOk() (*string, bool)`

GetAsyncRefundOk returns a tuple with the AsyncRefund field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsyncRefund

`func (o *RefundOrderRequestAdditionalInfo) SetAsyncRefund(v string)`

SetAsyncRefund sets AsyncRefund field to given value.

### HasAsyncRefund

`func (o *RefundOrderRequestAdditionalInfo) HasAsyncRefund() bool`

HasAsyncRefund returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


