package integrations

import (
	"context"
	restClient "github.com/magicbell/magicbell-go/pkg/user-client/internal/clients/rest"
	"github.com/magicbell/magicbell-go/pkg/user-client/internal/clients/rest/httptransport"
	"github.com/magicbell/magicbell-go/pkg/user-client/internal/configmanager"
	"github.com/magicbell/magicbell-go/pkg/user-client/clientconfig"
	"github.com/magicbell/magicbell-go/pkg/user-client/shared"
	"time"
)

type IntegrationsService struct {
	manager *configmanager.ConfigManager
}

func NewIntegrationsService() *IntegrationsService {
	return &IntegrationsService{
		manager: configmanager.NewConfigManager(clientconfig.Config{}),
	}
}

func (api *IntegrationsService) WithConfigManager(manager *configmanager.ConfigManager) *IntegrationsService {
	api.manager = manager
	return api
}

func (api *IntegrationsService) getConfig() *clientconfig.Config {
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

// Creates a new installation of a inbox integration for a user. This endpoint is used when an integration needs to be set up with user-specific credentials or configuration.
func (api *IntegrationsService) SaveInboxInstallation(ctx context.Context, inboxConfigPayload InboxConfigPayload) (*shared.ClientResponse[InboxConfigPayload], *shared.ClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("POST").
		WithPath("/integrations/inbox/installations").
		WithConfig(config).
		WithBody(inboxConfigPayload).
		AddHeader("CONTENT-TYPE", "application/json").
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[InboxConfigPayload](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewClientError[InboxConfigPayload](err)
	}

	return shared.NewClientResponse[InboxConfigPayload](resp), nil
}

// Initiates the installation flow for a inbox integration. This is the first step in a multi-step installation process where user authorization or external service configuration may be required.
func (api *IntegrationsService) StartInboxInstallation(ctx context.Context) (*shared.ClientResponse[InboxConfigPayload], *shared.ClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("POST").
		WithPath("/integrations/inbox/installations/start").
		WithConfig(config).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[InboxConfigPayload](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewClientError[InboxConfigPayload](err)
	}

	return shared.NewClientResponse[InboxConfigPayload](resp), nil
}

// Creates a new installation of a slack integration for a user. This endpoint is used when an integration needs to be set up with user-specific credentials or configuration.
func (api *IntegrationsService) SaveSlackInstallation(ctx context.Context, slackInstallation SlackInstallation) (*shared.ClientResponse[SlackInstallation], *shared.ClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("POST").
		WithPath("/integrations/slack/installations").
		WithConfig(config).
		WithBody(slackInstallation).
		AddHeader("CONTENT-TYPE", "application/json").
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[SlackInstallation](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewClientError[SlackInstallation](err)
	}

	return shared.NewClientResponse[SlackInstallation](resp), nil
}

// Completes the installation flow for a slack integration. This endpoint is typically called after the user has completed any required authorization steps with slack.
func (api *IntegrationsService) FinishSlackInstallation(ctx context.Context, slackFinishInstallResponse SlackFinishInstallResponse) (*shared.ClientResponse[SlackInstallation], *shared.ClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("POST").
		WithPath("/integrations/slack/installations/finish").
		WithConfig(config).
		WithBody(slackFinishInstallResponse).
		AddHeader("CONTENT-TYPE", "application/json").
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[SlackInstallation](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewClientError[SlackInstallation](err)
	}

	return shared.NewClientResponse[SlackInstallation](resp), nil
}

// Initiates the installation flow for a slack integration. This is the first step in a multi-step installation process where user authorization or external service configuration may be required.
func (api *IntegrationsService) StartSlackInstallation(ctx context.Context, slackStartInstall SlackStartInstall) (*shared.ClientResponse[SlackStartInstallResponseContent], *shared.ClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("POST").
		WithPath("/integrations/slack/installations/start").
		WithConfig(config).
		WithBody(slackStartInstall).
		AddHeader("CONTENT-TYPE", "application/json").
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[SlackStartInstallResponseContent](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewClientError[SlackStartInstallResponseContent](err)
	}

	return shared.NewClientResponse[SlackStartInstallResponseContent](resp), nil
}

// Creates a new installation of a templates integration for a user. This endpoint is used when an integration needs to be set up with user-specific credentials or configuration.
func (api *IntegrationsService) SaveTemplatesInstallation(ctx context.Context, templatesInstallation TemplatesInstallation) (*shared.ClientResponse[TemplatesInstallation], *shared.ClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("POST").
		WithPath("/integrations/templates/installations").
		WithConfig(config).
		WithBody(templatesInstallation).
		AddHeader("CONTENT-TYPE", "application/json").
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[TemplatesInstallation](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewClientError[TemplatesInstallation](err)
	}

	return shared.NewClientResponse[TemplatesInstallation](resp), nil
}

// Creates a new installation of a web_push integration for a user. This endpoint is used when an integration needs to be set up with user-specific credentials or configuration.
func (api *IntegrationsService) SaveWebPushInstallation(ctx context.Context, webPushTokenPayload shared.WebPushTokenPayload) (*shared.ClientResponse[shared.WebPushTokenPayload], *shared.ClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("POST").
		WithPath("/integrations/web_push/installations").
		WithConfig(config).
		WithBody(webPushTokenPayload).
		AddHeader("CONTENT-TYPE", "application/json").
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[shared.WebPushTokenPayload](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewClientError[shared.WebPushTokenPayload](err)
	}

	return shared.NewClientResponse[shared.WebPushTokenPayload](resp), nil
}

// Initiates the installation flow for a web_push integration. This is the first step in a multi-step installation process where user authorization or external service configuration may be required.
func (api *IntegrationsService) StartWebPushInstallation(ctx context.Context) (*shared.ClientResponse[WebPushStartInstallationResponse], *shared.ClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("POST").
		WithPath("/integrations/web_push/installations/start").
		WithConfig(config).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[WebPushStartInstallationResponse](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewClientError[WebPushStartInstallationResponse](err)
	}

	return shared.NewClientResponse[WebPushStartInstallationResponse](resp), nil
}
