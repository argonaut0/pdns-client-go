# \TsigkeyAPI

All URIs are relative to */api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTSIGKey**](TsigkeyAPI.md#CreateTSIGKey) | **Post** /servers/{server_id}/tsigkeys | Add a TSIG key
[**DeleteTSIGKey**](TsigkeyAPI.md#DeleteTSIGKey) | **Delete** /servers/{server_id}/tsigkeys/{tsigkey_id} | Delete the TSIGKey with tsigkey_id
[**GetTSIGKey**](TsigkeyAPI.md#GetTSIGKey) | **Get** /servers/{server_id}/tsigkeys/{tsigkey_id} | Get a specific TSIGKeys on the server, including the actual key
[**ListTSIGKeys**](TsigkeyAPI.md#ListTSIGKeys) | **Get** /servers/{server_id}/tsigkeys | Get all TSIGKeys on the server, except the actual key
[**PutTSIGKey**](TsigkeyAPI.md#PutTSIGKey) | **Put** /servers/{server_id}/tsigkeys/{tsigkey_id} | 



## CreateTSIGKey

> TSIGKey CreateTSIGKey(ctx, serverId).Tsigkey(tsigkey).Execute()

Add a TSIG key



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
	serverId := "serverId_example" // string | The id of the server
	tsigkey := *openapiclient.NewTSIGKey() // TSIGKey | The TSIGKey to add

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TsigkeyAPI.CreateTSIGKey(context.Background(), serverId).Tsigkey(tsigkey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TsigkeyAPI.CreateTSIGKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTSIGKey`: TSIGKey
	fmt.Fprintf(os.Stdout, "Response from `TsigkeyAPI.CreateTSIGKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The id of the server | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTSIGKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tsigkey** | [**TSIGKey**](TSIGKey.md) | The TSIGKey to add | 

### Return type

[**TSIGKey**](TSIGKey.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTSIGKey

> DeleteTSIGKey(ctx, serverId, tsigkeyId).Execute()

Delete the TSIGKey with tsigkey_id

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
	serverId := "serverId_example" // string | The id of the server to retrieve the key from
	tsigkeyId := "tsigkeyId_example" // string | The id of the TSIGkey. Should match the \"id\" field in the TSIGKey object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TsigkeyAPI.DeleteTSIGKey(context.Background(), serverId, tsigkeyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TsigkeyAPI.DeleteTSIGKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The id of the server to retrieve the key from | 
**tsigkeyId** | **string** | The id of the TSIGkey. Should match the \&quot;id\&quot; field in the TSIGKey object | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTSIGKeyRequest struct via the builder pattern


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


## GetTSIGKey

> TSIGKey GetTSIGKey(ctx, serverId, tsigkeyId).Execute()

Get a specific TSIGKeys on the server, including the actual key

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
	serverId := "serverId_example" // string | The id of the server to retrieve the key from
	tsigkeyId := "tsigkeyId_example" // string | The id of the TSIGkey. Should match the \"id\" field in the TSIGKey object

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TsigkeyAPI.GetTSIGKey(context.Background(), serverId, tsigkeyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TsigkeyAPI.GetTSIGKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTSIGKey`: TSIGKey
	fmt.Fprintf(os.Stdout, "Response from `TsigkeyAPI.GetTSIGKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The id of the server to retrieve the key from | 
**tsigkeyId** | **string** | The id of the TSIGkey. Should match the \&quot;id\&quot; field in the TSIGKey object | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTSIGKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**TSIGKey**](TSIGKey.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTSIGKeys

> []TSIGKey ListTSIGKeys(ctx, serverId).Execute()

Get all TSIGKeys on the server, except the actual key

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
	serverId := "serverId_example" // string | The id of the server

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TsigkeyAPI.ListTSIGKeys(context.Background(), serverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TsigkeyAPI.ListTSIGKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTSIGKeys`: []TSIGKey
	fmt.Fprintf(os.Stdout, "Response from `TsigkeyAPI.ListTSIGKeys`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The id of the server | 

### Other Parameters

Other parameters are passed through a pointer to a apiListTSIGKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]TSIGKey**](TSIGKey.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutTSIGKey

> TSIGKey PutTSIGKey(ctx, serverId, tsigkeyId).Tsigkey(tsigkey).Execute()





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
	serverId := "serverId_example" // string | The id of the server to retrieve the key from
	tsigkeyId := "tsigkeyId_example" // string | The id of the TSIGkey. Should match the \"id\" field in the TSIGKey object
	tsigkey := *openapiclient.NewTSIGKey() // TSIGKey | A (possibly stripped down) TSIGKey object with the new values

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TsigkeyAPI.PutTSIGKey(context.Background(), serverId, tsigkeyId).Tsigkey(tsigkey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TsigkeyAPI.PutTSIGKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutTSIGKey`: TSIGKey
	fmt.Fprintf(os.Stdout, "Response from `TsigkeyAPI.PutTSIGKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The id of the server to retrieve the key from | 
**tsigkeyId** | **string** | The id of the TSIGkey. Should match the \&quot;id\&quot; field in the TSIGKey object | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutTSIGKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **tsigkey** | [**TSIGKey**](TSIGKey.md) | A (possibly stripped down) TSIGKey object with the new values | 

### Return type

[**TSIGKey**](TSIGKey.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

