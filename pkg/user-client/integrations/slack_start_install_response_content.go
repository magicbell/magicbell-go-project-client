package integrations

import "encoding/json"

type SlackStartInstallResponseContent struct {
	AppId   *string  `json:"app_id,omitempty"`
	AuthUrl *string  `json:"auth_url,omitempty"`
	Scopes  []string `json:"scopes,omitempty"`
}

func (s *SlackStartInstallResponseContent) GetAppId() *string {
	if s == nil {
		return nil
	}
	return s.AppId
}

func (s *SlackStartInstallResponseContent) SetAppId(appId string) {
	s.AppId = &appId
}

func (s *SlackStartInstallResponseContent) GetAuthUrl() *string {
	if s == nil {
		return nil
	}
	return s.AuthUrl
}

func (s *SlackStartInstallResponseContent) SetAuthUrl(authUrl string) {
	s.AuthUrl = &authUrl
}

func (s *SlackStartInstallResponseContent) GetScopes() []string {
	if s == nil {
		return nil
	}
	return s.Scopes
}

func (s *SlackStartInstallResponseContent) SetScopes(scopes []string) {
	s.Scopes = scopes
}

func (s SlackStartInstallResponseContent) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SlackStartInstallResponseContent to string"
	}
	return string(jsonData)
}
