package integrations

import (
	"encoding/json"
)

type ExpoConfig struct {
	AccessToken *string `json:"access_token,omitempty" required:"true" minLength:"1"`
	touched     map[string]bool
}

func (e *ExpoConfig) GetAccessToken() *string {
	if e == nil {
		return nil
	}
	return e.AccessToken
}

func (e *ExpoConfig) SetAccessToken(accessToken string) {
	if e.touched == nil {
		e.touched = map[string]bool{}
	}
	e.touched["AccessToken"] = true
	e.AccessToken = &accessToken
}

func (e *ExpoConfig) SetAccessTokenNil() {
	if e.touched == nil {
		e.touched = map[string]bool{}
	}
	e.touched["AccessToken"] = true
	e.AccessToken = nil
}
func (e ExpoConfig) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if e.touched["AccessToken"] && e.AccessToken == nil {
		data["access_token"] = nil
	} else if e.AccessToken != nil {
		data["access_token"] = e.AccessToken
	}

	return json.Marshal(data)
}
