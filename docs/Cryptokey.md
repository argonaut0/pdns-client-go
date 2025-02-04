# Cryptokey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | set to \&quot;Cryptokey\&quot; | [optional] 
**Id** | Pointer to **int32** | The internal identifier, read only | [optional] 
**Keytype** | Pointer to **string** |  | [optional] 
**Active** | Pointer to **bool** | Whether or not the key is in active use | [optional] 
**Published** | Pointer to **bool** | Whether or not the DNSKEY record is published in the zone | [optional] 
**Dnskey** | Pointer to **string** | The DNSKEY record for this key | [optional] 
**Ds** | Pointer to **[]string** | An array of DS records for this key | [optional] 
**Cds** | Pointer to **[]string** | An array of DS records for this key, filtered by CDS publication settings | [optional] 
**Privatekey** | Pointer to **string** | The private key in ISC format | [optional] 
**Algorithm** | Pointer to **string** | The name of the algorithm of the key, should be a mnemonic | [optional] 
**Bits** | Pointer to **int32** | The size of the key | [optional] 

## Methods

### NewCryptokey

`func NewCryptokey() *Cryptokey`

NewCryptokey instantiates a new Cryptokey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCryptokeyWithDefaults

`func NewCryptokeyWithDefaults() *Cryptokey`

NewCryptokeyWithDefaults instantiates a new Cryptokey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Cryptokey) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Cryptokey) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Cryptokey) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Cryptokey) HasType() bool`

HasType returns a boolean if a field has been set.

### GetId

`func (o *Cryptokey) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Cryptokey) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Cryptokey) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Cryptokey) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKeytype

`func (o *Cryptokey) GetKeytype() string`

GetKeytype returns the Keytype field if non-nil, zero value otherwise.

### GetKeytypeOk

`func (o *Cryptokey) GetKeytypeOk() (*string, bool)`

GetKeytypeOk returns a tuple with the Keytype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeytype

`func (o *Cryptokey) SetKeytype(v string)`

SetKeytype sets Keytype field to given value.

### HasKeytype

`func (o *Cryptokey) HasKeytype() bool`

HasKeytype returns a boolean if a field has been set.

### GetActive

`func (o *Cryptokey) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *Cryptokey) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *Cryptokey) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *Cryptokey) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetPublished

`func (o *Cryptokey) GetPublished() bool`

GetPublished returns the Published field if non-nil, zero value otherwise.

### GetPublishedOk

`func (o *Cryptokey) GetPublishedOk() (*bool, bool)`

GetPublishedOk returns a tuple with the Published field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublished

`func (o *Cryptokey) SetPublished(v bool)`

SetPublished sets Published field to given value.

### HasPublished

`func (o *Cryptokey) HasPublished() bool`

HasPublished returns a boolean if a field has been set.

### GetDnskey

`func (o *Cryptokey) GetDnskey() string`

GetDnskey returns the Dnskey field if non-nil, zero value otherwise.

### GetDnskeyOk

`func (o *Cryptokey) GetDnskeyOk() (*string, bool)`

GetDnskeyOk returns a tuple with the Dnskey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnskey

`func (o *Cryptokey) SetDnskey(v string)`

SetDnskey sets Dnskey field to given value.

### HasDnskey

`func (o *Cryptokey) HasDnskey() bool`

HasDnskey returns a boolean if a field has been set.

### GetDs

`func (o *Cryptokey) GetDs() []string`

GetDs returns the Ds field if non-nil, zero value otherwise.

### GetDsOk

`func (o *Cryptokey) GetDsOk() (*[]string, bool)`

GetDsOk returns a tuple with the Ds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDs

`func (o *Cryptokey) SetDs(v []string)`

SetDs sets Ds field to given value.

### HasDs

`func (o *Cryptokey) HasDs() bool`

HasDs returns a boolean if a field has been set.

### GetCds

`func (o *Cryptokey) GetCds() []string`

GetCds returns the Cds field if non-nil, zero value otherwise.

### GetCdsOk

`func (o *Cryptokey) GetCdsOk() (*[]string, bool)`

GetCdsOk returns a tuple with the Cds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCds

`func (o *Cryptokey) SetCds(v []string)`

SetCds sets Cds field to given value.

### HasCds

`func (o *Cryptokey) HasCds() bool`

HasCds returns a boolean if a field has been set.

### GetPrivatekey

`func (o *Cryptokey) GetPrivatekey() string`

GetPrivatekey returns the Privatekey field if non-nil, zero value otherwise.

### GetPrivatekeyOk

`func (o *Cryptokey) GetPrivatekeyOk() (*string, bool)`

GetPrivatekeyOk returns a tuple with the Privatekey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivatekey

`func (o *Cryptokey) SetPrivatekey(v string)`

SetPrivatekey sets Privatekey field to given value.

### HasPrivatekey

`func (o *Cryptokey) HasPrivatekey() bool`

HasPrivatekey returns a boolean if a field has been set.

### GetAlgorithm

`func (o *Cryptokey) GetAlgorithm() string`

GetAlgorithm returns the Algorithm field if non-nil, zero value otherwise.

### GetAlgorithmOk

`func (o *Cryptokey) GetAlgorithmOk() (*string, bool)`

GetAlgorithmOk returns a tuple with the Algorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithm

`func (o *Cryptokey) SetAlgorithm(v string)`

SetAlgorithm sets Algorithm field to given value.

### HasAlgorithm

`func (o *Cryptokey) HasAlgorithm() bool`

HasAlgorithm returns a boolean if a field has been set.

### GetBits

`func (o *Cryptokey) GetBits() int32`

GetBits returns the Bits field if non-nil, zero value otherwise.

### GetBitsOk

`func (o *Cryptokey) GetBitsOk() (*int32, bool)`

GetBitsOk returns a tuple with the Bits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBits

`func (o *Cryptokey) SetBits(v int32)`

SetBits sets Bits field to given value.

### HasBits

`func (o *Cryptokey) HasBits() bool`

HasBits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


