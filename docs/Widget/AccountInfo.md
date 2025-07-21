# AccountInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BalanceType** | Pointer to **string** | Account information of balance type to specify which balance type expected to be returned. Will return all available balance type if this parameter empty | [optional] 
**Amount** | Pointer to [**Money**](Money.md) | Account information of amount which include the net active amount. Contains two sub-fields:<br /> 1. Value: Amount, including the cents<br /> 2. Currency: Currency code based on ISO  | [optional] 
**FloatAmount** | Pointer to [**Money**](Money.md) | Account information of float amount which include the inactive amount due to cut off period. Contains two sub-fields:<br /> 1. Value: Amount, including the cents<br /> 2. Currency: Currency code based on ISO  | [optional] 
**HoldAmount** | Pointer to [**Money**](Money.md) | Account information of hold amount which include the unusable amount due to certain type of transaction. Contains two sub-fields:<br /> 1. Value: Amount, including the cents<br /> 2. Currency: Currency code based on ISO  | [optional] 
**AvailableBalance** | Pointer to [**Money**](Money.md) | Account information of available balance which include the active amount that can be used for transaction. Contains two sub-fields:<br /> 1. Value: Amount, including the cents<br /> 2. Currency: Currency code based on ISO  | [optional] 
**LedgerBalance** | Pointer to [**Money**](Money.md) | Account information of ledger balance which include the starting balance for this day. Contains two sub-fields:<br /> 1. Value: Amount, including the cents<br /> 2. Currency: Currency code based on ISO  | [optional] 
**CurrentMultilateralLimit** | Pointer to [**Money**](Money.md) | Account information of current multilateral limit. Contains two sub-fields:<br /> 1. Value: Amount, including the cents<br /> 2. Currency: Currency code based on ISO  | [optional] 
**RegistrationStatusCode** | Pointer to **string** | Account information of customer registration status | [optional] 
**Status** | Pointer to **string** | Account information of status. The values include:<br /> 1 &#x3D; Active Account<br /> 2 &#x3D; Closed Account<br /> 4 &#x3D; New Account<br /> 6 &#x3D; Restricted Account<br /> 7 &#x3D; Frozen Account  | [optional] 

## Methods

### NewAccountInfo

`func NewAccountInfo() *AccountInfo`

NewAccountInfo instantiates a new AccountInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountInfoWithDefaults

`func NewAccountInfoWithDefaults() *AccountInfo`

NewAccountInfoWithDefaults instantiates a new AccountInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBalanceType

`func (o *AccountInfo) GetBalanceType() string`

GetBalanceType returns the BalanceType field if non-nil, zero value otherwise.

### GetBalanceTypeOk

`func (o *AccountInfo) GetBalanceTypeOk() (*string, bool)`

GetBalanceTypeOk returns a tuple with the BalanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceType

`func (o *AccountInfo) SetBalanceType(v string)`

SetBalanceType sets BalanceType field to given value.

### HasBalanceType

`func (o *AccountInfo) HasBalanceType() bool`

HasBalanceType returns a boolean if a field has been set.

### GetAmount

`func (o *AccountInfo) GetAmount() Money`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *AccountInfo) GetAmountOk() (*Money, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *AccountInfo) SetAmount(v Money)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *AccountInfo) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetFloatAmount

`func (o *AccountInfo) GetFloatAmount() Money`

GetFloatAmount returns the FloatAmount field if non-nil, zero value otherwise.

### GetFloatAmountOk

`func (o *AccountInfo) GetFloatAmountOk() (*Money, bool)`

GetFloatAmountOk returns a tuple with the FloatAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFloatAmount

`func (o *AccountInfo) SetFloatAmount(v Money)`

SetFloatAmount sets FloatAmount field to given value.

### HasFloatAmount

`func (o *AccountInfo) HasFloatAmount() bool`

HasFloatAmount returns a boolean if a field has been set.

### GetHoldAmount

`func (o *AccountInfo) GetHoldAmount() Money`

GetHoldAmount returns the HoldAmount field if non-nil, zero value otherwise.

### GetHoldAmountOk

`func (o *AccountInfo) GetHoldAmountOk() (*Money, bool)`

GetHoldAmountOk returns a tuple with the HoldAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldAmount

`func (o *AccountInfo) SetHoldAmount(v Money)`

SetHoldAmount sets HoldAmount field to given value.

### HasHoldAmount

`func (o *AccountInfo) HasHoldAmount() bool`

HasHoldAmount returns a boolean if a field has been set.

### GetAvailableBalance

`func (o *AccountInfo) GetAvailableBalance() Money`

GetAvailableBalance returns the AvailableBalance field if non-nil, zero value otherwise.

### GetAvailableBalanceOk

`func (o *AccountInfo) GetAvailableBalanceOk() (*Money, bool)`

GetAvailableBalanceOk returns a tuple with the AvailableBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableBalance

`func (o *AccountInfo) SetAvailableBalance(v Money)`

SetAvailableBalance sets AvailableBalance field to given value.

### HasAvailableBalance

`func (o *AccountInfo) HasAvailableBalance() bool`

HasAvailableBalance returns a boolean if a field has been set.

### GetLedgerBalance

`func (o *AccountInfo) GetLedgerBalance() Money`

GetLedgerBalance returns the LedgerBalance field if non-nil, zero value otherwise.

### GetLedgerBalanceOk

`func (o *AccountInfo) GetLedgerBalanceOk() (*Money, bool)`

GetLedgerBalanceOk returns a tuple with the LedgerBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLedgerBalance

`func (o *AccountInfo) SetLedgerBalance(v Money)`

SetLedgerBalance sets LedgerBalance field to given value.

### HasLedgerBalance

`func (o *AccountInfo) HasLedgerBalance() bool`

HasLedgerBalance returns a boolean if a field has been set.

### GetCurrentMultilateralLimit

`func (o *AccountInfo) GetCurrentMultilateralLimit() Money`

GetCurrentMultilateralLimit returns the CurrentMultilateralLimit field if non-nil, zero value otherwise.

### GetCurrentMultilateralLimitOk

`func (o *AccountInfo) GetCurrentMultilateralLimitOk() (*Money, bool)`

GetCurrentMultilateralLimitOk returns a tuple with the CurrentMultilateralLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentMultilateralLimit

`func (o *AccountInfo) SetCurrentMultilateralLimit(v Money)`

SetCurrentMultilateralLimit sets CurrentMultilateralLimit field to given value.

### HasCurrentMultilateralLimit

`func (o *AccountInfo) HasCurrentMultilateralLimit() bool`

HasCurrentMultilateralLimit returns a boolean if a field has been set.

### GetRegistrationStatusCode

`func (o *AccountInfo) GetRegistrationStatusCode() string`

GetRegistrationStatusCode returns the RegistrationStatusCode field if non-nil, zero value otherwise.

### GetRegistrationStatusCodeOk

`func (o *AccountInfo) GetRegistrationStatusCodeOk() (*string, bool)`

GetRegistrationStatusCodeOk returns a tuple with the RegistrationStatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationStatusCode

`func (o *AccountInfo) SetRegistrationStatusCode(v string)`

SetRegistrationStatusCode sets RegistrationStatusCode field to given value.

### HasRegistrationStatusCode

`func (o *AccountInfo) HasRegistrationStatusCode() bool`

HasRegistrationStatusCode returns a boolean if a field has been set.

### GetStatus

`func (o *AccountInfo) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AccountInfo) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AccountInfo) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AccountInfo) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


