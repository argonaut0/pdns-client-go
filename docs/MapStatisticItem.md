# MapStatisticItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Item name | [optional] 
**Type** | Pointer to **string** | Set to \&quot;MapStatisticItem\&quot; | [optional] 
**Value** | Pointer to [**[]SimpleStatisticItem**](SimpleStatisticItem.md) | Named values | [optional] 

## Methods

### NewMapStatisticItem

`func NewMapStatisticItem() *MapStatisticItem`

NewMapStatisticItem instantiates a new MapStatisticItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMapStatisticItemWithDefaults

`func NewMapStatisticItemWithDefaults() *MapStatisticItem`

NewMapStatisticItemWithDefaults instantiates a new MapStatisticItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *MapStatisticItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MapStatisticItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MapStatisticItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MapStatisticItem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *MapStatisticItem) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MapStatisticItem) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MapStatisticItem) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *MapStatisticItem) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValue

`func (o *MapStatisticItem) GetValue() []SimpleStatisticItem`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *MapStatisticItem) GetValueOk() (*[]SimpleStatisticItem, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *MapStatisticItem) SetValue(v []SimpleStatisticItem)`

SetValue sets Value field to given value.

### HasValue

`func (o *MapStatisticItem) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


