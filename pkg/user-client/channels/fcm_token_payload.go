package channels

import "encoding/json"

type FcmTokenPayload struct {
	DeviceToken    *string                        `json:"device_token,omitempty" required:"true" minLength:"64"`
	InstallationId *FcmTokenPayloadInstallationId `json:"installation_id,omitempty"`
}

func (f *FcmTokenPayload) GetDeviceToken() *string {
	if f == nil {
		return nil
	}
	return f.DeviceToken
}

func (f *FcmTokenPayload) SetDeviceToken(deviceToken string) {
	f.DeviceToken = &deviceToken
}

func (f *FcmTokenPayload) GetInstallationId() *FcmTokenPayloadInstallationId {
	if f == nil {
		return nil
	}
	return f.InstallationId
}

func (f *FcmTokenPayload) SetInstallationId(installationId FcmTokenPayloadInstallationId) {
	f.InstallationId = &installationId
}

func (f FcmTokenPayload) String() string {
	jsonData, err := json.MarshalIndent(f, "", "  ")
	if err != nil {
		return "error converting struct: FcmTokenPayload to string"
	}
	return string(jsonData)
}

type FcmTokenPayloadInstallationId string

const (
	FCM_TOKEN_PAYLOAD_INSTALLATION_ID_DEVELOPMENT FcmTokenPayloadInstallationId = "development"
	FCM_TOKEN_PAYLOAD_INSTALLATION_ID_PRODUCTION  FcmTokenPayloadInstallationId = "production"
)
