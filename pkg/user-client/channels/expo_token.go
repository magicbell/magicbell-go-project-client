package channels

import (
	"encoding/json"
	"github.com/magicbell/magicbell-go/pkg/user-client/util"
)

type ExpoToken struct {
	CreatedAt   *string                `json:"created_at,omitempty" required:"true"`
	DeviceToken *string                `json:"device_token,omitempty" required:"true" minLength:"1"`
	DiscardedAt *util.Nullable[string] `json:"discarded_at,omitempty"`
	Id          *string                `json:"id,omitempty" required:"true"`
	UpdatedAt   *util.Nullable[string] `json:"updated_at,omitempty"`
}

func (e *ExpoToken) GetCreatedAt() *string {
	if e == nil {
		return nil
	}
	return e.CreatedAt
}

func (e *ExpoToken) SetCreatedAt(createdAt string) {
	e.CreatedAt = &createdAt
}

func (e *ExpoToken) GetDeviceToken() *string {
	if e == nil {
		return nil
	}
	return e.DeviceToken
}

func (e *ExpoToken) SetDeviceToken(deviceToken string) {
	e.DeviceToken = &deviceToken
}

func (e *ExpoToken) GetDiscardedAt() *util.Nullable[string] {
	if e == nil {
		return nil
	}
	return e.DiscardedAt
}

func (e *ExpoToken) SetDiscardedAt(discardedAt util.Nullable[string]) {
	e.DiscardedAt = &discardedAt
}

func (e *ExpoToken) SetDiscardedAtNull() {
	e.DiscardedAt = &util.Nullable[string]{IsNull: true}
}

func (e *ExpoToken) GetId() *string {
	if e == nil {
		return nil
	}
	return e.Id
}

func (e *ExpoToken) SetId(id string) {
	e.Id = &id
}

func (e *ExpoToken) GetUpdatedAt() *util.Nullable[string] {
	if e == nil {
		return nil
	}
	return e.UpdatedAt
}

func (e *ExpoToken) SetUpdatedAt(updatedAt util.Nullable[string]) {
	e.UpdatedAt = &updatedAt
}

func (e *ExpoToken) SetUpdatedAtNull() {
	e.UpdatedAt = &util.Nullable[string]{IsNull: true}
}

func (e ExpoToken) String() string {
	jsonData, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		return "error converting struct: ExpoToken to string"
	}
	return string(jsonData)
}
