package notifications

import (
	"context"
	restClient "github.com/magicbell/magicbell-go/pkg/project-client/internal/clients/rest"
	"github.com/magicbell/magicbell-go/pkg/project-client/internal/clients/rest/httptransport"
	"github.com/magicbell/magicbell-go/pkg/project-client/internal/configmanager"
	"github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
	"github.com/magicbell/magicbell-go/pkg/project-client/shared"
	"time"
)

type NotificationsService struct {
	manager *configmanager.ConfigManager
}

func NewNotificationsService() *NotificationsService {
	return &NotificationsService{
		manager: configmanager.NewConfigManager(clientconfig.Config{}),
	}
}

func (api *NotificationsService) WithConfigManager(manager *configmanager.ConfigManager) *NotificationsService {
	api.manager = manager
	return api
}

func (api *NotificationsService) getConfig() *clientconfig.Config {
	return api.manager.GetNotifications()
}

func (api *NotificationsService) SetBaseUrl(baseUrl string) {
	config := api.getConfig()
	config.SetBaseUrl(baseUrl)
}

func (api *NotificationsService) SetTimeout(timeout time.Duration) {
	config := api.getConfig()
	config.SetTimeout(timeout)
}

func (api *NotificationsService) SetAccessToken(accessToken string) {
	config := api.getConfig()
	config.SetAccessToken(accessToken)
}

// Get the delivery plan for a notification.
func (api *NotificationsService) GetDeliveryplan(ctx context.Context, notificationId string) (*shared.ClientResponse[DeliveryPlanCollection], *shared.ClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/notifications/{notification_id}/deliveryplan").
		WithConfig(config).
		AddPathParam("notification_id", notificationId).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[DeliveryPlanCollection](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewClientError[DeliveryPlanCollection](err)
	}

	return shared.NewClientResponse[DeliveryPlanCollection](resp), nil
}
