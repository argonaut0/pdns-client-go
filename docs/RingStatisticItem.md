# RingStatisticItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Item name | [optional] 
**Type** | Pointer to **string** | Set to \&quot;RingStatisticItem\&quot; | [optional] 
**Size** | Pointer to **int32** | Ring size | [optional] 
**Value** | Pointer to [**[]SimpleStatisticItem**](SimpleStatisticItem.md) | Named values | [optional] 

## Methods

### NewRingStatisticItem

`func NewRingStatisticItem() *RingStatisticItem`

NewRingStatisticItem instantiates a new RingStatisticItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRingStatisticItemWithDefaults

`func NewRingStatisticItemWithDefaults() *RingStatisticItem`

NewRingStatisticItemWithDefaults instantiates a new RingStatisticItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RingStatisticItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RingStatisticItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RingStatisticItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RingStatisticItem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *RingStatisticItem) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RingStatisticItem) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RingStatisticItem) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *RingStatisticItem) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSize

`func (o *RingStatisticItem) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *RingStatisticItem) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *RingStatisticItem) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *RingStatisticItem) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetValue

`func (o *RingStatisticItem) GetValue() []SimpleStatisticItem`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *RingStatisticItem) GetValueOk() (*[]SimpleStatisticItem, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *RingStatisticItem) SetValue(v []SimpleStatisticItem)`

SetValue sets Value field to given value.

### HasValue

`func (o *RingStatisticItem) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


