# \ZonecryptokeyAPI

All URIs are relative to */api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCryptokey**](ZonecryptokeyAPI.md#CreateCryptokey) | **Post** /servers/{server_id}/zones/{zone_id}/cryptokeys | Creates a Cryptokey
[**DeleteCryptokey**](ZonecryptokeyAPI.md#DeleteCryptokey) | **Delete** /servers/{server_id}/zones/{zone_id}/cryptokeys/{cryptokey_id} | This method deletes a key specified by cryptokey_id.
[**GetCryptokey**](ZonecryptokeyAPI.md#GetCryptokey) | **Get** /servers/{server_id}/zones/{zone_id}/cryptokeys/{cryptokey_id} | Returns all data about the CryptoKey, including the privatekey.
[**ListCryptokeys**](ZonecryptokeyAPI.md#ListCryptokeys) | **Get** /servers/{server_id}/zones/{zone_id}/cryptokeys | Get all CryptoKeys for a zone, except the privatekey
[**ModifyCryptokey**](ZonecryptokeyAPI.md#ModifyCryptokey) | **Put** /servers/{server_id}/zones/{zone_id}/cryptokeys/{cryptokey_id} | This method (de)activates a key from zone_name specified by cryptokey_id



## CreateCryptokey

> Cryptokey CreateCryptokey(ctx, serverId, zoneId).Cryptokey(cryptokey).Execute()

Creates a Cryptokey



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
	cryptokey := *openapiclient.NewCryptokey() // Cryptokey | Add a Cryptokey

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ZonecryptokeyAPI.CreateCryptokey(context.Background(), serverId, zoneId).Cryptokey(cryptokey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ZonecryptokeyAPI.CreateCryptokey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCryptokey`: Cryptokey
	fmt.Fprintf(os.Stdout, "Response from `ZonecryptokeyAPI.CreateCryptokey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The id of the server to retrieve | 
**zoneId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCryptokeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cryptokey** | [**Cryptokey**](Cryptokey.md) | Add a Cryptokey | 

### Return type

[**Cryptokey**](Cryptokey.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCryptokey

> DeleteCryptokey(ctx, serverId, zoneId, cryptokeyId).Execute()

This method deletes a key specified by cryptokey_id.

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
	cryptokeyId := "cryptokeyId_example" // string | The id value of the Cryptokey

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ZonecryptokeyAPI.DeleteCryptokey(context.Background(), serverId, zoneId, cryptokeyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ZonecryptokeyAPI.DeleteCryptokey``: %v\n", err)
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
**cryptokeyId** | **string** | The id value of the Cryptokey | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCryptokeyRequest struct via the builder pattern


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


## GetCryptokey

> Cryptokey GetCryptokey(ctx, serverId, zoneId, cryptokeyId).Execute()

Returns all data about the CryptoKey, including the privatekey.

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
	cryptokeyId := "cryptokeyId_example" // string | The id value of the CryptoKey

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ZonecryptokeyAPI.GetCryptokey(context.Background(), serverId, zoneId, cryptokeyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ZonecryptokeyAPI.GetCryptokey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCryptokey`: Cryptokey
	fmt.Fprintf(os.Stdout, "Response from `ZonecryptokeyAPI.GetCryptokey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The id of the server to retrieve | 
**zoneId** | **string** | The id of the zone to retrieve | 
**cryptokeyId** | **string** | The id value of the CryptoKey | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCryptokeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Cryptokey**](Cryptokey.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCryptokeys

> []Cryptokey ListCryptokeys(ctx, serverId, zoneId).Execute()

Get all CryptoKeys for a zone, except the privatekey

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
	resp, r, err := apiClient.ZonecryptokeyAPI.ListCryptokeys(context.Background(), serverId, zoneId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ZonecryptokeyAPI.ListCryptokeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCryptokeys`: []Cryptokey
	fmt.Fprintf(os.Stdout, "Response from `ZonecryptokeyAPI.ListCryptokeys`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverId** | **string** | The id of the server to retrieve | 
**zoneId** | **string** | The id of the zone to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiListCryptokeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]Cryptokey**](Cryptokey.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyCryptokey

> ModifyCryptokey(ctx, serverId, zoneId, cryptokeyId).Cryptokey(cryptokey).Execute()

This method (de)activates a key from zone_name specified by cryptokey_id

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
	cryptokeyId := "cryptokeyId_example" // string | Cryptokey to manipulate
	cryptokey := *openapiclient.NewCryptokey() // Cryptokey | the Cryptokey

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ZonecryptokeyAPI.ModifyCryptokey(context.Background(), serverId, zoneId, cryptokeyId).Cryptokey(cryptokey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ZonecryptokeyAPI.ModifyCryptokey``: %v\n", err)
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
**cryptokeyId** | **string** | Cryptokey to manipulate | 

### Other Parameters

Other parameters are passed through a pointer to a apiModifyCryptokeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **cryptokey** | [**Cryptokey**](Cryptokey.md) | the Cryptokey | 

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

