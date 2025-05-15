package channels

import (
	"encoding/json"
	"github.com/magicbell/magicbell-go/pkg/project-client/util"
)

type SlackToken struct {
	CreatedAt   *string                `json:"created_at,omitempty" required:"true"`
	DiscardedAt *util.Nullable[string] `json:"discarded_at,omitempty"`
	Id          *string                `json:"id,omitempty" required:"true"`
	Oauth       *Oauth                 `json:"oauth,omitempty"`
	UpdatedAt   *util.Nullable[string] `json:"updated_at,omitempty"`
	// Obtained directly from the incoming_webhook object in the installation response from the Slack API.
	Webhook *SlackTokenWebhook `json:"webhook,omitempty"`
}

func (s *SlackToken) GetCreatedAt() *string {
	if s == nil {
		return nil
	}
	return s.CreatedAt
}

func (s *SlackToken) SetCreatedAt(createdAt string) {
	s.CreatedAt = &createdAt
}

func (s *SlackToken) GetDiscardedAt() *util.Nullable[string] {
	if s == nil {
		return nil
	}
	return s.DiscardedAt
}

func (s *SlackToken) SetDiscardedAt(discardedAt util.Nullable[string]) {
	s.DiscardedAt = &discardedAt
}

func (s *SlackToken) SetDiscardedAtNull() {
	s.DiscardedAt = &util.Nullable[string]{IsNull: true}
}

func (s *SlackToken) GetId() *string {
	if s == nil {
		return nil
	}
	return s.Id
}

func (s *SlackToken) SetId(id string) {
	s.Id = &id
}

func (s *SlackToken) GetOauth() *Oauth {
	if s == nil {
		return nil
	}
	return s.Oauth
}

func (s *SlackToken) SetOauth(oauth Oauth) {
	s.Oauth = &oauth
}

func (s *SlackToken) GetUpdatedAt() *util.Nullable[string] {
	if s == nil {
		return nil
	}
	return s.UpdatedAt
}

func (s *SlackToken) SetUpdatedAt(updatedAt util.Nullable[string]) {
	s.UpdatedAt = &updatedAt
}

func (s *SlackToken) SetUpdatedAtNull() {
	s.UpdatedAt = &util.Nullable[string]{IsNull: true}
}

func (s *SlackToken) GetWebhook() *SlackTokenWebhook {
	if s == nil {
		return nil
	}
	return s.Webhook
}

func (s *SlackToken) SetWebhook(webhook SlackTokenWebhook) {
	s.Webhook = &webhook
}

func (s SlackToken) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SlackToken to string"
	}
	return string(jsonData)
}

type Oauth struct {
	ChannelId      *string `json:"channel_id,omitempty" required:"true"`
	InstallationId *string `json:"installation_id,omitempty" required:"true"`
	Scope          *string `json:"scope,omitempty"`
}

func (o *Oauth) GetChannelId() *string {
	if o == nil {
		return nil
	}
	return o.ChannelId
}

func (o *Oauth) SetChannelId(channelId string) {
	o.ChannelId = &channelId
}

func (o *Oauth) GetInstallationId() *string {
	if o == nil {
		return nil
	}
	return o.InstallationId
}

func (o *Oauth) SetInstallationId(installationId string) {
	o.InstallationId = &installationId
}

func (o *Oauth) GetScope() *string {
	if o == nil {
		return nil
	}
	return o.Scope
}

func (o *Oauth) SetScope(scope string) {
	o.Scope = &scope
}

func (o Oauth) String() string {
	jsonData, err := json.MarshalIndent(o, "", "  ")
	if err != nil {
		return "error converting struct: Oauth to string"
	}
	return string(jsonData)
}

// Obtained directly from the incoming_webhook object in the installation response from the Slack API.
type SlackTokenWebhook struct {
	Url *string `json:"url,omitempty" required:"true" minLength:"1"`
}

func (s *SlackTokenWebhook) GetUrl() *string {
	if s == nil {
		return nil
	}
	return s.Url
}

func (s *SlackTokenWebhook) SetUrl(url string) {
	s.Url = &url
}

func (s SlackTokenWebhook) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SlackTokenWebhook to string"
	}
	return string(jsonData)
}
