# MemberAssetResultInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResultStatus** | **string** | Request status (&#x60;ResultStatus&#x60;). S&#x3D;Success, F&#x3D;Failure, U&#x3D;Unknown | 
**ResultCodeId** | **string** | Result code identifier (&#x60;ResultCodeId&#x60;) | 
**ResultCode** | **string** | Result code (&#x60;ResultCode&#x60;) | 
**ResultMsg** | Pointer to **string** | Result message (&#x60;ResultMsg&#x60;). Optional per API contract | [optional] 

## Methods

### NewMemberAssetResultInfo

`func NewMemberAssetResultInfo(resultStatus string, resultCodeId string, resultCode string, ) *MemberAssetResultInfo`

NewMemberAssetResultInfo instantiates a new MemberAssetResultInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemberAssetResultInfoWithDefaults

`func NewMemberAssetResultInfoWithDefaults() *MemberAssetResultInfo`

NewMemberAssetResultInfoWithDefaults instantiates a new MemberAssetResultInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResultStatus

`func (o *MemberAssetResultInfo) GetResultStatus() string`

GetResultStatus returns the ResultStatus field if non-nil, zero value otherwise.

### GetResultStatusOk

`func (o *MemberAssetResultInfo) GetResultStatusOk() (*string, bool)`

GetResultStatusOk returns a tuple with the ResultStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultStatus

`func (o *MemberAssetResultInfo) SetResultStatus(v string)`

SetResultStatus sets ResultStatus field to given value.


### GetResultCodeId

`func (o *MemberAssetResultInfo) GetResultCodeId() string`

GetResultCodeId returns the ResultCodeId field if non-nil, zero value otherwise.

### GetResultCodeIdOk

`func (o *MemberAssetResultInfo) GetResultCodeIdOk() (*string, bool)`

GetResultCodeIdOk returns a tuple with the ResultCodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultCodeId

`func (o *MemberAssetResultInfo) SetResultCodeId(v string)`

SetResultCodeId sets ResultCodeId field to given value.


### GetResultCode

`func (o *MemberAssetResultInfo) GetResultCode() string`

GetResultCode returns the ResultCode field if non-nil, zero value otherwise.

### GetResultCodeOk

`func (o *MemberAssetResultInfo) GetResultCodeOk() (*string, bool)`

GetResultCodeOk returns a tuple with the ResultCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultCode

`func (o *MemberAssetResultInfo) SetResultCode(v string)`

SetResultCode sets ResultCode field to given value.


### GetResultMsg

`func (o *MemberAssetResultInfo) GetResultMsg() string`

GetResultMsg returns the ResultMsg field if non-nil, zero value otherwise.

### GetResultMsgOk

`func (o *MemberAssetResultInfo) GetResultMsgOk() (*string, bool)`

GetResultMsgOk returns a tuple with the ResultMsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultMsg

`func (o *MemberAssetResultInfo) SetResultMsg(v string)`

SetResultMsg sets ResultMsg field to given value.

### HasResultMsg

`func (o *MemberAssetResultInfo) HasResultMsg() bool`

HasResultMsg returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


