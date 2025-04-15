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

func NewIntegrationsService() *IntegrationsService {
	return &IntegrationsService{
		manager: configmanager.NewConfigManager(magicbellprojectclientconfig.Config{}),
	}
}

func (api *IntegrationsService) WithConfigManager(manager *configmanager.ConfigManager) *IntegrationsService {
	api.manager = manager
	return api
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
func (api *IntegrationsService) ListIntegrations(ctx context.Context, params ListIntegrationsRequestParams) (*shared.MagicbellProjectClientResponse[IntegrationConfigCollection], *shared.MagicbellProjectClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/integrations").
		WithConfig(config).
		WithOptions(params).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[IntegrationConfigCollection](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewMagicbellProjectClientError[IntegrationConfigCollection](err)
	}

	return shared.NewMagicbellProjectClientResponse[IntegrationConfigCollection](resp), nil
}

// Retrieves the current apns integration configurations for a specific integration type in the project. Returns configuration details and status information.
func (api *IntegrationsService) GetApnsIntegration(ctx context.Context) (*shared.MagicbellProjectClientResponse[ApnsConfigCollection], *shared.MagicbellProjectClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/integrations/apns").
		WithConfig(config).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[ApnsConfigCollection](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewMagicbellProjectClientError[ApnsConfigCollection](err)
	}

	return shared.NewMagicbellProjectClientResponse[ApnsConfigCollection](resp), nil
}

// Creates or updates a apns integration for the project. Only administrators can configure integrations.
func (api *IntegrationsService) SaveApnsIntegration(ctx context.Context, apnsConfigPayload ApnsConfigPayload) (*shared.MagicbellProjectClientResponse[ApnsConfigPayload], *shared.MagicbellProjectClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("PUT").
		WithPath("/integrations/apns").
		WithConfig(config).
		WithBody(apnsConfigPayload).
		AddHeader("CONTENT-TYPE", "application/json").
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[ApnsConfigPayload](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewMagicbellProjectClientError[ApnsConfigPayload](err)
	}

	return shared.NewMagicbellProjectClientResponse[ApnsConfigPayload](resp), nil
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
func (api *IntegrationsService) GetAwssnsIntegration(ctx context.Context) (*shared.MagicbellProjectClientResponse[AwssnsConfigCollection], *shared.MagicbellProjectClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/integrations/awssns").
		WithConfig(config).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[AwssnsConfigCollection](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewMagicbellProjectClientError[AwssnsConfigCollection](err)
	}

	return shared.NewMagicbellProjectClientResponse[AwssnsConfigCollection](resp), nil
}

// Creates or updates a awssns integration for the project. Only administrators can configure integrations.
func (api *IntegrationsService) SaveAwssnsIntegration(ctx context.Context, awssnsConfigPayload AwssnsConfigPayload) (*shared.MagicbellProjectClientResponse[AwssnsConfigPayload], *shared.MagicbellProjectClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("PUT").
		WithPath("/integrations/awssns").
		WithConfig(config).
		WithBody(awssnsConfigPayload).
		AddHeader("CONTENT-TYPE", "application/json").
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[AwssnsConfigPayload](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewMagicbellProjectClientError[AwssnsConfigPayload](err)
	}

	return shared.NewMagicbellProjectClientResponse[AwssnsConfigPayload](resp), nil
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

// Retrieves the current eventsource integration configurations for a specific integration type in the project. Returns configuration details and status information.
func (api *IntegrationsService) GetEventsourceIntegration(ctx context.Context) (*shared.MagicbellProjectClientResponse[EventSourceConfigCollection], *shared.MagicbellProjectClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/integrations/eventsource").
		WithConfig(config).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[EventSourceConfigCollection](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewMagicbellProjectClientError[EventSourceConfigCollection](err)
	}

	return shared.NewMagicbellProjectClientResponse[EventSourceConfigCollection](resp), nil
}

// Creates or updates a eventsource integration for the project. Only administrators can configure integrations.
func (api *IntegrationsService) SaveEventsourceIntegration(ctx context.Context, eventSourceConfigPayload EventSourceConfigPayload) (*shared.MagicbellProjectClientResponse[EventSourceConfigPayload], *shared.MagicbellProjectClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("PUT").
		WithPath("/integrations/eventsource").
		WithConfig(config).
		WithBody(eventSourceConfigPayload).
		AddHeader("CONTENT-TYPE", "application/json").
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[EventSourceConfigPayload](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewMagicbellProjectClientError[EventSourceConfigPayload](err)
	}

	return shared.NewMagicbellProjectClientResponse[EventSourceConfigPayload](resp), nil
}

// Removes a eventsource integration configuration from the project. This will disable the integration's functionality within the project.
func (api *IntegrationsService) DeleteEventsourceIntegration(ctx context.Context) (*shared.MagicbellProjectClientResponse[any], *shared.MagicbellProjectClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("DELETE").
		WithPath("/integrations/eventsource").
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

// Removes a specific eventsource integration instance by ID from the project.
func (api *IntegrationsService) DeleteEventsourceIntegrationById(ctx context.Context, id string) (*shared.MagicbellProjectClientResponse[any], *shared.MagicbellProjectClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("DELETE").
		WithPath("/integrations/eventsource/{id}").
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
func (api *IntegrationsService) GetExpoIntegration(ctx context.Context) (*shared.MagicbellProjectClientResponse[ExpoConfigCollection], *shared.MagicbellProjectClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/integrations/expo").
		WithConfig(config).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[ExpoConfigCollection](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewMagicbellProjectClientError[ExpoConfigCollection](err)
	}

	return shared.NewMagicbellProjectClientResponse[ExpoConfigCollection](resp), nil
}

// Creates or updates a expo integration for the project. Only administrators can configure integrations.
func (api *IntegrationsService) SaveExpoIntegration(ctx context.Context, expoConfigPayload ExpoConfigPayload) (*shared.MagicbellProjectClientResponse[ExpoConfigPayload], *shared.MagicbellProjectClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("PUT").
		WithPath("/integrations/expo").
		WithConfig(config).
		WithBody(expoConfigPayload).
		AddHeader("CONTENT-TYPE", "application/json").
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[ExpoConfigPayload](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewMagicbellProjectClientError[ExpoConfigPayload](err)
	}

	return shared.NewMagicbellProjectClientResponse[ExpoConfigPayload](resp), nil
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
func (api *IntegrationsService) GetFcmIntegration(ctx context.Context) (*shared.MagicbellProjectClientResponse[FcmConfigCollection], *shared.MagicbellProjectClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/integrations/fcm").
		WithConfig(config).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[FcmConfigCollection](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewMagicbellProjectClientError[FcmConfigCollection](err)
	}

	return shared.NewMagicbellProjectClientResponse[FcmConfigCollection](resp), nil
}

// Creates or updates a fcm integration for the project. Only administrators can configure integrations.
func (api *IntegrationsService) SaveFcmIntegration(ctx context.Context, fcmConfigPayload FcmConfigPayload) (*shared.MagicbellProjectClientResponse[FcmConfigPayload], *shared.MagicbellProjectClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("PUT").
		WithPath("/integrations/fcm").
		WithConfig(config).
		WithBody(fcmConfigPayload).
		AddHeader("CONTENT-TYPE", "application/json").
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[FcmConfigPayload](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewMagicbellProjectClientError[FcmConfigPayload](err)
	}

	return shared.NewMagicbellProjectClientResponse[FcmConfigPayload](resp), nil
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
func (api *IntegrationsService) GetGithubIntegration(ctx context.Context) (*shared.MagicbellProjectClientResponse[GithubConfigCollection], *shared.MagicbellProjectClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/integrations/github").
		WithConfig(config).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[GithubConfigCollection](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewMagicbellProjectClientError[GithubConfigCollection](err)
	}

	return shared.NewMagicbellProjectClientResponse[GithubConfigCollection](resp), nil
}

// Creates or updates a github integration for the project. Only administrators can configure integrations.
func (api *IntegrationsService) SaveGithubIntegration(ctx context.Context, githubConfigPayload GithubConfigPayload) (*shared.MagicbellProjectClientResponse[GithubConfigPayload], *shared.MagicbellProjectClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("PUT").
		WithPath("/integrations/github").
		WithConfig(config).
		WithBody(githubConfigPayload).
		AddHeader("CONTENT-TYPE", "application/json").
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[GithubConfigPayload](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewMagicbellProjectClientError[GithubConfigPayload](err)
	}

	return shared.NewMagicbellProjectClientResponse[GithubConfigPayload](resp), nil
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
func (api *IntegrationsService) GetInboxIntegration(ctx context.Context) (*shared.MagicbellProjectClientResponse[InboxConfigCollection], *shared.MagicbellProjectClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/integrations/inbox").
		WithConfig(config).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[InboxConfigCollection](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewMagicbellProjectClientError[InboxConfigCollection](err)
	}

	return shared.NewMagicbellProjectClientResponse[InboxConfigCollection](resp), nil
}

// Creates or updates a inbox integration for the project. Only administrators can configure integrations.
func (api *IntegrationsService) SaveInboxIntegration(ctx context.Context, inboxConfigPayload InboxConfigPayload) (*shared.MagicbellProjectClientResponse[InboxConfigPayload], *shared.MagicbellProjectClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("PUT").
		WithPath("/integrations/inbox").
		WithConfig(config).
		WithBody(inboxConfigPayload).
		AddHeader("CONTENT-TYPE", "application/json").
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[InboxConfigPayload](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewMagicbellProjectClientError[InboxConfigPayload](err)
	}

	return shared.NewMagicbellProjectClientResponse[InboxConfigPayload](resp), nil
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
func (api *IntegrationsService) GetMailgunIntegration(ctx context.Context) (*shared.MagicbellProjectClientResponse[MailgunConfigCollection], *shared.MagicbellProjectClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/integrations/mailgun").
		WithConfig(config).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[MailgunConfigCollection](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewMagicbellProjectClientError[MailgunConfigCollection](err)
	}

	return shared.NewMagicbellProjectClientResponse[MailgunConfigCollection](resp), nil
}

// Creates or updates a mailgun integration for the project. Only administrators can configure integrations.
func (api *IntegrationsService) SaveMailgunIntegration(ctx context.Context, mailgunConfigPayload MailgunConfigPayload) (*shared.MagicbellProjectClientResponse[MailgunConfigPayload], *shared.MagicbellProjectClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("PUT").
		WithPath("/integrations/mailgun").
		WithConfig(config).
		WithBody(mailgunConfigPayload).
		AddHeader("CONTENT-TYPE", "application/json").
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[MailgunConfigPayload](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewMagicbellProjectClientError[MailgunConfigPayload](err)
	}

	return shared.NewMagicbellProjectClientResponse[MailgunConfigPayload](resp), nil
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
func (api *IntegrationsService) GetPingEmailIntegration(ctx context.Context) (*shared.MagicbellProjectClientResponse[PingConfigCollection], *shared.MagicbellProjectClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/integrations/ping_email").
		WithConfig(config).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[PingConfigCollection](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewMagicbellProjectClientError[PingConfigCollection](err)
	}

	return shared.NewMagicbellProjectClientResponse[PingConfigCollection](resp), nil
}

// Creates or updates a ping_email integration for the project. Only administrators can configure integrations.
func (api *IntegrationsService) SavePingEmailIntegration(ctx context.Context, pingConfigPayload PingConfigPayload) (*shared.MagicbellProjectClientResponse[PingConfigPayload], *shared.MagicbellProjectClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("PUT").
		WithPath("/integrations/ping_email").
		WithConfig(config).
		WithBody(pingConfigPayload).
		AddHeader("CONTENT-TYPE", "application/json").
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[PingConfigPayload](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewMagicbellProjectClientError[PingConfigPayload](err)
	}

	return shared.NewMagicbellProjectClientResponse[PingConfigPayload](resp), nil
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
func (api *IntegrationsService) GetSendgridIntegration(ctx context.Context) (*shared.MagicbellProjectClientResponse[SendgridConfigCollection], *shared.MagicbellProjectClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/integrations/sendgrid").
		WithConfig(config).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[SendgridConfigCollection](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewMagicbellProjectClientError[SendgridConfigCollection](err)
	}

	return shared.NewMagicbellProjectClientResponse[SendgridConfigCollection](resp), nil
}

// Creates or updates a sendgrid integration for the project. Only administrators can configure integrations.
func (api *IntegrationsService) SaveSendgridIntegration(ctx context.Context, sendgridConfigPayload SendgridConfigPayload) (*shared.MagicbellProjectClientResponse[SendgridConfigPayload], *shared.MagicbellProjectClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("PUT").
		WithPath("/integrations/sendgrid").
		WithConfig(config).
		WithBody(sendgridConfigPayload).
		AddHeader("CONTENT-TYPE", "application/json").
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[SendgridConfigPayload](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewMagicbellProjectClientError[SendgridConfigPayload](err)
	}

	return shared.NewMagicbellProjectClientResponse[SendgridConfigPayload](resp), nil
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
func (api *IntegrationsService) GetSesIntegration(ctx context.Context) (*shared.MagicbellProjectClientResponse[SesConfigCollection], *shared.MagicbellProjectClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/integrations/ses").
		WithConfig(config).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[SesConfigCollection](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewMagicbellProjectClientError[SesConfigCollection](err)
	}

	return shared.NewMagicbellProjectClientResponse[SesConfigCollection](resp), nil
}

// Creates or updates a ses integration for the project. Only administrators can configure integrations.
func (api *IntegrationsService) SaveSesIntegration(ctx context.Context, sesConfigPayload SesConfigPayload) (*shared.MagicbellProjectClientResponse[SesConfigPayload], *shared.MagicbellProjectClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("PUT").
		WithPath("/integrations/ses").
		WithConfig(config).
		WithBody(sesConfigPayload).
		AddHeader("CONTENT-TYPE", "application/json").
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[SesConfigPayload](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewMagicbellProjectClientError[SesConfigPayload](err)
	}

	return shared.NewMagicbellProjectClientResponse[SesConfigPayload](resp), nil
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
func (api *IntegrationsService) GetSlackIntegration(ctx context.Context) (*shared.MagicbellProjectClientResponse[SlackConfigCollection], *shared.MagicbellProjectClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/integrations/slack").
		WithConfig(config).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[SlackConfigCollection](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewMagicbellProjectClientError[SlackConfigCollection](err)
	}

	return shared.NewMagicbellProjectClientResponse[SlackConfigCollection](resp), nil
}

// Creates or updates a slack integration for the project. Only administrators can configure integrations.
func (api *IntegrationsService) SaveSlackIntegration(ctx context.Context, slackConfigPayload SlackConfigPayload) (*shared.MagicbellProjectClientResponse[SlackConfigPayload], *shared.MagicbellProjectClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("PUT").
		WithPath("/integrations/slack").
		WithConfig(config).
		WithBody(slackConfigPayload).
		AddHeader("CONTENT-TYPE", "application/json").
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[SlackConfigPayload](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewMagicbellProjectClientError[SlackConfigPayload](err)
	}

	return shared.NewMagicbellProjectClientResponse[SlackConfigPayload](resp), nil
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
func (api *IntegrationsService) GetStripeIntegration(ctx context.Context) (*shared.MagicbellProjectClientResponse[StripeConfigCollection], *shared.MagicbellProjectClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/integrations/stripe").
		WithConfig(config).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[StripeConfigCollection](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewMagicbellProjectClientError[StripeConfigCollection](err)
	}

	return shared.NewMagicbellProjectClientResponse[StripeConfigCollection](resp), nil
}

// Creates or updates a stripe integration for the project. Only administrators can configure integrations.
func (api *IntegrationsService) SaveStripeIntegration(ctx context.Context, stripeConfigPayload StripeConfigPayload) (*shared.MagicbellProjectClientResponse[StripeConfigPayload], *shared.MagicbellProjectClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("PUT").
		WithPath("/integrations/stripe").
		WithConfig(config).
		WithBody(stripeConfigPayload).
		AddHeader("CONTENT-TYPE", "application/json").
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[StripeConfigPayload](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewMagicbellProjectClientError[StripeConfigPayload](err)
	}

	return shared.NewMagicbellProjectClientResponse[StripeConfigPayload](resp), nil
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
func (api *IntegrationsService) GetTemplatesIntegration(ctx context.Context) (*shared.MagicbellProjectClientResponse[TemplatesConfigCollection], *shared.MagicbellProjectClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/integrations/templates").
		WithConfig(config).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[TemplatesConfigCollection](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewMagicbellProjectClientError[TemplatesConfigCollection](err)
	}

	return shared.NewMagicbellProjectClientResponse[TemplatesConfigCollection](resp), nil
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
func (api *IntegrationsService) GetTwilioIntegration(ctx context.Context) (*shared.MagicbellProjectClientResponse[TwilioConfigCollection], *shared.MagicbellProjectClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/integrations/twilio").
		WithConfig(config).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[TwilioConfigCollection](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewMagicbellProjectClientError[TwilioConfigCollection](err)
	}

	return shared.NewMagicbellProjectClientResponse[TwilioConfigCollection](resp), nil
}

// Creates or updates a twilio integration for the project. Only administrators can configure integrations.
func (api *IntegrationsService) SaveTwilioIntegration(ctx context.Context, twilioConfigPayload TwilioConfigPayload) (*shared.MagicbellProjectClientResponse[TwilioConfigPayload], *shared.MagicbellProjectClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("PUT").
		WithPath("/integrations/twilio").
		WithConfig(config).
		WithBody(twilioConfigPayload).
		AddHeader("CONTENT-TYPE", "application/json").
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[TwilioConfigPayload](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewMagicbellProjectClientError[TwilioConfigPayload](err)
	}

	return shared.NewMagicbellProjectClientResponse[TwilioConfigPayload](resp), nil
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
func (api *IntegrationsService) GetWebPushIntegration(ctx context.Context) (*shared.MagicbellProjectClientResponse[WebpushConfigCollection], *shared.MagicbellProjectClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/integrations/web_push").
		WithConfig(config).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[WebpushConfigCollection](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewMagicbellProjectClientError[WebpushConfigCollection](err)
	}

	return shared.NewMagicbellProjectClientResponse[WebpushConfigCollection](resp), nil
}

// Creates or updates a web_push integration for the project. Only administrators can configure integrations.
func (api *IntegrationsService) SaveWebPushIntegration(ctx context.Context, webpushConfigPayload WebpushConfigPayload) (*shared.MagicbellProjectClientResponse[WebpushConfigPayload], *shared.MagicbellProjectClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("PUT").
		WithPath("/integrations/web_push").
		WithConfig(config).
		WithBody(webpushConfigPayload).
		AddHeader("CONTENT-TYPE", "application/json").
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[WebpushConfigPayload](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewMagicbellProjectClientError[WebpushConfigPayload](err)
	}

	return shared.NewMagicbellProjectClientResponse[WebpushConfigPayload](resp), nil
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
