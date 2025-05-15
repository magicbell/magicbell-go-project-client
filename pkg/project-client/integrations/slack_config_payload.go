package integrations

import "encoding/json"

type SlackConfigPayload struct {
	// The Slack app ID that can be found in the app's settings page of the Slack API dashboard.
	AppId *string `json:"app_id,omitempty" required:"true" pattern:"^[0-9A-Z]+$"`
	// The Slack client ID that can be found in the app's settings page of the Slack API dashboard.
	ClientId *string `json:"client_id,omitempty" required:"true" pattern:"^[0-9]+\.[0-9]+$"`
	// The Slack client secret that can be found in the app's settings page of the Slack API dashboard.
	ClientSecret *string `json:"client_secret,omitempty" required:"true" maxLength:"32" minLength:"32"`
	// The Slack signing secret that can be found in the app's settings page of the Slack API dashboard.
	SigningSecret *string `json:"signing_secret,omitempty" required:"true" maxLength:"32" minLength:"32"`
}

func (s *SlackConfigPayload) GetAppId() *string {
	if s == nil {
		return nil
	}
	return s.AppId
}

func (s *SlackConfigPayload) SetAppId(appId string) {
	s.AppId = &appId
}

func (s *SlackConfigPayload) GetClientId() *string {
	if s == nil {
		return nil
	}
	return s.ClientId
}

func (s *SlackConfigPayload) SetClientId(clientId string) {
	s.ClientId = &clientId
}

func (s *SlackConfigPayload) GetClientSecret() *string {
	if s == nil {
		return nil
	}
	return s.ClientSecret
}

func (s *SlackConfigPayload) SetClientSecret(clientSecret string) {
	s.ClientSecret = &clientSecret
}

func (s *SlackConfigPayload) GetSigningSecret() *string {
	if s == nil {
		return nil
	}
	return s.SigningSecret
}

func (s *SlackConfigPayload) SetSigningSecret(signingSecret string) {
	s.SigningSecret = &signingSecret
}

func (s SlackConfigPayload) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SlackConfigPayload to string"
	}
	return string(jsonData)
}
