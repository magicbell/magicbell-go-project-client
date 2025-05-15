package channels

import (
	"context"
	restClient "github.com/magicbell/magicbell-go/pkg/user-client/internal/clients/rest"
	"github.com/magicbell/magicbell-go/pkg/user-client/internal/clients/rest/httptransport"
	"github.com/magicbell/magicbell-go/pkg/user-client/internal/configmanager"
	"github.com/magicbell/magicbell-go/pkg/user-client/shared"
	"github.com/magicbell/magicbell-go/pkg/user-client/userclientconfig"
	"time"
)

type ChannelsService struct {
	manager *configmanager.ConfigManager
}

func NewChannelsService() *ChannelsService {
	return &ChannelsService{
		manager: configmanager.NewConfigManager(userclientconfig.Config{}),
	}
}

func (api *ChannelsService) WithConfigManager(manager *configmanager.ConfigManager) *ChannelsService {
	api.manager = manager
	return api
}

func (api *ChannelsService) getConfig() *userclientconfig.Config {
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

// Lists all in_app tokens belonging to the authenticated user. Returns a paginated list of tokens, including their status, creation dates, and associated metadata.
func (api *ChannelsService) GetInAppInboxTokens(ctx context.Context, params GetInAppInboxTokensRequestParams) (*shared.UserClientResponse[InboxTokenResponseCollection], *shared.UserClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/channels/in_app/inbox/tokens").
		WithConfig(config).
		WithOptions(params).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[InboxTokenResponseCollection](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewUserClientError[InboxTokenResponseCollection](err)
	}

	return shared.NewUserClientResponse[InboxTokenResponseCollection](resp), nil
}

// Saves a in_app token for the authenticated user. This token serves as a credential for accessing channel-specific functionality. Each token is unique to the user and channel combination, allowing for direct communication with the user via the channel.
func (api *ChannelsService) SaveInAppInboxToken(ctx context.Context, inboxToken InboxToken) (*shared.UserClientResponse[InboxToken], *shared.UserClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("POST").
		WithPath("/channels/in_app/inbox/tokens").
		WithConfig(config).
		WithBody(inboxToken).
		AddHeader("CONTENT-TYPE", "application/json").
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[InboxToken](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewUserClientError[InboxToken](err)
	}

	return shared.NewUserClientResponse[InboxToken](resp), nil
}

// Retrieves details of a specific in_app token belonging to the authenticated user. Returns information about the token's status, creation date, and any associated metadata. Users can only access their own tokens.
func (api *ChannelsService) GetInAppInboxToken(ctx context.Context, tokenId string) (*shared.UserClientResponse[InboxTokenResponse], *shared.UserClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/channels/in_app/inbox/tokens/{token_id}").
		WithConfig(config).
		AddPathParam("token_id", tokenId).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[InboxTokenResponse](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewUserClientError[InboxTokenResponse](err)
	}

	return shared.NewUserClientResponse[InboxTokenResponse](resp), nil
}

// Revokes one of the authenticated user's in_app tokens. This permanently invalidates the specified token, preventing it from being used for future channel access. This action cannot be undone. Users can only revoke their own tokens.
func (api *ChannelsService) DiscardInAppInboxToken(ctx context.Context, tokenId string) (*shared.UserClientResponse[DiscardResult], *shared.UserClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("DELETE").
		WithPath("/channels/in_app/inbox/tokens/{token_id}").
		WithConfig(config).
		AddPathParam("token_id", tokenId).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[DiscardResult](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewUserClientError[DiscardResult](err)
	}

	return shared.NewUserClientResponse[DiscardResult](resp), nil
}

// Lists all mobile_push tokens belonging to the authenticated user. Returns a paginated list of tokens, including their status, creation dates, and associated metadata.
func (api *ChannelsService) GetMobilePushApnsTokens(ctx context.Context, params GetMobilePushApnsTokensRequestParams) (*shared.UserClientResponse[ApnsTokenCollection], *shared.UserClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/channels/mobile_push/apns/tokens").
		WithConfig(config).
		WithOptions(params).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[ApnsTokenCollection](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewUserClientError[ApnsTokenCollection](err)
	}

	return shared.NewUserClientResponse[ApnsTokenCollection](resp), nil
}

// Saves a mobile_push token for the authenticated user. This token serves as a credential for accessing channel-specific functionality. Each token is unique to the user and channel combination, allowing for direct communication with the user via the channel.
func (api *ChannelsService) SaveMobilePushApnsToken(ctx context.Context, apnsTokenPayload ApnsTokenPayload) (*shared.UserClientResponse[ApnsTokenPayload], *shared.UserClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("POST").
		WithPath("/channels/mobile_push/apns/tokens").
		WithConfig(config).
		WithBody(apnsTokenPayload).
		AddHeader("CONTENT-TYPE", "application/json").
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[ApnsTokenPayload](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewUserClientError[ApnsTokenPayload](err)
	}

	return shared.NewUserClientResponse[ApnsTokenPayload](resp), nil
}

// Retrieves details of a specific mobile_push token belonging to the authenticated user. Returns information about the token's status, creation date, and any associated metadata. Users can only access their own tokens.
func (api *ChannelsService) GetMobilePushApnsToken(ctx context.Context, tokenId string) (*shared.UserClientResponse[ApnsToken], *shared.UserClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/channels/mobile_push/apns/tokens/{token_id}").
		WithConfig(config).
		AddPathParam("token_id", tokenId).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[ApnsToken](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewUserClientError[ApnsToken](err)
	}

	return shared.NewUserClientResponse[ApnsToken](resp), nil
}

// Revokes one of the authenticated user's mobile_push tokens. This permanently invalidates the specified token, preventing it from being used for future channel access. This action cannot be undone. Users can only revoke their own tokens.
func (api *ChannelsService) DiscardMobilePushApnsToken(ctx context.Context, tokenId string) (*shared.UserClientResponse[DiscardResult], *shared.UserClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("DELETE").
		WithPath("/channels/mobile_push/apns/tokens/{token_id}").
		WithConfig(config).
		AddPathParam("token_id", tokenId).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[DiscardResult](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewUserClientError[DiscardResult](err)
	}

	return shared.NewUserClientResponse[DiscardResult](resp), nil
}

// Lists all mobile_push tokens belonging to the authenticated user. Returns a paginated list of tokens, including their status, creation dates, and associated metadata.
func (api *ChannelsService) GetMobilePushExpoTokens(ctx context.Context, params GetMobilePushExpoTokensRequestParams) (*shared.UserClientResponse[ExpoTokenCollection], *shared.UserClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/channels/mobile_push/expo/tokens").
		WithConfig(config).
		WithOptions(params).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[ExpoTokenCollection](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewUserClientError[ExpoTokenCollection](err)
	}

	return shared.NewUserClientResponse[ExpoTokenCollection](resp), nil
}

// Saves a mobile_push token for the authenticated user. This token serves as a credential for accessing channel-specific functionality. Each token is unique to the user and channel combination, allowing for direct communication with the user via the channel.
func (api *ChannelsService) SaveMobilePushExpoToken(ctx context.Context, expoTokenPayload ExpoTokenPayload) (*shared.UserClientResponse[ExpoTokenPayload], *shared.UserClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("POST").
		WithPath("/channels/mobile_push/expo/tokens").
		WithConfig(config).
		WithBody(expoTokenPayload).
		AddHeader("CONTENT-TYPE", "application/json").
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[ExpoTokenPayload](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewUserClientError[ExpoTokenPayload](err)
	}

	return shared.NewUserClientResponse[ExpoTokenPayload](resp), nil
}

// Retrieves details of a specific mobile_push token belonging to the authenticated user. Returns information about the token's status, creation date, and any associated metadata. Users can only access their own tokens.
func (api *ChannelsService) GetMobilePushExpoToken(ctx context.Context, tokenId string) (*shared.UserClientResponse[ExpoToken], *shared.UserClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/channels/mobile_push/expo/tokens/{token_id}").
		WithConfig(config).
		AddPathParam("token_id", tokenId).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[ExpoToken](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewUserClientError[ExpoToken](err)
	}

	return shared.NewUserClientResponse[ExpoToken](resp), nil
}

// Revokes one of the authenticated user's mobile_push tokens. This permanently invalidates the specified token, preventing it from being used for future channel access. This action cannot be undone. Users can only revoke their own tokens.
func (api *ChannelsService) DiscardMobilePushExpoToken(ctx context.Context, tokenId string) (*shared.UserClientResponse[DiscardResult], *shared.UserClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("DELETE").
		WithPath("/channels/mobile_push/expo/tokens/{token_id}").
		WithConfig(config).
		AddPathParam("token_id", tokenId).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[DiscardResult](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewUserClientError[DiscardResult](err)
	}

	return shared.NewUserClientResponse[DiscardResult](resp), nil
}

// Lists all mobile_push tokens belonging to the authenticated user. Returns a paginated list of tokens, including their status, creation dates, and associated metadata.
func (api *ChannelsService) GetMobilePushFcmTokens(ctx context.Context, params GetMobilePushFcmTokensRequestParams) (*shared.UserClientResponse[FcmTokenCollection], *shared.UserClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/channels/mobile_push/fcm/tokens").
		WithConfig(config).
		WithOptions(params).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[FcmTokenCollection](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewUserClientError[FcmTokenCollection](err)
	}

	return shared.NewUserClientResponse[FcmTokenCollection](resp), nil
}

// Saves a mobile_push token for the authenticated user. This token serves as a credential for accessing channel-specific functionality. Each token is unique to the user and channel combination, allowing for direct communication with the user via the channel.
func (api *ChannelsService) SaveMobilePushFcmToken(ctx context.Context, fcmTokenPayload FcmTokenPayload) (*shared.UserClientResponse[FcmTokenPayload], *shared.UserClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("POST").
		WithPath("/channels/mobile_push/fcm/tokens").
		WithConfig(config).
		WithBody(fcmTokenPayload).
		AddHeader("CONTENT-TYPE", "application/json").
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[FcmTokenPayload](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewUserClientError[FcmTokenPayload](err)
	}

	return shared.NewUserClientResponse[FcmTokenPayload](resp), nil
}

// Retrieves details of a specific mobile_push token belonging to the authenticated user. Returns information about the token's status, creation date, and any associated metadata. Users can only access their own tokens.
func (api *ChannelsService) GetMobilePushFcmToken(ctx context.Context, tokenId string) (*shared.UserClientResponse[FcmToken], *shared.UserClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/channels/mobile_push/fcm/tokens/{token_id}").
		WithConfig(config).
		AddPathParam("token_id", tokenId).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[FcmToken](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewUserClientError[FcmToken](err)
	}

	return shared.NewUserClientResponse[FcmToken](resp), nil
}

// Revokes one of the authenticated user's mobile_push tokens. This permanently invalidates the specified token, preventing it from being used for future channel access. This action cannot be undone. Users can only revoke their own tokens.
func (api *ChannelsService) DiscardMobilePushFcmToken(ctx context.Context, tokenId string) (*shared.UserClientResponse[DiscardResult], *shared.UserClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("DELETE").
		WithPath("/channels/mobile_push/fcm/tokens/{token_id}").
		WithConfig(config).
		AddPathParam("token_id", tokenId).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[DiscardResult](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewUserClientError[DiscardResult](err)
	}

	return shared.NewUserClientResponse[DiscardResult](resp), nil
}

// Lists all slack tokens belonging to the authenticated user. Returns a paginated list of tokens, including their status, creation dates, and associated metadata.
func (api *ChannelsService) GetSlackTokens(ctx context.Context, params GetSlackTokensRequestParams) (*shared.UserClientResponse[SlackTokenCollection], *shared.UserClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/channels/slack/tokens").
		WithConfig(config).
		WithOptions(params).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[SlackTokenCollection](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewUserClientError[SlackTokenCollection](err)
	}

	return shared.NewUserClientResponse[SlackTokenCollection](resp), nil
}

// Saves a slack token for the authenticated user. This token serves as a credential for accessing channel-specific functionality. Each token is unique to the user and channel combination, allowing for direct communication with the user via the channel.
func (api *ChannelsService) SaveSlackToken(ctx context.Context, slackTokenPayload SlackTokenPayload) (*shared.UserClientResponse[SlackTokenPayload], *shared.UserClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("POST").
		WithPath("/channels/slack/tokens").
		WithConfig(config).
		WithBody(slackTokenPayload).
		AddHeader("CONTENT-TYPE", "application/json").
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[SlackTokenPayload](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewUserClientError[SlackTokenPayload](err)
	}

	return shared.NewUserClientResponse[SlackTokenPayload](resp), nil
}

// Retrieves details of a specific slack token belonging to the authenticated user. Returns information about the token's status, creation date, and any associated metadata. Users can only access their own tokens.
func (api *ChannelsService) GetSlackToken(ctx context.Context, tokenId string) (*shared.UserClientResponse[SlackToken], *shared.UserClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/channels/slack/tokens/{token_id}").
		WithConfig(config).
		AddPathParam("token_id", tokenId).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[SlackToken](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewUserClientError[SlackToken](err)
	}

	return shared.NewUserClientResponse[SlackToken](resp), nil
}

// Revokes one of the authenticated user's slack tokens. This permanently invalidates the specified token, preventing it from being used for future channel access. This action cannot be undone. Users can only revoke their own tokens.
func (api *ChannelsService) DiscardSlackToken(ctx context.Context, tokenId string) (*shared.UserClientResponse[DiscardResult], *shared.UserClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("DELETE").
		WithPath("/channels/slack/tokens/{token_id}").
		WithConfig(config).
		AddPathParam("token_id", tokenId).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[DiscardResult](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewUserClientError[DiscardResult](err)
	}

	return shared.NewUserClientResponse[DiscardResult](resp), nil
}

// Lists all teams tokens belonging to the authenticated user. Returns a paginated list of tokens, including their status, creation dates, and associated metadata.
func (api *ChannelsService) GetTeamsTokens(ctx context.Context, params GetTeamsTokensRequestParams) (*shared.UserClientResponse[TeamsTokenCollection], *shared.UserClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/channels/teams/tokens").
		WithConfig(config).
		WithOptions(params).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[TeamsTokenCollection](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewUserClientError[TeamsTokenCollection](err)
	}

	return shared.NewUserClientResponse[TeamsTokenCollection](resp), nil
}

// Saves a teams token for the authenticated user. This token serves as a credential for accessing channel-specific functionality. Each token is unique to the user and channel combination, allowing for direct communication with the user via the channel.
func (api *ChannelsService) SaveTeamsToken(ctx context.Context, teamsTokenPayload TeamsTokenPayload) (*shared.UserClientResponse[TeamsTokenPayload], *shared.UserClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("POST").
		WithPath("/channels/teams/tokens").
		WithConfig(config).
		WithBody(teamsTokenPayload).
		AddHeader("CONTENT-TYPE", "application/json").
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[TeamsTokenPayload](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewUserClientError[TeamsTokenPayload](err)
	}

	return shared.NewUserClientResponse[TeamsTokenPayload](resp), nil
}

// Retrieves details of a specific teams token belonging to the authenticated user. Returns information about the token's status, creation date, and any associated metadata. Users can only access their own tokens.
func (api *ChannelsService) GetTeamsToken(ctx context.Context, tokenId string) (*shared.UserClientResponse[TeamsToken], *shared.UserClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/channels/teams/tokens/{token_id}").
		WithConfig(config).
		AddPathParam("token_id", tokenId).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[TeamsToken](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewUserClientError[TeamsToken](err)
	}

	return shared.NewUserClientResponse[TeamsToken](resp), nil
}

// Revokes one of the authenticated user's teams tokens. This permanently invalidates the specified token, preventing it from being used for future channel access. This action cannot be undone. Users can only revoke their own tokens.
func (api *ChannelsService) DiscardTeamsToken(ctx context.Context, tokenId string) (*shared.UserClientResponse[DiscardResult], *shared.UserClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("DELETE").
		WithPath("/channels/teams/tokens/{token_id}").
		WithConfig(config).
		AddPathParam("token_id", tokenId).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[DiscardResult](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewUserClientError[DiscardResult](err)
	}

	return shared.NewUserClientResponse[DiscardResult](resp), nil
}

// Lists all web_push tokens belonging to the authenticated user. Returns a paginated list of tokens, including their status, creation dates, and associated metadata.
func (api *ChannelsService) GetWebPushTokens(ctx context.Context, params GetWebPushTokensRequestParams) (*shared.UserClientResponse[WebPushTokenCollection], *shared.UserClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/channels/web_push/tokens").
		WithConfig(config).
		WithOptions(params).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[WebPushTokenCollection](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewUserClientError[WebPushTokenCollection](err)
	}

	return shared.NewUserClientResponse[WebPushTokenCollection](resp), nil
}

// Saves a web_push token for the authenticated user. This token serves as a credential for accessing channel-specific functionality. Each token is unique to the user and channel combination, allowing for direct communication with the user via the channel.
func (api *ChannelsService) SaveWebPushToken(ctx context.Context, webPushTokenPayload shared.WebPushTokenPayload) (*shared.UserClientResponse[shared.WebPushTokenPayload], *shared.UserClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("POST").
		WithPath("/channels/web_push/tokens").
		WithConfig(config).
		WithBody(webPushTokenPayload).
		AddHeader("CONTENT-TYPE", "application/json").
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[shared.WebPushTokenPayload](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewUserClientError[shared.WebPushTokenPayload](err)
	}

	return shared.NewUserClientResponse[shared.WebPushTokenPayload](resp), nil
}

// Retrieves details of a specific web_push token belonging to the authenticated user. Returns information about the token's status, creation date, and any associated metadata. Users can only access their own tokens.
func (api *ChannelsService) GetWebPushToken(ctx context.Context, tokenId string) (*shared.UserClientResponse[WebPushToken], *shared.UserClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/channels/web_push/tokens/{token_id}").
		WithConfig(config).
		AddPathParam("token_id", tokenId).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[WebPushToken](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewUserClientError[WebPushToken](err)
	}

	return shared.NewUserClientResponse[WebPushToken](resp), nil
}

// Revokes one of the authenticated user's web_push tokens. This permanently invalidates the specified token, preventing it from being used for future channel access. This action cannot be undone. Users can only revoke their own tokens.
func (api *ChannelsService) DiscardWebPushToken(ctx context.Context, tokenId string) (*shared.UserClientResponse[DiscardResult], *shared.UserClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("DELETE").
		WithPath("/channels/web_push/tokens/{token_id}").
		WithConfig(config).
		AddPathParam("token_id", tokenId).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[DiscardResult](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewUserClientError[DiscardResult](err)
	}

	return shared.NewUserClientResponse[DiscardResult](resp), nil
}
