# \ConfigAPI

All URIs are relative to */api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetConfig**](ConfigAPI.md#GetConfig) | **Get** /servers/{server_id}/config | Returns all ConfigSettings for a single server
[**GetConfigSetting**](ConfigAPI.md#GetConfigSetting) | **Get** /servers/{server_id}/config/{config_setting_name} | Returns a specific ConfigSetting for a single server



## GetConfig

> []ConfigSetting GetConfig(ctx, serverId).Execute()

Returns all ConfigSettings for a single server

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
	resp, r, err := apiClient.ConfigAPI.GetConfig(context.Background(), serverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigAPI.GetConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConfig`: []ConfigSetting
	fmt.Fprintf(os.Stdout, "Response from `ConfigAPI.GetConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The id of the server to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ConfigSetting**](ConfigSetting.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConfigSetting

> ConfigSetting GetConfigSetting(ctx, serverId, configSettingName).Execute()

Returns a specific ConfigSetting for a single server



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
	configSettingName := "configSettingName_example" // string | The name of the setting to retrieve

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigAPI.GetConfigSetting(context.Background(), serverId, configSettingName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigAPI.GetConfigSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConfigSetting`: ConfigSetting
	fmt.Fprintf(os.Stdout, "Response from `ConfigAPI.GetConfigSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The id of the server to retrieve | 
**configSettingName** | **string** | The name of the setting to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConfigSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ConfigSetting**](ConfigSetting.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

