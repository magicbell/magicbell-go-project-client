package configmanager

import (
	"time"

	"github.com/magicbell/magicbell-go-project-client/pkg/magicbellprojectclientconfig"
)

type ConfigManager struct {
	Broadcasts   magicbellprojectclientconfig.Config
	Channels     magicbellprojectclientconfig.Config
	Events       magicbellprojectclientconfig.Config
	Integrations magicbellprojectclientconfig.Config
	Jwt          magicbellprojectclientconfig.Config
}

func NewConfigManager(config magicbellprojectclientconfig.Config) *ConfigManager {
	return &ConfigManager{
		Broadcasts:   config,
		Channels:     config,
		Events:       config,
		Integrations: config,
		Jwt:          config,
	}
}

func (c *ConfigManager) SetBaseUrl(baseUrl string) {
	c.Broadcasts.SetBaseUrl(baseUrl)
	c.Channels.SetBaseUrl(baseUrl)
	c.Events.SetBaseUrl(baseUrl)
	c.Integrations.SetBaseUrl(baseUrl)
	c.Jwt.SetBaseUrl(baseUrl)
}

func (c *ConfigManager) SetTimeout(timeout time.Duration) {
	c.Broadcasts.SetTimeout(timeout)
	c.Channels.SetTimeout(timeout)
	c.Events.SetTimeout(timeout)
	c.Integrations.SetTimeout(timeout)
	c.Jwt.SetTimeout(timeout)
}

func (c *ConfigManager) SetAccessToken(accessToken string) {
	c.Broadcasts.SetAccessToken(accessToken)
	c.Channels.SetAccessToken(accessToken)
	c.Events.SetAccessToken(accessToken)
	c.Integrations.SetAccessToken(accessToken)
	c.Jwt.SetAccessToken(accessToken)
}

func (c *ConfigManager) UpdateAccessToken(originalValue string, newValue string) {

	if c.Broadcasts.AccessToken != nil && *c.Broadcasts.AccessToken == originalValue {
		c.Broadcasts.SetAccessToken(newValue)
	}

	if c.Channels.AccessToken != nil && *c.Channels.AccessToken == originalValue {
		c.Channels.SetAccessToken(newValue)
	}

	if c.Events.AccessToken != nil && *c.Events.AccessToken == originalValue {
		c.Events.SetAccessToken(newValue)
	}

	if c.Integrations.AccessToken != nil && *c.Integrations.AccessToken == originalValue {
		c.Integrations.SetAccessToken(newValue)
	}

	if c.Jwt.AccessToken != nil && *c.Jwt.AccessToken == originalValue {
		c.Jwt.SetAccessToken(newValue)
	}
}

func (c *ConfigManager) GetBroadcasts() *magicbellprojectclientconfig.Config {
	return &c.Broadcasts
}
func (c *ConfigManager) GetChannels() *magicbellprojectclientconfig.Config {
	return &c.Channels
}
func (c *ConfigManager) GetEvents() *magicbellprojectclientconfig.Config {
	return &c.Events
}
func (c *ConfigManager) GetIntegrations() *magicbellprojectclientconfig.Config {
	return &c.Integrations
}
func (c *ConfigManager) GetJwt() *magicbellprojectclientconfig.Config {
	return &c.Jwt
}
