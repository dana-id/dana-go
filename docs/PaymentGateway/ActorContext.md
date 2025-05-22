# ActorContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActorId** | Pointer to **string** | Actor identifier | [optional] 
**ActorType** | Pointer to **string** | Actor type. The enums:<br /> * USER - User<br /> * MERCHANT - Merchant&lt;br * MERCHANT_OPERATOR - Merchant operator<br /> * BACK_OFFICE - Back office<br /> * SYSTEM - System<br />  | [optional] 

## Methods

### NewActorContext

`func NewActorContext() *ActorContext`

NewActorContext instantiates a new ActorContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActorContextWithDefaults

`func NewActorContextWithDefaults() *ActorContext`

NewActorContextWithDefaults instantiates a new ActorContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActorId

`func (o *ActorContext) GetActorId() string`

GetActorId returns the ActorId field if non-nil, zero value otherwise.

### GetActorIdOk

`func (o *ActorContext) GetActorIdOk() (*string, bool)`

GetActorIdOk returns a tuple with the ActorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorId

`func (o *ActorContext) SetActorId(v string)`

SetActorId sets ActorId field to given value.

### HasActorId

`func (o *ActorContext) HasActorId() bool`

HasActorId returns a boolean if a field has been set.

### GetActorType

`func (o *ActorContext) GetActorType() string`

GetActorType returns the ActorType field if non-nil, zero value otherwise.

### GetActorTypeOk

`func (o *ActorContext) GetActorTypeOk() (*string, bool)`

GetActorTypeOk returns a tuple with the ActorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorType

`func (o *ActorContext) SetActorType(v string)`

SetActorType sets ActorType field to given value.

### HasActorType

`func (o *ActorContext) HasActorType() bool`

HasActorType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


