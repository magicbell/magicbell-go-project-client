package client

import (
	"github.com/magicbell/magicbell-go/pkg/project-client/internal/configmanager"
	"github.com/magicbell/magicbell-go/pkg/project-client/broadcasts"
	"github.com/magicbell/magicbell-go/pkg/project-client/channels"
	"github.com/magicbell/magicbell-go/pkg/project-client/clientconfig"
	"github.com/magicbell/magicbell-go/pkg/project-client/events"
	"github.com/magicbell/magicbell-go/pkg/project-client/integrations"
	"github.com/magicbell/magicbell-go/pkg/project-client/jwt"
	"github.com/magicbell/magicbell-go/pkg/project-client/notifications"
	"github.com/magicbell/magicbell-go/pkg/project-client/users"
	"time"
)

type Client struct {
	Broadcasts    *broadcasts.BroadcastsService
	Channels      *channels.ChannelsService
	Events        *events.EventsService
	Integrations  *integrations.IntegrationsService
	Jwt           *jwt.JwtService
	Notifications *notifications.NotificationsService
	Users         *users.UsersService
	manager       *configmanager.ConfigManager
}

func NewClient(config clientconfig.Config) *Client {
	broadcasts := broadcasts.NewBroadcastsService()
	channels := channels.NewChannelsService()
	events := events.NewEventsService()
	integrations := integrations.NewIntegrationsService()
	jwt := jwt.NewJwtService()
	notifications := notifications.NewNotificationsService()
	users := users.NewUsersService()

	manager := configmanager.NewConfigManager(config)
	broadcasts.WithConfigManager(manager)
	channels.WithConfigManager(manager)
	events.WithConfigManager(manager)
	integrations.WithConfigManager(manager)
	jwt.WithConfigManager(manager)
	notifications.WithConfigManager(manager)
	users.WithConfigManager(manager)

	return &Client{
		Broadcasts:    broadcasts,
		Channels:      channels,
		Events:        events,
		Integrations:  integrations,
		Jwt:           jwt,
		Notifications: notifications,
		Users:         users,
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
