# \ZonemetadataAPI

All URIs are relative to */api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMetadata**](ZonemetadataAPI.md#CreateMetadata) | **Post** /servers/{server_id}/zones/{zone_id}/metadata | Creates a set of metadata entries
[**DeleteMetadata**](ZonemetadataAPI.md#DeleteMetadata) | **Delete** /servers/{server_id}/zones/{zone_id}/metadata/{metadata_kind} | Delete all items of a single kind of domain metadata.
[**GetMetadata**](ZonemetadataAPI.md#GetMetadata) | **Get** /servers/{server_id}/zones/{zone_id}/metadata/{metadata_kind} | Get the content of a single kind of domain metadata as a Metadata object.
[**ListMetadata**](ZonemetadataAPI.md#ListMetadata) | **Get** /servers/{server_id}/zones/{zone_id}/metadata | Get all the Metadata associated with the zone.
[**ModifyMetadata**](ZonemetadataAPI.md#ModifyMetadata) | **Put** /servers/{server_id}/zones/{zone_id}/metadata/{metadata_kind} | Replace the content of a single kind of domain metadata.



## CreateMetadata

> CreateMetadata(ctx, serverId, zoneId).Metadata(metadata).Execute()

Creates a set of metadata entries



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
	metadata := *openapiclient.NewMetadata() // Metadata | Metadata object with list of values to create

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ZonemetadataAPI.CreateMetadata(context.Background(), serverId, zoneId).Metadata(metadata).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ZonemetadataAPI.CreateMetadata``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCreateMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **metadata** | [**Metadata**](Metadata.md) | Metadata object with list of values to create | 

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


## DeleteMetadata

> DeleteMetadata(ctx, serverId, zoneId, metadataKind).Execute()

Delete all items of a single kind of domain metadata.

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
	metadataKind := "metadataKind_example" // string | The kind of metadata

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ZonemetadataAPI.DeleteMetadata(context.Background(), serverId, zoneId, metadataKind).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ZonemetadataAPI.DeleteMetadata``: %v\n", err)
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
**metadataKind** | **string** | The kind of metadata | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMetadataRequest struct via the builder pattern


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


## GetMetadata

> Metadata GetMetadata(ctx, serverId, zoneId, metadataKind).Execute()

Get the content of a single kind of domain metadata as a Metadata object.

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
	metadataKind := "metadataKind_example" // string | The kind of metadata

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ZonemetadataAPI.GetMetadata(context.Background(), serverId, zoneId, metadataKind).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ZonemetadataAPI.GetMetadata``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMetadata`: Metadata
	fmt.Fprintf(os.Stdout, "Response from `ZonemetadataAPI.GetMetadata`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The id of the server to retrieve | 
**zoneId** | **string** | The id of the zone to retrieve | 
**metadataKind** | **string** | The kind of metadata | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Metadata**](Metadata.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMetadata

> []Metadata ListMetadata(ctx, serverId, zoneId).Execute()

Get all the Metadata associated with the zone.

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
	resp, r, err := apiClient.ZonemetadataAPI.ListMetadata(context.Background(), serverId, zoneId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ZonemetadataAPI.ListMetadata``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMetadata`: []Metadata
	fmt.Fprintf(os.Stdout, "Response from `ZonemetadataAPI.ListMetadata`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The id of the server to retrieve | 
**zoneId** | **string** | The id of the zone to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiListMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]Metadata**](Metadata.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyMetadata

> Metadata ModifyMetadata(ctx, serverId, zoneId, metadataKind).Metadata(metadata).Execute()

Replace the content of a single kind of domain metadata.



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
	metadataKind := "metadataKind_example" // string | The kind of metadata
	metadata := *openapiclient.NewMetadata() // Metadata | metadata to add/create

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ZonemetadataAPI.ModifyMetadata(context.Background(), serverId, zoneId, metadataKind).Metadata(metadata).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ZonemetadataAPI.ModifyMetadata``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ModifyMetadata`: Metadata
	fmt.Fprintf(os.Stdout, "Response from `ZonemetadataAPI.ModifyMetadata`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The id of the server to retrieve | 
**zoneId** | **string** |  | 
**metadataKind** | **string** | The kind of metadata | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **metadata** | [**Metadata**](Metadata.md) | metadata to add/create | 

### Return type

[**Metadata**](Metadata.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

