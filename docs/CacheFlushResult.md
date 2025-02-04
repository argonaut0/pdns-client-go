# CacheFlushResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **float32** | Amount of entries flushed | [optional] 
**Result** | Pointer to **string** | A message about the result like \&quot;Flushed cache\&quot; | [optional] 

## Methods

### NewCacheFlushResult

`func NewCacheFlushResult() *CacheFlushResult`

NewCacheFlushResult instantiates a new CacheFlushResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCacheFlushResultWithDefaults

`func NewCacheFlushResultWithDefaults() *CacheFlushResult`

NewCacheFlushResultWithDefaults instantiates a new CacheFlushResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *CacheFlushResult) GetCount() float32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *CacheFlushResult) GetCountOk() (*float32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *CacheFlushResult) SetCount(v float32)`

SetCount sets Count field to given value.

### HasCount

`func (o *CacheFlushResult) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetResult

`func (o *CacheFlushResult) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *CacheFlushResult) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *CacheFlushResult) SetResult(v string)`

SetResult sets Result field to given value.

### HasResult

`func (o *CacheFlushResult) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


