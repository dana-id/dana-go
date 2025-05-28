# ShippingInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MerchantShippingId** | **string** | Merchant shipping identifier | 
**TrackingNo** | Pointer to **string** | Number of tracking | [optional] 
**Carrier** | Pointer to **string** | Information of carrier | [optional] 
**ChargeAmount** | Pointer to [**Money**](Money.md) | Promo amount. Contains two sub-fields:<br /> 1. Value: Transaction amount, including the cents<br /> 2. Currency: Currency code based on ISO<br />  | [optional] 
**CountryName** | **string** | Name of country | 
**StateName** | **string** | Name of state | 
**CityName** | **string** | Name of city | 
**AreaName** | Pointer to **string** | Name of area | [optional] 
**Address1** | **string** | Information of address 1 | 
**Address2** | Pointer to **string** | Information of address 2 | [optional] 
**FirstName** | **string** | First name | 
**LastName** | **string** | Last name | 
**MobileNo** | Pointer to **string** | Mobile number | [optional] 
**PhoneNo** | Pointer to **string** | Phone number | [optional] 
**ZipCode** | **string** | Zip code | 
**Email** | Pointer to **string** | Email | [optional] 
**FaxNo** | Pointer to **string** | Fax number | [optional] 

## Methods

### NewShippingInfo

`func NewShippingInfo(merchantShippingId string, countryName string, stateName string, cityName string, address1 string, firstName string, lastName string, zipCode string, ) *ShippingInfo`

NewShippingInfo instantiates a new ShippingInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShippingInfoWithDefaults

`func NewShippingInfoWithDefaults() *ShippingInfo`

NewShippingInfoWithDefaults instantiates a new ShippingInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMerchantShippingId

`func (o *ShippingInfo) GetMerchantShippingId() string`

GetMerchantShippingId returns the MerchantShippingId field if non-nil, zero value otherwise.

### GetMerchantShippingIdOk

`func (o *ShippingInfo) GetMerchantShippingIdOk() (*string, bool)`

GetMerchantShippingIdOk returns a tuple with the MerchantShippingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantShippingId

`func (o *ShippingInfo) SetMerchantShippingId(v string)`

SetMerchantShippingId sets MerchantShippingId field to given value.


### GetTrackingNo

`func (o *ShippingInfo) GetTrackingNo() string`

GetTrackingNo returns the TrackingNo field if non-nil, zero value otherwise.

### GetTrackingNoOk

`func (o *ShippingInfo) GetTrackingNoOk() (*string, bool)`

GetTrackingNoOk returns a tuple with the TrackingNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingNo

`func (o *ShippingInfo) SetTrackingNo(v string)`

SetTrackingNo sets TrackingNo field to given value.

### HasTrackingNo

`func (o *ShippingInfo) HasTrackingNo() bool`

HasTrackingNo returns a boolean if a field has been set.

### GetCarrier

`func (o *ShippingInfo) GetCarrier() string`

GetCarrier returns the Carrier field if non-nil, zero value otherwise.

### GetCarrierOk

`func (o *ShippingInfo) GetCarrierOk() (*string, bool)`

GetCarrierOk returns a tuple with the Carrier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrier

`func (o *ShippingInfo) SetCarrier(v string)`

SetCarrier sets Carrier field to given value.

### HasCarrier

`func (o *ShippingInfo) HasCarrier() bool`

HasCarrier returns a boolean if a field has been set.

### GetChargeAmount

`func (o *ShippingInfo) GetChargeAmount() Money`

GetChargeAmount returns the ChargeAmount field if non-nil, zero value otherwise.

### GetChargeAmountOk

`func (o *ShippingInfo) GetChargeAmountOk() (*Money, bool)`

GetChargeAmountOk returns a tuple with the ChargeAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeAmount

`func (o *ShippingInfo) SetChargeAmount(v Money)`

SetChargeAmount sets ChargeAmount field to given value.

### HasChargeAmount

`func (o *ShippingInfo) HasChargeAmount() bool`

HasChargeAmount returns a boolean if a field has been set.

### GetCountryName

`func (o *ShippingInfo) GetCountryName() string`

GetCountryName returns the CountryName field if non-nil, zero value otherwise.

### GetCountryNameOk

`func (o *ShippingInfo) GetCountryNameOk() (*string, bool)`

GetCountryNameOk returns a tuple with the CountryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryName

`func (o *ShippingInfo) SetCountryName(v string)`

SetCountryName sets CountryName field to given value.


### GetStateName

`func (o *ShippingInfo) GetStateName() string`

GetStateName returns the StateName field if non-nil, zero value otherwise.

### GetStateNameOk

`func (o *ShippingInfo) GetStateNameOk() (*string, bool)`

GetStateNameOk returns a tuple with the StateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateName

`func (o *ShippingInfo) SetStateName(v string)`

SetStateName sets StateName field to given value.


### GetCityName

`func (o *ShippingInfo) GetCityName() string`

GetCityName returns the CityName field if non-nil, zero value otherwise.

### GetCityNameOk

`func (o *ShippingInfo) GetCityNameOk() (*string, bool)`

GetCityNameOk returns a tuple with the CityName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCityName

`func (o *ShippingInfo) SetCityName(v string)`

SetCityName sets CityName field to given value.


### GetAreaName

`func (o *ShippingInfo) GetAreaName() string`

GetAreaName returns the AreaName field if non-nil, zero value otherwise.

### GetAreaNameOk

`func (o *ShippingInfo) GetAreaNameOk() (*string, bool)`

GetAreaNameOk returns a tuple with the AreaName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAreaName

`func (o *ShippingInfo) SetAreaName(v string)`

SetAreaName sets AreaName field to given value.

### HasAreaName

`func (o *ShippingInfo) HasAreaName() bool`

HasAreaName returns a boolean if a field has been set.

### GetAddress1

`func (o *ShippingInfo) GetAddress1() string`

GetAddress1 returns the Address1 field if non-nil, zero value otherwise.

### GetAddress1Ok

`func (o *ShippingInfo) GetAddress1Ok() (*string, bool)`

GetAddress1Ok returns a tuple with the Address1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress1

`func (o *ShippingInfo) SetAddress1(v string)`

SetAddress1 sets Address1 field to given value.


### GetAddress2

`func (o *ShippingInfo) GetAddress2() string`

GetAddress2 returns the Address2 field if non-nil, zero value otherwise.

### GetAddress2Ok

`func (o *ShippingInfo) GetAddress2Ok() (*string, bool)`

GetAddress2Ok returns a tuple with the Address2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress2

`func (o *ShippingInfo) SetAddress2(v string)`

SetAddress2 sets Address2 field to given value.

### HasAddress2

`func (o *ShippingInfo) HasAddress2() bool`

HasAddress2 returns a boolean if a field has been set.

### GetFirstName

`func (o *ShippingInfo) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *ShippingInfo) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *ShippingInfo) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetLastName

`func (o *ShippingInfo) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *ShippingInfo) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *ShippingInfo) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetMobileNo

`func (o *ShippingInfo) GetMobileNo() string`

GetMobileNo returns the MobileNo field if non-nil, zero value otherwise.

### GetMobileNoOk

`func (o *ShippingInfo) GetMobileNoOk() (*string, bool)`

GetMobileNoOk returns a tuple with the MobileNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobileNo

`func (o *ShippingInfo) SetMobileNo(v string)`

SetMobileNo sets MobileNo field to given value.

### HasMobileNo

`func (o *ShippingInfo) HasMobileNo() bool`

HasMobileNo returns a boolean if a field has been set.

### GetPhoneNo

`func (o *ShippingInfo) GetPhoneNo() string`

GetPhoneNo returns the PhoneNo field if non-nil, zero value otherwise.

### GetPhoneNoOk

`func (o *ShippingInfo) GetPhoneNoOk() (*string, bool)`

GetPhoneNoOk returns a tuple with the PhoneNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNo

`func (o *ShippingInfo) SetPhoneNo(v string)`

SetPhoneNo sets PhoneNo field to given value.

### HasPhoneNo

`func (o *ShippingInfo) HasPhoneNo() bool`

HasPhoneNo returns a boolean if a field has been set.

### GetZipCode

`func (o *ShippingInfo) GetZipCode() string`

GetZipCode returns the ZipCode field if non-nil, zero value otherwise.

### GetZipCodeOk

`func (o *ShippingInfo) GetZipCodeOk() (*string, bool)`

GetZipCodeOk returns a tuple with the ZipCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZipCode

`func (o *ShippingInfo) SetZipCode(v string)`

SetZipCode sets ZipCode field to given value.


### GetEmail

`func (o *ShippingInfo) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ShippingInfo) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ShippingInfo) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *ShippingInfo) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetFaxNo

`func (o *ShippingInfo) GetFaxNo() string`

GetFaxNo returns the FaxNo field if non-nil, zero value otherwise.

### GetFaxNoOk

`func (o *ShippingInfo) GetFaxNoOk() (*string, bool)`

GetFaxNoOk returns a tuple with the FaxNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaxNo

`func (o *ShippingInfo) SetFaxNo(v string)`

SetFaxNo sets FaxNo field to given value.

### HasFaxNo

`func (o *ShippingInfo) HasFaxNo() bool`

HasFaxNo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


