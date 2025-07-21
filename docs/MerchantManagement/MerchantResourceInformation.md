# MerchantResourceInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceType** | **string** | Type of merchant resource | 
**Value** | **string** | JSON string containing amount and currency information | 

## Methods

### NewMerchantResourceInformation

`func NewMerchantResourceInformation(resourceType string, value string, ) *MerchantResourceInformation`

NewMerchantResourceInformation instantiates a new MerchantResourceInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMerchantResourceInformationWithDefaults

`func NewMerchantResourceInformationWithDefaults() *MerchantResourceInformation`

NewMerchantResourceInformationWithDefaults instantiates a new MerchantResourceInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceType

`func (o *MerchantResourceInformation) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *MerchantResourceInformation) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *MerchantResourceInformation) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.


### GetValue

`func (o *MerchantResourceInformation) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *MerchantResourceInformation) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *MerchantResourceInformation) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


