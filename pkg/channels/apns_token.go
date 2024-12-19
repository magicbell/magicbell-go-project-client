package channels

import (
	"encoding/json"
)

type ApnsToken struct {
	// (Optional) The bundle identifier of the application that is registering this token. Use this field to override the default identifier specified in the projects APNs integration.
	AppId       *string `json:"app_id,omitempty" pattern:"^[a-zA-Z0-9]+(.[a-zA-Z0-9]+)*$"`
	DeviceToken *string `json:"device_token,omitempty" required:"true" minLength:"64"`
	// (Optional) The APNs environment the token is registered for. If none is provided we assume the token is used in `production`.
	InstallationId *ApnsTokenInstallationId `json:"installation_id,omitempty"`
	touched        map[string]bool
}

func (a *ApnsToken) GetAppId() *string {
	if a == nil {
		return nil
	}
	return a.AppId
}

func (a *ApnsToken) SetAppId(appId string) {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["AppId"] = true
	a.AppId = &appId
}

func (a *ApnsToken) SetAppIdNil() {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["AppId"] = true
	a.AppId = nil
}

func (a *ApnsToken) GetDeviceToken() *string {
	if a == nil {
		return nil
	}
	return a.DeviceToken
}

func (a *ApnsToken) SetDeviceToken(deviceToken string) {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["DeviceToken"] = true
	a.DeviceToken = &deviceToken
}

func (a *ApnsToken) SetDeviceTokenNil() {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["DeviceToken"] = true
	a.DeviceToken = nil
}

func (a *ApnsToken) GetInstallationId() *ApnsTokenInstallationId {
	if a == nil {
		return nil
	}
	return a.InstallationId
}

func (a *ApnsToken) SetInstallationId(installationId ApnsTokenInstallationId) {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["InstallationId"] = true
	a.InstallationId = &installationId
}

func (a *ApnsToken) SetInstallationIdNil() {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["InstallationId"] = true
	a.InstallationId = nil
}
func (a ApnsToken) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if a.touched["AppId"] && a.AppId == nil {
		data["app_id"] = nil
	} else if a.AppId != nil {
		data["app_id"] = a.AppId
	}

	if a.touched["DeviceToken"] && a.DeviceToken == nil {
		data["device_token"] = nil
	} else if a.DeviceToken != nil {
		data["device_token"] = a.DeviceToken
	}

	if a.touched["InstallationId"] && a.InstallationId == nil {
		data["installation_id"] = nil
	} else if a.InstallationId != nil {
		data["installation_id"] = a.InstallationId
	}

	return json.Marshal(data)
}

// (Optional) The APNs environment the token is registered for. If none is provided we assume the token is used in `production`.
type ApnsTokenInstallationId string

const (
	APNS_TOKEN_INSTALLATION_ID_DEVELOPMENT ApnsTokenInstallationId = "development"
	APNS_TOKEN_INSTALLATION_ID_PRODUCTION  ApnsTokenInstallationId = "production"
)
