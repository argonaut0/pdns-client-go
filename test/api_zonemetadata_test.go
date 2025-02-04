/*
PowerDNS Authoritative HTTP API

Testing ZonemetadataAPIService

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

func Test_openapi_ZonemetadataAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ZonemetadataAPIService CreateMetadata", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var serverId string
		var zoneId string

		httpRes, err := apiClient.ZonemetadataAPI.CreateMetadata(context.Background(), serverId, zoneId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ZonemetadataAPIService DeleteMetadata", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var serverId string
		var zoneId string
		var metadataKind string

		httpRes, err := apiClient.ZonemetadataAPI.DeleteMetadata(context.Background(), serverId, zoneId, metadataKind).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ZonemetadataAPIService GetMetadata", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var serverId string
		var zoneId string
		var metadataKind string

		resp, httpRes, err := apiClient.ZonemetadataAPI.GetMetadata(context.Background(), serverId, zoneId, metadataKind).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ZonemetadataAPIService ListMetadata", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var serverId string
		var zoneId string

		resp, httpRes, err := apiClient.ZonemetadataAPI.ListMetadata(context.Background(), serverId, zoneId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ZonemetadataAPIService ModifyMetadata", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var serverId string
		var zoneId string
		var metadataKind string

		resp, httpRes, err := apiClient.ZonemetadataAPI.ModifyMetadata(context.Background(), serverId, zoneId, metadataKind).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
