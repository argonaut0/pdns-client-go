# \ZonesAPI

All URIs are relative to */api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AxfrExportZone**](ZonesAPI.md#AxfrExportZone) | **Get** /servers/{server_id}/zones/{zone_id}/export | Returns the zone in AXFR format.
[**AxfrRetrieveZone**](ZonesAPI.md#AxfrRetrieveZone) | **Put** /servers/{server_id}/zones/{zone_id}/axfr-retrieve | Retrieve slave zone from its master.
[**CreateZone**](ZonesAPI.md#CreateZone) | **Post** /servers/{server_id}/zones | Creates a new domain, returns the Zone on creation.
[**DeleteZone**](ZonesAPI.md#DeleteZone) | **Delete** /servers/{server_id}/zones/{zone_id} | Deletes this zone, all attached metadata and rrsets.
[**ListZone**](ZonesAPI.md#ListZone) | **Get** /servers/{server_id}/zones/{zone_id} | zone managed by a server
[**ListZones**](ZonesAPI.md#ListZones) | **Get** /servers/{server_id}/zones | List all Zones in a server
[**NotifyZone**](ZonesAPI.md#NotifyZone) | **Put** /servers/{server_id}/zones/{zone_id}/notify | Send a DNS NOTIFY to all slaves.
[**PatchZone**](ZonesAPI.md#PatchZone) | **Patch** /servers/{server_id}/zones/{zone_id} | Creates/modifies/deletes RRsets present in the payload and their comments. Returns 204 No Content on success.
[**PutZone**](ZonesAPI.md#PutZone) | **Put** /servers/{server_id}/zones/{zone_id} | Modifies basic zone data.
[**RectifyZone**](ZonesAPI.md#RectifyZone) | **Put** /servers/{server_id}/zones/{zone_id}/rectify | Rectify the zone data.



## AxfrExportZone

> string AxfrExportZone(ctx, serverId, zoneId).Execute()

Returns the zone in AXFR format.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/argonaut0/pdns-client-go"
)

func main() {
	serverId := "serverId_example" // string | The id of the server to retrieve
	zoneId := "zoneId_example" // string | The id of the zone to retrieve

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ZonesAPI.AxfrExportZone(context.Background(), serverId, zoneId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ZonesAPI.AxfrExportZone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AxfrExportZone`: string
	fmt.Fprintf(os.Stdout, "Response from `ZonesAPI.AxfrExportZone`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The id of the server to retrieve | 
**zoneId** | **string** | The id of the zone to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiAxfrExportZoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**string**

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AxfrRetrieveZone

> AxfrRetrieveZone(ctx, serverId, zoneId).Execute()

Retrieve slave zone from its master.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/argonaut0/pdns-client-go"
)

func main() {
	serverId := "serverId_example" // string | The id of the server to retrieve
	zoneId := "zoneId_example" // string | The id of the zone to retrieve

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ZonesAPI.AxfrRetrieveZone(context.Background(), serverId, zoneId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ZonesAPI.AxfrRetrieveZone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The id of the server to retrieve | 
**zoneId** | **string** | The id of the zone to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiAxfrRetrieveZoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateZone

> Zone CreateZone(ctx, serverId).ZoneStruct(zoneStruct).Rrsets(rrsets).Execute()

Creates a new domain, returns the Zone on creation.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/argonaut0/pdns-client-go"
)

func main() {
	serverId := "serverId_example" // string | The id of the server to retrieve
	zoneStruct := *openapiclient.NewZone() // Zone | The zone struct to patch with
	rrsets := true // bool | “true” (default) or “false”, whether to include the “rrsets” in the response Zone object. (optional) (default to true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ZonesAPI.CreateZone(context.Background(), serverId).ZoneStruct(zoneStruct).Rrsets(rrsets).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ZonesAPI.CreateZone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateZone`: Zone
	fmt.Fprintf(os.Stdout, "Response from `ZonesAPI.CreateZone`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The id of the server to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateZoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **zoneStruct** | [**Zone**](Zone.md) | The zone struct to patch with | 
 **rrsets** | **bool** | “true” (default) or “false”, whether to include the “rrsets” in the response Zone object. | [default to true]

### Return type

[**Zone**](Zone.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteZone

> DeleteZone(ctx, serverId, zoneId).Execute()

Deletes this zone, all attached metadata and rrsets.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/argonaut0/pdns-client-go"
)

func main() {
	serverId := "serverId_example" // string | The id of the server to retrieve
	zoneId := "zoneId_example" // string | The id of the zone to retrieve

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ZonesAPI.DeleteZone(context.Background(), serverId, zoneId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ZonesAPI.DeleteZone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The id of the server to retrieve | 
**zoneId** | **string** | The id of the zone to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteZoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListZone

> Zone ListZone(ctx, serverId, zoneId).Rrsets(rrsets).RrsetName(rrsetName).RrsetType(rrsetType).Execute()

zone managed by a server

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/argonaut0/pdns-client-go"
)

func main() {
	serverId := "serverId_example" // string | The id of the server to retrieve
	zoneId := "zoneId_example" // string | The id of the zone to retrieve
	rrsets := true // bool | “true” (default) or “false”, whether to include the “rrsets” in the response Zone object. (optional) (default to true)
	rrsetName := "rrsetName_example" // string | Limit output to RRsets for this name. (optional)
	rrsetType := "rrsetType_example" // string | Limit output to the RRset of this type. Can only be used together with rrset_name. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ZonesAPI.ListZone(context.Background(), serverId, zoneId).Rrsets(rrsets).RrsetName(rrsetName).RrsetType(rrsetType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ZonesAPI.ListZone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListZone`: Zone
	fmt.Fprintf(os.Stdout, "Response from `ZonesAPI.ListZone`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The id of the server to retrieve | 
**zoneId** | **string** | The id of the zone to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiListZoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **rrsets** | **bool** | “true” (default) or “false”, whether to include the “rrsets” in the response Zone object. | [default to true]
 **rrsetName** | **string** | Limit output to RRsets for this name. | 
 **rrsetType** | **string** | Limit output to the RRset of this type. Can only be used together with rrset_name. | 

### Return type

[**Zone**](Zone.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListZones

> []Zone ListZones(ctx, serverId).Zone(zone).Dnssec(dnssec).Execute()

List all Zones in a server

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/argonaut0/pdns-client-go"
)

func main() {
	serverId := "serverId_example" // string | The id of the server to retrieve
	zone := "zone_example" // string | When set to the name of a zone, only this zone is returned. If no zone with that name exists, the response is an empty array. This can e.g. be used to check if a zone exists in the database without having to guess/encode the zone's id or to check if a zone exists.  (optional)
	dnssec := true // bool | “true” (default) or “false”, whether to include the “dnssec” and ”edited_serial” fields in the Zone objects. Setting this to ”false” will make the query a lot faster. (optional) (default to true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ZonesAPI.ListZones(context.Background(), serverId).Zone(zone).Dnssec(dnssec).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ZonesAPI.ListZones``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListZones`: []Zone
	fmt.Fprintf(os.Stdout, "Response from `ZonesAPI.ListZones`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The id of the server to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiListZonesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **zone** | **string** | When set to the name of a zone, only this zone is returned. If no zone with that name exists, the response is an empty array. This can e.g. be used to check if a zone exists in the database without having to guess/encode the zone&#39;s id or to check if a zone exists.  | 
 **dnssec** | **bool** | “true” (default) or “false”, whether to include the “dnssec” and ”edited_serial” fields in the Zone objects. Setting this to ”false” will make the query a lot faster. | [default to true]

### Return type

[**[]Zone**](Zone.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NotifyZone

> NotifyZone(ctx, serverId, zoneId).Execute()

Send a DNS NOTIFY to all slaves.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/argonaut0/pdns-client-go"
)

func main() {
	serverId := "serverId_example" // string | The id of the server to retrieve
	zoneId := "zoneId_example" // string | The id of the zone to retrieve

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ZonesAPI.NotifyZone(context.Background(), serverId, zoneId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ZonesAPI.NotifyZone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The id of the server to retrieve | 
**zoneId** | **string** | The id of the zone to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiNotifyZoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchZone

> PatchZone(ctx, serverId, zoneId).ZoneStruct(zoneStruct).Execute()

Creates/modifies/deletes RRsets present in the payload and their comments. Returns 204 No Content on success.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/argonaut0/pdns-client-go"
)

func main() {
	serverId := "serverId_example" // string | The id of the server to retrieve
	zoneId := "zoneId_example" // string | 
	zoneStruct := *openapiclient.NewZone() // Zone | The zone struct to patch with

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ZonesAPI.PatchZone(context.Background(), serverId, zoneId).ZoneStruct(zoneStruct).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ZonesAPI.PatchZone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The id of the server to retrieve | 
**zoneId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchZoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **zoneStruct** | [**Zone**](Zone.md) | The zone struct to patch with | 

### Return type

 (empty response body)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutZone

> PutZone(ctx, serverId, zoneId).ZoneStruct(zoneStruct).Execute()

Modifies basic zone data.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/argonaut0/pdns-client-go"
)

func main() {
	serverId := "serverId_example" // string | The id of the server to retrieve
	zoneId := "zoneId_example" // string | 
	zoneStruct := *openapiclient.NewZone() // Zone | The zone struct to patch with

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ZonesAPI.PutZone(context.Background(), serverId, zoneId).ZoneStruct(zoneStruct).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ZonesAPI.PutZone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The id of the server to retrieve | 
**zoneId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutZoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **zoneStruct** | [**Zone**](Zone.md) | The zone struct to patch with | 

### Return type

 (empty response body)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RectifyZone

> string RectifyZone(ctx, serverId, zoneId).Execute()

Rectify the zone data.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/argonaut0/pdns-client-go"
)

func main() {
	serverId := "serverId_example" // string | The id of the server to retrieve
	zoneId := "zoneId_example" // string | The id of the zone to retrieve

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ZonesAPI.RectifyZone(context.Background(), serverId, zoneId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ZonesAPI.RectifyZone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RectifyZone`: string
	fmt.Fprintf(os.Stdout, "Response from `ZonesAPI.RectifyZone`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The id of the server to retrieve | 
**zoneId** | **string** | The id of the zone to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiRectifyZoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**string**

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

