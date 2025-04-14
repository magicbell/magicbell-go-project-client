package users

import "encoding/json"

type User struct {
	CreatedAt        *string `json:"created_at,omitempty"`
	CustomAttributes any     `json:"custom_attributes,omitempty"`
	Email            *string `json:"email,omitempty"`
	ExternalId       *string `json:"external_id,omitempty"`
	FirstName        *string `json:"first_name,omitempty"`
	Id               *string `json:"id,omitempty"`
	LastName         *string `json:"last_name,omitempty"`
	LastNotifiedAt   *string `json:"last_notified_at,omitempty"`
	LastSeenAt       *string `json:"last_seen_at,omitempty"`
	UpdatedAt        *string `json:"updated_at,omitempty"`
}

func (u *User) GetCreatedAt() *string {
	if u == nil {
		return nil
	}
	return u.CreatedAt
}

func (u *User) SetCreatedAt(createdAt string) {
	u.CreatedAt = &createdAt
}

func (u *User) GetCustomAttributes() any {
	if u == nil {
		return nil
	}
	return u.CustomAttributes
}

func (u *User) SetCustomAttributes(customAttributes any) {
	u.CustomAttributes = &customAttributes
}

func (u *User) GetEmail() *string {
	if u == nil {
		return nil
	}
	return u.Email
}

func (u *User) SetEmail(email string) {
	u.Email = &email
}

func (u *User) GetExternalId() *string {
	if u == nil {
		return nil
	}
	return u.ExternalId
}

func (u *User) SetExternalId(externalId string) {
	u.ExternalId = &externalId
}

func (u *User) GetFirstName() *string {
	if u == nil {
		return nil
	}
	return u.FirstName
}

func (u *User) SetFirstName(firstName string) {
	u.FirstName = &firstName
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

func (u *User) GetLastName() *string {
	if u == nil {
		return nil
	}
	return u.LastName
}

func (u *User) SetLastName(lastName string) {
	u.LastName = &lastName
}

func (u *User) GetLastNotifiedAt() *string {
	if u == nil {
		return nil
	}
	return u.LastNotifiedAt
}

func (u *User) SetLastNotifiedAt(lastNotifiedAt string) {
	u.LastNotifiedAt = &lastNotifiedAt
}

func (u *User) GetLastSeenAt() *string {
	if u == nil {
		return nil
	}
	return u.LastSeenAt
}

func (u *User) SetLastSeenAt(lastSeenAt string) {
	u.LastSeenAt = &lastSeenAt
}

func (u *User) GetUpdatedAt() *string {
	if u == nil {
		return nil
	}
	return u.UpdatedAt
}

func (u *User) SetUpdatedAt(updatedAt string) {
	u.UpdatedAt = &updatedAt
}

func (u User) String() string {
	jsonData, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		return "error converting struct: User to string"
	}
	return string(jsonData)
}
