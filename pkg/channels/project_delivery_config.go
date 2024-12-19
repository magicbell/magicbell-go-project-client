package channels

import (
	"encoding/json"
)

type ProjectDeliveryConfig struct {
	Channels []ProjectDeliveryConfigChannels `json:"channels,omitempty" required:"true"`
	touched  map[string]bool
}

func (p *ProjectDeliveryConfig) GetChannels() []ProjectDeliveryConfigChannels {
	if p == nil {
		return nil
	}
	return p.Channels
}

func (p *ProjectDeliveryConfig) SetChannels(channels []ProjectDeliveryConfigChannels) {
	if p.touched == nil {
		p.touched = map[string]bool{}
	}
	p.touched["Channels"] = true
	p.Channels = channels
}

func (p *ProjectDeliveryConfig) SetChannelsNil() {
	if p.touched == nil {
		p.touched = map[string]bool{}
	}
	p.touched["Channels"] = true
	p.Channels = nil
}
func (p ProjectDeliveryConfig) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if p.touched["Channels"] && p.Channels == nil {
		data["channels"] = nil
	} else if p.Channels != nil {
		data["channels"] = p.Channels
	}

	return json.Marshal(data)
}

type ProjectDeliveryConfigChannels struct {
	Channel *ChannelsChannel1 `json:"channel,omitempty" required:"true"`
	// Delay (in seconds) since the last step, before the message is sent to the channel
	Delay    *int64  `json:"delay,omitempty" min:"0"`
	Disabled *bool   `json:"disabled,omitempty"`
	If_      *string `json:"if,omitempty"`
	Priority *int64  `json:"priority,omitempty" min:"0"`
	touched  map[string]bool
}

func (p *ProjectDeliveryConfigChannels) GetChannel() *ChannelsChannel1 {
	if p == nil {
		return nil
	}
	return p.Channel
}

func (p *ProjectDeliveryConfigChannels) SetChannel(channel ChannelsChannel1) {
	if p.touched == nil {
		p.touched = map[string]bool{}
	}
	p.touched["Channel"] = true
	p.Channel = &channel
}

func (p *ProjectDeliveryConfigChannels) SetChannelNil() {
	if p.touched == nil {
		p.touched = map[string]bool{}
	}
	p.touched["Channel"] = true
	p.Channel = nil
}

func (p *ProjectDeliveryConfigChannels) GetDelay() *int64 {
	if p == nil {
		return nil
	}
	return p.Delay
}

func (p *ProjectDeliveryConfigChannels) SetDelay(delay int64) {
	if p.touched == nil {
		p.touched = map[string]bool{}
	}
	p.touched["Delay"] = true
	p.Delay = &delay
}

func (p *ProjectDeliveryConfigChannels) SetDelayNil() {
	if p.touched == nil {
		p.touched = map[string]bool{}
	}
	p.touched["Delay"] = true
	p.Delay = nil
}

func (p *ProjectDeliveryConfigChannels) GetDisabled() *bool {
	if p == nil {
		return nil
	}
	return p.Disabled
}

func (p *ProjectDeliveryConfigChannels) SetDisabled(disabled bool) {
	if p.touched == nil {
		p.touched = map[string]bool{}
	}
	p.touched["Disabled"] = true
	p.Disabled = &disabled
}

func (p *ProjectDeliveryConfigChannels) SetDisabledNil() {
	if p.touched == nil {
		p.touched = map[string]bool{}
	}
	p.touched["Disabled"] = true
	p.Disabled = nil
}

func (p *ProjectDeliveryConfigChannels) GetIf_() *string {
	if p == nil {
		return nil
	}
	return p.If_
}

func (p *ProjectDeliveryConfigChannels) SetIf_(if_ string) {
	if p.touched == nil {
		p.touched = map[string]bool{}
	}
	p.touched["If_"] = true
	p.If_ = &if_
}

func (p *ProjectDeliveryConfigChannels) SetIf_Nil() {
	if p.touched == nil {
		p.touched = map[string]bool{}
	}
	p.touched["If_"] = true
	p.If_ = nil
}

func (p *ProjectDeliveryConfigChannels) GetPriority() *int64 {
	if p == nil {
		return nil
	}
	return p.Priority
}

func (p *ProjectDeliveryConfigChannels) SetPriority(priority int64) {
	if p.touched == nil {
		p.touched = map[string]bool{}
	}
	p.touched["Priority"] = true
	p.Priority = &priority
}

func (p *ProjectDeliveryConfigChannels) SetPriorityNil() {
	if p.touched == nil {
		p.touched = map[string]bool{}
	}
	p.touched["Priority"] = true
	p.Priority = nil
}
func (p ProjectDeliveryConfigChannels) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if p.touched["Channel"] && p.Channel == nil {
		data["channel"] = nil
	} else if p.Channel != nil {
		data["channel"] = p.Channel
	}

	if p.touched["Delay"] && p.Delay == nil {
		data["delay"] = nil
	} else if p.Delay != nil {
		data["delay"] = p.Delay
	}

	if p.touched["Disabled"] && p.Disabled == nil {
		data["disabled"] = nil
	} else if p.Disabled != nil {
		data["disabled"] = p.Disabled
	}

	if p.touched["If_"] && p.If_ == nil {
		data["if"] = nil
	} else if p.If_ != nil {
		data["if"] = p.If_
	}

	if p.touched["Priority"] && p.Priority == nil {
		data["priority"] = nil
	} else if p.Priority != nil {
		data["priority"] = p.Priority
	}

	return json.Marshal(data)
}

type ChannelsChannel1 string

const (
	CHANNELS_CHANNEL1_IN_APP      ChannelsChannel1 = "in_app"
	CHANNELS_CHANNEL1_SLACK       ChannelsChannel1 = "slack"
	CHANNELS_CHANNEL1_WEB_PUSH    ChannelsChannel1 = "web_push"
	CHANNELS_CHANNEL1_MOBILE_PUSH ChannelsChannel1 = "mobile_push"
	CHANNELS_CHANNEL1_TEAMS       ChannelsChannel1 = "teams"
	CHANNELS_CHANNEL1_EMAIL       ChannelsChannel1 = "email"
	CHANNELS_CHANNEL1_SMS         ChannelsChannel1 = "sms"
)
