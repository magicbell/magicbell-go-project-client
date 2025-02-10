package channels

import (
	"encoding/json"
)

type SlackTokenResponse struct {
	CreatedAt   *string                    `json:"created_at,omitempty" required:"true"`
	DiscardedAt *string                    `json:"discarded_at,omitempty"`
	Id          *string                    `json:"id,omitempty" required:"true"`
	Oauth       *Oauth                     `json:"oauth,omitempty"`
	UpdatedAt   *string                    `json:"updated_at,omitempty"`
	Webhook     *SlackTokenResponseWebhook `json:"webhook,omitempty"`
	touched     map[string]bool
}

func (s *SlackTokenResponse) GetCreatedAt() *string {
	if s == nil {
		return nil
	}
	return s.CreatedAt
}

func (s *SlackTokenResponse) SetCreatedAt(createdAt string) {
	if s.touched == nil {
		s.touched = map[string]bool{}
	}
	s.touched["CreatedAt"] = true
	s.CreatedAt = &createdAt
}

func (s *SlackTokenResponse) SetCreatedAtNil() {
	if s.touched == nil {
		s.touched = map[string]bool{}
	}
	s.touched["CreatedAt"] = true
	s.CreatedAt = nil
}

func (s *SlackTokenResponse) GetDiscardedAt() *string {
	if s == nil {
		return nil
	}
	return s.DiscardedAt
}

func (s *SlackTokenResponse) SetDiscardedAt(discardedAt string) {
	if s.touched == nil {
		s.touched = map[string]bool{}
	}
	s.touched["DiscardedAt"] = true
	s.DiscardedAt = &discardedAt
}

func (s *SlackTokenResponse) SetDiscardedAtNil() {
	if s.touched == nil {
		s.touched = map[string]bool{}
	}
	s.touched["DiscardedAt"] = true
	s.DiscardedAt = nil
}

func (s *SlackTokenResponse) GetId() *string {
	if s == nil {
		return nil
	}
	return s.Id
}

func (s *SlackTokenResponse) SetId(id string) {
	if s.touched == nil {
		s.touched = map[string]bool{}
	}
	s.touched["Id"] = true
	s.Id = &id
}

func (s *SlackTokenResponse) SetIdNil() {
	if s.touched == nil {
		s.touched = map[string]bool{}
	}
	s.touched["Id"] = true
	s.Id = nil
}

func (s *SlackTokenResponse) GetOauth() *Oauth {
	if s == nil {
		return nil
	}
	return s.Oauth
}

func (s *SlackTokenResponse) SetOauth(oauth Oauth) {
	if s.touched == nil {
		s.touched = map[string]bool{}
	}
	s.touched["Oauth"] = true
	s.Oauth = &oauth
}

func (s *SlackTokenResponse) SetOauthNil() {
	if s.touched == nil {
		s.touched = map[string]bool{}
	}
	s.touched["Oauth"] = true
	s.Oauth = nil
}

func (s *SlackTokenResponse) GetUpdatedAt() *string {
	if s == nil {
		return nil
	}
	return s.UpdatedAt
}

func (s *SlackTokenResponse) SetUpdatedAt(updatedAt string) {
	if s.touched == nil {
		s.touched = map[string]bool{}
	}
	s.touched["UpdatedAt"] = true
	s.UpdatedAt = &updatedAt
}

func (s *SlackTokenResponse) SetUpdatedAtNil() {
	if s.touched == nil {
		s.touched = map[string]bool{}
	}
	s.touched["UpdatedAt"] = true
	s.UpdatedAt = nil
}

func (s *SlackTokenResponse) GetWebhook() *SlackTokenResponseWebhook {
	if s == nil {
		return nil
	}
	return s.Webhook
}

func (s *SlackTokenResponse) SetWebhook(webhook SlackTokenResponseWebhook) {
	if s.touched == nil {
		s.touched = map[string]bool{}
	}
	s.touched["Webhook"] = true
	s.Webhook = &webhook
}

func (s *SlackTokenResponse) SetWebhookNil() {
	if s.touched == nil {
		s.touched = map[string]bool{}
	}
	s.touched["Webhook"] = true
	s.Webhook = nil
}

func (s SlackTokenResponse) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if s.touched["CreatedAt"] && s.CreatedAt == nil {
		data["created_at"] = nil
	} else if s.CreatedAt != nil {
		data["created_at"] = s.CreatedAt
	}

	if s.touched["DiscardedAt"] && s.DiscardedAt == nil {
		data["discarded_at"] = nil
	} else if s.DiscardedAt != nil {
		data["discarded_at"] = s.DiscardedAt
	}

	if s.touched["Id"] && s.Id == nil {
		data["id"] = nil
	} else if s.Id != nil {
		data["id"] = s.Id
	}

	if s.touched["Oauth"] && s.Oauth == nil {
		data["oauth"] = nil
	} else if s.Oauth != nil {
		data["oauth"] = s.Oauth
	}

	if s.touched["UpdatedAt"] && s.UpdatedAt == nil {
		data["updated_at"] = nil
	} else if s.UpdatedAt != nil {
		data["updated_at"] = s.UpdatedAt
	}

	if s.touched["Webhook"] && s.Webhook == nil {
		data["webhook"] = nil
	} else if s.Webhook != nil {
		data["webhook"] = s.Webhook
	}

	return json.Marshal(data)
}

func (s SlackTokenResponse) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SlackTokenResponse to string"
	}
	return string(jsonData)
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

func (o Oauth) String() string {
	jsonData, err := json.MarshalIndent(o, "", "  ")
	if err != nil {
		return "error converting struct: Oauth to string"
	}
	return string(jsonData)
}

type SlackTokenResponseWebhook struct {
	Url     *string `json:"url,omitempty" required:"true" minLength:"1"`
	touched map[string]bool
}

func (s *SlackTokenResponseWebhook) GetUrl() *string {
	if s == nil {
		return nil
	}
	return s.Url
}

func (s *SlackTokenResponseWebhook) SetUrl(url string) {
	if s.touched == nil {
		s.touched = map[string]bool{}
	}
	s.touched["Url"] = true
	s.Url = &url
}

func (s *SlackTokenResponseWebhook) SetUrlNil() {
	if s.touched == nil {
		s.touched = map[string]bool{}
	}
	s.touched["Url"] = true
	s.Url = nil
}

func (s SlackTokenResponseWebhook) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if s.touched["Url"] && s.Url == nil {
		data["url"] = nil
	} else if s.Url != nil {
		data["url"] = s.Url
	}

	return json.Marshal(data)
}

func (s SlackTokenResponseWebhook) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SlackTokenResponseWebhook to string"
	}
	return string(jsonData)
}
