package projectclient

import (
	"github.com/magicbell/magicbell-go/pkg/project-client/internal/configmanager"
	"github.com/magicbell/magicbell-go/pkg/project-client/broadcasts"
	"github.com/magicbell/magicbell-go/pkg/project-client/channels"
	"github.com/magicbell/magicbell-go/pkg/project-client/events"
	"github.com/magicbell/magicbell-go/pkg/project-client/integrations"
	"github.com/magicbell/magicbell-go/pkg/project-client/jwt"
	"github.com/magicbell/magicbell-go/pkg/project-client/notifications"
	"github.com/magicbell/magicbell-go/pkg/project-client/projectclientconfig"
	"github.com/magicbell/magicbell-go/pkg/project-client/users"
	"time"
)

type ProjectClient struct {
	Broadcasts    *broadcasts.BroadcastsService
	Channels      *channels.ChannelsService
	Events        *events.EventsService
	Integrations  *integrations.IntegrationsService
	Jwt           *jwt.JwtService
	Notifications *notifications.NotificationsService
	Users         *users.UsersService
	manager       *configmanager.ConfigManager
}

func NewProjectClient(config projectclientconfig.Config) *ProjectClient {
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

	return &ProjectClient{
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

func (p *ProjectClient) SetBaseUrl(baseUrl string) {
	p.manager.SetBaseUrl(baseUrl)
}

func (p *ProjectClient) SetTimeout(timeout time.Duration) {
	p.manager.SetTimeout(timeout)
}

func (p *ProjectClient) SetAccessToken(accessToken string) {
	p.manager.SetAccessToken(accessToken)
}

// c029837e0e474b76bc487506e8799df5e3335891efe4fb02bda7a1441840310c
