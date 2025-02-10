package channels

import (
	"encoding/json"
)

type FcmTokenResponse struct {
	CreatedAt      *string                         `json:"created_at,omitempty" required:"true"`
	DeviceToken    *string                         `json:"device_token,omitempty" required:"true" minLength:"64"`
	DiscardedAt    *string                         `json:"discarded_at,omitempty"`
	Id             *string                         `json:"id,omitempty" required:"true"`
	InstallationId *FcmTokenResponseInstallationId `json:"installation_id,omitempty"`
	UpdatedAt      *string                         `json:"updated_at,omitempty"`
	touched        map[string]bool
}

func (f *FcmTokenResponse) GetCreatedAt() *string {
	if f == nil {
		return nil
	}
	return f.CreatedAt
}

func (f *FcmTokenResponse) SetCreatedAt(createdAt string) {
	if f.touched == nil {
		f.touched = map[string]bool{}
	}
	f.touched["CreatedAt"] = true
	f.CreatedAt = &createdAt
}

func (f *FcmTokenResponse) SetCreatedAtNil() {
	if f.touched == nil {
		f.touched = map[string]bool{}
	}
	f.touched["CreatedAt"] = true
	f.CreatedAt = nil
}

func (f *FcmTokenResponse) GetDeviceToken() *string {
	if f == nil {
		return nil
	}
	return f.DeviceToken
}

func (f *FcmTokenResponse) SetDeviceToken(deviceToken string) {
	if f.touched == nil {
		f.touched = map[string]bool{}
	}
	f.touched["DeviceToken"] = true
	f.DeviceToken = &deviceToken
}

func (f *FcmTokenResponse) SetDeviceTokenNil() {
	if f.touched == nil {
		f.touched = map[string]bool{}
	}
	f.touched["DeviceToken"] = true
	f.DeviceToken = nil
}

func (f *FcmTokenResponse) GetDiscardedAt() *string {
	if f == nil {
		return nil
	}
	return f.DiscardedAt
}

func (f *FcmTokenResponse) SetDiscardedAt(discardedAt string) {
	if f.touched == nil {
		f.touched = map[string]bool{}
	}
	f.touched["DiscardedAt"] = true
	f.DiscardedAt = &discardedAt
}

func (f *FcmTokenResponse) SetDiscardedAtNil() {
	if f.touched == nil {
		f.touched = map[string]bool{}
	}
	f.touched["DiscardedAt"] = true
	f.DiscardedAt = nil
}

func (f *FcmTokenResponse) GetId() *string {
	if f == nil {
		return nil
	}
	return f.Id
}

func (f *FcmTokenResponse) SetId(id string) {
	if f.touched == nil {
		f.touched = map[string]bool{}
	}
	f.touched["Id"] = true
	f.Id = &id
}

func (f *FcmTokenResponse) SetIdNil() {
	if f.touched == nil {
		f.touched = map[string]bool{}
	}
	f.touched["Id"] = true
	f.Id = nil
}

func (f *FcmTokenResponse) GetInstallationId() *FcmTokenResponseInstallationId {
	if f == nil {
		return nil
	}
	return f.InstallationId
}

func (f *FcmTokenResponse) SetInstallationId(installationId FcmTokenResponseInstallationId) {
	if f.touched == nil {
		f.touched = map[string]bool{}
	}
	f.touched["InstallationId"] = true
	f.InstallationId = &installationId
}

func (f *FcmTokenResponse) SetInstallationIdNil() {
	if f.touched == nil {
		f.touched = map[string]bool{}
	}
	f.touched["InstallationId"] = true
	f.InstallationId = nil
}

func (f *FcmTokenResponse) GetUpdatedAt() *string {
	if f == nil {
		return nil
	}
	return f.UpdatedAt
}

func (f *FcmTokenResponse) SetUpdatedAt(updatedAt string) {
	if f.touched == nil {
		f.touched = map[string]bool{}
	}
	f.touched["UpdatedAt"] = true
	f.UpdatedAt = &updatedAt
}

func (f *FcmTokenResponse) SetUpdatedAtNil() {
	if f.touched == nil {
		f.touched = map[string]bool{}
	}
	f.touched["UpdatedAt"] = true
	f.UpdatedAt = nil
}

func (f FcmTokenResponse) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if f.touched["CreatedAt"] && f.CreatedAt == nil {
		data["created_at"] = nil
	} else if f.CreatedAt != nil {
		data["created_at"] = f.CreatedAt
	}

	if f.touched["DeviceToken"] && f.DeviceToken == nil {
		data["device_token"] = nil
	} else if f.DeviceToken != nil {
		data["device_token"] = f.DeviceToken
	}

	if f.touched["DiscardedAt"] && f.DiscardedAt == nil {
		data["discarded_at"] = nil
	} else if f.DiscardedAt != nil {
		data["discarded_at"] = f.DiscardedAt
	}

	if f.touched["Id"] && f.Id == nil {
		data["id"] = nil
	} else if f.Id != nil {
		data["id"] = f.Id
	}

	if f.touched["InstallationId"] && f.InstallationId == nil {
		data["installation_id"] = nil
	} else if f.InstallationId != nil {
		data["installation_id"] = f.InstallationId
	}

	if f.touched["UpdatedAt"] && f.UpdatedAt == nil {
		data["updated_at"] = nil
	} else if f.UpdatedAt != nil {
		data["updated_at"] = f.UpdatedAt
	}

	return json.Marshal(data)
}

func (f FcmTokenResponse) String() string {
	jsonData, err := json.MarshalIndent(f, "", "  ")
	if err != nil {
		return "error converting struct: FcmTokenResponse to string"
	}
	return string(jsonData)
}

type FcmTokenResponseInstallationId string

const (
	FCM_TOKEN_RESPONSE_INSTALLATION_ID_DEVELOPMENT FcmTokenResponseInstallationId = "development"
	FCM_TOKEN_RESPONSE_INSTALLATION_ID_PRODUCTION  FcmTokenResponseInstallationId = "production"
)
