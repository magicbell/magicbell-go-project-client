package integrations

import (
	"context"
	restClient "github.com/magicbell/magicbell-go-project-client/internal/clients/rest"
	"github.com/magicbell/magicbell-go-project-client/internal/clients/rest/httptransport"
	"github.com/magicbell/magicbell-go-project-client/internal/configmanager"
	"github.com/magicbell/magicbell-go-project-client/pkg/magicbellprojectclientconfig"
	"github.com/magicbell/magicbell-go-project-client/pkg/shared"
	"time"
)

type IntegrationsService struct {
	manager *configmanager.ConfigManager
}

func NewIntegrationsService(manager *configmanager.ConfigManager) *IntegrationsService {
	return &IntegrationsService{
		manager: manager,
	}
}

func (api *IntegrationsService) getConfig() *magicbellprojectclientconfig.Config {
	return api.manager.GetIntegrations()
}

func (api *IntegrationsService) SetBaseUrl(baseUrl string) {
	config := api.getConfig()
	config.SetBaseUrl(baseUrl)
}

func (api *IntegrationsService) SetTimeout(timeout time.Duration) {
	config := api.getConfig()
	config.SetTimeout(timeout)
}

func (api *IntegrationsService) SetAccessToken(accessToken string) {
	config := api.getConfig()
	config.SetAccessToken(accessToken)
}

// Lists all available and configured integrations for the project. Returns a summary of each integration including its type, status, and basic configuration information.
func (api *IntegrationsService) ListIntegrations(ctx context.Context, params ListIntegrationsRequestParams) (*shared.MagicbellProjectClientResponse[ArrayOfIntegrationObjects], *shared.MagicbellProjectClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/integrations").
		WithConfig(config).
		WithOptions(params).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[ArrayOfIntegrationObjects](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewMagicbellProjectClientError[ArrayOfIntegrationObjects](err)
	}

	return shared.NewMagicbellProjectClientResponse[ArrayOfIntegrationObjects](resp), nil
}

// Retrieves the current apns integration configurations for a specific integration type in the project. Returns configuration details and status information.
func (api *IntegrationsService) GetApnsIntegration(ctx context.Context) (*shared.MagicbellProjectClientResponse[ArrayOfApnsConfigObjects], *shared.MagicbellProjectClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/integrations/apns").
		WithConfig(config).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[ArrayOfApnsConfigObjects](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewMagicbellProjectClientError[ArrayOfApnsConfigObjects](err)
	}

	return shared.NewMagicbellProjectClientResponse[ArrayOfApnsConfigObjects](resp), nil
}

// Creates or updates a apns integration for the project. Only administrators can configure integrations.
func (api *IntegrationsService) SaveApnsIntegration(ctx context.Context, apnsConfig ApnsConfig) (*shared.MagicbellProjectClientResponse[ApnsConfig], *shared.MagicbellProjectClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("PUT").
		WithPath("/integrations/apns").
		WithConfig(config).
		WithBody(apnsConfig).
		AddHeader("CONTENT-TYPE", "application/json").
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[ApnsConfig](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewMagicbellProjectClientError[ApnsConfig](err)
	}

	return shared.NewMagicbellProjectClientResponse[ApnsConfig](resp), nil
}

// Removes a apns integration configuration from the project. This will disable the integration's functionality within the project.
func (api *IntegrationsService) DeleteApnsIntegration(ctx context.Context) (*shared.MagicbellProjectClientResponse[any], *shared.MagicbellProjectClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("DELETE").
		WithPath("/integrations/apns").
		WithConfig(config).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[any](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewMagicbellProjectClientError[any](err)
	}

	return shared.NewMagicbellProjectClientResponse[any](resp), nil
}

// Removes a specific apns integration instance by ID from the project.
func (api *IntegrationsService) DeleteApnsIntegrationById(ctx context.Context, id string) (*shared.MagicbellProjectClientResponse[any], *shared.MagicbellProjectClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("DELETE").
		WithPath("/integrations/apns/{id}").
		WithConfig(config).
		AddPathParam("id", id).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[any](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewMagicbellProjectClientError[any](err)
	}

	return shared.NewMagicbellProjectClientResponse[any](resp), nil
}

// Retrieves the current awssns integration configurations for a specific integration type in the project. Returns configuration details and status information.
func (api *IntegrationsService) GetAwssnsIntegration(ctx context.Context) (*shared.MagicbellProjectClientResponse[ArrayOfAwssnsConfigObjects], *shared.MagicbellProjectClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/integrations/awssns").
		WithConfig(config).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[ArrayOfAwssnsConfigObjects](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewMagicbellProjectClientError[ArrayOfAwssnsConfigObjects](err)
	}

	return shared.NewMagicbellProjectClientResponse[ArrayOfAwssnsConfigObjects](resp), nil
}

// Creates or updates a awssns integration for the project. Only administrators can configure integrations.
func (api *IntegrationsService) SaveAwssnsIntegration(ctx context.Context, awssnsConfig AwssnsConfig) (*shared.MagicbellProjectClientResponse[AwssnsConfig], *shared.MagicbellProjectClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("PUT").
		WithPath("/integrations/awssns").
		WithConfig(config).
		WithBody(awssnsConfig).
		AddHeader("CONTENT-TYPE", "application/json").
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[AwssnsConfig](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewMagicbellProjectClientError[AwssnsConfig](err)
	}

	return shared.NewMagicbellProjectClientResponse[AwssnsConfig](resp), nil
}

// Removes a awssns integration configuration from the project. This will disable the integration's functionality within the project.
func (api *IntegrationsService) DeleteAwssnsIntegration(ctx context.Context) (*shared.MagicbellProjectClientResponse[any], *shared.MagicbellProjectClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("DELETE").
		WithPath("/integrations/awssns").
		WithConfig(config).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[any](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewMagicbellProjectClientError[any](err)
	}

	return shared.NewMagicbellProjectClientResponse[any](resp), nil
}

// Removes a specific awssns integration instance by ID from the project.
func (api *IntegrationsService) DeleteAwssnsIntegrationById(ctx context.Context, id string) (*shared.MagicbellProjectClientResponse[any], *shared.MagicbellProjectClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("DELETE").
		WithPath("/integrations/awssns/{id}").
		WithConfig(config).
		AddPathParam("id", id).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[any](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewMagicbellProjectClientError[any](err)
	}

	return shared.NewMagicbellProjectClientResponse[any](resp), nil
}

// Retrieves the current expo integration configurations for a specific integration type in the project. Returns configuration details and status information.
func (api *IntegrationsService) GetExpoIntegration(ctx context.Context) (*shared.MagicbellProjectClientResponse[ArrayOfExpoConfigObjects], *shared.MagicbellProjectClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/integrations/expo").
		WithConfig(config).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[ArrayOfExpoConfigObjects](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewMagicbellProjectClientError[ArrayOfExpoConfigObjects](err)
	}

	return shared.NewMagicbellProjectClientResponse[ArrayOfExpoConfigObjects](resp), nil
}

// Creates or updates a expo integration for the project. Only administrators can configure integrations.
func (api *IntegrationsService) SaveExpoIntegration(ctx context.Context, expoConfig ExpoConfig) (*shared.MagicbellProjectClientResponse[ExpoConfig], *shared.MagicbellProjectClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("PUT").
		WithPath("/integrations/expo").
		WithConfig(config).
		WithBody(expoConfig).
		AddHeader("CONTENT-TYPE", "application/json").
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[ExpoConfig](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewMagicbellProjectClientError[ExpoConfig](err)
	}

	return shared.NewMagicbellProjectClientResponse[ExpoConfig](resp), nil
}

// Removes a expo integration configuration from the project. This will disable the integration's functionality within the project.
func (api *IntegrationsService) DeleteExpoIntegration(ctx context.Context) (*shared.MagicbellProjectClientResponse[any], *shared.MagicbellProjectClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("DELETE").
		WithPath("/integrations/expo").
		WithConfig(config).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[any](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewMagicbellProjectClientError[any](err)
	}

	return shared.NewMagicbellProjectClientResponse[any](resp), nil
}

// Removes a specific expo integration instance by ID from the project.
func (api *IntegrationsService) DeleteExpoIntegrationById(ctx context.Context, id string) (*shared.MagicbellProjectClientResponse[any], *shared.MagicbellProjectClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("DELETE").
		WithPath("/integrations/expo/{id}").
		WithConfig(config).
		AddPathParam("id", id).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[any](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewMagicbellProjectClientError[any](err)
	}

	return shared.NewMagicbellProjectClientResponse[any](resp), nil
}

// Retrieves the current fcm integration configurations for a specific integration type in the project. Returns configuration details and status information.
func (api *IntegrationsService) GetFcmIntegration(ctx context.Context) (*shared.MagicbellProjectClientResponse[ArrayOfFcmConfigObjects], *shared.MagicbellProjectClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/integrations/fcm").
		WithConfig(config).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[ArrayOfFcmConfigObjects](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewMagicbellProjectClientError[ArrayOfFcmConfigObjects](err)
	}

	return shared.NewMagicbellProjectClientResponse[ArrayOfFcmConfigObjects](resp), nil
}

// Creates or updates a fcm integration for the project. Only administrators can configure integrations.
func (api *IntegrationsService) SaveFcmIntegration(ctx context.Context, fcmConfig FcmConfig) (*shared.MagicbellProjectClientResponse[FcmConfig], *shared.MagicbellProjectClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("PUT").
		WithPath("/integrations/fcm").
		WithConfig(config).
		WithBody(fcmConfig).
		AddHeader("CONTENT-TYPE", "application/json").
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[FcmConfig](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewMagicbellProjectClientError[FcmConfig](err)
	}

	return shared.NewMagicbellProjectClientResponse[FcmConfig](resp), nil
}

// Removes a fcm integration configuration from the project. This will disable the integration's functionality within the project.
func (api *IntegrationsService) DeleteFcmIntegration(ctx context.Context) (*shared.MagicbellProjectClientResponse[any], *shared.MagicbellProjectClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("DELETE").
		WithPath("/integrations/fcm").
		WithConfig(config).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[any](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewMagicbellProjectClientError[any](err)
	}

	return shared.NewMagicbellProjectClientResponse[any](resp), nil
}

// Removes a specific fcm integration instance by ID from the project.
func (api *IntegrationsService) DeleteFcmIntegrationById(ctx context.Context, id string) (*shared.MagicbellProjectClientResponse[any], *shared.MagicbellProjectClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("DELETE").
		WithPath("/integrations/fcm/{id}").
		WithConfig(config).
		AddPathParam("id", id).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[any](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewMagicbellProjectClientError[any](err)
	}

	return shared.NewMagicbellProjectClientResponse[any](resp), nil
}

// Retrieves the current github integration configurations for a specific integration type in the project. Returns configuration details and status information.
func (api *IntegrationsService) GetGithubIntegration(ctx context.Context) (*shared.MagicbellProjectClientResponse[ArrayOfGithubConfigObjects], *shared.MagicbellProjectClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/integrations/github").
		WithConfig(config).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[ArrayOfGithubConfigObjects](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewMagicbellProjectClientError[ArrayOfGithubConfigObjects](err)
	}

	return shared.NewMagicbellProjectClientResponse[ArrayOfGithubConfigObjects](resp), nil
}

// Creates or updates a github integration for the project. Only administrators can configure integrations.
func (api *IntegrationsService) SaveGithubIntegration(ctx context.Context, githubConfig GithubConfig) (*shared.MagicbellProjectClientResponse[GithubConfig], *shared.MagicbellProjectClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("PUT").
		WithPath("/integrations/github").
		WithConfig(config).
		WithBody(githubConfig).
		AddHeader("CONTENT-TYPE", "application/json").
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[GithubConfig](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewMagicbellProjectClientError[GithubConfig](err)
	}

	return shared.NewMagicbellProjectClientResponse[GithubConfig](resp), nil
}

// Removes a github integration configuration from the project. This will disable the integration's functionality within the project.
func (api *IntegrationsService) DeleteGithubIntegration(ctx context.Context) (*shared.MagicbellProjectClientResponse[any], *shared.MagicbellProjectClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("DELETE").
		WithPath("/integrations/github").
		WithConfig(config).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[any](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewMagicbellProjectClientError[any](err)
	}

	return shared.NewMagicbellProjectClientResponse[any](resp), nil
}

// Removes a specific github integration instance by ID from the project.
func (api *IntegrationsService) DeleteGithubIntegrationById(ctx context.Context, id string) (*shared.MagicbellProjectClientResponse[any], *shared.MagicbellProjectClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("DELETE").
		WithPath("/integrations/github/{id}").
		WithConfig(config).
		AddPathParam("id", id).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[any](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewMagicbellProjectClientError[any](err)
	}

	return shared.NewMagicbellProjectClientResponse[any](resp), nil
}

// Retrieves the current inbox integration configurations for a specific integration type in the project. Returns configuration details and status information.
func (api *IntegrationsService) GetInboxIntegration(ctx context.Context) (*shared.MagicbellProjectClientResponse[ArrayOfInboxConfigObjects], *shared.MagicbellProjectClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/integrations/inbox").
		WithConfig(config).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[ArrayOfInboxConfigObjects](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewMagicbellProjectClientError[ArrayOfInboxConfigObjects](err)
	}

	return shared.NewMagicbellProjectClientResponse[ArrayOfInboxConfigObjects](resp), nil
}

// Creates or updates a inbox integration for the project. Only administrators can configure integrations.
func (api *IntegrationsService) SaveInboxIntegration(ctx context.Context, inboxConfig InboxConfig) (*shared.MagicbellProjectClientResponse[InboxConfig], *shared.MagicbellProjectClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("PUT").
		WithPath("/integrations/inbox").
		WithConfig(config).
		WithBody(inboxConfig).
		AddHeader("CONTENT-TYPE", "application/json").
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[InboxConfig](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewMagicbellProjectClientError[InboxConfig](err)
	}

	return shared.NewMagicbellProjectClientResponse[InboxConfig](resp), nil
}

// Removes a inbox integration configuration from the project. This will disable the integration's functionality within the project.
func (api *IntegrationsService) DeleteInboxIntegration(ctx context.Context) (*shared.MagicbellProjectClientResponse[any], *shared.MagicbellProjectClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("DELETE").
		WithPath("/integrations/inbox").
		WithConfig(config).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[any](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewMagicbellProjectClientError[any](err)
	}

	return shared.NewMagicbellProjectClientResponse[any](resp), nil
}

// Removes a specific inbox integration instance by ID from the project.
func (api *IntegrationsService) DeleteInboxIntegrationById(ctx context.Context, id string) (*shared.MagicbellProjectClientResponse[any], *shared.MagicbellProjectClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("DELETE").
		WithPath("/integrations/inbox/{id}").
		WithConfig(config).
		AddPathParam("id", id).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[any](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewMagicbellProjectClientError[any](err)
	}

	return shared.NewMagicbellProjectClientResponse[any](resp), nil
}

// Retrieves the current mailgun integration configurations for a specific integration type in the project. Returns configuration details and status information.
func (api *IntegrationsService) GetMailgunIntegration(ctx context.Context) (*shared.MagicbellProjectClientResponse[ArrayOfMailgunConfigObjects], *shared.MagicbellProjectClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/integrations/mailgun").
		WithConfig(config).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[ArrayOfMailgunConfigObjects](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewMagicbellProjectClientError[ArrayOfMailgunConfigObjects](err)
	}

	return shared.NewMagicbellProjectClientResponse[ArrayOfMailgunConfigObjects](resp), nil
}

// Creates or updates a mailgun integration for the project. Only administrators can configure integrations.
func (api *IntegrationsService) SaveMailgunIntegration(ctx context.Context, mailgunConfig MailgunConfig) (*shared.MagicbellProjectClientResponse[MailgunConfig], *shared.MagicbellProjectClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("PUT").
		WithPath("/integrations/mailgun").
		WithConfig(config).
		WithBody(mailgunConfig).
		AddHeader("CONTENT-TYPE", "application/json").
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[MailgunConfig](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewMagicbellProjectClientError[MailgunConfig](err)
	}

	return shared.NewMagicbellProjectClientResponse[MailgunConfig](resp), nil
}

// Removes a mailgun integration configuration from the project. This will disable the integration's functionality within the project.
func (api *IntegrationsService) DeleteMailgunIntegration(ctx context.Context) (*shared.MagicbellProjectClientResponse[any], *shared.MagicbellProjectClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("DELETE").
		WithPath("/integrations/mailgun").
		WithConfig(config).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[any](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewMagicbellProjectClientError[any](err)
	}

	return shared.NewMagicbellProjectClientResponse[any](resp), nil
}

// Removes a specific mailgun integration instance by ID from the project.
func (api *IntegrationsService) DeleteMailgunIntegrationById(ctx context.Context, id string) (*shared.MagicbellProjectClientResponse[any], *shared.MagicbellProjectClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("DELETE").
		WithPath("/integrations/mailgun/{id}").
		WithConfig(config).
		AddPathParam("id", id).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[any](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewMagicbellProjectClientError[any](err)
	}

	return shared.NewMagicbellProjectClientResponse[any](resp), nil
}

// Retrieves the current ping_email integration configurations for a specific integration type in the project. Returns configuration details and status information.
func (api *IntegrationsService) GetPingEmailIntegration(ctx context.Context) (*shared.MagicbellProjectClientResponse[ArrayOfPingConfigObjects], *shared.MagicbellProjectClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/integrations/ping_email").
		WithConfig(config).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[ArrayOfPingConfigObjects](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewMagicbellProjectClientError[ArrayOfPingConfigObjects](err)
	}

	return shared.NewMagicbellProjectClientResponse[ArrayOfPingConfigObjects](resp), nil
}

// Creates or updates a ping_email integration for the project. Only administrators can configure integrations.
func (api *IntegrationsService) SavePingEmailIntegration(ctx context.Context, pingConfig PingConfig) (*shared.MagicbellProjectClientResponse[PingConfig], *shared.MagicbellProjectClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("PUT").
		WithPath("/integrations/ping_email").
		WithConfig(config).
		WithBody(pingConfig).
		AddHeader("CONTENT-TYPE", "application/json").
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[PingConfig](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewMagicbellProjectClientError[PingConfig](err)
	}

	return shared.NewMagicbellProjectClientResponse[PingConfig](resp), nil
}

// Removes a ping_email integration configuration from the project. This will disable the integration's functionality within the project.
func (api *IntegrationsService) DeletePingEmailIntegration(ctx context.Context) (*shared.MagicbellProjectClientResponse[any], *shared.MagicbellProjectClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("DELETE").
		WithPath("/integrations/ping_email").
		WithConfig(config).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[any](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewMagicbellProjectClientError[any](err)
	}

	return shared.NewMagicbellProjectClientResponse[any](resp), nil
}

// Removes a specific ping_email integration instance by ID from the project.
func (api *IntegrationsService) DeletePingEmailIntegrationById(ctx context.Context, id string) (*shared.MagicbellProjectClientResponse[any], *shared.MagicbellProjectClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("DELETE").
		WithPath("/integrations/ping_email/{id}").
		WithConfig(config).
		AddPathParam("id", id).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[any](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewMagicbellProjectClientError[any](err)
	}

	return shared.NewMagicbellProjectClientResponse[any](resp), nil
}

// Retrieves the current sendgrid integration configurations for a specific integration type in the project. Returns configuration details and status information.
func (api *IntegrationsService) GetSendgridIntegration(ctx context.Context) (*shared.MagicbellProjectClientResponse[ArrayOfSendgridConfigObjects], *shared.MagicbellProjectClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/integrations/sendgrid").
		WithConfig(config).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[ArrayOfSendgridConfigObjects](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewMagicbellProjectClientError[ArrayOfSendgridConfigObjects](err)
	}

	return shared.NewMagicbellProjectClientResponse[ArrayOfSendgridConfigObjects](resp), nil
}

// Creates or updates a sendgrid integration for the project. Only administrators can configure integrations.
func (api *IntegrationsService) SaveSendgridIntegration(ctx context.Context, sendgridConfig SendgridConfig) (*shared.MagicbellProjectClientResponse[SendgridConfig], *shared.MagicbellProjectClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("PUT").
		WithPath("/integrations/sendgrid").
		WithConfig(config).
		WithBody(sendgridConfig).
		AddHeader("CONTENT-TYPE", "application/json").
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[SendgridConfig](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewMagicbellProjectClientError[SendgridConfig](err)
	}

	return shared.NewMagicbellProjectClientResponse[SendgridConfig](resp), nil
}

// Removes a sendgrid integration configuration from the project. This will disable the integration's functionality within the project.
func (api *IntegrationsService) DeleteSendgridIntegration(ctx context.Context) (*shared.MagicbellProjectClientResponse[any], *shared.MagicbellProjectClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("DELETE").
		WithPath("/integrations/sendgrid").
		WithConfig(config).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[any](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewMagicbellProjectClientError[any](err)
	}

	return shared.NewMagicbellProjectClientResponse[any](resp), nil
}

// Removes a specific sendgrid integration instance by ID from the project.
func (api *IntegrationsService) DeleteSendgridIntegrationById(ctx context.Context, id string) (*shared.MagicbellProjectClientResponse[any], *shared.MagicbellProjectClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("DELETE").
		WithPath("/integrations/sendgrid/{id}").
		WithConfig(config).
		AddPathParam("id", id).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[any](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewMagicbellProjectClientError[any](err)
	}

	return shared.NewMagicbellProjectClientResponse[any](resp), nil
}

// Retrieves the current ses integration configurations for a specific integration type in the project. Returns configuration details and status information.
func (api *IntegrationsService) GetSesIntegration(ctx context.Context) (*shared.MagicbellProjectClientResponse[ArrayOfSesConfigObjects], *shared.MagicbellProjectClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/integrations/ses").
		WithConfig(config).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[ArrayOfSesConfigObjects](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewMagicbellProjectClientError[ArrayOfSesConfigObjects](err)
	}

	return shared.NewMagicbellProjectClientResponse[ArrayOfSesConfigObjects](resp), nil
}

// Creates or updates a ses integration for the project. Only administrators can configure integrations.
func (api *IntegrationsService) SaveSesIntegration(ctx context.Context, sesConfig SesConfig) (*shared.MagicbellProjectClientResponse[SesConfig], *shared.MagicbellProjectClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("PUT").
		WithPath("/integrations/ses").
		WithConfig(config).
		WithBody(sesConfig).
		AddHeader("CONTENT-TYPE", "application/json").
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[SesConfig](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewMagicbellProjectClientError[SesConfig](err)
	}

	return shared.NewMagicbellProjectClientResponse[SesConfig](resp), nil
}

// Removes a ses integration configuration from the project. This will disable the integration's functionality within the project.
func (api *IntegrationsService) DeleteSesIntegration(ctx context.Context) (*shared.MagicbellProjectClientResponse[any], *shared.MagicbellProjectClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("DELETE").
		WithPath("/integrations/ses").
		WithConfig(config).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[any](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewMagicbellProjectClientError[any](err)
	}

	return shared.NewMagicbellProjectClientResponse[any](resp), nil
}

// Removes a specific ses integration instance by ID from the project.
func (api *IntegrationsService) DeleteSesIntegrationById(ctx context.Context, id string) (*shared.MagicbellProjectClientResponse[any], *shared.MagicbellProjectClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("DELETE").
		WithPath("/integrations/ses/{id}").
		WithConfig(config).
		AddPathParam("id", id).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[any](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewMagicbellProjectClientError[any](err)
	}

	return shared.NewMagicbellProjectClientResponse[any](resp), nil
}

// Retrieves the current slack integration configurations for a specific integration type in the project. Returns configuration details and status information.
func (api *IntegrationsService) GetSlackIntegration(ctx context.Context) (*shared.MagicbellProjectClientResponse[ArrayOfSlackConfigObjects], *shared.MagicbellProjectClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/integrations/slack").
		WithConfig(config).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[ArrayOfSlackConfigObjects](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewMagicbellProjectClientError[ArrayOfSlackConfigObjects](err)
	}

	return shared.NewMagicbellProjectClientResponse[ArrayOfSlackConfigObjects](resp), nil
}

// Creates or updates a slack integration for the project. Only administrators can configure integrations.
func (api *IntegrationsService) SaveSlackIntegration(ctx context.Context, slackConfig SlackConfig) (*shared.MagicbellProjectClientResponse[SlackConfig], *shared.MagicbellProjectClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("PUT").
		WithPath("/integrations/slack").
		WithConfig(config).
		WithBody(slackConfig).
		AddHeader("CONTENT-TYPE", "application/json").
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[SlackConfig](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewMagicbellProjectClientError[SlackConfig](err)
	}

	return shared.NewMagicbellProjectClientResponse[SlackConfig](resp), nil
}

// Removes a slack integration configuration from the project. This will disable the integration's functionality within the project.
func (api *IntegrationsService) DeleteSlackIntegration(ctx context.Context) (*shared.MagicbellProjectClientResponse[any], *shared.MagicbellProjectClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("DELETE").
		WithPath("/integrations/slack").
		WithConfig(config).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[any](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewMagicbellProjectClientError[any](err)
	}

	return shared.NewMagicbellProjectClientResponse[any](resp), nil
}

// Removes a specific slack integration instance by ID from the project.
func (api *IntegrationsService) DeleteSlackIntegrationById(ctx context.Context, id string) (*shared.MagicbellProjectClientResponse[any], *shared.MagicbellProjectClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("DELETE").
		WithPath("/integrations/slack/{id}").
		WithConfig(config).
		AddPathParam("id", id).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[any](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewMagicbellProjectClientError[any](err)
	}

	return shared.NewMagicbellProjectClientResponse[any](resp), nil
}

// Retrieves the current stripe integration configurations for a specific integration type in the project. Returns configuration details and status information.
func (api *IntegrationsService) GetStripeIntegration(ctx context.Context) (*shared.MagicbellProjectClientResponse[ArrayOfStripeConfigObjects], *shared.MagicbellProjectClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/integrations/stripe").
		WithConfig(config).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[ArrayOfStripeConfigObjects](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewMagicbellProjectClientError[ArrayOfStripeConfigObjects](err)
	}

	return shared.NewMagicbellProjectClientResponse[ArrayOfStripeConfigObjects](resp), nil
}

// Creates or updates a stripe integration for the project. Only administrators can configure integrations.
func (api *IntegrationsService) SaveStripeIntegration(ctx context.Context, stripeConfig StripeConfig) (*shared.MagicbellProjectClientResponse[StripeConfig], *shared.MagicbellProjectClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("PUT").
		WithPath("/integrations/stripe").
		WithConfig(config).
		WithBody(stripeConfig).
		AddHeader("CONTENT-TYPE", "application/json").
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[StripeConfig](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewMagicbellProjectClientError[StripeConfig](err)
	}

	return shared.NewMagicbellProjectClientResponse[StripeConfig](resp), nil
}

// Removes a stripe integration configuration from the project. This will disable the integration's functionality within the project.
func (api *IntegrationsService) DeleteStripeIntegration(ctx context.Context) (*shared.MagicbellProjectClientResponse[any], *shared.MagicbellProjectClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("DELETE").
		WithPath("/integrations/stripe").
		WithConfig(config).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[any](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewMagicbellProjectClientError[any](err)
	}

	return shared.NewMagicbellProjectClientResponse[any](resp), nil
}

// Removes a specific stripe integration instance by ID from the project.
func (api *IntegrationsService) DeleteStripeIntegrationById(ctx context.Context, id string) (*shared.MagicbellProjectClientResponse[any], *shared.MagicbellProjectClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("DELETE").
		WithPath("/integrations/stripe/{id}").
		WithConfig(config).
		AddPathParam("id", id).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[any](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewMagicbellProjectClientError[any](err)
	}

	return shared.NewMagicbellProjectClientResponse[any](resp), nil
}

// Retrieves the current templates integration configurations for a specific integration type in the project. Returns configuration details and status information.
func (api *IntegrationsService) GetTemplatesIntegration(ctx context.Context) (*shared.MagicbellProjectClientResponse[ArrayOfTemplatesConfigObjects], *shared.MagicbellProjectClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/integrations/templates").
		WithConfig(config).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[ArrayOfTemplatesConfigObjects](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewMagicbellProjectClientError[ArrayOfTemplatesConfigObjects](err)
	}

	return shared.NewMagicbellProjectClientResponse[ArrayOfTemplatesConfigObjects](resp), nil
}

// Creates or updates a templates integration for the project. Only administrators can configure integrations.
func (api *IntegrationsService) SaveTemplatesIntegration(ctx context.Context) (*shared.MagicbellProjectClientResponse[[]byte], *shared.MagicbellProjectClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("PUT").
		WithPath("/integrations/templates").
		WithConfig(config).
		AddHeader("CONTENT-TYPE", "application/json").
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[[]byte](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewMagicbellProjectClientError[[]byte](err)
	}

	return shared.NewMagicbellProjectClientResponse[[]byte](resp), nil
}

// Removes a templates integration configuration from the project. This will disable the integration's functionality within the project.
func (api *IntegrationsService) DeleteTemplatesIntegration(ctx context.Context) (*shared.MagicbellProjectClientResponse[any], *shared.MagicbellProjectClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("DELETE").
		WithPath("/integrations/templates").
		WithConfig(config).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[any](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewMagicbellProjectClientError[any](err)
	}

	return shared.NewMagicbellProjectClientResponse[any](resp), nil
}

// Removes a specific templates integration instance by ID from the project.
func (api *IntegrationsService) DeleteTemplatesIntegrationById(ctx context.Context, id string) (*shared.MagicbellProjectClientResponse[any], *shared.MagicbellProjectClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("DELETE").
		WithPath("/integrations/templates/{id}").
		WithConfig(config).
		AddPathParam("id", id).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[any](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewMagicbellProjectClientError[any](err)
	}

	return shared.NewMagicbellProjectClientResponse[any](resp), nil
}

// Retrieves the current twilio integration configurations for a specific integration type in the project. Returns configuration details and status information.
func (api *IntegrationsService) GetTwilioIntegration(ctx context.Context) (*shared.MagicbellProjectClientResponse[ArrayOfTwilioConfigObjects], *shared.MagicbellProjectClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/integrations/twilio").
		WithConfig(config).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[ArrayOfTwilioConfigObjects](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewMagicbellProjectClientError[ArrayOfTwilioConfigObjects](err)
	}

	return shared.NewMagicbellProjectClientResponse[ArrayOfTwilioConfigObjects](resp), nil
}

// Creates or updates a twilio integration for the project. Only administrators can configure integrations.
func (api *IntegrationsService) SaveTwilioIntegration(ctx context.Context, twilioConfig TwilioConfig) (*shared.MagicbellProjectClientResponse[TwilioConfig], *shared.MagicbellProjectClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("PUT").
		WithPath("/integrations/twilio").
		WithConfig(config).
		WithBody(twilioConfig).
		AddHeader("CONTENT-TYPE", "application/json").
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[TwilioConfig](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewMagicbellProjectClientError[TwilioConfig](err)
	}

	return shared.NewMagicbellProjectClientResponse[TwilioConfig](resp), nil
}

// Removes a twilio integration configuration from the project. This will disable the integration's functionality within the project.
func (api *IntegrationsService) DeleteTwilioIntegration(ctx context.Context) (*shared.MagicbellProjectClientResponse[any], *shared.MagicbellProjectClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("DELETE").
		WithPath("/integrations/twilio").
		WithConfig(config).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[any](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewMagicbellProjectClientError[any](err)
	}

	return shared.NewMagicbellProjectClientResponse[any](resp), nil
}

// Removes a specific twilio integration instance by ID from the project.
func (api *IntegrationsService) DeleteTwilioIntegrationById(ctx context.Context, id string) (*shared.MagicbellProjectClientResponse[any], *shared.MagicbellProjectClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("DELETE").
		WithPath("/integrations/twilio/{id}").
		WithConfig(config).
		AddPathParam("id", id).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[any](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewMagicbellProjectClientError[any](err)
	}

	return shared.NewMagicbellProjectClientResponse[any](resp), nil
}

// Retrieves the current web_push integration configurations for a specific integration type in the project. Returns configuration details and status information.
func (api *IntegrationsService) GetWebPushIntegration(ctx context.Context) (*shared.MagicbellProjectClientResponse[ArrayOfWebpushConfigObjects], *shared.MagicbellProjectClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/integrations/web_push").
		WithConfig(config).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[ArrayOfWebpushConfigObjects](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewMagicbellProjectClientError[ArrayOfWebpushConfigObjects](err)
	}

	return shared.NewMagicbellProjectClientResponse[ArrayOfWebpushConfigObjects](resp), nil
}

// Creates or updates a web_push integration for the project. Only administrators can configure integrations.
func (api *IntegrationsService) SaveWebPushIntegration(ctx context.Context, webpushConfig WebpushConfig) (*shared.MagicbellProjectClientResponse[WebpushConfig], *shared.MagicbellProjectClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("PUT").
		WithPath("/integrations/web_push").
		WithConfig(config).
		WithBody(webpushConfig).
		AddHeader("CONTENT-TYPE", "application/json").
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[WebpushConfig](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewMagicbellProjectClientError[WebpushConfig](err)
	}

	return shared.NewMagicbellProjectClientResponse[WebpushConfig](resp), nil
}

// Removes a web_push integration configuration from the project. This will disable the integration's functionality within the project.
func (api *IntegrationsService) DeleteWebPushIntegration(ctx context.Context) (*shared.MagicbellProjectClientResponse[any], *shared.MagicbellProjectClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("DELETE").
		WithPath("/integrations/web_push").
		WithConfig(config).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[any](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewMagicbellProjectClientError[any](err)
	}

	return shared.NewMagicbellProjectClientResponse[any](resp), nil
}

// Removes a specific web_push integration instance by ID from the project.
func (api *IntegrationsService) DeleteWebPushIntegrationById(ctx context.Context, id string) (*shared.MagicbellProjectClientResponse[any], *shared.MagicbellProjectClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("DELETE").
		WithPath("/integrations/web_push/{id}").
		WithConfig(config).
		AddPathParam("id", id).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[any](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewMagicbellProjectClientError[any](err)
	}

	return shared.NewMagicbellProjectClientResponse[any](resp), nil
}
