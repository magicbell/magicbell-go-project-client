package channels

import (
	"encoding/json"
)

type WebPushToken struct {
	Endpoint *string `json:"endpoint,omitempty" required:"true"`
	Keys     *Keys   `json:"keys,omitempty" required:"true"`
	touched  map[string]bool
}

func (w *WebPushToken) GetEndpoint() *string {
	if w == nil {
		return nil
	}
	return w.Endpoint
}

func (w *WebPushToken) SetEndpoint(endpoint string) {
	if w.touched == nil {
		w.touched = map[string]bool{}
	}
	w.touched["Endpoint"] = true
	w.Endpoint = &endpoint
}

func (w *WebPushToken) SetEndpointNil() {
	if w.touched == nil {
		w.touched = map[string]bool{}
	}
	w.touched["Endpoint"] = true
	w.Endpoint = nil
}

func (w *WebPushToken) GetKeys() *Keys {
	if w == nil {
		return nil
	}
	return w.Keys
}

func (w *WebPushToken) SetKeys(keys Keys) {
	if w.touched == nil {
		w.touched = map[string]bool{}
	}
	w.touched["Keys"] = true
	w.Keys = &keys
}

func (w *WebPushToken) SetKeysNil() {
	if w.touched == nil {
		w.touched = map[string]bool{}
	}
	w.touched["Keys"] = true
	w.Keys = nil
}
func (w WebPushToken) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if w.touched["Endpoint"] && w.Endpoint == nil {
		data["endpoint"] = nil
	} else if w.Endpoint != nil {
		data["endpoint"] = w.Endpoint
	}

	if w.touched["Keys"] && w.Keys == nil {
		data["keys"] = nil
	} else if w.Keys != nil {
		data["keys"] = w.Keys
	}

	return json.Marshal(data)
}

type Keys struct {
	Auth    *string `json:"auth,omitempty" required:"true"`
	P256dh  *string `json:"p256dh,omitempty" required:"true"`
	touched map[string]bool
}

func (k *Keys) GetAuth() *string {
	if k == nil {
		return nil
	}
	return k.Auth
}

func (k *Keys) SetAuth(auth string) {
	if k.touched == nil {
		k.touched = map[string]bool{}
	}
	k.touched["Auth"] = true
	k.Auth = &auth
}

func (k *Keys) SetAuthNil() {
	if k.touched == nil {
		k.touched = map[string]bool{}
	}
	k.touched["Auth"] = true
	k.Auth = nil
}

func (k *Keys) GetP256dh() *string {
	if k == nil {
		return nil
	}
	return k.P256dh
}

func (k *Keys) SetP256dh(p256dh string) {
	if k.touched == nil {
		k.touched = map[string]bool{}
	}
	k.touched["P256dh"] = true
	k.P256dh = &p256dh
}

func (k *Keys) SetP256dhNil() {
	if k.touched == nil {
		k.touched = map[string]bool{}
	}
	k.touched["P256dh"] = true
	k.P256dh = nil
}
func (k Keys) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if k.touched["Auth"] && k.Auth == nil {
		data["auth"] = nil
	} else if k.Auth != nil {
		data["auth"] = k.Auth
	}

	if k.touched["P256dh"] && k.P256dh == nil {
		data["p256dh"] = nil
	} else if k.P256dh != nil {
		data["p256dh"] = k.P256dh
	}

	return json.Marshal(data)
}
