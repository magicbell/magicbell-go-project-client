package channels

import "encoding/json"

type ExpoTokenPayload struct {
	DeviceToken *string `json:"device_token,omitempty" required:"true" minLength:"1"`
}

func (e *ExpoTokenPayload) GetDeviceToken() *string {
	if e == nil {
		return nil
	}
	return e.DeviceToken
}

func (e *ExpoTokenPayload) SetDeviceToken(deviceToken string) {
	e.DeviceToken = &deviceToken
}

func (e ExpoTokenPayload) String() string {
	jsonData, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		return "error converting struct: ExpoTokenPayload to string"
	}
	return string(jsonData)
}
