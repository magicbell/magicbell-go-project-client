package integrations

import (
	"encoding/json"
)

type WebpushConfig struct {
	PrivateKey *string `json:"private_key,omitempty" required:"true" maxLength:"128" minLength:"8"`
	PublicKey  *string `json:"public_key,omitempty" required:"true" maxLength:"128" minLength:"8"`
	touched    map[string]bool
}

func (w *WebpushConfig) GetPrivateKey() *string {
	if w == nil {
		return nil
	}
	return w.PrivateKey
}

func (w *WebpushConfig) SetPrivateKey(privateKey string) {
	if w.touched == nil {
		w.touched = map[string]bool{}
	}
	w.touched["PrivateKey"] = true
	w.PrivateKey = &privateKey
}

func (w *WebpushConfig) SetPrivateKeyNil() {
	if w.touched == nil {
		w.touched = map[string]bool{}
	}
	w.touched["PrivateKey"] = true
	w.PrivateKey = nil
}

func (w *WebpushConfig) GetPublicKey() *string {
	if w == nil {
		return nil
	}
	return w.PublicKey
}

func (w *WebpushConfig) SetPublicKey(publicKey string) {
	if w.touched == nil {
		w.touched = map[string]bool{}
	}
	w.touched["PublicKey"] = true
	w.PublicKey = &publicKey
}

func (w *WebpushConfig) SetPublicKeyNil() {
	if w.touched == nil {
		w.touched = map[string]bool{}
	}
	w.touched["PublicKey"] = true
	w.PublicKey = nil
}
func (w WebpushConfig) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if w.touched["PrivateKey"] && w.PrivateKey == nil {
		data["private_key"] = nil
	} else if w.PrivateKey != nil {
		data["private_key"] = w.PrivateKey
	}

	if w.touched["PublicKey"] && w.PublicKey == nil {
		data["public_key"] = nil
	} else if w.PublicKey != nil {
		data["public_key"] = w.PublicKey
	}

	return json.Marshal(data)
}
