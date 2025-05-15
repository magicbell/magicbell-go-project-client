package configmanager

import (
	"github.com/magicbell/magicbell-go/pkg/project-client/projectclientconfig"
	"time"
)

type ConfigManager struct {
	Broadcasts    projectclientconfig.Config
	Channels      projectclientconfig.Config
	Events        projectclientconfig.Config
	Integrations  projectclientconfig.Config
	Jwt           projectclientconfig.Config
	Notifications projectclientconfig.Config
	Users         projectclientconfig.Config
}

func NewConfigManager(config projectclientconfig.Config) *ConfigManager {
	return &ConfigManager{
		Broadcasts:    config,
		Channels:      config,
		Events:        config,
		Integrations:  config,
		Jwt:           config,
		Notifications: config,
		Users:         config,
	}
}

func (c *ConfigManager) SetBaseUrl(baseUrl string) {
	c.Broadcasts.SetBaseUrl(baseUrl)
	c.Channels.SetBaseUrl(baseUrl)
	c.Events.SetBaseUrl(baseUrl)
	c.Integrations.SetBaseUrl(baseUrl)
	c.Jwt.SetBaseUrl(baseUrl)
	c.Notifications.SetBaseUrl(baseUrl)
	c.Users.SetBaseUrl(baseUrl)
}

func (c *ConfigManager) SetTimeout(timeout time.Duration) {
	c.Broadcasts.SetTimeout(timeout)
	c.Channels.SetTimeout(timeout)
	c.Events.SetTimeout(timeout)
	c.Integrations.SetTimeout(timeout)
	c.Jwt.SetTimeout(timeout)
	c.Notifications.SetTimeout(timeout)
	c.Users.SetTimeout(timeout)
}

func (c *ConfigManager) SetAccessToken(accessToken string) {
	c.Broadcasts.SetAccessToken(accessToken)
	c.Channels.SetAccessToken(accessToken)
	c.Events.SetAccessToken(accessToken)
	c.Integrations.SetAccessToken(accessToken)
	c.Jwt.SetAccessToken(accessToken)
	c.Notifications.SetAccessToken(accessToken)
	c.Users.SetAccessToken(accessToken)
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

	if c.Notifications.AccessToken != nil && *c.Notifications.AccessToken == originalValue {
		c.Notifications.SetAccessToken(newValue)
	}

	if c.Users.AccessToken != nil && *c.Users.AccessToken == originalValue {
		c.Users.SetAccessToken(newValue)
	}
}

func (c *ConfigManager) GetBroadcasts() *projectclientconfig.Config {
	return &c.Broadcasts
}
func (c *ConfigManager) GetChannels() *projectclientconfig.Config {
	return &c.Channels
}
func (c *ConfigManager) GetEvents() *projectclientconfig.Config {
	return &c.Events
}
func (c *ConfigManager) GetIntegrations() *projectclientconfig.Config {
	return &c.Integrations
}
func (c *ConfigManager) GetJwt() *projectclientconfig.Config {
	return &c.Jwt
}
func (c *ConfigManager) GetNotifications() *projectclientconfig.Config {
	return &c.Notifications
}
func (c *ConfigManager) GetUsers() *projectclientconfig.Config {
	return &c.Users
}
