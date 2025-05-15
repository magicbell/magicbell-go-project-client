package channels

import (
	"encoding/json"
	"github.com/magicbell/magicbell-go/pkg/project-client/util"
)

type WebPushToken struct {
	CreatedAt   *string                `json:"created_at,omitempty" required:"true"`
	DiscardedAt *util.Nullable[string] `json:"discarded_at,omitempty"`
	// The push subscription URL obtained from PushSubscription.endpoint after calling registration.pushManager.subscribe(). This is the unique URL for this device that push messages will be sent to.
	Endpoint *string `json:"endpoint,omitempty" required:"true"`
	Id       *string `json:"id,omitempty" required:"true"`
	// The encryption keys from the PushSubscription.getKey() method, needed to encrypt push messages for this subscription.
	Keys      *Keys                  `json:"keys,omitempty" required:"true"`
	UpdatedAt *util.Nullable[string] `json:"updated_at,omitempty"`
}

func (w *WebPushToken) GetCreatedAt() *string {
	if w == nil {
		return nil
	}
	return w.CreatedAt
}

func (w *WebPushToken) SetCreatedAt(createdAt string) {
	w.CreatedAt = &createdAt
}

func (w *WebPushToken) GetDiscardedAt() *util.Nullable[string] {
	if w == nil {
		return nil
	}
	return w.DiscardedAt
}

func (w *WebPushToken) SetDiscardedAt(discardedAt util.Nullable[string]) {
	w.DiscardedAt = &discardedAt
}

func (w *WebPushToken) SetDiscardedAtNull() {
	w.DiscardedAt = &util.Nullable[string]{IsNull: true}
}

func (w *WebPushToken) GetEndpoint() *string {
	if w == nil {
		return nil
	}
	return w.Endpoint
}

func (w *WebPushToken) SetEndpoint(endpoint string) {
	w.Endpoint = &endpoint
}

func (w *WebPushToken) GetId() *string {
	if w == nil {
		return nil
	}
	return w.Id
}

func (w *WebPushToken) SetId(id string) {
	w.Id = &id
}

func (w *WebPushToken) GetKeys() *Keys {
	if w == nil {
		return nil
	}
	return w.Keys
}

func (w *WebPushToken) SetKeys(keys Keys) {
	w.Keys = &keys
}

func (w *WebPushToken) GetUpdatedAt() *util.Nullable[string] {
	if w == nil {
		return nil
	}
	return w.UpdatedAt
}

func (w *WebPushToken) SetUpdatedAt(updatedAt util.Nullable[string]) {
	w.UpdatedAt = &updatedAt
}

func (w *WebPushToken) SetUpdatedAtNull() {
	w.UpdatedAt = &util.Nullable[string]{IsNull: true}
}

func (w WebPushToken) String() string {
	jsonData, err := json.MarshalIndent(w, "", "  ")
	if err != nil {
		return "error converting struct: WebPushToken to string"
	}
	return string(jsonData)
}

// The encryption keys from the PushSubscription.getKey() method, needed to encrypt push messages for this subscription.
type Keys struct {
	// The authentication secret obtained from PushSubscription.getKey('auth'). Used to encrypt push messages for this subscription.
	Auth *string `json:"auth,omitempty" required:"true"`
	// The P-256 ECDH public key obtained from PushSubscription.getKey('p256dh'). Used to encrypt push messages for this subscription.
	P256dh *string `json:"p256dh,omitempty" required:"true"`
}

func (k *Keys) GetAuth() *string {
	if k == nil {
		return nil
	}
	return k.Auth
}

func (k *Keys) SetAuth(auth string) {
	k.Auth = &auth
}

func (k *Keys) GetP256dh() *string {
	if k == nil {
		return nil
	}
	return k.P256dh
}

func (k *Keys) SetP256dh(p256dh string) {
	k.P256dh = &p256dh
}

func (k Keys) String() string {
	jsonData, err := json.MarshalIndent(k, "", "  ")
	if err != nil {
		return "error converting struct: Keys to string"
	}
	return string(jsonData)
}
