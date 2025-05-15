package channels

import (
	"context"
	restClient "github.com/magicbell/magicbell-go/pkg/project-client/internal/clients/rest"
	"github.com/magicbell/magicbell-go/pkg/project-client/internal/clients/rest/httptransport"
	"github.com/magicbell/magicbell-go/pkg/project-client/internal/configmanager"
	"github.com/magicbell/magicbell-go/pkg/project-client/projectclientconfig"
	"github.com/magicbell/magicbell-go/pkg/project-client/shared"
	"time"
)

type ChannelsService struct {
	manager *configmanager.ConfigManager
}

func NewChannelsService() *ChannelsService {
	return &ChannelsService{
		manager: configmanager.NewConfigManager(projectclientconfig.Config{}),
	}
}

func (api *ChannelsService) WithConfigManager(manager *configmanager.ConfigManager) *ChannelsService {
	api.manager = manager
	return api
}

func (api *ChannelsService) getConfig() *projectclientconfig.Config {
	return api.manager.GetChannels()
}

func (api *ChannelsService) SetBaseUrl(baseUrl string) {
	config := api.getConfig()
	config.SetBaseUrl(baseUrl)
}

func (api *ChannelsService) SetTimeout(timeout time.Duration) {
	config := api.getConfig()
	config.SetTimeout(timeout)
}

func (api *ChannelsService) SetAccessToken(accessToken string) {
	config := api.getConfig()
	config.SetAccessToken(accessToken)
}

func (api *ChannelsService) GetDeliveryconfig(ctx context.Context, params GetDeliveryconfigRequestParams) (*shared.ProjectClientResponse[CategoryDeliveryConfig], *shared.ProjectClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/channels/deliveryconfig").
		WithConfig(config).
		WithOptions(params).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[CategoryDeliveryConfig](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewProjectClientError[CategoryDeliveryConfig](err)
	}

	return shared.NewProjectClientResponse[CategoryDeliveryConfig](resp), nil
}

func (api *ChannelsService) SaveDeliveryconfig(ctx context.Context, categoryDeliveryConfig CategoryDeliveryConfig) (*shared.ProjectClientResponse[CategoryDeliveryConfig], *shared.ProjectClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("PUT").
		WithPath("/channels/deliveryconfig").
		WithConfig(config).
		WithBody(categoryDeliveryConfig).
		AddHeader("CONTENT-TYPE", "application/json").
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[CategoryDeliveryConfig](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewProjectClientError[CategoryDeliveryConfig](err)
	}

	return shared.NewProjectClientResponse[CategoryDeliveryConfig](resp), nil
}

// Lists all in_app tokens associated with a specific user. This endpoint is available to project administrators and returns a paginated list of tokens, including both active and revoked tokens.
func (api *ChannelsService) GetInAppInboxUserTokens(ctx context.Context, userId string, params GetInAppInboxUserTokensRequestParams) (*shared.ProjectClientResponse[InboxTokenResponseCollection], *shared.ProjectClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/users/{user_id}/channels/in_app/inbox/tokens").
		WithConfig(config).
		AddPathParam("user_id", userId).
		WithOptions(params).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[InboxTokenResponseCollection](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewProjectClientError[InboxTokenResponseCollection](err)
	}

	return shared.NewProjectClientResponse[InboxTokenResponseCollection](resp), nil
}

// Retrieves a specific in_app token by its ID for a given user. This endpoint is available to project administrators and requires project-level authentication. Use this to inspect token details including its status, creation date, and associated metadata.
func (api *ChannelsService) GetInAppInboxUserToken(ctx context.Context, userId string, tokenId string) (*shared.ProjectClientResponse[InboxTokenResponse], *shared.ProjectClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/users/{user_id}/channels/in_app/inbox/tokens/{token_id}").
		WithConfig(config).
		AddPathParam("user_id", userId).
		AddPathParam("token_id", tokenId).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[InboxTokenResponse](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewProjectClientError[InboxTokenResponse](err)
	}

	return shared.NewProjectClientResponse[InboxTokenResponse](resp), nil
}

// Revokes a specific user's in_app token. This endpoint is available to project administrators and permanently invalidates the specified token. Once revoked, the token can no longer be used to access channel features. This action cannot be undone.
func (api *ChannelsService) DiscardInAppInboxUserToken(ctx context.Context, userId string, tokenId string) (*shared.ProjectClientResponse[DiscardResult], *shared.ProjectClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("DELETE").
		WithPath("/users/{user_id}/channels/in_app/inbox/tokens/{token_id}").
		WithConfig(config).
		AddPathParam("user_id", userId).
		AddPathParam("token_id", tokenId).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[DiscardResult](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewProjectClientError[DiscardResult](err)
	}

	return shared.NewProjectClientResponse[DiscardResult](resp), nil
}

// Lists all mobile_push tokens associated with a specific user. This endpoint is available to project administrators and returns a paginated list of tokens, including both active and revoked tokens.
func (api *ChannelsService) GetMobilePushApnsUserTokens(ctx context.Context, userId string, params GetMobilePushApnsUserTokensRequestParams) (*shared.ProjectClientResponse[ApnsTokenCollection], *shared.ProjectClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/users/{user_id}/channels/mobile_push/apns/tokens").
		WithConfig(config).
		AddPathParam("user_id", userId).
		WithOptions(params).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[ApnsTokenCollection](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewProjectClientError[ApnsTokenCollection](err)
	}

	return shared.NewProjectClientResponse[ApnsTokenCollection](resp), nil
}

// Retrieves a specific mobile_push token by its ID for a given user. This endpoint is available to project administrators and requires project-level authentication. Use this to inspect token details including its status, creation date, and associated metadata.
func (api *ChannelsService) GetMobilePushApnsUserToken(ctx context.Context, userId string, tokenId string) (*shared.ProjectClientResponse[ApnsToken], *shared.ProjectClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/users/{user_id}/channels/mobile_push/apns/tokens/{token_id}").
		WithConfig(config).
		AddPathParam("user_id", userId).
		AddPathParam("token_id", tokenId).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[ApnsToken](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewProjectClientError[ApnsToken](err)
	}

	return shared.NewProjectClientResponse[ApnsToken](resp), nil
}

// Revokes a specific user's mobile_push token. This endpoint is available to project administrators and permanently invalidates the specified token. Once revoked, the token can no longer be used to access channel features. This action cannot be undone.
func (api *ChannelsService) DiscardMobilePushApnsUserToken(ctx context.Context, userId string, tokenId string) (*shared.ProjectClientResponse[DiscardResult], *shared.ProjectClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("DELETE").
		WithPath("/users/{user_id}/channels/mobile_push/apns/tokens/{token_id}").
		WithConfig(config).
		AddPathParam("user_id", userId).
		AddPathParam("token_id", tokenId).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[DiscardResult](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewProjectClientError[DiscardResult](err)
	}

	return shared.NewProjectClientResponse[DiscardResult](resp), nil
}

// Lists all mobile_push tokens associated with a specific user. This endpoint is available to project administrators and returns a paginated list of tokens, including both active and revoked tokens.
func (api *ChannelsService) GetMobilePushExpoUserTokens(ctx context.Context, userId string, params GetMobilePushExpoUserTokensRequestParams) (*shared.ProjectClientResponse[ExpoTokenCollection], *shared.ProjectClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/users/{user_id}/channels/mobile_push/expo/tokens").
		WithConfig(config).
		AddPathParam("user_id", userId).
		WithOptions(params).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[ExpoTokenCollection](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewProjectClientError[ExpoTokenCollection](err)
	}

	return shared.NewProjectClientResponse[ExpoTokenCollection](resp), nil
}

// Retrieves a specific mobile_push token by its ID for a given user. This endpoint is available to project administrators and requires project-level authentication. Use this to inspect token details including its status, creation date, and associated metadata.
func (api *ChannelsService) GetMobilePushExpoUserToken(ctx context.Context, userId string, tokenId string) (*shared.ProjectClientResponse[ExpoToken], *shared.ProjectClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/users/{user_id}/channels/mobile_push/expo/tokens/{token_id}").
		WithConfig(config).
		AddPathParam("user_id", userId).
		AddPathParam("token_id", tokenId).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[ExpoToken](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewProjectClientError[ExpoToken](err)
	}

	return shared.NewProjectClientResponse[ExpoToken](resp), nil
}

// Revokes a specific user's mobile_push token. This endpoint is available to project administrators and permanently invalidates the specified token. Once revoked, the token can no longer be used to access channel features. This action cannot be undone.
func (api *ChannelsService) DiscardMobilePushExpoUserToken(ctx context.Context, userId string, tokenId string) (*shared.ProjectClientResponse[DiscardResult], *shared.ProjectClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("DELETE").
		WithPath("/users/{user_id}/channels/mobile_push/expo/tokens/{token_id}").
		WithConfig(config).
		AddPathParam("user_id", userId).
		AddPathParam("token_id", tokenId).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[DiscardResult](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewProjectClientError[DiscardResult](err)
	}

	return shared.NewProjectClientResponse[DiscardResult](resp), nil
}

// Lists all mobile_push tokens associated with a specific user. This endpoint is available to project administrators and returns a paginated list of tokens, including both active and revoked tokens.
func (api *ChannelsService) GetMobilePushFcmUserTokens(ctx context.Context, userId string, params GetMobilePushFcmUserTokensRequestParams) (*shared.ProjectClientResponse[FcmTokenCollection], *shared.ProjectClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/users/{user_id}/channels/mobile_push/fcm/tokens").
		WithConfig(config).
		AddPathParam("user_id", userId).
		WithOptions(params).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[FcmTokenCollection](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewProjectClientError[FcmTokenCollection](err)
	}

	return shared.NewProjectClientResponse[FcmTokenCollection](resp), nil
}

// Retrieves a specific mobile_push token by its ID for a given user. This endpoint is available to project administrators and requires project-level authentication. Use this to inspect token details including its status, creation date, and associated metadata.
func (api *ChannelsService) GetMobilePushFcmUserToken(ctx context.Context, userId string, tokenId string) (*shared.ProjectClientResponse[FcmToken], *shared.ProjectClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/users/{user_id}/channels/mobile_push/fcm/tokens/{token_id}").
		WithConfig(config).
		AddPathParam("user_id", userId).
		AddPathParam("token_id", tokenId).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[FcmToken](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewProjectClientError[FcmToken](err)
	}

	return shared.NewProjectClientResponse[FcmToken](resp), nil
}

// Revokes a specific user's mobile_push token. This endpoint is available to project administrators and permanently invalidates the specified token. Once revoked, the token can no longer be used to access channel features. This action cannot be undone.
func (api *ChannelsService) DiscardMobilePushFcmUserToken(ctx context.Context, userId string, tokenId string) (*shared.ProjectClientResponse[DiscardResult], *shared.ProjectClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("DELETE").
		WithPath("/users/{user_id}/channels/mobile_push/fcm/tokens/{token_id}").
		WithConfig(config).
		AddPathParam("user_id", userId).
		AddPathParam("token_id", tokenId).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[DiscardResult](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewProjectClientError[DiscardResult](err)
	}

	return shared.NewProjectClientResponse[DiscardResult](resp), nil
}

// Lists all slack tokens associated with a specific user. This endpoint is available to project administrators and returns a paginated list of tokens, including both active and revoked tokens.
func (api *ChannelsService) GetSlackUserTokens(ctx context.Context, userId string, params GetSlackUserTokensRequestParams) (*shared.ProjectClientResponse[SlackTokenCollection], *shared.ProjectClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/users/{user_id}/channels/slack/tokens").
		WithConfig(config).
		AddPathParam("user_id", userId).
		WithOptions(params).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[SlackTokenCollection](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewProjectClientError[SlackTokenCollection](err)
	}

	return shared.NewProjectClientResponse[SlackTokenCollection](resp), nil
}

// Retrieves a specific slack token by its ID for a given user. This endpoint is available to project administrators and requires project-level authentication. Use this to inspect token details including its status, creation date, and associated metadata.
func (api *ChannelsService) GetSlackUserToken(ctx context.Context, userId string, tokenId string) (*shared.ProjectClientResponse[SlackToken], *shared.ProjectClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/users/{user_id}/channels/slack/tokens/{token_id}").
		WithConfig(config).
		AddPathParam("user_id", userId).
		AddPathParam("token_id", tokenId).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[SlackToken](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewProjectClientError[SlackToken](err)
	}

	return shared.NewProjectClientResponse[SlackToken](resp), nil
}

// Revokes a specific user's slack token. This endpoint is available to project administrators and permanently invalidates the specified token. Once revoked, the token can no longer be used to access channel features. This action cannot be undone.
func (api *ChannelsService) DiscardSlackUserToken(ctx context.Context, userId string, tokenId string) (*shared.ProjectClientResponse[DiscardResult], *shared.ProjectClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("DELETE").
		WithPath("/users/{user_id}/channels/slack/tokens/{token_id}").
		WithConfig(config).
		AddPathParam("user_id", userId).
		AddPathParam("token_id", tokenId).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[DiscardResult](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewProjectClientError[DiscardResult](err)
	}

	return shared.NewProjectClientResponse[DiscardResult](resp), nil
}

// Lists all teams tokens associated with a specific user. This endpoint is available to project administrators and returns a paginated list of tokens, including both active and revoked tokens.
func (api *ChannelsService) GetTeamsUserTokens(ctx context.Context, userId string, params GetTeamsUserTokensRequestParams) (*shared.ProjectClientResponse[TeamsTokenCollection], *shared.ProjectClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/users/{user_id}/channels/teams/tokens").
		WithConfig(config).
		AddPathParam("user_id", userId).
		WithOptions(params).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[TeamsTokenCollection](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewProjectClientError[TeamsTokenCollection](err)
	}

	return shared.NewProjectClientResponse[TeamsTokenCollection](resp), nil
}

// Retrieves a specific teams token by its ID for a given user. This endpoint is available to project administrators and requires project-level authentication. Use this to inspect token details including its status, creation date, and associated metadata.
func (api *ChannelsService) GetTeamsUserToken(ctx context.Context, userId string, tokenId string) (*shared.ProjectClientResponse[TeamsToken], *shared.ProjectClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/users/{user_id}/channels/teams/tokens/{token_id}").
		WithConfig(config).
		AddPathParam("user_id", userId).
		AddPathParam("token_id", tokenId).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[TeamsToken](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewProjectClientError[TeamsToken](err)
	}

	return shared.NewProjectClientResponse[TeamsToken](resp), nil
}

// Revokes a specific user's teams token. This endpoint is available to project administrators and permanently invalidates the specified token. Once revoked, the token can no longer be used to access channel features. This action cannot be undone.
func (api *ChannelsService) DiscardTeamsUserToken(ctx context.Context, userId string, tokenId string) (*shared.ProjectClientResponse[DiscardResult], *shared.ProjectClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("DELETE").
		WithPath("/users/{user_id}/channels/teams/tokens/{token_id}").
		WithConfig(config).
		AddPathParam("user_id", userId).
		AddPathParam("token_id", tokenId).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[DiscardResult](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewProjectClientError[DiscardResult](err)
	}

	return shared.NewProjectClientResponse[DiscardResult](resp), nil
}

// Lists all web_push tokens associated with a specific user. This endpoint is available to project administrators and returns a paginated list of tokens, including both active and revoked tokens.
func (api *ChannelsService) GetWebPushUserTokens(ctx context.Context, userId string, params GetWebPushUserTokensRequestParams) (*shared.ProjectClientResponse[WebPushTokenCollection], *shared.ProjectClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/users/{user_id}/channels/web_push/tokens").
		WithConfig(config).
		AddPathParam("user_id", userId).
		WithOptions(params).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[WebPushTokenCollection](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewProjectClientError[WebPushTokenCollection](err)
	}

	return shared.NewProjectClientResponse[WebPushTokenCollection](resp), nil
}

// Retrieves a specific web_push token by its ID for a given user. This endpoint is available to project administrators and requires project-level authentication. Use this to inspect token details including its status, creation date, and associated metadata.
func (api *ChannelsService) GetWebPushUserToken(ctx context.Context, userId string, tokenId string) (*shared.ProjectClientResponse[WebPushToken], *shared.ProjectClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/users/{user_id}/channels/web_push/tokens/{token_id}").
		WithConfig(config).
		AddPathParam("user_id", userId).
		AddPathParam("token_id", tokenId).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[WebPushToken](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewProjectClientError[WebPushToken](err)
	}

	return shared.NewProjectClientResponse[WebPushToken](resp), nil
}

// Revokes a specific user's web_push token. This endpoint is available to project administrators and permanently invalidates the specified token. Once revoked, the token can no longer be used to access channel features. This action cannot be undone.
func (api *ChannelsService) DiscardWebPushUserToken(ctx context.Context, userId string, tokenId string) (*shared.ProjectClientResponse[DiscardResult], *shared.ProjectClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("DELETE").
		WithPath("/users/{user_id}/channels/web_push/tokens/{token_id}").
		WithConfig(config).
		AddPathParam("user_id", userId).
		AddPathParam("token_id", tokenId).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[DiscardResult](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewProjectClientError[DiscardResult](err)
	}

	return shared.NewProjectClientResponse[DiscardResult](resp), nil
}
