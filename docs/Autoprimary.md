# Autoprimary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ip** | Pointer to **string** | IP address of the autoprimary server | [optional] 
**Nameserver** | Pointer to **string** | DNS name of the autoprimary server | [optional] 
**Account** | Pointer to **string** | Account name for the autoprimary server | [optional] 

## Methods

### NewAutoprimary

`func NewAutoprimary() *Autoprimary`

NewAutoprimary instantiates a new Autoprimary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoprimaryWithDefaults

`func NewAutoprimaryWithDefaults() *Autoprimary`

NewAutoprimaryWithDefaults instantiates a new Autoprimary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIp

`func (o *Autoprimary) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *Autoprimary) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *Autoprimary) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *Autoprimary) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetNameserver

`func (o *Autoprimary) GetNameserver() string`

GetNameserver returns the Nameserver field if non-nil, zero value otherwise.

### GetNameserverOk

`func (o *Autoprimary) GetNameserverOk() (*string, bool)`

GetNameserverOk returns a tuple with the Nameserver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameserver

`func (o *Autoprimary) SetNameserver(v string)`

SetNameserver sets Nameserver field to given value.

### HasNameserver

`func (o *Autoprimary) HasNameserver() bool`

HasNameserver returns a boolean if a field has been set.

### GetAccount

`func (o *Autoprimary) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *Autoprimary) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *Autoprimary) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *Autoprimary) HasAccount() bool`

HasAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


