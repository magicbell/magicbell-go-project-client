package magicbellprojectclient

import (
	"github.com/magicbell/magicbell-go-project-client/internal/configmanager"
	"github.com/magicbell/magicbell-go-project-client/pkg/broadcasts"
	"github.com/magicbell/magicbell-go-project-client/pkg/channels"
	"github.com/magicbell/magicbell-go-project-client/pkg/events"
	"github.com/magicbell/magicbell-go-project-client/pkg/integrations"
	"github.com/magicbell/magicbell-go-project-client/pkg/jwt"
	"github.com/magicbell/magicbell-go-project-client/pkg/magicbellprojectclientconfig"
	"time"
)

type MagicbellProjectClient struct {
	Broadcasts   *broadcasts.BroadcastsService
	Channels     *channels.ChannelsService
	Events       *events.EventsService
	Integrations *integrations.IntegrationsService
	Jwt          *jwt.JwtService
	manager      *configmanager.ConfigManager
}

func NewMagicbellProjectClient(config magicbellprojectclientconfig.Config) *MagicbellProjectClient {
	manager := configmanager.NewConfigManager(config)
	return &MagicbellProjectClient{
		Broadcasts:   broadcasts.NewBroadcastsService(manager),
		Channels:     channels.NewChannelsService(manager),
		Events:       events.NewEventsService(manager),
		Integrations: integrations.NewIntegrationsService(manager),
		Jwt:          jwt.NewJwtService(manager),
		manager:      manager,
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
