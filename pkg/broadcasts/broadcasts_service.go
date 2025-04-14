package broadcasts

import (
	"context"
	restClient "github.com/magicbell/magicbell-go-project-client/internal/clients/rest"
	"github.com/magicbell/magicbell-go-project-client/internal/clients/rest/httptransport"
	"github.com/magicbell/magicbell-go-project-client/internal/configmanager"
	"github.com/magicbell/magicbell-go-project-client/pkg/magicbellprojectclientconfig"
	"github.com/magicbell/magicbell-go-project-client/pkg/shared"
	"time"
)

type BroadcastsService struct {
	manager *configmanager.ConfigManager
}

func NewBroadcastsService() *BroadcastsService {
	return &BroadcastsService{
		manager: configmanager.NewConfigManager(magicbellprojectclientconfig.Config{}),
	}
}

func (api *BroadcastsService) WithConfigManager(manager *configmanager.ConfigManager) *BroadcastsService {
	api.manager = manager
	return api
}

func (api *BroadcastsService) getConfig() *magicbellprojectclientconfig.Config {
	return api.manager.GetBroadcasts()
}

func (api *BroadcastsService) SetBaseUrl(baseUrl string) {
	config := api.getConfig()
	config.SetBaseUrl(baseUrl)
}

func (api *BroadcastsService) SetTimeout(timeout time.Duration) {
	config := api.getConfig()
	config.SetTimeout(timeout)
}

func (api *BroadcastsService) SetAccessToken(accessToken string) {
	config := api.getConfig()
	config.SetAccessToken(accessToken)
}

// Retrieves a paginated list of broadcasts for the project. Returns basic information about each broadcast including its creation time and status.
func (api *BroadcastsService) ListBroadcasts(ctx context.Context, params ListBroadcastsRequestParams) (*shared.MagicbellProjectClientResponse[BroadcastCollection], *shared.MagicbellProjectClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/broadcasts").
		WithConfig(config).
		WithOptions(params).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[BroadcastCollection](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewMagicbellProjectClientError[BroadcastCollection](err)
	}

	return shared.NewMagicbellProjectClientResponse[BroadcastCollection](resp), nil
}

// Creates a new broadcast message. When a broadcast is created, it generates individual notifications for relevant users within the project.
func (api *BroadcastsService) CreateBroadcast(ctx context.Context, broadcast Broadcast) (*shared.MagicbellProjectClientResponse[Broadcast], *shared.MagicbellProjectClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("POST").
		WithPath("/broadcasts").
		WithConfig(config).
		WithBody(broadcast).
		AddHeader("CONTENT-TYPE", "application/json").
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[Broadcast](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewMagicbellProjectClientError[Broadcast](err)
	}

	return shared.NewMagicbellProjectClientResponse[Broadcast](resp), nil
}

// Retrieves detailed information about a specific broadcast by its ID. Includes the broadcast's configuration and current status.
func (api *BroadcastsService) FetchBroadcast(ctx context.Context, broadcastId string) (*shared.MagicbellProjectClientResponse[Broadcast], *shared.MagicbellProjectClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/broadcasts/{broadcast_id}").
		WithConfig(config).
		AddPathParam("broadcast_id", broadcastId).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[Broadcast](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewMagicbellProjectClientError[Broadcast](err)
	}

	return shared.NewMagicbellProjectClientResponse[Broadcast](resp), nil
}
