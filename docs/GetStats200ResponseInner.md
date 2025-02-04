# GetStats200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Item name | [optional] 
**Type** | Pointer to **string** | Set to \&quot;RingStatisticItem\&quot; | [optional] 
**Value** | Pointer to [**[]SimpleStatisticItem**](SimpleStatisticItem.md) | Named values | [optional] 
**Size** | Pointer to **int32** | Ring size | [optional] 

## Methods

### NewGetStats200ResponseInner

`func NewGetStats200ResponseInner() *GetStats200ResponseInner`

NewGetStats200ResponseInner instantiates a new GetStats200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetStats200ResponseInnerWithDefaults

`func NewGetStats200ResponseInnerWithDefaults() *GetStats200ResponseInner`

NewGetStats200ResponseInnerWithDefaults instantiates a new GetStats200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GetStats200ResponseInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetStats200ResponseInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetStats200ResponseInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetStats200ResponseInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *GetStats200ResponseInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetStats200ResponseInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetStats200ResponseInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetStats200ResponseInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValue

`func (o *GetStats200ResponseInner) GetValue() []SimpleStatisticItem`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *GetStats200ResponseInner) GetValueOk() (*[]SimpleStatisticItem, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *GetStats200ResponseInner) SetValue(v []SimpleStatisticItem)`

SetValue sets Value field to given value.

### HasValue

`func (o *GetStats200ResponseInner) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetSize

`func (o *GetStats200ResponseInner) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *GetStats200ResponseInner) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *GetStats200ResponseInner) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *GetStats200ResponseInner) HasSize() bool`

HasSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


