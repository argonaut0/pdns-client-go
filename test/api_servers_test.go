/*
PowerDNS Authoritative HTTP API

Testing ServersAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	"testing"

	openapiclient "github.com/argonaut0/pdns-client-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_openapi_ServersAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ServersAPIService CacheFlushByName", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var serverId string

		resp, httpRes, err := apiClient.ServersAPI.CacheFlushByName(context.Background(), serverId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServersAPIService ListServer", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var serverId string

		resp, httpRes, err := apiClient.ServersAPI.ListServer(context.Background(), serverId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServersAPIService ListServers", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ServersAPI.ListServers(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
