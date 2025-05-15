package integrations

import "encoding/json"

type SlackStartInstall struct {
	AppId       *string  `json:"app_id,omitempty" required:"true"`
	AuthUrl     *string  `json:"auth_url,omitempty"`
	ExtraScopes []string `json:"extra_scopes,omitempty"`
	RedirectUrl *string  `json:"redirect_url,omitempty"`
}

func (s *SlackStartInstall) GetAppId() *string {
	if s == nil {
		return nil
	}
	return s.AppId
}

func (s *SlackStartInstall) SetAppId(appId string) {
	s.AppId = &appId
}

func (s *SlackStartInstall) GetAuthUrl() *string {
	if s == nil {
		return nil
	}
	return s.AuthUrl
}

func (s *SlackStartInstall) SetAuthUrl(authUrl string) {
	s.AuthUrl = &authUrl
}

func (s *SlackStartInstall) GetExtraScopes() []string {
	if s == nil {
		return nil
	}
	return s.ExtraScopes
}

func (s *SlackStartInstall) SetExtraScopes(extraScopes []string) {
	s.ExtraScopes = extraScopes
}

func (s *SlackStartInstall) GetRedirectUrl() *string {
	if s == nil {
		return nil
	}
	return s.RedirectUrl
}

func (s *SlackStartInstall) SetRedirectUrl(redirectUrl string) {
	s.RedirectUrl = &redirectUrl
}

func (s SlackStartInstall) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SlackStartInstall to string"
	}
	return string(jsonData)
}
