package userclient

import (
	"github.com/magicbell/magicbell-go/pkg/user-client/internal/configmanager"
	"github.com/magicbell/magicbell-go/pkg/user-client/channels"
	"github.com/magicbell/magicbell-go/pkg/user-client/integrations"
	"github.com/magicbell/magicbell-go/pkg/user-client/notifications"
	"github.com/magicbell/magicbell-go/pkg/user-client/userclientconfig"
	"time"
)

type UserClient struct {
	Channels      *channels.ChannelsService
	Integrations  *integrations.IntegrationsService
	Notifications *notifications.NotificationsService
	manager       *configmanager.ConfigManager
}

func NewUserClient(config userclientconfig.Config) *UserClient {
	channels := channels.NewChannelsService()
	integrations := integrations.NewIntegrationsService()
	notifications := notifications.NewNotificationsService()

	manager := configmanager.NewConfigManager(config)
	channels.WithConfigManager(manager)
	integrations.WithConfigManager(manager)
	notifications.WithConfigManager(manager)

	return &UserClient{
		Channels:      channels,
		Integrations:  integrations,
		Notifications: notifications,
		manager:       manager,
	}
}

func (u *UserClient) SetBaseUrl(baseUrl string) {
	u.manager.SetBaseUrl(baseUrl)
}

func (u *UserClient) SetTimeout(timeout time.Duration) {
	u.manager.SetTimeout(timeout)
}

func (u *UserClient) SetAccessToken(accessToken string) {
	u.manager.SetAccessToken(accessToken)
}

// c029837e0e474b76bc487506e8799df5e3335891efe4fb02bda7a1441840310c
