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

func NewEventsService(manager *configmanager.ConfigManager) *EventsService {
	return &EventsService{
		manager: manager,
	}
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
func (api *EventsService) GetEvents(ctx context.Context, params GetEventsRequestParams) (*shared.MagicbellProjectClientResponse[ArrayOfEvents], *shared.MagicbellProjectClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/events").
		WithConfig(config).
		WithOptions(params).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[ArrayOfEvents](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewMagicbellProjectClientError[ArrayOfEvents](err)
	}

	return shared.NewMagicbellProjectClientResponse[ArrayOfEvents](resp), nil
}
