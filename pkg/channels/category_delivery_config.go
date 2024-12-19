package channels

import (
	"encoding/json"
)

type CategoryDeliveryConfig struct {
	Category *string                          `json:"category,omitempty" required:"true" minLength:"3" pattern:"^[a-zA-Z0-9_]+$"`
	Channels []CategoryDeliveryConfigChannels `json:"channels,omitempty" required:"true"`
	Disabled *bool                            `json:"disabled,omitempty"`
	touched  map[string]bool
}

func (c *CategoryDeliveryConfig) GetCategory() *string {
	if c == nil {
		return nil
	}
	return c.Category
}

func (c *CategoryDeliveryConfig) SetCategory(category string) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Category"] = true
	c.Category = &category
}

func (c *CategoryDeliveryConfig) SetCategoryNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Category"] = true
	c.Category = nil
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
func (c CategoryDeliveryConfig) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if c.touched["Category"] && c.Category == nil {
		data["category"] = nil
	} else if c.Category != nil {
		data["category"] = c.Category
	}

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

	return json.Marshal(data)
}

type CategoryDeliveryConfigChannels struct {
	Channel  *ChannelsChannel2 `json:"channel,omitempty" required:"true"`
	Delay    *int64            `json:"delay,omitempty" min:"0"`
	Disabled *bool             `json:"disabled,omitempty"`
	If_      *string           `json:"if,omitempty"`
	Priority *int64            `json:"priority,omitempty" min:"0"`
	touched  map[string]bool
}

func (c *CategoryDeliveryConfigChannels) GetChannel() *ChannelsChannel2 {
	if c == nil {
		return nil
	}
	return c.Channel
}

func (c *CategoryDeliveryConfigChannels) SetChannel(channel ChannelsChannel2) {
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

func (c *CategoryDeliveryConfigChannels) GetDisabled() *bool {
	if c == nil {
		return nil
	}
	return c.Disabled
}

func (c *CategoryDeliveryConfigChannels) SetDisabled(disabled bool) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Disabled"] = true
	c.Disabled = &disabled
}

func (c *CategoryDeliveryConfigChannels) SetDisabledNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Disabled"] = true
	c.Disabled = nil
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

func (c *CategoryDeliveryConfigChannels) GetPriority() *int64 {
	if c == nil {
		return nil
	}
	return c.Priority
}

func (c *CategoryDeliveryConfigChannels) SetPriority(priority int64) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Priority"] = true
	c.Priority = &priority
}

func (c *CategoryDeliveryConfigChannels) SetPriorityNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Priority"] = true
	c.Priority = nil
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

	if c.touched["Disabled"] && c.Disabled == nil {
		data["disabled"] = nil
	} else if c.Disabled != nil {
		data["disabled"] = c.Disabled
	}

	if c.touched["If_"] && c.If_ == nil {
		data["if"] = nil
	} else if c.If_ != nil {
		data["if"] = c.If_
	}

	if c.touched["Priority"] && c.Priority == nil {
		data["priority"] = nil
	} else if c.Priority != nil {
		data["priority"] = c.Priority
	}

	return json.Marshal(data)
}

type ChannelsChannel2 string

const (
	CHANNELS_CHANNEL2_IN_APP      ChannelsChannel2 = "in_app"
	CHANNELS_CHANNEL2_SLACK       ChannelsChannel2 = "slack"
	CHANNELS_CHANNEL2_WEB_PUSH    ChannelsChannel2 = "web_push"
	CHANNELS_CHANNEL2_MOBILE_PUSH ChannelsChannel2 = "mobile_push"
	CHANNELS_CHANNEL2_TEAMS       ChannelsChannel2 = "teams"
	CHANNELS_CHANNEL2_EMAIL       ChannelsChannel2 = "email"
)
