package events

import (
	"context"
	restClient "github.com/magicbell/magicbell-go-project-client/internal/clients/rest"
	"github.com/magicbell/magicbell-go-project-client/internal/clients/rest/httptransport"
	"github.com/magicbell/magicbell-go-project-client/internal/configmanager"
	"github.com/magicbell/magicbell-go-project-client/pkg/magicbellprojectclientconfig"
	"github.com/magicbell/magicbell-go-project-client/pkg/shared"
	"time"
)

type EventsService struct {
	manager *configmanager.ConfigManager
}

func NewEventsService() *EventsService {
	return &EventsService{
		manager: configmanager.NewConfigManager(magicbellprojectclientconfig.Config{}),
	}
}

func (api *EventsService) WithConfigManager(manager *configmanager.ConfigManager) *EventsService {
	api.manager = manager
	return api
}

func (api *EventsService) getConfig() *magicbellprojectclientconfig.Config {
	return api.manager.GetEvents()
}

func (api *EventsService) SetBaseUrl(baseUrl string) {
	config := api.getConfig()
	config.SetBaseUrl(baseUrl)
}

func (api *EventsService) SetTimeout(timeout time.Duration) {
	config := api.getConfig()
	config.SetTimeout(timeout)
}

func (api *EventsService) SetAccessToken(accessToken string) {
	config := api.getConfig()
	config.SetAccessToken(accessToken)
}

// Retrieves a paginated list of events for the project.
func (api *EventsService) ListEvents(ctx context.Context, params ListEventsRequestParams) (*shared.MagicbellProjectClientResponse[EventCollection], *shared.MagicbellProjectClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/events").
		WithConfig(config).
		WithOptions(params).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[EventCollection](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewMagicbellProjectClientError[EventCollection](err)
	}

	return shared.NewMagicbellProjectClientResponse[EventCollection](resp), nil
}

// Retrieves a project event by its ID.
func (api *EventsService) GetEvent(ctx context.Context, id string) (*shared.MagicbellProjectClientResponse[Event], *shared.MagicbellProjectClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/events/{id}").
		WithConfig(config).
		AddPathParam("id", id).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[Event](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewMagicbellProjectClientError[Event](err)
	}

	return shared.NewMagicbellProjectClientResponse[Event](resp), nil
}
