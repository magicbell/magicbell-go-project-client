package channels

import (
	"encoding/json"
)

type CategoryDeliveryConfig struct {
	Channels []CategoryDeliveryConfigChannels `json:"channels,omitempty" required:"true"`
	Disabled *bool                            `json:"disabled,omitempty"`
	Key      *string                          `json:"key,omitempty" required:"true" minLength:"3" pattern:"^[A-Za-z0-9_\.\-:]+$"`
	touched  map[string]bool
}

func (c *CategoryDeliveryConfig) GetChannels() []CategoryDeliveryConfigChannels {
	if c == nil {
		return nil
	}
	return c.Channels
}

func (c *CategoryDeliveryConfig) SetChannels(channels []CategoryDeliveryConfigChannels) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Channels"] = true
	c.Channels = channels
}

func (c *CategoryDeliveryConfig) SetChannelsNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Channels"] = true
	c.Channels = nil
}

func (c *CategoryDeliveryConfig) GetDisabled() *bool {
	if c == nil {
		return nil
	}
	return c.Disabled
}

func (c *CategoryDeliveryConfig) SetDisabled(disabled bool) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Disabled"] = true
	c.Disabled = &disabled
}

func (c *CategoryDeliveryConfig) SetDisabledNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Disabled"] = true
	c.Disabled = nil
}

func (c *CategoryDeliveryConfig) GetKey() *string {
	if c == nil {
		return nil
	}
	return c.Key
}

func (c *CategoryDeliveryConfig) SetKey(key string) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Key"] = true
	c.Key = &key
}

func (c *CategoryDeliveryConfig) SetKeyNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Key"] = true
	c.Key = nil
}

func (c CategoryDeliveryConfig) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if c.touched["Channels"] && c.Channels == nil {
		data["channels"] = nil
	} else if c.Channels != nil {
		data["channels"] = c.Channels
	}

	if c.touched["Disabled"] && c.Disabled == nil {
		data["disabled"] = nil
	} else if c.Disabled != nil {
		data["disabled"] = c.Disabled
	}

	if c.touched["Key"] && c.Key == nil {
		data["key"] = nil
	} else if c.Key != nil {
		data["key"] = c.Key
	}

	return json.Marshal(data)
}

func (c CategoryDeliveryConfig) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CategoryDeliveryConfig to string"
	}
	return string(jsonData)
}

type CategoryDeliveryConfigChannels struct {
	Channel *Channel `json:"channel,omitempty" required:"true"`
	Delay   *int64   `json:"delay,omitempty" min:"0"`
	If_     *string  `json:"if,omitempty"`
	touched map[string]bool
}

func (c *CategoryDeliveryConfigChannels) GetChannel() *Channel {
	if c == nil {
		return nil
	}
	return c.Channel
}

func (c *CategoryDeliveryConfigChannels) SetChannel(channel Channel) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Channel"] = true
	c.Channel = &channel
}

func (c *CategoryDeliveryConfigChannels) SetChannelNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Channel"] = true
	c.Channel = nil
}

func (c *CategoryDeliveryConfigChannels) GetDelay() *int64 {
	if c == nil {
		return nil
	}
	return c.Delay
}

func (c *CategoryDeliveryConfigChannels) SetDelay(delay int64) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Delay"] = true
	c.Delay = &delay
}

func (c *CategoryDeliveryConfigChannels) SetDelayNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Delay"] = true
	c.Delay = nil
}

func (c *CategoryDeliveryConfigChannels) GetIf_() *string {
	if c == nil {
		return nil
	}
	return c.If_
}

func (c *CategoryDeliveryConfigChannels) SetIf_(if_ string) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["If_"] = true
	c.If_ = &if_
}

func (c *CategoryDeliveryConfigChannels) SetIf_Nil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["If_"] = true
	c.If_ = nil
}

func (c CategoryDeliveryConfigChannels) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if c.touched["Channel"] && c.Channel == nil {
		data["channel"] = nil
	} else if c.Channel != nil {
		data["channel"] = c.Channel
	}

	if c.touched["Delay"] && c.Delay == nil {
		data["delay"] = nil
	} else if c.Delay != nil {
		data["delay"] = c.Delay
	}

	if c.touched["If_"] && c.If_ == nil {
		data["if"] = nil
	} else if c.If_ != nil {
		data["if"] = c.If_
	}

	return json.Marshal(data)
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
