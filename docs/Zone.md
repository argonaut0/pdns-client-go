# Zone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Opaque zone id (string), assigned by the server, should not be interpreted by the application. Guaranteed to be safe for embedding in URLs. | [optional] 
**Name** | Pointer to **string** | Name of the zone (e.g. “example.com.”) MUST have a trailing dot | [optional] 
**Type** | Pointer to **string** | Set to “Zone” | [optional] 
**Url** | Pointer to **string** | API endpoint for this zone | [optional] 
**Kind** | Pointer to **string** | Zone kind, one of “Native”, “Master”, “Slave”, “Producer”, “Consumer” | [optional] 
**Rrsets** | Pointer to [**[]RRSet**](RRSet.md) | RRSets in this zone (for zones/{zone_id} endpoint only; omitted during GET on the .../zones list endpoint) | [optional] 
**Serial** | Pointer to **int32** | The SOA serial number | [optional] 
**NotifiedSerial** | Pointer to **int32** | The SOA serial notifications have been sent out for | [optional] 
**EditedSerial** | Pointer to **int32** | The SOA serial as seen in query responses. Calculated using the SOA-EDIT metadata, default-soa-edit and default-soa-edit-signed settings | [optional] 
**Masters** | Pointer to **[]string** |  List of IP addresses configured as a master for this zone (“Slave” type zones only) | [optional] 
**Dnssec** | Pointer to **bool** | Whether or not this zone is DNSSEC signed (inferred from presigned being true XOR presence of at least one cryptokey with active being true) | [optional] 
**Nsec3param** | Pointer to **string** | The NSEC3PARAM record | [optional] 
**Nsec3narrow** | Pointer to **bool** | Whether or not the zone uses NSEC3 narrow | [optional] 
**Presigned** | Pointer to **bool** | Whether or not the zone is pre-signed | [optional] 
**SoaEdit** | Pointer to **string** | The SOA-EDIT metadata item | [optional] 
**SoaEditApi** | Pointer to **string** | The SOA-EDIT-API metadata item | [optional] 
**ApiRectify** | Pointer to **bool** | Whether or not the zone will be rectified on data changes via the API | [optional] 
**Zone** | Pointer to **string** | MAY contain a BIND-style zone file when creating a zone | [optional] 
**Catalog** | Pointer to **string** | The catalog this zone is a member of | [optional] 
**Account** | Pointer to **string** | MAY be set. Its value is defined by local policy | [optional] 
**Nameservers** | Pointer to **[]string** | MAY be sent in client bodies during creation, and MUST NOT be sent by the server. Simple list of strings of nameserver names, including the trailing dot. Not required for slave zones. | [optional] 
**MasterTsigKeyIds** | Pointer to **[]string** | The id of the TSIG keys used for master operation in this zone | [optional] 
**SlaveTsigKeyIds** | Pointer to **[]string** | The id of the TSIG keys used for slave operation in this zone | [optional] 

## Methods

### NewZone

`func NewZone() *Zone`

NewZone instantiates a new Zone object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewZoneWithDefaults

`func NewZoneWithDefaults() *Zone`

NewZoneWithDefaults instantiates a new Zone object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Zone) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Zone) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Zone) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Zone) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Zone) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Zone) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Zone) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Zone) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *Zone) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Zone) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Zone) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Zone) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUrl

`func (o *Zone) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Zone) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Zone) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *Zone) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetKind

`func (o *Zone) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *Zone) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *Zone) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *Zone) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetRrsets

`func (o *Zone) GetRrsets() []RRSet`

GetRrsets returns the Rrsets field if non-nil, zero value otherwise.

### GetRrsetsOk

`func (o *Zone) GetRrsetsOk() (*[]RRSet, bool)`

GetRrsetsOk returns a tuple with the Rrsets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRrsets

`func (o *Zone) SetRrsets(v []RRSet)`

SetRrsets sets Rrsets field to given value.

### HasRrsets

`func (o *Zone) HasRrsets() bool`

HasRrsets returns a boolean if a field has been set.

### GetSerial

`func (o *Zone) GetSerial() int32`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *Zone) GetSerialOk() (*int32, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *Zone) SetSerial(v int32)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *Zone) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetNotifiedSerial

`func (o *Zone) GetNotifiedSerial() int32`

GetNotifiedSerial returns the NotifiedSerial field if non-nil, zero value otherwise.

### GetNotifiedSerialOk

`func (o *Zone) GetNotifiedSerialOk() (*int32, bool)`

GetNotifiedSerialOk returns a tuple with the NotifiedSerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifiedSerial

`func (o *Zone) SetNotifiedSerial(v int32)`

SetNotifiedSerial sets NotifiedSerial field to given value.

### HasNotifiedSerial

`func (o *Zone) HasNotifiedSerial() bool`

HasNotifiedSerial returns a boolean if a field has been set.

### GetEditedSerial

`func (o *Zone) GetEditedSerial() int32`

GetEditedSerial returns the EditedSerial field if non-nil, zero value otherwise.

### GetEditedSerialOk

`func (o *Zone) GetEditedSerialOk() (*int32, bool)`

GetEditedSerialOk returns a tuple with the EditedSerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditedSerial

`func (o *Zone) SetEditedSerial(v int32)`

SetEditedSerial sets EditedSerial field to given value.

### HasEditedSerial

`func (o *Zone) HasEditedSerial() bool`

HasEditedSerial returns a boolean if a field has been set.

### GetMasters

`func (o *Zone) GetMasters() []string`

GetMasters returns the Masters field if non-nil, zero value otherwise.

### GetMastersOk

`func (o *Zone) GetMastersOk() (*[]string, bool)`

GetMastersOk returns a tuple with the Masters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasters

`func (o *Zone) SetMasters(v []string)`

SetMasters sets Masters field to given value.

### HasMasters

`func (o *Zone) HasMasters() bool`

HasMasters returns a boolean if a field has been set.

### GetDnssec

`func (o *Zone) GetDnssec() bool`

GetDnssec returns the Dnssec field if non-nil, zero value otherwise.

### GetDnssecOk

`func (o *Zone) GetDnssecOk() (*bool, bool)`

GetDnssecOk returns a tuple with the Dnssec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnssec

`func (o *Zone) SetDnssec(v bool)`

SetDnssec sets Dnssec field to given value.

### HasDnssec

`func (o *Zone) HasDnssec() bool`

HasDnssec returns a boolean if a field has been set.

### GetNsec3param

`func (o *Zone) GetNsec3param() string`

GetNsec3param returns the Nsec3param field if non-nil, zero value otherwise.

### GetNsec3paramOk

`func (o *Zone) GetNsec3paramOk() (*string, bool)`

GetNsec3paramOk returns a tuple with the Nsec3param field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsec3param

`func (o *Zone) SetNsec3param(v string)`

SetNsec3param sets Nsec3param field to given value.

### HasNsec3param

`func (o *Zone) HasNsec3param() bool`

HasNsec3param returns a boolean if a field has been set.

### GetNsec3narrow

`func (o *Zone) GetNsec3narrow() bool`

GetNsec3narrow returns the Nsec3narrow field if non-nil, zero value otherwise.

### GetNsec3narrowOk

`func (o *Zone) GetNsec3narrowOk() (*bool, bool)`

GetNsec3narrowOk returns a tuple with the Nsec3narrow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsec3narrow

`func (o *Zone) SetNsec3narrow(v bool)`

SetNsec3narrow sets Nsec3narrow field to given value.

### HasNsec3narrow

`func (o *Zone) HasNsec3narrow() bool`

HasNsec3narrow returns a boolean if a field has been set.

### GetPresigned

`func (o *Zone) GetPresigned() bool`

GetPresigned returns the Presigned field if non-nil, zero value otherwise.

### GetPresignedOk

`func (o *Zone) GetPresignedOk() (*bool, bool)`

GetPresignedOk returns a tuple with the Presigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresigned

`func (o *Zone) SetPresigned(v bool)`

SetPresigned sets Presigned field to given value.

### HasPresigned

`func (o *Zone) HasPresigned() bool`

HasPresigned returns a boolean if a field has been set.

### GetSoaEdit

`func (o *Zone) GetSoaEdit() string`

GetSoaEdit returns the SoaEdit field if non-nil, zero value otherwise.

### GetSoaEditOk

`func (o *Zone) GetSoaEditOk() (*string, bool)`

GetSoaEditOk returns a tuple with the SoaEdit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoaEdit

`func (o *Zone) SetSoaEdit(v string)`

SetSoaEdit sets SoaEdit field to given value.

### HasSoaEdit

`func (o *Zone) HasSoaEdit() bool`

HasSoaEdit returns a boolean if a field has been set.

### GetSoaEditApi

`func (o *Zone) GetSoaEditApi() string`

GetSoaEditApi returns the SoaEditApi field if non-nil, zero value otherwise.

### GetSoaEditApiOk

`func (o *Zone) GetSoaEditApiOk() (*string, bool)`

GetSoaEditApiOk returns a tuple with the SoaEditApi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoaEditApi

`func (o *Zone) SetSoaEditApi(v string)`

SetSoaEditApi sets SoaEditApi field to given value.

### HasSoaEditApi

`func (o *Zone) HasSoaEditApi() bool`

HasSoaEditApi returns a boolean if a field has been set.

### GetApiRectify

`func (o *Zone) GetApiRectify() bool`

GetApiRectify returns the ApiRectify field if non-nil, zero value otherwise.

### GetApiRectifyOk

`func (o *Zone) GetApiRectifyOk() (*bool, bool)`

GetApiRectifyOk returns a tuple with the ApiRectify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiRectify

`func (o *Zone) SetApiRectify(v bool)`

SetApiRectify sets ApiRectify field to given value.

### HasApiRectify

`func (o *Zone) HasApiRectify() bool`

HasApiRectify returns a boolean if a field has been set.

### GetZone

`func (o *Zone) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *Zone) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *Zone) SetZone(v string)`

SetZone sets Zone field to given value.

### HasZone

`func (o *Zone) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetCatalog

`func (o *Zone) GetCatalog() string`

GetCatalog returns the Catalog field if non-nil, zero value otherwise.

### GetCatalogOk

`func (o *Zone) GetCatalogOk() (*string, bool)`

GetCatalogOk returns a tuple with the Catalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalog

`func (o *Zone) SetCatalog(v string)`

SetCatalog sets Catalog field to given value.

### HasCatalog

`func (o *Zone) HasCatalog() bool`

HasCatalog returns a boolean if a field has been set.

### GetAccount

`func (o *Zone) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *Zone) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *Zone) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *Zone) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetNameservers

`func (o *Zone) GetNameservers() []string`

GetNameservers returns the Nameservers field if non-nil, zero value otherwise.

### GetNameserversOk

`func (o *Zone) GetNameserversOk() (*[]string, bool)`

GetNameserversOk returns a tuple with the Nameservers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameservers

`func (o *Zone) SetNameservers(v []string)`

SetNameservers sets Nameservers field to given value.

### HasNameservers

`func (o *Zone) HasNameservers() bool`

HasNameservers returns a boolean if a field has been set.

### GetMasterTsigKeyIds

`func (o *Zone) GetMasterTsigKeyIds() []string`

GetMasterTsigKeyIds returns the MasterTsigKeyIds field if non-nil, zero value otherwise.

### GetMasterTsigKeyIdsOk

`func (o *Zone) GetMasterTsigKeyIdsOk() (*[]string, bool)`

GetMasterTsigKeyIdsOk returns a tuple with the MasterTsigKeyIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterTsigKeyIds

`func (o *Zone) SetMasterTsigKeyIds(v []string)`

SetMasterTsigKeyIds sets MasterTsigKeyIds field to given value.

### HasMasterTsigKeyIds

`func (o *Zone) HasMasterTsigKeyIds() bool`

HasMasterTsigKeyIds returns a boolean if a field has been set.

### GetSlaveTsigKeyIds

`func (o *Zone) GetSlaveTsigKeyIds() []string`

GetSlaveTsigKeyIds returns the SlaveTsigKeyIds field if non-nil, zero value otherwise.

### GetSlaveTsigKeyIdsOk

`func (o *Zone) GetSlaveTsigKeyIdsOk() (*[]string, bool)`

GetSlaveTsigKeyIdsOk returns a tuple with the SlaveTsigKeyIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlaveTsigKeyIds

`func (o *Zone) SetSlaveTsigKeyIds(v []string)`

SetSlaveTsigKeyIds sets SlaveTsigKeyIds field to given value.

### HasSlaveTsigKeyIds

`func (o *Zone) HasSlaveTsigKeyIds() bool`

HasSlaveTsigKeyIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


