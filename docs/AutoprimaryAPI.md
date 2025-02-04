# \AutoprimaryAPI

All URIs are relative to */api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAutoprimary**](AutoprimaryAPI.md#CreateAutoprimary) | **Post** /servers/{server_id}/autoprimaries | Add an autoprimary
[**DeleteAutoprimary**](AutoprimaryAPI.md#DeleteAutoprimary) | **Delete** /servers/{server_id}/autoprimaries/{ip}/{nameserver} | Delete the autoprimary entry
[**GetAutoprimaries**](AutoprimaryAPI.md#GetAutoprimaries) | **Get** /servers/{server_id}/autoprimaries | Get a list of autoprimaries



## CreateAutoprimary

> CreateAutoprimary(ctx, serverId).Autoprimary(autoprimary).Execute()

Add an autoprimary



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
	serverId := "serverId_example" // string | The id of the server to manage the list of autoprimaries on
	autoprimary := *openapiclient.NewAutoprimary() // Autoprimary | autoprimary entry to add

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AutoprimaryAPI.CreateAutoprimary(context.Background(), serverId).Autoprimary(autoprimary).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoprimaryAPI.CreateAutoprimary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The id of the server to manage the list of autoprimaries on | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAutoprimaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **autoprimary** | [**Autoprimary**](Autoprimary.md) | autoprimary entry to add | 

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


## DeleteAutoprimary

> DeleteAutoprimary(ctx, serverId, ip, nameserver).Execute()

Delete the autoprimary entry

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
	serverId := "serverId_example" // string | The id of the server to delete the autoprimary from
	ip := "ip_example" // string | IP address of autoprimary
	nameserver := "nameserver_example" // string | DNS name of the autoprimary

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AutoprimaryAPI.DeleteAutoprimary(context.Background(), serverId, ip, nameserver).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoprimaryAPI.DeleteAutoprimary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The id of the server to delete the autoprimary from | 
**ip** | **string** | IP address of autoprimary | 
**nameserver** | **string** | DNS name of the autoprimary | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAutoprimaryRequest struct via the builder pattern


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


## GetAutoprimaries

> Autoprimary GetAutoprimaries(ctx, serverId).Execute()

Get a list of autoprimaries

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
	serverId := "serverId_example" // string | The id of the server to manage the list of autoprimaries on

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AutoprimaryAPI.GetAutoprimaries(context.Background(), serverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoprimaryAPI.GetAutoprimaries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAutoprimaries`: Autoprimary
	fmt.Fprintf(os.Stdout, "Response from `AutoprimaryAPI.GetAutoprimaries`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The id of the server to manage the list of autoprimaries on | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAutoprimariesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Autoprimary**](Autoprimary.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

