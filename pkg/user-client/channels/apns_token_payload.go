package channels

import "encoding/json"

type ApnsTokenPayload struct {
	// (Optional) The bundle identifier of the application that is registering this token. Use this field to override the default identifier specified in the projects APNs integration.
	AppId       *string `json:"app_id,omitempty" pattern:"^[a-zA-Z0-9]+(.[a-zA-Z0-9]+)*$"`
	DeviceToken *string `json:"device_token,omitempty" required:"true" minLength:"64"`
	// (Optional) The APNs environment the token is registered for. If none is provided we assume the token is used in `production`.
	InstallationId *ApnsTokenPayloadInstallationId `json:"installation_id,omitempty"`
}

func (a *ApnsTokenPayload) GetAppId() *string {
	if a == nil {
		return nil
	}
	return a.AppId
}

func (a *ApnsTokenPayload) SetAppId(appId string) {
	a.AppId = &appId
}

func (a *ApnsTokenPayload) GetDeviceToken() *string {
	if a == nil {
		return nil
	}
	return a.DeviceToken
}

func (a *ApnsTokenPayload) SetDeviceToken(deviceToken string) {
	a.DeviceToken = &deviceToken
}

func (a *ApnsTokenPayload) GetInstallationId() *ApnsTokenPayloadInstallationId {
	if a == nil {
		return nil
	}
	return a.InstallationId
}

func (a *ApnsTokenPayload) SetInstallationId(installationId ApnsTokenPayloadInstallationId) {
	a.InstallationId = &installationId
}

func (a ApnsTokenPayload) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: ApnsTokenPayload to string"
	}
	return string(jsonData)
}

// (Optional) The APNs environment the token is registered for. If none is provided we assume the token is used in `production`.
type ApnsTokenPayloadInstallationId string

const (
	APNS_TOKEN_PAYLOAD_INSTALLATION_ID_DEVELOPMENT ApnsTokenPayloadInstallationId = "development"
	APNS_TOKEN_PAYLOAD_INSTALLATION_ID_PRODUCTION  ApnsTokenPayloadInstallationId = "production"
)
