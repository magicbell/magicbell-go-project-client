package channels

import (
	"encoding/json"
)

type ExpoToken struct {
	DeviceToken *string `json:"device_token,omitempty" required:"true" minLength:"1"`
	touched     map[string]bool
}

func (e *ExpoToken) GetDeviceToken() *string {
	if e == nil {
		return nil
	}
	return e.DeviceToken
}

func (e *ExpoToken) SetDeviceToken(deviceToken string) {
	if e.touched == nil {
		e.touched = map[string]bool{}
	}
	e.touched["DeviceToken"] = true
	e.DeviceToken = &deviceToken
}

func (e *ExpoToken) SetDeviceTokenNil() {
	if e.touched == nil {
		e.touched = map[string]bool{}
	}
	e.touched["DeviceToken"] = true
	e.DeviceToken = nil
}
func (e ExpoToken) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if e.touched["DeviceToken"] && e.DeviceToken == nil {
		data["device_token"] = nil
	} else if e.DeviceToken != nil {
		data["device_token"] = e.DeviceToken
	}

	return json.Marshal(data)
}
