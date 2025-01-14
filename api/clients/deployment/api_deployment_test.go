/*
Dynatrace Environment API

Testing DeploymentAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package deployment

import (
	"context"
	envclient "github.com/0sewa0/dynatrace-configuration-as-code-core/gen/environment"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
)

func CreateDtClient(ctx context.Context, apiUrl, apiToken string) *envclient.APIClient {
	tokenKey := "Api-Token"
	configuration := envclient.NewConfiguration()
	configuration.Servers = envclient.ServerConfigurations{{URL: apiUrl}} // You can(/have-to) override the server URL here, as the schema has a URL in it already.
	configuration.HTTPClient = http.DefaultClient                         // You can mess with the http-client here, customize it as you like
	apiClient := envclient.NewAPIClient(configuration)
	ctx = context.WithValue(ctx, envclient.ContextAPIKeys, map[string]envclient.APIKey{
		tokenKey: {
			Prefix: tokenKey,
			Key:    apiToken,
		},
	})
	return apiClient
}

func Test_environment_DeploymentAPIService(t *testing.T) {
	url := "https://test.dev.dynatracelabs.com/api/v1"
	token := "asd"
	ctx := context.Background()
	apiClient := CreateDtClient(ctx, url, token)

	t.Run("Test DeploymentAPIService DownloadAgentInstallerWithVersion", func(t *testing.T) {

		//t.Skip("skip test") // remove to run test

		file, httpRes, err := apiClient.DeploymentAPI.DownloadAgentInstallerWithVersion(context.Background(), "unix", "paas", "1.284.0.20240104-081422").Flavor("default").SkipMetadata(true).Arch("all").Bitness("all").Execute()

		stat, err := file.Stat()
		assert.NotNil(t, stat)
		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeploymentAPIService DownloadAgentOrchestrationSignatureWithVersion", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orchestrationType string
		var version string

		httpRes, err := apiClient.DeploymentAPI.DownloadAgentOrchestrationSignatureWithVersion(context.Background(), orchestrationType, version).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeploymentAPIService DownloadAgentOrchestrationWithVersion", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orchestrationType string
		var version string

		httpRes, err := apiClient.DeploymentAPI.DownloadAgentOrchestrationWithVersion(context.Background(), orchestrationType, version).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeploymentAPIService DownloadBoshReleaseWithVersion", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var osType string
		var version string

		httpRes, err := apiClient.DeploymentAPI.DownloadBoshReleaseWithVersion(context.Background(), osType, version).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeploymentAPIService DownloadGatewayInstallerWithVersion", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var osType string
		var version string

		httpRes, err := apiClient.DeploymentAPI.DownloadGatewayInstallerWithVersion(context.Background(), osType, version).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeploymentAPIService DownloadLatestAgentInstaller", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var osType string
		var installerType string

		httpRes, err := apiClient.DeploymentAPI.DownloadLatestAgentInstaller(context.Background(), osType, installerType).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeploymentAPIService DownloadLatestAgentOrchestration", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orchestrationType string

		httpRes, err := apiClient.DeploymentAPI.DownloadLatestAgentOrchestration(context.Background(), orchestrationType).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeploymentAPIService DownloadLatestAgentOrchestrationSignature", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orchestrationType string

		httpRes, err := apiClient.DeploymentAPI.DownloadLatestAgentOrchestrationSignature(context.Background(), orchestrationType).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeploymentAPIService DownloadLatestGatewayInstaller", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var osType string

		httpRes, err := apiClient.DeploymentAPI.DownloadLatestGatewayInstaller(context.Background(), osType).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeploymentAPIService GetActiveGateInstallerAvailableVersions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var osType string

		resp, httpRes, err := apiClient.DeploymentAPI.GetActiveGateInstallerAvailableVersions(context.Background(), osType).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeploymentAPIService GetActiveGateInstallerConnectionInfo", func(t *testing.T) {

		resp, httpRes, err := apiClient.DeploymentAPI.GetActiveGateInstallerConnectionInfo(ctx).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeploymentAPIService GetAgentInstallerAvailableVersions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var osType string
		var installerType string

		resp, httpRes, err := apiClient.DeploymentAPI.GetAgentInstallerAvailableVersions(context.Background(), osType, installerType).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeploymentAPIService GetAgentInstallerConnectionInfo", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.DeploymentAPI.GetAgentInstallerConnectionInfo(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeploymentAPIService GetAgentInstallerConnectionInfoEndpoints", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.DeploymentAPI.GetAgentInstallerConnectionInfoEndpoints(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeploymentAPIService GetAgentInstallerMetaInfo", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var osType string
		var installerType string

		resp, httpRes, err := apiClient.DeploymentAPI.GetAgentInstallerMetaInfo(context.Background(), osType, installerType).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeploymentAPIService GetAgentInstallerWithVersionChecksum", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var osType string
		var installerType string
		var version string

		resp, httpRes, err := apiClient.DeploymentAPI.GetAgentInstallerWithVersionChecksum(context.Background(), osType, installerType, version).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeploymentAPIService GetAgentProcessModuleConfig", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.DeploymentAPI.GetAgentProcessModuleConfig(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeploymentAPIService GetBoshReleaseAvailableVersions", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var osType string

		resp, httpRes, err := apiClient.DeploymentAPI.GetBoshReleaseAvailableVersions(context.Background(), osType).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeploymentAPIService GetBoshReleaseChecksum", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var osType string
		var version string

		resp, httpRes, err := apiClient.DeploymentAPI.GetBoshReleaseChecksum(context.Background(), osType, version).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeploymentAPIService GetGatewayInstallerMetaInfo", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var osType string

		resp, httpRes, err := apiClient.DeploymentAPI.GetGatewayInstallerMetaInfo(context.Background(), osType).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeploymentAPIService GetLatestActiveGateImage", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.DeploymentAPI.GetLatestActiveGateImage(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeploymentAPIService GetLatestAgentImage", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var agentImageType string

		resp, httpRes, err := apiClient.DeploymentAPI.GetLatestAgentImage(context.Background(), agentImageType).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeploymentAPIService GetLatestLambdaBuildUnits", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.DeploymentAPI.GetLatestLambdaBuildUnits(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeploymentAPIService GetLatestSyntheticImage", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var imageKey string

		resp, httpRes, err := apiClient.DeploymentAPI.GetLatestSyntheticImage(context.Background(), imageKey).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
