# Server

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | Set to “Server” | [optional] 
**Id** | Pointer to **string** | The id of the server, “localhost” | [optional] 
**DaemonType** | Pointer to **string** | “recursor” for the PowerDNS Recursor and “authoritative” for the Authoritative Server | [optional] 
**Version** | Pointer to **string** | The version of the server software | [optional] 
**Url** | Pointer to **string** | The API endpoint for this server | [optional] 
**ConfigUrl** | Pointer to **string** | The API endpoint for this server’s configuration | [optional] 
**ZonesUrl** | Pointer to **string** | The API endpoint for this server’s zones | [optional] 

## Methods

### NewServer

`func NewServer() *Server`

NewServer instantiates a new Server object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerWithDefaults

`func NewServerWithDefaults() *Server`

NewServerWithDefaults instantiates a new Server object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Server) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Server) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Server) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Server) HasType() bool`

HasType returns a boolean if a field has been set.

### GetId

`func (o *Server) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Server) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Server) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Server) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDaemonType

`func (o *Server) GetDaemonType() string`

GetDaemonType returns the DaemonType field if non-nil, zero value otherwise.

### GetDaemonTypeOk

`func (o *Server) GetDaemonTypeOk() (*string, bool)`

GetDaemonTypeOk returns a tuple with the DaemonType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaemonType

`func (o *Server) SetDaemonType(v string)`

SetDaemonType sets DaemonType field to given value.

### HasDaemonType

`func (o *Server) HasDaemonType() bool`

HasDaemonType returns a boolean if a field has been set.

### GetVersion

`func (o *Server) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Server) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Server) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Server) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetUrl

`func (o *Server) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Server) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Server) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *Server) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetConfigUrl

`func (o *Server) GetConfigUrl() string`

GetConfigUrl returns the ConfigUrl field if non-nil, zero value otherwise.

### GetConfigUrlOk

`func (o *Server) GetConfigUrlOk() (*string, bool)`

GetConfigUrlOk returns a tuple with the ConfigUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigUrl

`func (o *Server) SetConfigUrl(v string)`

SetConfigUrl sets ConfigUrl field to given value.

### HasConfigUrl

`func (o *Server) HasConfigUrl() bool`

HasConfigUrl returns a boolean if a field has been set.

### GetZonesUrl

`func (o *Server) GetZonesUrl() string`

GetZonesUrl returns the ZonesUrl field if non-nil, zero value otherwise.

### GetZonesUrlOk

`func (o *Server) GetZonesUrlOk() (*string, bool)`

GetZonesUrlOk returns a tuple with the ZonesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZonesUrl

`func (o *Server) SetZonesUrl(v string)`

SetZonesUrl sets ZonesUrl field to given value.

### HasZonesUrl

`func (o *Server) HasZonesUrl() bool`

HasZonesUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


