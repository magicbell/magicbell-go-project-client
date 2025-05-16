package client

import (
	"github.com/magicbell/magicbell-go/pkg/user-client/internal/configmanager"
	"github.com/magicbell/magicbell-go/pkg/user-client/channels"
	"github.com/magicbell/magicbell-go/pkg/user-client/clientconfig"
	"github.com/magicbell/magicbell-go/pkg/user-client/integrations"
	"github.com/magicbell/magicbell-go/pkg/user-client/notifications"
	"time"
)

type Client struct {
	Channels      *channels.ChannelsService
	Integrations  *integrations.IntegrationsService
	Notifications *notifications.NotificationsService
	manager       *configmanager.ConfigManager
}

func NewClient(config clientconfig.Config) *Client {
	channels := channels.NewChannelsService()
	integrations := integrations.NewIntegrationsService()
	notifications := notifications.NewNotificationsService()

	manager := configmanager.NewConfigManager(config)
	channels.WithConfigManager(manager)
	integrations.WithConfigManager(manager)
	notifications.WithConfigManager(manager)

	return &Client{
		Channels:      channels,
		Integrations:  integrations,
		Notifications: notifications,
		manager:       manager,
	}
}

func (c *Client) SetBaseUrl(baseUrl string) {
	c.manager.SetBaseUrl(baseUrl)
}

func (c *Client) SetTimeout(timeout time.Duration) {
	c.manager.SetTimeout(timeout)
}

func (c *Client) SetAccessToken(accessToken string) {
	c.manager.SetAccessToken(accessToken)
}

// c029837e0e474b76bc487506e8799df5e3335891efe4fb02bda7a1441840310c
