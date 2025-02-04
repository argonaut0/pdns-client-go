# \ServersAPI

All URIs are relative to */api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CacheFlushByName**](ServersAPI.md#CacheFlushByName) | **Put** /servers/{server_id}/cache/flush | Flush a cache-entry by name
[**ListServer**](ServersAPI.md#ListServer) | **Get** /servers/{server_id} | List a server
[**ListServers**](ServersAPI.md#ListServers) | **Get** /servers | List all servers



## CacheFlushByName

> CacheFlushResult CacheFlushByName(ctx, serverId).Domain(domain).Execute()

Flush a cache-entry by name

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
	domain := "domain_example" // string | The domain name to flush from the cache

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServersAPI.CacheFlushByName(context.Background(), serverId).Domain(domain).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServersAPI.CacheFlushByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CacheFlushByName`: CacheFlushResult
	fmt.Fprintf(os.Stdout, "Response from `ServersAPI.CacheFlushByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The id of the server to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiCacheFlushByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **domain** | **string** | The domain name to flush from the cache | 

### Return type

[**CacheFlushResult**](CacheFlushResult.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListServer

> Server ListServer(ctx, serverId).Execute()

List a server

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServersAPI.ListServer(context.Background(), serverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServersAPI.ListServer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListServer`: Server
	fmt.Fprintf(os.Stdout, "Response from `ServersAPI.ListServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The id of the server to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiListServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Server**](Server.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListServers

> []Server ListServers(ctx).Execute()

List all servers

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServersAPI.ListServers(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServersAPI.ListServers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListServers`: []Server
	fmt.Fprintf(os.Stdout, "Response from `ServersAPI.ListServers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListServersRequest struct via the builder pattern


### Return type

[**[]Server**](Server.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

