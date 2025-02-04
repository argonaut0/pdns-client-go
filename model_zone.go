/*
PowerDNS Authoritative HTTP API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the Zone type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Zone{}

// Zone This represents an authoritative DNS Zone.
type Zone struct {
	// Opaque zone id (string), assigned by the server, should not be interpreted by the application. Guaranteed to be safe for embedding in URLs.
	Id *string `json:"id,omitempty"`
	// Name of the zone (e.g. “example.com.”) MUST have a trailing dot
	Name *string `json:"name,omitempty"`
	// Set to “Zone”
	Type *string `json:"type,omitempty"`
	// API endpoint for this zone
	Url *string `json:"url,omitempty"`
	// Zone kind, one of “Native”, “Master”, “Slave”, “Producer”, “Consumer”
	Kind *string `json:"kind,omitempty"`
	// RRSets in this zone (for zones/{zone_id} endpoint only; omitted during GET on the .../zones list endpoint)
	Rrsets []RRSet `json:"rrsets,omitempty"`
	// The SOA serial number
	Serial *int32 `json:"serial,omitempty"`
	// The SOA serial notifications have been sent out for
	NotifiedSerial *int32 `json:"notified_serial,omitempty"`
	// The SOA serial as seen in query responses. Calculated using the SOA-EDIT metadata, default-soa-edit and default-soa-edit-signed settings
	EditedSerial *int32 `json:"edited_serial,omitempty"`
	//  List of IP addresses configured as a master for this zone (“Slave” type zones only)
	Masters []string `json:"masters,omitempty"`
	// Whether or not this zone is DNSSEC signed (inferred from presigned being true XOR presence of at least one cryptokey with active being true)
	Dnssec *bool `json:"dnssec,omitempty"`
	// The NSEC3PARAM record
	Nsec3param *string `json:"nsec3param,omitempty"`
	// Whether or not the zone uses NSEC3 narrow
	Nsec3narrow *bool `json:"nsec3narrow,omitempty"`
	// Whether or not the zone is pre-signed
	Presigned *bool `json:"presigned,omitempty"`
	// The SOA-EDIT metadata item
	SoaEdit *string `json:"soa_edit,omitempty"`
	// The SOA-EDIT-API metadata item
	SoaEditApi *string `json:"soa_edit_api,omitempty"`
	// Whether or not the zone will be rectified on data changes via the API
	ApiRectify *bool `json:"api_rectify,omitempty"`
	// MAY contain a BIND-style zone file when creating a zone
	Zone *string `json:"zone,omitempty"`
	// The catalog this zone is a member of
	Catalog *string `json:"catalog,omitempty"`
	// MAY be set. Its value is defined by local policy
	Account *string `json:"account,omitempty"`
	// MAY be sent in client bodies during creation, and MUST NOT be sent by the server. Simple list of strings of nameserver names, including the trailing dot. Not required for slave zones.
	Nameservers []string `json:"nameservers,omitempty"`
	// The id of the TSIG keys used for master operation in this zone
	MasterTsigKeyIds []string `json:"master_tsig_key_ids,omitempty"`
	// The id of the TSIG keys used for slave operation in this zone
	SlaveTsigKeyIds []string `json:"slave_tsig_key_ids,omitempty"`
}

// NewZone instantiates a new Zone object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewZone() *Zone {
	this := Zone{}
	return &this
}

// NewZoneWithDefaults instantiates a new Zone object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewZoneWithDefaults() *Zone {
	this := Zone{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Zone) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Zone) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Zone) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Zone) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Zone) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Zone) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Zone) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Zone) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Zone) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Zone) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Zone) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Zone) SetType(v string) {
	o.Type = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *Zone) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Zone) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *Zone) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *Zone) SetUrl(v string) {
	o.Url = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *Zone) GetKind() string {
	if o == nil || IsNil(o.Kind) {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Zone) GetKindOk() (*string, bool) {
	if o == nil || IsNil(o.Kind) {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *Zone) HasKind() bool {
	if o != nil && !IsNil(o.Kind) {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *Zone) SetKind(v string) {
	o.Kind = &v
}

// GetRrsets returns the Rrsets field value if set, zero value otherwise.
func (o *Zone) GetRrsets() []RRSet {
	if o == nil || IsNil(o.Rrsets) {
		var ret []RRSet
		return ret
	}
	return o.Rrsets
}

// GetRrsetsOk returns a tuple with the Rrsets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Zone) GetRrsetsOk() ([]RRSet, bool) {
	if o == nil || IsNil(o.Rrsets) {
		return nil, false
	}
	return o.Rrsets, true
}

// HasRrsets returns a boolean if a field has been set.
func (o *Zone) HasRrsets() bool {
	if o != nil && !IsNil(o.Rrsets) {
		return true
	}

	return false
}

// SetRrsets gets a reference to the given []RRSet and assigns it to the Rrsets field.
func (o *Zone) SetRrsets(v []RRSet) {
	o.Rrsets = v
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *Zone) GetSerial() int32 {
	if o == nil || IsNil(o.Serial) {
		var ret int32
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Zone) GetSerialOk() (*int32, bool) {
	if o == nil || IsNil(o.Serial) {
		return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *Zone) HasSerial() bool {
	if o != nil && !IsNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given int32 and assigns it to the Serial field.
func (o *Zone) SetSerial(v int32) {
	o.Serial = &v
}

// GetNotifiedSerial returns the NotifiedSerial field value if set, zero value otherwise.
func (o *Zone) GetNotifiedSerial() int32 {
	if o == nil || IsNil(o.NotifiedSerial) {
		var ret int32
		return ret
	}
	return *o.NotifiedSerial
}

// GetNotifiedSerialOk returns a tuple with the NotifiedSerial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Zone) GetNotifiedSerialOk() (*int32, bool) {
	if o == nil || IsNil(o.NotifiedSerial) {
		return nil, false
	}
	return o.NotifiedSerial, true
}

// HasNotifiedSerial returns a boolean if a field has been set.
func (o *Zone) HasNotifiedSerial() bool {
	if o != nil && !IsNil(o.NotifiedSerial) {
		return true
	}

	return false
}

// SetNotifiedSerial gets a reference to the given int32 and assigns it to the NotifiedSerial field.
func (o *Zone) SetNotifiedSerial(v int32) {
	o.NotifiedSerial = &v
}

// GetEditedSerial returns the EditedSerial field value if set, zero value otherwise.
func (o *Zone) GetEditedSerial() int32 {
	if o == nil || IsNil(o.EditedSerial) {
		var ret int32
		return ret
	}
	return *o.EditedSerial
}

// GetEditedSerialOk returns a tuple with the EditedSerial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Zone) GetEditedSerialOk() (*int32, bool) {
	if o == nil || IsNil(o.EditedSerial) {
		return nil, false
	}
	return o.EditedSerial, true
}

// HasEditedSerial returns a boolean if a field has been set.
func (o *Zone) HasEditedSerial() bool {
	if o != nil && !IsNil(o.EditedSerial) {
		return true
	}

	return false
}

// SetEditedSerial gets a reference to the given int32 and assigns it to the EditedSerial field.
func (o *Zone) SetEditedSerial(v int32) {
	o.EditedSerial = &v
}

// GetMasters returns the Masters field value if set, zero value otherwise.
func (o *Zone) GetMasters() []string {
	if o == nil || IsNil(o.Masters) {
		var ret []string
		return ret
	}
	return o.Masters
}

// GetMastersOk returns a tuple with the Masters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Zone) GetMastersOk() ([]string, bool) {
	if o == nil || IsNil(o.Masters) {
		return nil, false
	}
	return o.Masters, true
}

// HasMasters returns a boolean if a field has been set.
func (o *Zone) HasMasters() bool {
	if o != nil && !IsNil(o.Masters) {
		return true
	}

	return false
}

// SetMasters gets a reference to the given []string and assigns it to the Masters field.
func (o *Zone) SetMasters(v []string) {
	o.Masters = v
}

// GetDnssec returns the Dnssec field value if set, zero value otherwise.
func (o *Zone) GetDnssec() bool {
	if o == nil || IsNil(o.Dnssec) {
		var ret bool
		return ret
	}
	return *o.Dnssec
}

// GetDnssecOk returns a tuple with the Dnssec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Zone) GetDnssecOk() (*bool, bool) {
	if o == nil || IsNil(o.Dnssec) {
		return nil, false
	}
	return o.Dnssec, true
}

// HasDnssec returns a boolean if a field has been set.
func (o *Zone) HasDnssec() bool {
	if o != nil && !IsNil(o.Dnssec) {
		return true
	}

	return false
}

// SetDnssec gets a reference to the given bool and assigns it to the Dnssec field.
func (o *Zone) SetDnssec(v bool) {
	o.Dnssec = &v
}

// GetNsec3param returns the Nsec3param field value if set, zero value otherwise.
func (o *Zone) GetNsec3param() string {
	if o == nil || IsNil(o.Nsec3param) {
		var ret string
		return ret
	}
	return *o.Nsec3param
}

// GetNsec3paramOk returns a tuple with the Nsec3param field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Zone) GetNsec3paramOk() (*string, bool) {
	if o == nil || IsNil(o.Nsec3param) {
		return nil, false
	}
	return o.Nsec3param, true
}

// HasNsec3param returns a boolean if a field has been set.
func (o *Zone) HasNsec3param() bool {
	if o != nil && !IsNil(o.Nsec3param) {
		return true
	}

	return false
}

// SetNsec3param gets a reference to the given string and assigns it to the Nsec3param field.
func (o *Zone) SetNsec3param(v string) {
	o.Nsec3param = &v
}

// GetNsec3narrow returns the Nsec3narrow field value if set, zero value otherwise.
func (o *Zone) GetNsec3narrow() bool {
	if o == nil || IsNil(o.Nsec3narrow) {
		var ret bool
		return ret
	}
	return *o.Nsec3narrow
}

// GetNsec3narrowOk returns a tuple with the Nsec3narrow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Zone) GetNsec3narrowOk() (*bool, bool) {
	if o == nil || IsNil(o.Nsec3narrow) {
		return nil, false
	}
	return o.Nsec3narrow, true
}

// HasNsec3narrow returns a boolean if a field has been set.
func (o *Zone) HasNsec3narrow() bool {
	if o != nil && !IsNil(o.Nsec3narrow) {
		return true
	}

	return false
}

// SetNsec3narrow gets a reference to the given bool and assigns it to the Nsec3narrow field.
func (o *Zone) SetNsec3narrow(v bool) {
	o.Nsec3narrow = &v
}

// GetPresigned returns the Presigned field value if set, zero value otherwise.
func (o *Zone) GetPresigned() bool {
	if o == nil || IsNil(o.Presigned) {
		var ret bool
		return ret
	}
	return *o.Presigned
}

// GetPresignedOk returns a tuple with the Presigned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Zone) GetPresignedOk() (*bool, bool) {
	if o == nil || IsNil(o.Presigned) {
		return nil, false
	}
	return o.Presigned, true
}

// HasPresigned returns a boolean if a field has been set.
func (o *Zone) HasPresigned() bool {
	if o != nil && !IsNil(o.Presigned) {
		return true
	}

	return false
}

// SetPresigned gets a reference to the given bool and assigns it to the Presigned field.
func (o *Zone) SetPresigned(v bool) {
	o.Presigned = &v
}

// GetSoaEdit returns the SoaEdit field value if set, zero value otherwise.
func (o *Zone) GetSoaEdit() string {
	if o == nil || IsNil(o.SoaEdit) {
		var ret string
		return ret
	}
	return *o.SoaEdit
}

// GetSoaEditOk returns a tuple with the SoaEdit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Zone) GetSoaEditOk() (*string, bool) {
	if o == nil || IsNil(o.SoaEdit) {
		return nil, false
	}
	return o.SoaEdit, true
}

// HasSoaEdit returns a boolean if a field has been set.
func (o *Zone) HasSoaEdit() bool {
	if o != nil && !IsNil(o.SoaEdit) {
		return true
	}

	return false
}

// SetSoaEdit gets a reference to the given string and assigns it to the SoaEdit field.
func (o *Zone) SetSoaEdit(v string) {
	o.SoaEdit = &v
}

// GetSoaEditApi returns the SoaEditApi field value if set, zero value otherwise.
func (o *Zone) GetSoaEditApi() string {
	if o == nil || IsNil(o.SoaEditApi) {
		var ret string
		return ret
	}
	return *o.SoaEditApi
}

// GetSoaEditApiOk returns a tuple with the SoaEditApi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Zone) GetSoaEditApiOk() (*string, bool) {
	if o == nil || IsNil(o.SoaEditApi) {
		return nil, false
	}
	return o.SoaEditApi, true
}

// HasSoaEditApi returns a boolean if a field has been set.
func (o *Zone) HasSoaEditApi() bool {
	if o != nil && !IsNil(o.SoaEditApi) {
		return true
	}

	return false
}

// SetSoaEditApi gets a reference to the given string and assigns it to the SoaEditApi field.
func (o *Zone) SetSoaEditApi(v string) {
	o.SoaEditApi = &v
}

// GetApiRectify returns the ApiRectify field value if set, zero value otherwise.
func (o *Zone) GetApiRectify() bool {
	if o == nil || IsNil(o.ApiRectify) {
		var ret bool
		return ret
	}
	return *o.ApiRectify
}

// GetApiRectifyOk returns a tuple with the ApiRectify field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Zone) GetApiRectifyOk() (*bool, bool) {
	if o == nil || IsNil(o.ApiRectify) {
		return nil, false
	}
	return o.ApiRectify, true
}

// HasApiRectify returns a boolean if a field has been set.
func (o *Zone) HasApiRectify() bool {
	if o != nil && !IsNil(o.ApiRectify) {
		return true
	}

	return false
}

// SetApiRectify gets a reference to the given bool and assigns it to the ApiRectify field.
func (o *Zone) SetApiRectify(v bool) {
	o.ApiRectify = &v
}

// GetZone returns the Zone field value if set, zero value otherwise.
func (o *Zone) GetZone() string {
	if o == nil || IsNil(o.Zone) {
		var ret string
		return ret
	}
	return *o.Zone
}

// GetZoneOk returns a tuple with the Zone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Zone) GetZoneOk() (*string, bool) {
	if o == nil || IsNil(o.Zone) {
		return nil, false
	}
	return o.Zone, true
}

// HasZone returns a boolean if a field has been set.
func (o *Zone) HasZone() bool {
	if o != nil && !IsNil(o.Zone) {
		return true
	}

	return false
}

// SetZone gets a reference to the given string and assigns it to the Zone field.
func (o *Zone) SetZone(v string) {
	o.Zone = &v
}

// GetCatalog returns the Catalog field value if set, zero value otherwise.
func (o *Zone) GetCatalog() string {
	if o == nil || IsNil(o.Catalog) {
		var ret string
		return ret
	}
	return *o.Catalog
}

// GetCatalogOk returns a tuple with the Catalog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Zone) GetCatalogOk() (*string, bool) {
	if o == nil || IsNil(o.Catalog) {
		return nil, false
	}
	return o.Catalog, true
}

// HasCatalog returns a boolean if a field has been set.
func (o *Zone) HasCatalog() bool {
	if o != nil && !IsNil(o.Catalog) {
		return true
	}

	return false
}

// SetCatalog gets a reference to the given string and assigns it to the Catalog field.
func (o *Zone) SetCatalog(v string) {
	o.Catalog = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *Zone) GetAccount() string {
	if o == nil || IsNil(o.Account) {
		var ret string
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Zone) GetAccountOk() (*string, bool) {
	if o == nil || IsNil(o.Account) {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *Zone) HasAccount() bool {
	if o != nil && !IsNil(o.Account) {
		return true
	}

	return false
}

// SetAccount gets a reference to the given string and assigns it to the Account field.
func (o *Zone) SetAccount(v string) {
	o.Account = &v
}

// GetNameservers returns the Nameservers field value if set, zero value otherwise.
func (o *Zone) GetNameservers() []string {
	if o == nil || IsNil(o.Nameservers) {
		var ret []string
		return ret
	}
	return o.Nameservers
}

// GetNameserversOk returns a tuple with the Nameservers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Zone) GetNameserversOk() ([]string, bool) {
	if o == nil || IsNil(o.Nameservers) {
		return nil, false
	}
	return o.Nameservers, true
}

// HasNameservers returns a boolean if a field has been set.
func (o *Zone) HasNameservers() bool {
	if o != nil && !IsNil(o.Nameservers) {
		return true
	}

	return false
}

// SetNameservers gets a reference to the given []string and assigns it to the Nameservers field.
func (o *Zone) SetNameservers(v []string) {
	o.Nameservers = v
}

// GetMasterTsigKeyIds returns the MasterTsigKeyIds field value if set, zero value otherwise.
func (o *Zone) GetMasterTsigKeyIds() []string {
	if o == nil || IsNil(o.MasterTsigKeyIds) {
		var ret []string
		return ret
	}
	return o.MasterTsigKeyIds
}

// GetMasterTsigKeyIdsOk returns a tuple with the MasterTsigKeyIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Zone) GetMasterTsigKeyIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.MasterTsigKeyIds) {
		return nil, false
	}
	return o.MasterTsigKeyIds, true
}

// HasMasterTsigKeyIds returns a boolean if a field has been set.
func (o *Zone) HasMasterTsigKeyIds() bool {
	if o != nil && !IsNil(o.MasterTsigKeyIds) {
		return true
	}

	return false
}

// SetMasterTsigKeyIds gets a reference to the given []string and assigns it to the MasterTsigKeyIds field.
func (o *Zone) SetMasterTsigKeyIds(v []string) {
	o.MasterTsigKeyIds = v
}

// GetSlaveTsigKeyIds returns the SlaveTsigKeyIds field value if set, zero value otherwise.
func (o *Zone) GetSlaveTsigKeyIds() []string {
	if o == nil || IsNil(o.SlaveTsigKeyIds) {
		var ret []string
		return ret
	}
	return o.SlaveTsigKeyIds
}

// GetSlaveTsigKeyIdsOk returns a tuple with the SlaveTsigKeyIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Zone) GetSlaveTsigKeyIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.SlaveTsigKeyIds) {
		return nil, false
	}
	return o.SlaveTsigKeyIds, true
}

// HasSlaveTsigKeyIds returns a boolean if a field has been set.
func (o *Zone) HasSlaveTsigKeyIds() bool {
	if o != nil && !IsNil(o.SlaveTsigKeyIds) {
		return true
	}

	return false
}

// SetSlaveTsigKeyIds gets a reference to the given []string and assigns it to the SlaveTsigKeyIds field.
func (o *Zone) SetSlaveTsigKeyIds(v []string) {
	o.SlaveTsigKeyIds = v
}

func (o Zone) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Zone) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.Kind) {
		toSerialize["kind"] = o.Kind
	}
	if !IsNil(o.Rrsets) {
		toSerialize["rrsets"] = o.Rrsets
	}
	if !IsNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	if !IsNil(o.NotifiedSerial) {
		toSerialize["notified_serial"] = o.NotifiedSerial
	}
	if !IsNil(o.EditedSerial) {
		toSerialize["edited_serial"] = o.EditedSerial
	}
	if !IsNil(o.Masters) {
		toSerialize["masters"] = o.Masters
	}
	if !IsNil(o.Dnssec) {
		toSerialize["dnssec"] = o.Dnssec
	}
	if !IsNil(o.Nsec3param) {
		toSerialize["nsec3param"] = o.Nsec3param
	}
	if !IsNil(o.Nsec3narrow) {
		toSerialize["nsec3narrow"] = o.Nsec3narrow
	}
	if !IsNil(o.Presigned) {
		toSerialize["presigned"] = o.Presigned
	}
	if !IsNil(o.SoaEdit) {
		toSerialize["soa_edit"] = o.SoaEdit
	}
	if !IsNil(o.SoaEditApi) {
		toSerialize["soa_edit_api"] = o.SoaEditApi
	}
	if !IsNil(o.ApiRectify) {
		toSerialize["api_rectify"] = o.ApiRectify
	}
	if !IsNil(o.Zone) {
		toSerialize["zone"] = o.Zone
	}
	if !IsNil(o.Catalog) {
		toSerialize["catalog"] = o.Catalog
	}
	if !IsNil(o.Account) {
		toSerialize["account"] = o.Account
	}
	if !IsNil(o.Nameservers) {
		toSerialize["nameservers"] = o.Nameservers
	}
	if !IsNil(o.MasterTsigKeyIds) {
		toSerialize["master_tsig_key_ids"] = o.MasterTsigKeyIds
	}
	if !IsNil(o.SlaveTsigKeyIds) {
		toSerialize["slave_tsig_key_ids"] = o.SlaveTsigKeyIds
	}
	return toSerialize, nil
}

type NullableZone struct {
	value *Zone
	isSet bool
}

func (v NullableZone) Get() *Zone {
	return v.value
}

func (v *NullableZone) Set(val *Zone) {
	v.value = val
	v.isSet = true
}

func (v NullableZone) IsSet() bool {
	return v.isSet
}

func (v *NullableZone) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableZone(val *Zone) *NullableZone {
	return &NullableZone{value: val, isSet: true}
}

func (v NullableZone) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableZone) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


