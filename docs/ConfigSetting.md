# ConfigSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | set to \&quot;ConfigSetting\&quot; | [optional] 
**Type** | Pointer to **string** | The name of this setting (e.g. ‘webserver-port’) | [optional] 
**Value** | Pointer to **string** | The value of setting name | [optional] 

## Methods

### NewConfigSetting

`func NewConfigSetting() *ConfigSetting`

NewConfigSetting instantiates a new ConfigSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigSettingWithDefaults

`func NewConfigSettingWithDefaults() *ConfigSetting`

NewConfigSettingWithDefaults instantiates a new ConfigSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ConfigSetting) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConfigSetting) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConfigSetting) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ConfigSetting) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *ConfigSetting) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ConfigSetting) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ConfigSetting) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ConfigSetting) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValue

`func (o *ConfigSetting) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ConfigSetting) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ConfigSetting) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *ConfigSetting) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


