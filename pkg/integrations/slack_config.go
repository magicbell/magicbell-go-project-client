package integrations

import (
	"encoding/json"
)

type SlackConfig struct {
	// The Slack app ID that can be found in the app's settings page of the Slack API dashboard.
	AppId *string `json:"app_id,omitempty" required:"true" pattern:"^[0-9A-Z]+$"`
	// The Slack client ID that can be found in the app's settings page of the Slack API dashboard.
	ClientId *string `json:"client_id,omitempty" required:"true" pattern:"^[0-9]+\.[0-9]+$"`
	// The Slack client secret that can be found in the app's settings page of the Slack API dashboard.
	ClientSecret *string `json:"client_secret,omitempty" required:"true" maxLength:"32" minLength:"32"`
	// The Slack signing secret that can be found in the app's settings page of the Slack API dashboard.
	SigningSecret *string `json:"signing_secret,omitempty" required:"true" maxLength:"32" minLength:"32"`
	touched       map[string]bool
}

func (s *SlackConfig) GetAppId() *string {
	if s == nil {
		return nil
	}
	return s.AppId
}

func (s *SlackConfig) SetAppId(appId string) {
	if s.touched == nil {
		s.touched = map[string]bool{}
	}
	s.touched["AppId"] = true
	s.AppId = &appId
}

func (s *SlackConfig) SetAppIdNil() {
	if s.touched == nil {
		s.touched = map[string]bool{}
	}
	s.touched["AppId"] = true
	s.AppId = nil
}

func (s *SlackConfig) GetClientId() *string {
	if s == nil {
		return nil
	}
	return s.ClientId
}

func (s *SlackConfig) SetClientId(clientId string) {
	if s.touched == nil {
		s.touched = map[string]bool{}
	}
	s.touched["ClientId"] = true
	s.ClientId = &clientId
}

func (s *SlackConfig) SetClientIdNil() {
	if s.touched == nil {
		s.touched = map[string]bool{}
	}
	s.touched["ClientId"] = true
	s.ClientId = nil
}

func (s *SlackConfig) GetClientSecret() *string {
	if s == nil {
		return nil
	}
	return s.ClientSecret
}

func (s *SlackConfig) SetClientSecret(clientSecret string) {
	if s.touched == nil {
		s.touched = map[string]bool{}
	}
	s.touched["ClientSecret"] = true
	s.ClientSecret = &clientSecret
}

func (s *SlackConfig) SetClientSecretNil() {
	if s.touched == nil {
		s.touched = map[string]bool{}
	}
	s.touched["ClientSecret"] = true
	s.ClientSecret = nil
}

func (s *SlackConfig) GetSigningSecret() *string {
	if s == nil {
		return nil
	}
	return s.SigningSecret
}

func (s *SlackConfig) SetSigningSecret(signingSecret string) {
	if s.touched == nil {
		s.touched = map[string]bool{}
	}
	s.touched["SigningSecret"] = true
	s.SigningSecret = &signingSecret
}

func (s *SlackConfig) SetSigningSecretNil() {
	if s.touched == nil {
		s.touched = map[string]bool{}
	}
	s.touched["SigningSecret"] = true
	s.SigningSecret = nil
}
func (s SlackConfig) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if s.touched["AppId"] && s.AppId == nil {
		data["app_id"] = nil
	} else if s.AppId != nil {
		data["app_id"] = s.AppId
	}

	if s.touched["ClientId"] && s.ClientId == nil {
		data["client_id"] = nil
	} else if s.ClientId != nil {
		data["client_id"] = s.ClientId
	}

	if s.touched["ClientSecret"] && s.ClientSecret == nil {
		data["client_secret"] = nil
	} else if s.ClientSecret != nil {
		data["client_secret"] = s.ClientSecret
	}

	if s.touched["SigningSecret"] && s.SigningSecret == nil {
		data["signing_secret"] = nil
	} else if s.SigningSecret != nil {
		data["signing_secret"] = s.SigningSecret
	}

	return json.Marshal(data)
}
