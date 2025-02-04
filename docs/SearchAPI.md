# \SearchAPI

All URIs are relative to */api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SearchData**](SearchAPI.md#SearchData) | **Get** /servers/{server_id}/search-data | Search the data inside PowerDNS



## SearchData

> []SearchResult SearchData(ctx, serverId).Q(q).Max(max).ObjectType(objectType).Execute()

Search the data inside PowerDNS



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
	q := "q_example" // string | The string to search for
	max := int32(56) // int32 | Maximum number of entries to return
	objectType := "objectType_example" // string | Type of data to search for, one of “all”, “zone”, “record”, “comment” (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchAPI.SearchData(context.Background(), serverId).Q(q).Max(max).ObjectType(objectType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.SearchData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchData`: []SearchResult
	fmt.Fprintf(os.Stdout, "Response from `SearchAPI.SearchData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The id of the server to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **q** | **string** | The string to search for | 
 **max** | **int32** | Maximum number of entries to return | 
 **objectType** | **string** | Type of data to search for, one of “all”, “zone”, “record”, “comment” | 

### Return type

[**[]SearchResult**](SearchResult.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

