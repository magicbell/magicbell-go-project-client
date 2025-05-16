package configmanager

import (
	"github.com/magicbell/magicbell-go/pkg/user-client/clientconfig"
	"time"
)

type ConfigManager struct {
	Channels      clientconfig.Config
	Integrations  clientconfig.Config
	Notifications clientconfig.Config
}

func NewConfigManager(config clientconfig.Config) *ConfigManager {
	return &ConfigManager{
		Channels:      config,
		Integrations:  config,
		Notifications: config,
	}
}

func (c *ConfigManager) SetBaseUrl(baseUrl string) {
	c.Channels.SetBaseUrl(baseUrl)
	c.Integrations.SetBaseUrl(baseUrl)
	c.Notifications.SetBaseUrl(baseUrl)
}

func (c *ConfigManager) SetTimeout(timeout time.Duration) {
	c.Channels.SetTimeout(timeout)
	c.Integrations.SetTimeout(timeout)
	c.Notifications.SetTimeout(timeout)
}

func (c *ConfigManager) SetAccessToken(accessToken string) {
	c.Channels.SetAccessToken(accessToken)
	c.Integrations.SetAccessToken(accessToken)
	c.Notifications.SetAccessToken(accessToken)
}

func (c *ConfigManager) UpdateAccessToken(originalValue string, newValue string) {

	if c.Channels.AccessToken != nil && *c.Channels.AccessToken == originalValue {
		c.Channels.SetAccessToken(newValue)
	}

	if c.Integrations.AccessToken != nil && *c.Integrations.AccessToken == originalValue {
		c.Integrations.SetAccessToken(newValue)
	}

	if c.Notifications.AccessToken != nil && *c.Notifications.AccessToken == originalValue {
		c.Notifications.SetAccessToken(newValue)
	}
}

func (c *ConfigManager) GetChannels() *clientconfig.Config {
	return &c.Channels
}
func (c *ConfigManager) GetIntegrations() *clientconfig.Config {
	return &c.Integrations
}
func (c *ConfigManager) GetNotifications() *clientconfig.Config {
	return &c.Notifications
}
