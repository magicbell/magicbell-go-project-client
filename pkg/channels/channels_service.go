package channels

import (
	"context"
	restClient "github.com/magicbell/magicbell-go-project-client/internal/clients/rest"
	"github.com/magicbell/magicbell-go-project-client/internal/clients/rest/httptransport"
	"github.com/magicbell/magicbell-go-project-client/internal/configmanager"
	"github.com/magicbell/magicbell-go-project-client/pkg/magicbellprojectclientconfig"
	"github.com/magicbell/magicbell-go-project-client/pkg/shared"
	"time"
)

type ChannelsService struct {
	manager *configmanager.ConfigManager
}

func NewChannelsService(manager *configmanager.ConfigManager) *ChannelsService {
	return &ChannelsService{
		manager: manager,
	}
}

func (api *ChannelsService) getConfig() *magicbellprojectclientconfig.Config {
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

func (api *ChannelsService) GetProjectDeliveryconfig(ctx context.Context) (*shared.MagicbellProjectClientResponse[ProjectDeliveryConfig], *shared.MagicbellProjectClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/channels/deliveryconfig").
		WithConfig(config).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[ProjectDeliveryConfig](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewMagicbellProjectClientError[ProjectDeliveryConfig](err)
	}

	return shared.NewMagicbellProjectClientResponse[ProjectDeliveryConfig](resp), nil
}

func (api *ChannelsService) SaveProjectDeliveryconfig(ctx context.Context, projectDeliveryConfig ProjectDeliveryConfig) (*shared.MagicbellProjectClientResponse[ProjectDeliveryConfig], *shared.MagicbellProjectClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("PUT").
		WithPath("/channels/deliveryconfig").
		WithConfig(config).
		WithBody(projectDeliveryConfig).
		AddHeader("CONTENT-TYPE", "application/json").
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[ProjectDeliveryConfig](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewMagicbellProjectClientError[ProjectDeliveryConfig](err)
	}

	return shared.NewMagicbellProjectClientResponse[ProjectDeliveryConfig](resp), nil
}

func (api *ChannelsService) SaveCategoriesDeliveryconfig(ctx context.Context, categoryDeliveryConfig CategoryDeliveryConfig) (*shared.MagicbellProjectClientResponse[CategoryDeliveryConfig], *shared.MagicbellProjectClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("POST").
		WithPath("/channels/deliveryconfig/categories").
		WithConfig(config).
		WithBody(categoryDeliveryConfig).
		AddHeader("CONTENT-TYPE", "application/json").
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[CategoryDeliveryConfig](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewMagicbellProjectClientError[CategoryDeliveryConfig](err)
	}

	return shared.NewMagicbellProjectClientResponse[CategoryDeliveryConfig](resp), nil
}

// Lists all mobile_push tokens associated with a specific user. This endpoint is available to project administrators and returns a paginated list of tokens, including both active and revoked tokens.
func (api *ChannelsService) GetMobilePushApnsUserTokens(ctx context.Context, userId string, params GetMobilePushApnsUserTokensRequestParams) (*shared.MagicbellProjectClientResponse[ArrayOfMetadataApnsTokens], *shared.MagicbellProjectClientError) {
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

	client := restClient.NewRestClient[ArrayOfMetadataApnsTokens](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewMagicbellProjectClientError[ArrayOfMetadataApnsTokens](err)
	}

	return shared.NewMagicbellProjectClientResponse[ArrayOfMetadataApnsTokens](resp), nil
}

// Retrieves a specific mobile_push token by its ID for a given user. This endpoint is available to project administrators and requires project-level authentication. Use this to inspect token details including its status, creation date, and associated metadata.
func (api *ChannelsService) GetMobilePushApnsUserToken(ctx context.Context, userId string, tokenId string) (*shared.MagicbellProjectClientResponse[MetadataApnsToken], *shared.MagicbellProjectClientError) {
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

	client := restClient.NewRestClient[MetadataApnsToken](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewMagicbellProjectClientError[MetadataApnsToken](err)
	}

	return shared.NewMagicbellProjectClientResponse[MetadataApnsToken](resp), nil
}

// Revokes a specific user's mobile_push token. This endpoint is available to project administrators and permanently invalidates the specified token. Once revoked, the token can no longer be used to access channel features. This action cannot be undone.
func (api *ChannelsService) DiscardMobilePushApnsUserToken(ctx context.Context, userId string, tokenId string) (*shared.MagicbellProjectClientResponse[DiscardResult], *shared.MagicbellProjectClientError) {
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
		return nil, shared.NewMagicbellProjectClientError[DiscardResult](err)
	}

	return shared.NewMagicbellProjectClientResponse[DiscardResult](resp), nil
}

// Lists all mobile_push tokens associated with a specific user. This endpoint is available to project administrators and returns a paginated list of tokens, including both active and revoked tokens.
func (api *ChannelsService) GetMobilePushExpoUserTokens(ctx context.Context, userId string, params GetMobilePushExpoUserTokensRequestParams) (*shared.MagicbellProjectClientResponse[ArrayOfMetadataExpoTokens], *shared.MagicbellProjectClientError) {
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

	client := restClient.NewRestClient[ArrayOfMetadataExpoTokens](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewMagicbellProjectClientError[ArrayOfMetadataExpoTokens](err)
	}

	return shared.NewMagicbellProjectClientResponse[ArrayOfMetadataExpoTokens](resp), nil
}

// Retrieves a specific mobile_push token by its ID for a given user. This endpoint is available to project administrators and requires project-level authentication. Use this to inspect token details including its status, creation date, and associated metadata.
func (api *ChannelsService) GetMobilePushExpoUserToken(ctx context.Context, userId string, tokenId string) (*shared.MagicbellProjectClientResponse[MetadataExpoToken], *shared.MagicbellProjectClientError) {
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

	client := restClient.NewRestClient[MetadataExpoToken](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewMagicbellProjectClientError[MetadataExpoToken](err)
	}

	return shared.NewMagicbellProjectClientResponse[MetadataExpoToken](resp), nil
}

// Revokes a specific user's mobile_push token. This endpoint is available to project administrators and permanently invalidates the specified token. Once revoked, the token can no longer be used to access channel features. This action cannot be undone.
func (api *ChannelsService) DiscardMobilePushExpoUserToken(ctx context.Context, userId string, tokenId string) (*shared.MagicbellProjectClientResponse[DiscardResult], *shared.MagicbellProjectClientError) {
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
		return nil, shared.NewMagicbellProjectClientError[DiscardResult](err)
	}

	return shared.NewMagicbellProjectClientResponse[DiscardResult](resp), nil
}

// Lists all mobile_push tokens associated with a specific user. This endpoint is available to project administrators and returns a paginated list of tokens, including both active and revoked tokens.
func (api *ChannelsService) GetMobilePushFcmUserTokens(ctx context.Context, userId string, params GetMobilePushFcmUserTokensRequestParams) (*shared.MagicbellProjectClientResponse[ArrayOfMetadataFcmTokens], *shared.MagicbellProjectClientError) {
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

	client := restClient.NewRestClient[ArrayOfMetadataFcmTokens](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewMagicbellProjectClientError[ArrayOfMetadataFcmTokens](err)
	}

	return shared.NewMagicbellProjectClientResponse[ArrayOfMetadataFcmTokens](resp), nil
}

// Retrieves a specific mobile_push token by its ID for a given user. This endpoint is available to project administrators and requires project-level authentication. Use this to inspect token details including its status, creation date, and associated metadata.
func (api *ChannelsService) GetMobilePushFcmUserToken(ctx context.Context, userId string, tokenId string) (*shared.MagicbellProjectClientResponse[MetadataFcmToken], *shared.MagicbellProjectClientError) {
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

	client := restClient.NewRestClient[MetadataFcmToken](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewMagicbellProjectClientError[MetadataFcmToken](err)
	}

	return shared.NewMagicbellProjectClientResponse[MetadataFcmToken](resp), nil
}

// Revokes a specific user's mobile_push token. This endpoint is available to project administrators and permanently invalidates the specified token. Once revoked, the token can no longer be used to access channel features. This action cannot be undone.
func (api *ChannelsService) DiscardMobilePushFcmUserToken(ctx context.Context, userId string, tokenId string) (*shared.MagicbellProjectClientResponse[DiscardResult], *shared.MagicbellProjectClientError) {
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
		return nil, shared.NewMagicbellProjectClientError[DiscardResult](err)
	}

	return shared.NewMagicbellProjectClientResponse[DiscardResult](resp), nil
}

// Lists all slack tokens associated with a specific user. This endpoint is available to project administrators and returns a paginated list of tokens, including both active and revoked tokens.
func (api *ChannelsService) GetSlackUserTokens(ctx context.Context, userId string, params GetSlackUserTokensRequestParams) (*shared.MagicbellProjectClientResponse[ArrayOfMetadataSlackTokens], *shared.MagicbellProjectClientError) {
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

	client := restClient.NewRestClient[ArrayOfMetadataSlackTokens](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewMagicbellProjectClientError[ArrayOfMetadataSlackTokens](err)
	}

	return shared.NewMagicbellProjectClientResponse[ArrayOfMetadataSlackTokens](resp), nil
}

// Retrieves a specific slack token by its ID for a given user. This endpoint is available to project administrators and requires project-level authentication. Use this to inspect token details including its status, creation date, and associated metadata.
func (api *ChannelsService) GetSlackUserToken(ctx context.Context, userId string, tokenId string) (*shared.MagicbellProjectClientResponse[MetadataSlackToken], *shared.MagicbellProjectClientError) {
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

	client := restClient.NewRestClient[MetadataSlackToken](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewMagicbellProjectClientError[MetadataSlackToken](err)
	}

	return shared.NewMagicbellProjectClientResponse[MetadataSlackToken](resp), nil
}

// Revokes a specific user's slack token. This endpoint is available to project administrators and permanently invalidates the specified token. Once revoked, the token can no longer be used to access channel features. This action cannot be undone.
func (api *ChannelsService) DiscardSlackUserToken(ctx context.Context, userId string, tokenId string) (*shared.MagicbellProjectClientResponse[DiscardResult], *shared.MagicbellProjectClientError) {
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
		return nil, shared.NewMagicbellProjectClientError[DiscardResult](err)
	}

	return shared.NewMagicbellProjectClientResponse[DiscardResult](resp), nil
}

// Lists all teams tokens associated with a specific user. This endpoint is available to project administrators and returns a paginated list of tokens, including both active and revoked tokens.
func (api *ChannelsService) GetTeamsUserTokens(ctx context.Context, userId string, params GetTeamsUserTokensRequestParams) (*shared.MagicbellProjectClientResponse[ArrayOfMetadataTeamsTokens], *shared.MagicbellProjectClientError) {
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

	client := restClient.NewRestClient[ArrayOfMetadataTeamsTokens](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewMagicbellProjectClientError[ArrayOfMetadataTeamsTokens](err)
	}

	return shared.NewMagicbellProjectClientResponse[ArrayOfMetadataTeamsTokens](resp), nil
}

// Retrieves a specific teams token by its ID for a given user. This endpoint is available to project administrators and requires project-level authentication. Use this to inspect token details including its status, creation date, and associated metadata.
func (api *ChannelsService) GetTeamsUserToken(ctx context.Context, userId string, tokenId string) (*shared.MagicbellProjectClientResponse[MetadataTeamsToken], *shared.MagicbellProjectClientError) {
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

	client := restClient.NewRestClient[MetadataTeamsToken](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewMagicbellProjectClientError[MetadataTeamsToken](err)
	}

	return shared.NewMagicbellProjectClientResponse[MetadataTeamsToken](resp), nil
}

// Revokes a specific user's teams token. This endpoint is available to project administrators and permanently invalidates the specified token. Once revoked, the token can no longer be used to access channel features. This action cannot be undone.
func (api *ChannelsService) DiscardTeamsUserToken(ctx context.Context, userId string, tokenId string) (*shared.MagicbellProjectClientResponse[DiscardResult], *shared.MagicbellProjectClientError) {
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
		return nil, shared.NewMagicbellProjectClientError[DiscardResult](err)
	}

	return shared.NewMagicbellProjectClientResponse[DiscardResult](resp), nil
}

// Lists all web_push tokens associated with a specific user. This endpoint is available to project administrators and returns a paginated list of tokens, including both active and revoked tokens.
func (api *ChannelsService) GetWebPushUserTokens(ctx context.Context, userId string, params GetWebPushUserTokensRequestParams) (*shared.MagicbellProjectClientResponse[ArrayOfMetadataWebPushTokens], *shared.MagicbellProjectClientError) {
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

	client := restClient.NewRestClient[ArrayOfMetadataWebPushTokens](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewMagicbellProjectClientError[ArrayOfMetadataWebPushTokens](err)
	}

	return shared.NewMagicbellProjectClientResponse[ArrayOfMetadataWebPushTokens](resp), nil
}

// Retrieves a specific web_push token by its ID for a given user. This endpoint is available to project administrators and requires project-level authentication. Use this to inspect token details including its status, creation date, and associated metadata.
func (api *ChannelsService) GetWebPushUserToken(ctx context.Context, userId string, tokenId string) (*shared.MagicbellProjectClientResponse[MetadataWebPushToken], *shared.MagicbellProjectClientError) {
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

	client := restClient.NewRestClient[MetadataWebPushToken](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewMagicbellProjectClientError[MetadataWebPushToken](err)
	}

	return shared.NewMagicbellProjectClientResponse[MetadataWebPushToken](resp), nil
}

// Revokes a specific user's web_push token. This endpoint is available to project administrators and permanently invalidates the specified token. Once revoked, the token can no longer be used to access channel features. This action cannot be undone.
func (api *ChannelsService) DiscardWebPushUserToken(ctx context.Context, userId string, tokenId string) (*shared.MagicbellProjectClientResponse[DiscardResult], *shared.MagicbellProjectClientError) {
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
		return nil, shared.NewMagicbellProjectClientError[DiscardResult](err)
	}

	return shared.NewMagicbellProjectClientResponse[DiscardResult](resp), nil
}
