# \StatsAPI

All URIs are relative to */api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetStats**](StatsAPI.md#GetStats) | **Get** /servers/{server_id}/statistics | Query statistics.



## GetStats

> []GetStats200ResponseInner GetStats(ctx, serverId).Statistic(statistic).Includerings(includerings).Execute()

Query statistics.



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
	statistic := "statistic_example" // string | When set to the name of a specific statistic, only this value is returned. If no statistic with that name exists, the response has a 422 status and an error message.  (optional)
	includerings := true // bool | “true” (default) or “false”, whether to include the Ring items, which can contain thousands of log messages or queried domains. Setting this to ”false” may make the response a lot smaller. (optional) (default to true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StatsAPI.GetStats(context.Background(), serverId).Statistic(statistic).Includerings(includerings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StatsAPI.GetStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStats`: []GetStats200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `StatsAPI.GetStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The id of the server to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **statistic** | **string** | When set to the name of a specific statistic, only this value is returned. If no statistic with that name exists, the response has a 422 status and an error message.  | 
 **includerings** | **bool** | “true” (default) or “false”, whether to include the Ring items, which can contain thousands of log messages or queried domains. Setting this to ”false” may make the response a lot smaller. | [default to true]

### Return type

[**[]GetStats200ResponseInner**](GetStats200ResponseInner.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

