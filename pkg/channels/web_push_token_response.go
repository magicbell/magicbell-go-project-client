package channels

import (
	"encoding/json"
)

type WebPushTokenResponse struct {
	CreatedAt   *string `json:"created_at,omitempty" required:"true"`
	DiscardedAt *string `json:"discarded_at,omitempty"`
	Endpoint    *string `json:"endpoint,omitempty" required:"true"`
	Id          *string `json:"id,omitempty" required:"true"`
	Keys        *Keys   `json:"keys,omitempty" required:"true"`
	UpdatedAt   *string `json:"updated_at,omitempty"`
	touched     map[string]bool
}

func (w *WebPushTokenResponse) GetCreatedAt() *string {
	if w == nil {
		return nil
	}
	return w.CreatedAt
}

func (w *WebPushTokenResponse) SetCreatedAt(createdAt string) {
	if w.touched == nil {
		w.touched = map[string]bool{}
	}
	w.touched["CreatedAt"] = true
	w.CreatedAt = &createdAt
}

func (w *WebPushTokenResponse) SetCreatedAtNil() {
	if w.touched == nil {
		w.touched = map[string]bool{}
	}
	w.touched["CreatedAt"] = true
	w.CreatedAt = nil
}

func (w *WebPushTokenResponse) GetDiscardedAt() *string {
	if w == nil {
		return nil
	}
	return w.DiscardedAt
}

func (w *WebPushTokenResponse) SetDiscardedAt(discardedAt string) {
	if w.touched == nil {
		w.touched = map[string]bool{}
	}
	w.touched["DiscardedAt"] = true
	w.DiscardedAt = &discardedAt
}

func (w *WebPushTokenResponse) SetDiscardedAtNil() {
	if w.touched == nil {
		w.touched = map[string]bool{}
	}
	w.touched["DiscardedAt"] = true
	w.DiscardedAt = nil
}

func (w *WebPushTokenResponse) GetEndpoint() *string {
	if w == nil {
		return nil
	}
	return w.Endpoint
}

func (w *WebPushTokenResponse) SetEndpoint(endpoint string) {
	if w.touched == nil {
		w.touched = map[string]bool{}
	}
	w.touched["Endpoint"] = true
	w.Endpoint = &endpoint
}

func (w *WebPushTokenResponse) SetEndpointNil() {
	if w.touched == nil {
		w.touched = map[string]bool{}
	}
	w.touched["Endpoint"] = true
	w.Endpoint = nil
}

func (w *WebPushTokenResponse) GetId() *string {
	if w == nil {
		return nil
	}
	return w.Id
}

func (w *WebPushTokenResponse) SetId(id string) {
	if w.touched == nil {
		w.touched = map[string]bool{}
	}
	w.touched["Id"] = true
	w.Id = &id
}

func (w *WebPushTokenResponse) SetIdNil() {
	if w.touched == nil {
		w.touched = map[string]bool{}
	}
	w.touched["Id"] = true
	w.Id = nil
}

func (w *WebPushTokenResponse) GetKeys() *Keys {
	if w == nil {
		return nil
	}
	return w.Keys
}

func (w *WebPushTokenResponse) SetKeys(keys Keys) {
	if w.touched == nil {
		w.touched = map[string]bool{}
	}
	w.touched["Keys"] = true
	w.Keys = &keys
}

func (w *WebPushTokenResponse) SetKeysNil() {
	if w.touched == nil {
		w.touched = map[string]bool{}
	}
	w.touched["Keys"] = true
	w.Keys = nil
}

func (w *WebPushTokenResponse) GetUpdatedAt() *string {
	if w == nil {
		return nil
	}
	return w.UpdatedAt
}

func (w *WebPushTokenResponse) SetUpdatedAt(updatedAt string) {
	if w.touched == nil {
		w.touched = map[string]bool{}
	}
	w.touched["UpdatedAt"] = true
	w.UpdatedAt = &updatedAt
}

func (w *WebPushTokenResponse) SetUpdatedAtNil() {
	if w.touched == nil {
		w.touched = map[string]bool{}
	}
	w.touched["UpdatedAt"] = true
	w.UpdatedAt = nil
}

func (w WebPushTokenResponse) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if w.touched["CreatedAt"] && w.CreatedAt == nil {
		data["created_at"] = nil
	} else if w.CreatedAt != nil {
		data["created_at"] = w.CreatedAt
	}

	if w.touched["DiscardedAt"] && w.DiscardedAt == nil {
		data["discarded_at"] = nil
	} else if w.DiscardedAt != nil {
		data["discarded_at"] = w.DiscardedAt
	}

	if w.touched["Endpoint"] && w.Endpoint == nil {
		data["endpoint"] = nil
	} else if w.Endpoint != nil {
		data["endpoint"] = w.Endpoint
	}

	if w.touched["Id"] && w.Id == nil {
		data["id"] = nil
	} else if w.Id != nil {
		data["id"] = w.Id
	}

	if w.touched["Keys"] && w.Keys == nil {
		data["keys"] = nil
	} else if w.Keys != nil {
		data["keys"] = w.Keys
	}

	if w.touched["UpdatedAt"] && w.UpdatedAt == nil {
		data["updated_at"] = nil
	} else if w.UpdatedAt != nil {
		data["updated_at"] = w.UpdatedAt
	}

	return json.Marshal(data)
}

func (w WebPushTokenResponse) String() string {
	jsonData, err := json.MarshalIndent(w, "", "  ")
	if err != nil {
		return "error converting struct: WebPushTokenResponse to string"
	}
	return string(jsonData)
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

func (k Keys) String() string {
	jsonData, err := json.MarshalIndent(k, "", "  ")
	if err != nil {
		return "error converting struct: Keys to string"
	}
	return string(jsonData)
}
