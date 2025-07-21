# BusinessDocs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DocType** | Pointer to **string** | Document type. \&quot;individu\&quot; entity can only use KTP and SIM. Other entities can use SIUP and NIB | [optional] 
**DocId** | Pointer to **string** | Document ID | [optional] 
**DocFile** | Pointer to **string** | Document file encoded in base64 | [optional] 

## Methods

### NewBusinessDocs

`func NewBusinessDocs() *BusinessDocs`

NewBusinessDocs instantiates a new BusinessDocs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBusinessDocsWithDefaults

`func NewBusinessDocsWithDefaults() *BusinessDocs`

NewBusinessDocsWithDefaults instantiates a new BusinessDocs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocType

`func (o *BusinessDocs) GetDocType() string`

GetDocType returns the DocType field if non-nil, zero value otherwise.

### GetDocTypeOk

`func (o *BusinessDocs) GetDocTypeOk() (*string, bool)`

GetDocTypeOk returns a tuple with the DocType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocType

`func (o *BusinessDocs) SetDocType(v string)`

SetDocType sets DocType field to given value.

### HasDocType

`func (o *BusinessDocs) HasDocType() bool`

HasDocType returns a boolean if a field has been set.

### GetDocId

`func (o *BusinessDocs) GetDocId() string`

GetDocId returns the DocId field if non-nil, zero value otherwise.

### GetDocIdOk

`func (o *BusinessDocs) GetDocIdOk() (*string, bool)`

GetDocIdOk returns a tuple with the DocId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocId

`func (o *BusinessDocs) SetDocId(v string)`

SetDocId sets DocId field to given value.

### HasDocId

`func (o *BusinessDocs) HasDocId() bool`

HasDocId returns a boolean if a field has been set.

### GetDocFile

`func (o *BusinessDocs) GetDocFile() string`

GetDocFile returns the DocFile field if non-nil, zero value otherwise.

### GetDocFileOk

`func (o *BusinessDocs) GetDocFileOk() (*string, bool)`

GetDocFileOk returns a tuple with the DocFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocFile

`func (o *BusinessDocs) SetDocFile(v string)`

SetDocFile sets DocFile field to given value.

### HasDocFile

`func (o *BusinessDocs) HasDocFile() bool`

HasDocFile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


