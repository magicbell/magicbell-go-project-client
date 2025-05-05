package users

import (
	"encoding/json"
	"github.com/magicbell/magicbell-go-project-client/pkg/util"
)

type User struct {
	CreatedAt        *util.Nullable[string] `json:"created_at,omitempty"`
	CustomAttributes *util.Nullable[any]    `json:"custom_attributes,omitempty"`
	Email            *util.Nullable[string] `json:"email,omitempty"`
	ExternalId       *util.Nullable[string] `json:"external_id,omitempty"`
	FirstName        *util.Nullable[string] `json:"first_name,omitempty"`
	Id               *string                `json:"id,omitempty"`
	LastName         *util.Nullable[string] `json:"last_name,omitempty"`
	LastNotifiedAt   *util.Nullable[string] `json:"last_notified_at,omitempty"`
	LastSeenAt       *util.Nullable[string] `json:"last_seen_at,omitempty"`
	UpdatedAt        *util.Nullable[string] `json:"updated_at,omitempty"`
}

func (u *User) GetCreatedAt() *util.Nullable[string] {
	if u == nil {
		return nil
	}
	return u.CreatedAt
}

func (u *User) SetCreatedAt(createdAt util.Nullable[string]) {
	u.CreatedAt = &createdAt
}

func (u *User) SetCreatedAtNull() {
	u.CreatedAt = &util.Nullable[string]{IsNull: true}
}

func (u *User) GetCustomAttributes() *util.Nullable[any] {
	if u == nil {
		return nil
	}
	return u.CustomAttributes
}

func (u *User) SetCustomAttributes(customAttributes util.Nullable[any]) {
	u.CustomAttributes = &customAttributes
}

func (u *User) SetCustomAttributesNull() {
	u.CustomAttributes = &util.Nullable[any]{IsNull: true}
}

func (u *User) GetEmail() *util.Nullable[string] {
	if u == nil {
		return nil
	}
	return u.Email
}

func (u *User) SetEmail(email util.Nullable[string]) {
	u.Email = &email
}

func (u *User) SetEmailNull() {
	u.Email = &util.Nullable[string]{IsNull: true}
}

func (u *User) GetExternalId() *util.Nullable[string] {
	if u == nil {
		return nil
	}
	return u.ExternalId
}

func (u *User) SetExternalId(externalId util.Nullable[string]) {
	u.ExternalId = &externalId
}

func (u *User) SetExternalIdNull() {
	u.ExternalId = &util.Nullable[string]{IsNull: true}
}

func (u *User) GetFirstName() *util.Nullable[string] {
	if u == nil {
		return nil
	}
	return u.FirstName
}

func (u *User) SetFirstName(firstName util.Nullable[string]) {
	u.FirstName = &firstName
}

func (u *User) SetFirstNameNull() {
	u.FirstName = &util.Nullable[string]{IsNull: true}
}

func (u *User) GetId() *string {
	if u == nil {
		return nil
	}
	return u.Id
}

func (u *User) SetId(id string) {
	u.Id = &id
}

func (u *User) GetLastName() *util.Nullable[string] {
	if u == nil {
		return nil
	}
	return u.LastName
}

func (u *User) SetLastName(lastName util.Nullable[string]) {
	u.LastName = &lastName
}

func (u *User) SetLastNameNull() {
	u.LastName = &util.Nullable[string]{IsNull: true}
}

func (u *User) GetLastNotifiedAt() *util.Nullable[string] {
	if u == nil {
		return nil
	}
	return u.LastNotifiedAt
}

func (u *User) SetLastNotifiedAt(lastNotifiedAt util.Nullable[string]) {
	u.LastNotifiedAt = &lastNotifiedAt
}

func (u *User) SetLastNotifiedAtNull() {
	u.LastNotifiedAt = &util.Nullable[string]{IsNull: true}
}

func (u *User) GetLastSeenAt() *util.Nullable[string] {
	if u == nil {
		return nil
	}
	return u.LastSeenAt
}

func (u *User) SetLastSeenAt(lastSeenAt util.Nullable[string]) {
	u.LastSeenAt = &lastSeenAt
}

func (u *User) SetLastSeenAtNull() {
	u.LastSeenAt = &util.Nullable[string]{IsNull: true}
}

func (u *User) GetUpdatedAt() *util.Nullable[string] {
	if u == nil {
		return nil
	}
	return u.UpdatedAt
}

func (u *User) SetUpdatedAt(updatedAt util.Nullable[string]) {
	u.UpdatedAt = &updatedAt
}

func (u *User) SetUpdatedAtNull() {
	u.UpdatedAt = &util.Nullable[string]{IsNull: true}
}

func (u User) String() string {
	jsonData, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		return "error converting struct: User to string"
	}
	return string(jsonData)
}
