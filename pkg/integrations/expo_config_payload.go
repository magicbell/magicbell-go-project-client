package integrations

import "encoding/json"

type ExpoConfigPayload struct {
	AccessToken *string `json:"access_token,omitempty" required:"true" minLength:"1"`
}

func (e *ExpoConfigPayload) GetAccessToken() *string {
	if e == nil {
		return nil
	}
	return e.AccessToken
}

func (e *ExpoConfigPayload) SetAccessToken(accessToken string) {
	e.AccessToken = &accessToken
}

func (e ExpoConfigPayload) String() string {
	jsonData, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		return "error converting struct: ExpoConfigPayload to string"
	}
	return string(jsonData)
}
