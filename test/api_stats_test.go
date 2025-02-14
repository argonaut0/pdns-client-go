/*
PowerDNS Authoritative HTTP API

Testing StatsAPIService

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

func Test_openapi_StatsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test StatsAPIService GetStats", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var serverId string

		resp, httpRes, err := apiClient.StatsAPI.GetStats(context.Background(), serverId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
