package channels

import (
	"encoding/json"
)

type SlackToken struct {
	Oauth   *Oauth             `json:"oauth,omitempty"`
	Webhook *SlackTokenWebhook `json:"webhook,omitempty"`
	touched map[string]bool
}

func (s *SlackToken) GetOauth() *Oauth {
	if s == nil {
		return nil
	}
	return s.Oauth
}

func (s *SlackToken) SetOauth(oauth Oauth) {
	if s.touched == nil {
		s.touched = map[string]bool{}
	}
	s.touched["Oauth"] = true
	s.Oauth = &oauth
}

func (s *SlackToken) SetOauthNil() {
	if s.touched == nil {
		s.touched = map[string]bool{}
	}
	s.touched["Oauth"] = true
	s.Oauth = nil
}

func (s *SlackToken) GetWebhook() *SlackTokenWebhook {
	if s == nil {
		return nil
	}
	return s.Webhook
}

func (s *SlackToken) SetWebhook(webhook SlackTokenWebhook) {
	if s.touched == nil {
		s.touched = map[string]bool{}
	}
	s.touched["Webhook"] = true
	s.Webhook = &webhook
}

func (s *SlackToken) SetWebhookNil() {
	if s.touched == nil {
		s.touched = map[string]bool{}
	}
	s.touched["Webhook"] = true
	s.Webhook = nil
}
func (s SlackToken) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if s.touched["Oauth"] && s.Oauth == nil {
		data["oauth"] = nil
	} else if s.Oauth != nil {
		data["oauth"] = s.Oauth
	}

	if s.touched["Webhook"] && s.Webhook == nil {
		data["webhook"] = nil
	} else if s.Webhook != nil {
		data["webhook"] = s.Webhook
	}

	return json.Marshal(data)
}

type Oauth struct {
	ChannelId      *string `json:"channel_id,omitempty" required:"true"`
	InstallationId *string `json:"installation_id,omitempty" required:"true"`
	Scope          *string `json:"scope,omitempty"`
	touched        map[string]bool
}

func (o *Oauth) GetChannelId() *string {
	if o == nil {
		return nil
	}
	return o.ChannelId
}

func (o *Oauth) SetChannelId(channelId string) {
	if o.touched == nil {
		o.touched = map[string]bool{}
	}
	o.touched["ChannelId"] = true
	o.ChannelId = &channelId
}

func (o *Oauth) SetChannelIdNil() {
	if o.touched == nil {
		o.touched = map[string]bool{}
	}
	o.touched["ChannelId"] = true
	o.ChannelId = nil
}

func (o *Oauth) GetInstallationId() *string {
	if o == nil {
		return nil
	}
	return o.InstallationId
}

func (o *Oauth) SetInstallationId(installationId string) {
	if o.touched == nil {
		o.touched = map[string]bool{}
	}
	o.touched["InstallationId"] = true
	o.InstallationId = &installationId
}

func (o *Oauth) SetInstallationIdNil() {
	if o.touched == nil {
		o.touched = map[string]bool{}
	}
	o.touched["InstallationId"] = true
	o.InstallationId = nil
}

func (o *Oauth) GetScope() *string {
	if o == nil {
		return nil
	}
	return o.Scope
}

func (o *Oauth) SetScope(scope string) {
	if o.touched == nil {
		o.touched = map[string]bool{}
	}
	o.touched["Scope"] = true
	o.Scope = &scope
}

func (o *Oauth) SetScopeNil() {
	if o.touched == nil {
		o.touched = map[string]bool{}
	}
	o.touched["Scope"] = true
	o.Scope = nil
}
func (o Oauth) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if o.touched["ChannelId"] && o.ChannelId == nil {
		data["channel_id"] = nil
	} else if o.ChannelId != nil {
		data["channel_id"] = o.ChannelId
	}

	if o.touched["InstallationId"] && o.InstallationId == nil {
		data["installation_id"] = nil
	} else if o.InstallationId != nil {
		data["installation_id"] = o.InstallationId
	}

	if o.touched["Scope"] && o.Scope == nil {
		data["scope"] = nil
	} else if o.Scope != nil {
		data["scope"] = o.Scope
	}

	return json.Marshal(data)
}

type SlackTokenWebhook struct {
	Url     *string `json:"url,omitempty" required:"true" minLength:"1"`
	touched map[string]bool
}

func (s *SlackTokenWebhook) GetUrl() *string {
	if s == nil {
		return nil
	}
	return s.Url
}

func (s *SlackTokenWebhook) SetUrl(url string) {
	if s.touched == nil {
		s.touched = map[string]bool{}
	}
	s.touched["Url"] = true
	s.Url = &url
}

func (s *SlackTokenWebhook) SetUrlNil() {
	if s.touched == nil {
		s.touched = map[string]bool{}
	}
	s.touched["Url"] = true
	s.Url = nil
}
func (s SlackTokenWebhook) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if s.touched["Url"] && s.Url == nil {
		data["url"] = nil
	} else if s.Url != nil {
		data["url"] = s.Url
	}

	return json.Marshal(data)
}
