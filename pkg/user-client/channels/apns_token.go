package channels

import (
	"encoding/json"
	"github.com/magicbell/magicbell-go/pkg/user-client/util"
)

type ApnsToken struct {
	// (Optional) The bundle identifier of the application that is registering this token. Use this field to override the default identifier specified in the projects APNs integration.
	AppId       *string                `json:"app_id,omitempty" pattern:"^[a-zA-Z0-9]+(.[a-zA-Z0-9]+)*$"`
	CreatedAt   *string                `json:"created_at,omitempty" required:"true"`
	DeviceToken *string                `json:"device_token,omitempty" required:"true" minLength:"64"`
	DiscardedAt *util.Nullable[string] `json:"discarded_at,omitempty"`
	Id          *string                `json:"id,omitempty" required:"true"`
	// (Optional) The APNs environment the token is registered for. If none is provided we assume the token is used in `production`.
	InstallationId *ApnsTokenInstallationId `json:"installation_id,omitempty"`
	UpdatedAt      *util.Nullable[string]   `json:"updated_at,omitempty"`
}

func (a *ApnsToken) GetAppId() *string {
	if a == nil {
		return nil
	}
	return a.AppId
}

func (a *ApnsToken) SetAppId(appId string) {
	a.AppId = &appId
}

func (a *ApnsToken) GetCreatedAt() *string {
	if a == nil {
		return nil
	}
	return a.CreatedAt
}

func (a *ApnsToken) SetCreatedAt(createdAt string) {
	a.CreatedAt = &createdAt
}

func (a *ApnsToken) GetDeviceToken() *string {
	if a == nil {
		return nil
	}
	return a.DeviceToken
}

func (a *ApnsToken) SetDeviceToken(deviceToken string) {
	a.DeviceToken = &deviceToken
}

func (a *ApnsToken) GetDiscardedAt() *util.Nullable[string] {
	if a == nil {
		return nil
	}
	return a.DiscardedAt
}

func (a *ApnsToken) SetDiscardedAt(discardedAt util.Nullable[string]) {
	a.DiscardedAt = &discardedAt
}

func (a *ApnsToken) SetDiscardedAtNull() {
	a.DiscardedAt = &util.Nullable[string]{IsNull: true}
}

func (a *ApnsToken) GetId() *string {
	if a == nil {
		return nil
	}
	return a.Id
}

func (a *ApnsToken) SetId(id string) {
	a.Id = &id
}

func (a *ApnsToken) GetInstallationId() *ApnsTokenInstallationId {
	if a == nil {
		return nil
	}
	return a.InstallationId
}

func (a *ApnsToken) SetInstallationId(installationId ApnsTokenInstallationId) {
	a.InstallationId = &installationId
}

func (a *ApnsToken) GetUpdatedAt() *util.Nullable[string] {
	if a == nil {
		return nil
	}
	return a.UpdatedAt
}

func (a *ApnsToken) SetUpdatedAt(updatedAt util.Nullable[string]) {
	a.UpdatedAt = &updatedAt
}

func (a *ApnsToken) SetUpdatedAtNull() {
	a.UpdatedAt = &util.Nullable[string]{IsNull: true}
}

func (a ApnsToken) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: ApnsToken to string"
	}
	return string(jsonData)
}

// (Optional) The APNs environment the token is registered for. If none is provided we assume the token is used in `production`.
type ApnsTokenInstallationId string

const (
	APNS_TOKEN_INSTALLATION_ID_DEVELOPMENT ApnsTokenInstallationId = "development"
	APNS_TOKEN_INSTALLATION_ID_PRODUCTION  ApnsTokenInstallationId = "production"
)
