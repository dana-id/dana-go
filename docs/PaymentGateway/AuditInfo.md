# AuditInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionReason** | Pointer to **string** | Action trigger reason | [optional] 
**ThirdClientId** | Pointer to **string** | Third party client identifier | [optional] 

## Methods

### NewAuditInfo

`func NewAuditInfo() *AuditInfo`

NewAuditInfo instantiates a new AuditInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditInfoWithDefaults

`func NewAuditInfoWithDefaults() *AuditInfo`

NewAuditInfoWithDefaults instantiates a new AuditInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionReason

`func (o *AuditInfo) GetActionReason() string`

GetActionReason returns the ActionReason field if non-nil, zero value otherwise.

### GetActionReasonOk

`func (o *AuditInfo) GetActionReasonOk() (*string, bool)`

GetActionReasonOk returns a tuple with the ActionReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionReason

`func (o *AuditInfo) SetActionReason(v string)`

SetActionReason sets ActionReason field to given value.

### HasActionReason

`func (o *AuditInfo) HasActionReason() bool`

HasActionReason returns a boolean if a field has been set.

### GetThirdClientId

`func (o *AuditInfo) GetThirdClientId() string`

GetThirdClientId returns the ThirdClientId field if non-nil, zero value otherwise.

### GetThirdClientIdOk

`func (o *AuditInfo) GetThirdClientIdOk() (*string, bool)`

GetThirdClientIdOk returns a tuple with the ThirdClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThirdClientId

`func (o *AuditInfo) SetThirdClientId(v string)`

SetThirdClientId sets ThirdClientId field to given value.

### HasThirdClientId

`func (o *AuditInfo) HasThirdClientId() bool`

HasThirdClientId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


