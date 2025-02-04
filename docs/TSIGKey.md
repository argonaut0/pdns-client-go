# TSIGKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the key | [optional] 
**Id** | Pointer to **string** | The ID for this key, used in the TSIGkey URL endpoint. | [optional] [readonly] 
**Algorithm** | Pointer to **string** | The algorithm of the TSIG key | [optional] 
**Key** | Pointer to **string** | The Base64 encoded secret key, empty when listing keys. MAY be empty when POSTing to have the server generate the key material | [optional] 
**Type** | Pointer to **string** | Set to \&quot;TSIGKey\&quot; | [optional] [readonly] 

## Methods

### NewTSIGKey

`func NewTSIGKey() *TSIGKey`

NewTSIGKey instantiates a new TSIGKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTSIGKeyWithDefaults

`func NewTSIGKeyWithDefaults() *TSIGKey`

NewTSIGKeyWithDefaults instantiates a new TSIGKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TSIGKey) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TSIGKey) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TSIGKey) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TSIGKey) HasName() bool`

HasName returns a boolean if a field has been set.

### GetId

`func (o *TSIGKey) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TSIGKey) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TSIGKey) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TSIGKey) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAlgorithm

`func (o *TSIGKey) GetAlgorithm() string`

GetAlgorithm returns the Algorithm field if non-nil, zero value otherwise.

### GetAlgorithmOk

`func (o *TSIGKey) GetAlgorithmOk() (*string, bool)`

GetAlgorithmOk returns a tuple with the Algorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithm

`func (o *TSIGKey) SetAlgorithm(v string)`

SetAlgorithm sets Algorithm field to given value.

### HasAlgorithm

`func (o *TSIGKey) HasAlgorithm() bool`

HasAlgorithm returns a boolean if a field has been set.

### GetKey

`func (o *TSIGKey) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *TSIGKey) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *TSIGKey) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *TSIGKey) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetType

`func (o *TSIGKey) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TSIGKey) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TSIGKey) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TSIGKey) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


