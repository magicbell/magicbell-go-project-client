package shared

import "encoding/json"

type WebPushTokenPayload struct {
	// The push subscription URL obtained from PushSubscription.endpoint after calling registration.pushManager.subscribe(). This is the unique URL for this device that push messages will be sent to.
	Endpoint *string `json:"endpoint,omitempty" required:"true"`
	// The encryption keys from the PushSubscription.getKey() method, needed to encrypt push messages for this subscription.
	Keys *WebPushTokenPayloadKeys `json:"keys,omitempty" required:"true"`
}

func (w *WebPushTokenPayload) GetEndpoint() *string {
	if w == nil {
		return nil
	}
	return w.Endpoint
}

func (w *WebPushTokenPayload) SetEndpoint(endpoint string) {
	w.Endpoint = &endpoint
}

func (w *WebPushTokenPayload) GetKeys() *WebPushTokenPayloadKeys {
	if w == nil {
		return nil
	}
	return w.Keys
}

func (w *WebPushTokenPayload) SetKeys(keys WebPushTokenPayloadKeys) {
	w.Keys = &keys
}

func (w WebPushTokenPayload) String() string {
	jsonData, err := json.MarshalIndent(w, "", "  ")
	if err != nil {
		return "error converting struct: WebPushTokenPayload to string"
	}
	return string(jsonData)
}

// The encryption keys from the PushSubscription.getKey() method, needed to encrypt push messages for this subscription.
type WebPushTokenPayloadKeys struct {
	// The authentication secret obtained from PushSubscription.getKey('auth'). Used to encrypt push messages for this subscription.
	Auth *string `json:"auth,omitempty" required:"true"`
	// The P-256 ECDH public key obtained from PushSubscription.getKey('p256dh'). Used to encrypt push messages for this subscription.
	P256dh *string `json:"p256dh,omitempty" required:"true"`
}

func (w *WebPushTokenPayloadKeys) GetAuth() *string {
	if w == nil {
		return nil
	}
	return w.Auth
}

func (w *WebPushTokenPayloadKeys) SetAuth(auth string) {
	w.Auth = &auth
}

func (w *WebPushTokenPayloadKeys) GetP256dh() *string {
	if w == nil {
		return nil
	}
	return w.P256dh
}

func (w *WebPushTokenPayloadKeys) SetP256dh(p256dh string) {
	w.P256dh = &p256dh
}

func (w WebPushTokenPayloadKeys) String() string {
	jsonData, err := json.MarshalIndent(w, "", "  ")
	if err != nil {
		return "error converting struct: WebPushTokenPayloadKeys to string"
	}
	return string(jsonData)
}
