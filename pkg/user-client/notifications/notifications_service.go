package notifications

import (
	"context"
	restClient "github.com/magicbell/magicbell-go/pkg/user-client/internal/clients/rest"
	"github.com/magicbell/magicbell-go/pkg/user-client/internal/clients/rest/httptransport"
	"github.com/magicbell/magicbell-go/pkg/user-client/internal/configmanager"
	"github.com/magicbell/magicbell-go/pkg/user-client/clientconfig"
	"github.com/magicbell/magicbell-go/pkg/user-client/shared"
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

// Lists all notifications for a user.
func (api *NotificationsService) ListNotifications(ctx context.Context, params ListNotificationsRequestParams) (*shared.ClientResponse[NotificationCollection], *shared.ClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/notifications").
		WithConfig(config).
		WithOptions(params).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[NotificationCollection](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewClientError[NotificationCollection](err)
	}

	return shared.NewClientResponse[NotificationCollection](resp), nil
}

// Archives a notification.
func (api *NotificationsService) ArchiveNotification(ctx context.Context, notificationId string) (*shared.ClientResponse[any], *shared.ClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("POST").
		WithPath("/notifications/{notification_id}/archive").
		WithConfig(config).
		AddPathParam("notification_id", notificationId).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[any](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewClientError[any](err)
	}

	return shared.NewClientResponse[any](resp), nil
}

// Marks a notification as read.
func (api *NotificationsService) MarkNotificationRead(ctx context.Context, notificationId string) (*shared.ClientResponse[any], *shared.ClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("POST").
		WithPath("/notifications/{notification_id}/read").
		WithConfig(config).
		AddPathParam("notification_id", notificationId).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[any](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewClientError[any](err)
	}

	return shared.NewClientResponse[any](resp), nil
}

// Unarchives a notification.
func (api *NotificationsService) UnarchiveNotification(ctx context.Context, notificationId string) (*shared.ClientResponse[any], *shared.ClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("POST").
		WithPath("/notifications/{notification_id}/unarchive").
		WithConfig(config).
		AddPathParam("notification_id", notificationId).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[any](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewClientError[any](err)
	}

	return shared.NewClientResponse[any](resp), nil
}

// Marks a notification as unread.
func (api *NotificationsService) MarkNotificationUnread(ctx context.Context, notificationId string) (*shared.ClientResponse[any], *shared.ClientError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("POST").
		WithPath("/notifications/{notification_id}/unread").
		WithConfig(config).
		AddPathParam("notification_id", notificationId).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[any](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewClientError[any](err)
	}

	return shared.NewClientResponse[any](resp), nil
}
