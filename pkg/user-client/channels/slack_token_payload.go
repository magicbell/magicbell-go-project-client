package channels

import "encoding/json"

type SlackTokenPayload struct {
	Oauth *SlackTokenPayloadOauth `json:"oauth,omitempty"`
	// Obtained directly from the incoming_webhook object in the installation response from the Slack API.
	Webhook *SlackTokenPayloadWebhook `json:"webhook,omitempty"`
}

func (s *SlackTokenPayload) GetOauth() *SlackTokenPayloadOauth {
	if s == nil {
		return nil
	}
	return s.Oauth
}

func (s *SlackTokenPayload) SetOauth(oauth SlackTokenPayloadOauth) {
	s.Oauth = &oauth
}

func (s *SlackTokenPayload) GetWebhook() *SlackTokenPayloadWebhook {
	if s == nil {
		return nil
	}
	return s.Webhook
}

func (s *SlackTokenPayload) SetWebhook(webhook SlackTokenPayloadWebhook) {
	s.Webhook = &webhook
}

func (s SlackTokenPayload) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SlackTokenPayload to string"
	}
	return string(jsonData)
}

type SlackTokenPayloadOauth struct {
	ChannelId      *string `json:"channel_id,omitempty" required:"true"`
	InstallationId *string `json:"installation_id,omitempty" required:"true"`
	Scope          *string `json:"scope,omitempty"`
}

func (s *SlackTokenPayloadOauth) GetChannelId() *string {
	if s == nil {
		return nil
	}
	return s.ChannelId
}

func (s *SlackTokenPayloadOauth) SetChannelId(channelId string) {
	s.ChannelId = &channelId
}

func (s *SlackTokenPayloadOauth) GetInstallationId() *string {
	if s == nil {
		return nil
	}
	return s.InstallationId
}

func (s *SlackTokenPayloadOauth) SetInstallationId(installationId string) {
	s.InstallationId = &installationId
}

func (s *SlackTokenPayloadOauth) GetScope() *string {
	if s == nil {
		return nil
	}
	return s.Scope
}

func (s *SlackTokenPayloadOauth) SetScope(scope string) {
	s.Scope = &scope
}

func (s SlackTokenPayloadOauth) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SlackTokenPayloadOauth to string"
	}
	return string(jsonData)
}

// Obtained directly from the incoming_webhook object in the installation response from the Slack API.
type SlackTokenPayloadWebhook struct {
	Url *string `json:"url,omitempty" required:"true" minLength:"1"`
}

func (s *SlackTokenPayloadWebhook) GetUrl() *string {
	if s == nil {
		return nil
	}
	return s.Url
}

func (s *SlackTokenPayloadWebhook) SetUrl(url string) {
	s.Url = &url
}

func (s SlackTokenPayloadWebhook) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SlackTokenPayloadWebhook to string"
	}
	return string(jsonData)
}
