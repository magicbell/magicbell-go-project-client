package channels

import (
	"encoding/json"
)

type FcmToken struct {
	DeviceToken    *string                 `json:"device_token,omitempty" required:"true" minLength:"64"`
	InstallationId *FcmTokenInstallationId `json:"installation_id,omitempty"`
	touched        map[string]bool
}

func (f *FcmToken) GetDeviceToken() *string {
	if f == nil {
		return nil
	}
	return f.DeviceToken
}

func (f *FcmToken) SetDeviceToken(deviceToken string) {
	if f.touched == nil {
		f.touched = map[string]bool{}
	}
	f.touched["DeviceToken"] = true
	f.DeviceToken = &deviceToken
}

func (f *FcmToken) SetDeviceTokenNil() {
	if f.touched == nil {
		f.touched = map[string]bool{}
	}
	f.touched["DeviceToken"] = true
	f.DeviceToken = nil
}

func (f *FcmToken) GetInstallationId() *FcmTokenInstallationId {
	if f == nil {
		return nil
	}
	return f.InstallationId
}

func (f *FcmToken) SetInstallationId(installationId FcmTokenInstallationId) {
	if f.touched == nil {
		f.touched = map[string]bool{}
	}
	f.touched["InstallationId"] = true
	f.InstallationId = &installationId
}

func (f *FcmToken) SetInstallationIdNil() {
	if f.touched == nil {
		f.touched = map[string]bool{}
	}
	f.touched["InstallationId"] = true
	f.InstallationId = nil
}
func (f FcmToken) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if f.touched["DeviceToken"] && f.DeviceToken == nil {
		data["device_token"] = nil
	} else if f.DeviceToken != nil {
		data["device_token"] = f.DeviceToken
	}

	if f.touched["InstallationId"] && f.InstallationId == nil {
		data["installation_id"] = nil
	} else if f.InstallationId != nil {
		data["installation_id"] = f.InstallationId
	}

	return json.Marshal(data)
}

type FcmTokenInstallationId string

const (
	FCM_TOKEN_INSTALLATION_ID_DEVELOPMENT FcmTokenInstallationId = "development"
	FCM_TOKEN_INSTALLATION_ID_PRODUCTION  FcmTokenInstallationId = "production"
)
