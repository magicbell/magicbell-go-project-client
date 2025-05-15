package integrations

import "encoding/json"

type SlackFinishInstallResponse struct {
	// The app ID of the Slack app that was originally configured at the project-level.
	AppId *string `json:"app_id,omitempty" required:"true"`
	// The code that was returned from the OAuth flow, and found in the query string of the redirect URL.
	Code        *string `json:"code,omitempty" required:"true"`
	RedirectUrl *string `json:"redirect_url,omitempty"`
}

func (s *SlackFinishInstallResponse) GetAppId() *string {
	if s == nil {
		return nil
	}
	return s.AppId
}

func (s *SlackFinishInstallResponse) SetAppId(appId string) {
	s.AppId = &appId
}

func (s *SlackFinishInstallResponse) GetCode() *string {
	if s == nil {
		return nil
	}
	return s.Code
}

func (s *SlackFinishInstallResponse) SetCode(code string) {
	s.Code = &code
}

func (s *SlackFinishInstallResponse) GetRedirectUrl() *string {
	if s == nil {
		return nil
	}
	return s.RedirectUrl
}

func (s *SlackFinishInstallResponse) SetRedirectUrl(redirectUrl string) {
	s.RedirectUrl = &redirectUrl
}

func (s SlackFinishInstallResponse) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SlackFinishInstallResponse to string"
	}
	return string(jsonData)
}
