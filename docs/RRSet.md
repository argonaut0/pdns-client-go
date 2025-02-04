# RRSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name for record set (e.g. “www.powerdns.com.”) | 
**Type** | **string** | Type of this record (e.g. “A”, “PTR”, “MX”) | 
**Ttl** | **int32** | DNS TTL of the records, in seconds. MUST NOT be included when changetype is set to “DELETE”. | 
**Changetype** | **string** | MUST be added when updating the RRSet. Must be REPLACE or DELETE. With DELETE, all existing RRs matching name and type will be deleted, including all comments. With REPLACE: when records is present, all existing RRs matching name and type will be deleted, and then new records given in records will be created. If no records are left, any existing comments will be deleted as well. When comments is present, all existing comments for the RRs matching name and type will be deleted, and then new comments given in comments will be created. | 
**Records** | [**[]Record**](Record.md) | All records in this RRSet. When updating Records, this is the list of new records (replacing the old ones). Must be empty when changetype is set to DELETE. An empty list results in deletion of all records (and comments). | 
**Comments** | Pointer to [**[]Comment**](Comment.md) | List of Comment. Must be empty when changetype is set to DELETE. An empty list results in deletion of all comments. modified_at is optional and defaults to the current server time. | [optional] 

## Methods

### NewRRSet

`func NewRRSet(name string, type_ string, ttl int32, changetype string, records []Record, ) *RRSet`

NewRRSet instantiates a new RRSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRRSetWithDefaults

`func NewRRSetWithDefaults() *RRSet`

NewRRSetWithDefaults instantiates a new RRSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RRSet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RRSet) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RRSet) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *RRSet) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RRSet) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RRSet) SetType(v string)`

SetType sets Type field to given value.


### GetTtl

`func (o *RRSet) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *RRSet) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *RRSet) SetTtl(v int32)`

SetTtl sets Ttl field to given value.


### GetChangetype

`func (o *RRSet) GetChangetype() string`

GetChangetype returns the Changetype field if non-nil, zero value otherwise.

### GetChangetypeOk

`func (o *RRSet) GetChangetypeOk() (*string, bool)`

GetChangetypeOk returns a tuple with the Changetype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangetype

`func (o *RRSet) SetChangetype(v string)`

SetChangetype sets Changetype field to given value.


### GetRecords

`func (o *RRSet) GetRecords() []Record`

GetRecords returns the Records field if non-nil, zero value otherwise.

### GetRecordsOk

`func (o *RRSet) GetRecordsOk() (*[]Record, bool)`

GetRecordsOk returns a tuple with the Records field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecords

`func (o *RRSet) SetRecords(v []Record)`

SetRecords sets Records field to given value.


### GetComments

`func (o *RRSet) GetComments() []Comment`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *RRSet) GetCommentsOk() (*[]Comment, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *RRSet) SetComments(v []Comment)`

SetComments sets Comments field to given value.

### HasComments

`func (o *RRSet) HasComments() bool`

HasComments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


