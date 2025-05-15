package integrations

import "encoding/json"

type WebPushStartInstallationResponse struct {
	AuthToken *string `json:"auth_token,omitempty" required:"true" minLength:"8"`
	PublicKey *string `json:"public_key,omitempty" required:"true" maxLength:"128" minLength:"8"`
}

func (w *WebPushStartInstallationResponse) GetAuthToken() *string {
	if w == nil {
		return nil
	}
	return w.AuthToken
}

func (w *WebPushStartInstallationResponse) SetAuthToken(authToken string) {
	w.AuthToken = &authToken
}

func (w *WebPushStartInstallationResponse) GetPublicKey() *string {
	if w == nil {
		return nil
	}
	return w.PublicKey
}

func (w *WebPushStartInstallationResponse) SetPublicKey(publicKey string) {
	w.PublicKey = &publicKey
}

func (w WebPushStartInstallationResponse) String() string {
	jsonData, err := json.MarshalIndent(w, "", "  ")
	if err != nil {
		return "error converting struct: WebPushStartInstallationResponse to string"
	}
	return string(jsonData)
}
