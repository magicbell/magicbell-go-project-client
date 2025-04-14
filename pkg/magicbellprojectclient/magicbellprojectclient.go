package magicbellprojectclient

import (
	"github.com/magicbell/magicbell-go-project-client/internal/configmanager"
	"github.com/magicbell/magicbell-go-project-client/pkg/broadcasts"
	"github.com/magicbell/magicbell-go-project-client/pkg/channels"
	"github.com/magicbell/magicbell-go-project-client/pkg/events"
	"github.com/magicbell/magicbell-go-project-client/pkg/integrations"
	"github.com/magicbell/magicbell-go-project-client/pkg/jwt"
	"github.com/magicbell/magicbell-go-project-client/pkg/magicbellprojectclientconfig"
	"github.com/magicbell/magicbell-go-project-client/pkg/notifications"
	"github.com/magicbell/magicbell-go-project-client/pkg/users"
	"time"
)

type MagicbellProjectClient struct {
	Broadcasts    *broadcasts.BroadcastsService
	Channels      *channels.ChannelsService
	Events        *events.EventsService
	Integrations  *integrations.IntegrationsService
	Jwt           *jwt.JwtService
	Notifications *notifications.NotificationsService
	Users         *users.UsersService
	manager       *configmanager.ConfigManager
}

func NewMagicbellProjectClient(config magicbellprojectclientconfig.Config) *MagicbellProjectClient {
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

	return &MagicbellProjectClient{
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

func (m *MagicbellProjectClient) SetBaseUrl(baseUrl string) {
	m.manager.SetBaseUrl(baseUrl)
}

func (m *MagicbellProjectClient) SetTimeout(timeout time.Duration) {
	m.manager.SetTimeout(timeout)
}

func (m *MagicbellProjectClient) SetAccessToken(accessToken string) {
	m.manager.SetAccessToken(accessToken)
}

// c029837e0e474b76bc487506e8799df5e3335891efe4fb02bda7a1441840310c
