package channels

import (
	"encoding/json"
	"github.com/magicbell/magicbell-go/pkg/project-client/util"
)

type CategoryDeliveryConfig struct {
	Channels []CategoryDeliveryConfigChannels `json:"channels,omitempty" required:"true"`
	Disabled *bool                            `json:"disabled,omitempty"`
	Key      *string                          `json:"key,omitempty" required:"true" minLength:"3" pattern:"^[A-Za-z0-9_.\-:/]+$"`
}

func (c *CategoryDeliveryConfig) GetChannels() []CategoryDeliveryConfigChannels {
	if c == nil {
		return nil
	}
	return c.Channels
}

func (c *CategoryDeliveryConfig) SetChannels(channels []CategoryDeliveryConfigChannels) {
	c.Channels = channels
}

func (c *CategoryDeliveryConfig) GetDisabled() *bool {
	if c == nil {
		return nil
	}
	return c.Disabled
}

func (c *CategoryDeliveryConfig) SetDisabled(disabled bool) {
	c.Disabled = &disabled
}

func (c *CategoryDeliveryConfig) GetKey() *string {
	if c == nil {
		return nil
	}
	return c.Key
}

func (c *CategoryDeliveryConfig) SetKey(key string) {
	c.Key = &key
}

func (c CategoryDeliveryConfig) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CategoryDeliveryConfig to string"
	}
	return string(jsonData)
}

type CategoryDeliveryConfigChannels struct {
	Channel *Channel               `json:"channel,omitempty" required:"true"`
	Delay   *int64                 `json:"delay,omitempty" min:"0"`
	If_     *util.Nullable[string] `json:"if,omitempty"`
}

func (c *CategoryDeliveryConfigChannels) GetChannel() *Channel {
	if c == nil {
		return nil
	}
	return c.Channel
}

func (c *CategoryDeliveryConfigChannels) SetChannel(channel Channel) {
	c.Channel = &channel
}

func (c *CategoryDeliveryConfigChannels) GetDelay() *int64 {
	if c == nil {
		return nil
	}
	return c.Delay
}

func (c *CategoryDeliveryConfigChannels) SetDelay(delay int64) {
	c.Delay = &delay
}

func (c *CategoryDeliveryConfigChannels) GetIf_() *util.Nullable[string] {
	if c == nil {
		return nil
	}
	return c.If_
}

func (c *CategoryDeliveryConfigChannels) SetIf_(if_ util.Nullable[string]) {
	c.If_ = &if_
}

func (c *CategoryDeliveryConfigChannels) SetIf_Null() {
	c.If_ = &util.Nullable[string]{IsNull: true}
}

func (c CategoryDeliveryConfigChannels) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CategoryDeliveryConfigChannels to string"
	}
	return string(jsonData)
}

type Channel string

const (
	CHANNEL_IN_APP      Channel = "in_app"
	CHANNEL_SLACK       Channel = "slack"
	CHANNEL_WEB_PUSH    Channel = "web_push"
	CHANNEL_MOBILE_PUSH Channel = "mobile_push"
	CHANNEL_TEAMS       Channel = "teams"
	CHANNEL_EMAIL       Channel = "email"
	CHANNEL_SMS         Channel = "sms"
)
