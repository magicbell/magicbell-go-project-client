package integrations

import "encoding/json"

type WebpushConfigPayload struct {
	// VAPID private key - from the pair you generated.
	PrivateKey *string `json:"private_key,omitempty" required:"true" maxLength:"128" minLength:"8"`
	// VAPID public key - generate one at https://tools.reactpwa.com/vapid.
	PublicKey *string `json:"public_key,omitempty" required:"true" maxLength:"128" minLength:"8"`
}

func (w *WebpushConfigPayload) GetPrivateKey() *string {
	if w == nil {
		return nil
	}
	return w.PrivateKey
}

func (w *WebpushConfigPayload) SetPrivateKey(privateKey string) {
	w.PrivateKey = &privateKey
}

func (w *WebpushConfigPayload) GetPublicKey() *string {
	if w == nil {
		return nil
	}
	return w.PublicKey
}

func (w *WebpushConfigPayload) SetPublicKey(publicKey string) {
	w.PublicKey = &publicKey
}

func (w WebpushConfigPayload) String() string {
	jsonData, err := json.MarshalIndent(w, "", "  ")
	if err != nil {
		return "error converting struct: WebpushConfigPayload to string"
	}
	return string(jsonData)
}
