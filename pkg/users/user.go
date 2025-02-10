package users

import (
	"encoding/json"
)

type User struct {
	CreatedAt        *string  `json:"created_at,omitempty" required:"true"`
	CustomAttributes any      `json:"custom_attributes,omitempty" required:"true"`
	Email            *string  `json:"email,omitempty"`
	ExternalId       *string  `json:"external_id,omitempty"`
	FirstName        *string  `json:"first_name,omitempty"`
	Id               *string  `json:"id,omitempty" required:"true"`
	LastName         *string  `json:"last_name,omitempty"`
	LastNotifiedAt   *string  `json:"last_notified_at,omitempty"`
	LastSeenAt       *string  `json:"last_seen_at,omitempty"`
	PhoneNumbers     []string `json:"phone_numbers,omitempty"`
	ProjectId        *int64   `json:"project_id,omitempty" required:"true"`
	UpdatedAt        *string  `json:"updated_at,omitempty" required:"true"`
	touched          map[string]bool
}

func (u *User) GetCreatedAt() *string {
	if u == nil {
		return nil
	}
	return u.CreatedAt
}

func (u *User) SetCreatedAt(createdAt string) {
	if u.touched == nil {
		u.touched = map[string]bool{}
	}
	u.touched["CreatedAt"] = true
	u.CreatedAt = &createdAt
}

func (u *User) SetCreatedAtNil() {
	if u.touched == nil {
		u.touched = map[string]bool{}
	}
	u.touched["CreatedAt"] = true
	u.CreatedAt = nil
}

func (u *User) GetCustomAttributes() any {
	if u == nil {
		return nil
	}
	return u.CustomAttributes
}

func (u *User) SetCustomAttributes(customAttributes any) {
	if u.touched == nil {
		u.touched = map[string]bool{}
	}
	u.touched["CustomAttributes"] = true
	u.CustomAttributes = customAttributes
}

func (u *User) SetCustomAttributesNil() {
	if u.touched == nil {
		u.touched = map[string]bool{}
	}
	u.touched["CustomAttributes"] = true
	u.CustomAttributes = nil
}

func (u *User) GetEmail() *string {
	if u == nil {
		return nil
	}
	return u.Email
}

func (u *User) SetEmail(email string) {
	if u.touched == nil {
		u.touched = map[string]bool{}
	}
	u.touched["Email"] = true
	u.Email = &email
}

func (u *User) SetEmailNil() {
	if u.touched == nil {
		u.touched = map[string]bool{}
	}
	u.touched["Email"] = true
	u.Email = nil
}

func (u *User) GetExternalId() *string {
	if u == nil {
		return nil
	}
	return u.ExternalId
}

func (u *User) SetExternalId(externalId string) {
	if u.touched == nil {
		u.touched = map[string]bool{}
	}
	u.touched["ExternalId"] = true
	u.ExternalId = &externalId
}

func (u *User) SetExternalIdNil() {
	if u.touched == nil {
		u.touched = map[string]bool{}
	}
	u.touched["ExternalId"] = true
	u.ExternalId = nil
}

func (u *User) GetFirstName() *string {
	if u == nil {
		return nil
	}
	return u.FirstName
}

func (u *User) SetFirstName(firstName string) {
	if u.touched == nil {
		u.touched = map[string]bool{}
	}
	u.touched["FirstName"] = true
	u.FirstName = &firstName
}

func (u *User) SetFirstNameNil() {
	if u.touched == nil {
		u.touched = map[string]bool{}
	}
	u.touched["FirstName"] = true
	u.FirstName = nil
}

func (u *User) GetId() *string {
	if u == nil {
		return nil
	}
	return u.Id
}

func (u *User) SetId(id string) {
	if u.touched == nil {
		u.touched = map[string]bool{}
	}
	u.touched["Id"] = true
	u.Id = &id
}

func (u *User) SetIdNil() {
	if u.touched == nil {
		u.touched = map[string]bool{}
	}
	u.touched["Id"] = true
	u.Id = nil
}

func (u *User) GetLastName() *string {
	if u == nil {
		return nil
	}
	return u.LastName
}

func (u *User) SetLastName(lastName string) {
	if u.touched == nil {
		u.touched = map[string]bool{}
	}
	u.touched["LastName"] = true
	u.LastName = &lastName
}

func (u *User) SetLastNameNil() {
	if u.touched == nil {
		u.touched = map[string]bool{}
	}
	u.touched["LastName"] = true
	u.LastName = nil
}

func (u *User) GetLastNotifiedAt() *string {
	if u == nil {
		return nil
	}
	return u.LastNotifiedAt
}

func (u *User) SetLastNotifiedAt(lastNotifiedAt string) {
	if u.touched == nil {
		u.touched = map[string]bool{}
	}
	u.touched["LastNotifiedAt"] = true
	u.LastNotifiedAt = &lastNotifiedAt
}

func (u *User) SetLastNotifiedAtNil() {
	if u.touched == nil {
		u.touched = map[string]bool{}
	}
	u.touched["LastNotifiedAt"] = true
	u.LastNotifiedAt = nil
}

func (u *User) GetLastSeenAt() *string {
	if u == nil {
		return nil
	}
	return u.LastSeenAt
}

func (u *User) SetLastSeenAt(lastSeenAt string) {
	if u.touched == nil {
		u.touched = map[string]bool{}
	}
	u.touched["LastSeenAt"] = true
	u.LastSeenAt = &lastSeenAt
}

func (u *User) SetLastSeenAtNil() {
	if u.touched == nil {
		u.touched = map[string]bool{}
	}
	u.touched["LastSeenAt"] = true
	u.LastSeenAt = nil
}

func (u *User) GetPhoneNumbers() []string {
	if u == nil {
		return nil
	}
	return u.PhoneNumbers
}

func (u *User) SetPhoneNumbers(phoneNumbers []string) {
	if u.touched == nil {
		u.touched = map[string]bool{}
	}
	u.touched["PhoneNumbers"] = true
	u.PhoneNumbers = phoneNumbers
}

func (u *User) SetPhoneNumbersNil() {
	if u.touched == nil {
		u.touched = map[string]bool{}
	}
	u.touched["PhoneNumbers"] = true
	u.PhoneNumbers = nil
}

func (u *User) GetProjectId() *int64 {
	if u == nil {
		return nil
	}
	return u.ProjectId
}

func (u *User) SetProjectId(projectId int64) {
	if u.touched == nil {
		u.touched = map[string]bool{}
	}
	u.touched["ProjectId"] = true
	u.ProjectId = &projectId
}

func (u *User) SetProjectIdNil() {
	if u.touched == nil {
		u.touched = map[string]bool{}
	}
	u.touched["ProjectId"] = true
	u.ProjectId = nil
}

func (u *User) GetUpdatedAt() *string {
	if u == nil {
		return nil
	}
	return u.UpdatedAt
}

func (u *User) SetUpdatedAt(updatedAt string) {
	if u.touched == nil {
		u.touched = map[string]bool{}
	}
	u.touched["UpdatedAt"] = true
	u.UpdatedAt = &updatedAt
}

func (u *User) SetUpdatedAtNil() {
	if u.touched == nil {
		u.touched = map[string]bool{}
	}
	u.touched["UpdatedAt"] = true
	u.UpdatedAt = nil
}

func (u User) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if u.touched["CreatedAt"] && u.CreatedAt == nil {
		data["created_at"] = nil
	} else if u.CreatedAt != nil {
		data["created_at"] = u.CreatedAt
	}

	if u.touched["CustomAttributes"] && u.CustomAttributes == nil {
		data["custom_attributes"] = nil
	} else if u.CustomAttributes != nil {
		data["custom_attributes"] = u.CustomAttributes
	}

	if u.touched["Email"] && u.Email == nil {
		data["email"] = nil
	} else if u.Email != nil {
		data["email"] = u.Email
	}

	if u.touched["ExternalId"] && u.ExternalId == nil {
		data["external_id"] = nil
	} else if u.ExternalId != nil {
		data["external_id"] = u.ExternalId
	}

	if u.touched["FirstName"] && u.FirstName == nil {
		data["first_name"] = nil
	} else if u.FirstName != nil {
		data["first_name"] = u.FirstName
	}

	if u.touched["Id"] && u.Id == nil {
		data["id"] = nil
	} else if u.Id != nil {
		data["id"] = u.Id
	}

	if u.touched["LastName"] && u.LastName == nil {
		data["last_name"] = nil
	} else if u.LastName != nil {
		data["last_name"] = u.LastName
	}

	if u.touched["LastNotifiedAt"] && u.LastNotifiedAt == nil {
		data["last_notified_at"] = nil
	} else if u.LastNotifiedAt != nil {
		data["last_notified_at"] = u.LastNotifiedAt
	}

	if u.touched["LastSeenAt"] && u.LastSeenAt == nil {
		data["last_seen_at"] = nil
	} else if u.LastSeenAt != nil {
		data["last_seen_at"] = u.LastSeenAt
	}

	if u.touched["PhoneNumbers"] && u.PhoneNumbers == nil {
		data["phone_numbers"] = nil
	} else if u.PhoneNumbers != nil {
		data["phone_numbers"] = u.PhoneNumbers
	}

	if u.touched["ProjectId"] && u.ProjectId == nil {
		data["project_id"] = nil
	} else if u.ProjectId != nil {
		data["project_id"] = u.ProjectId
	}

	if u.touched["UpdatedAt"] && u.UpdatedAt == nil {
		data["updated_at"] = nil
	} else if u.UpdatedAt != nil {
		data["updated_at"] = u.UpdatedAt
	}

	return json.Marshal(data)
}

func (u User) String() string {
	jsonData, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		return "error converting struct: User to string"
	}
	return string(jsonData)
}
