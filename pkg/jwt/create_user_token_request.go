package jwt

import (
	"encoding/json"
)

type CreateUserTokenRequest struct {
	// The user's email.
	Email *string `json:"email,omitempty" maxLength:"255"`
	// The duration for which the token is valid (in seconds)
	Expiry *int64 `json:"expiry,omitempty" min:"1"`
	// A unique string that MagicBell can utilize to identify the user uniquely. We recommend setting this attribute to the ID of the user in your database. Provide the external id if the user's email is unavailable.
	ExternalId *string `json:"external_id,omitempty" maxLength:"255"`
	// The name of the token.
	Name    *string `json:"name,omitempty" maxLength:"255"`
	touched map[string]bool
}

func (c *CreateUserTokenRequest) GetEmail() *string {
	if c == nil {
		return nil
	}
	return c.Email
}

func (c *CreateUserTokenRequest) SetEmail(email string) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Email"] = true
	c.Email = &email
}

func (c *CreateUserTokenRequest) SetEmailNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Email"] = true
	c.Email = nil
}

func (c *CreateUserTokenRequest) GetExpiry() *int64 {
	if c == nil {
		return nil
	}
	return c.Expiry
}

func (c *CreateUserTokenRequest) SetExpiry(expiry int64) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Expiry"] = true
	c.Expiry = &expiry
}

func (c *CreateUserTokenRequest) SetExpiryNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Expiry"] = true
	c.Expiry = nil
}

func (c *CreateUserTokenRequest) GetExternalId() *string {
	if c == nil {
		return nil
	}
	return c.ExternalId
}

func (c *CreateUserTokenRequest) SetExternalId(externalId string) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["ExternalId"] = true
	c.ExternalId = &externalId
}

func (c *CreateUserTokenRequest) SetExternalIdNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["ExternalId"] = true
	c.ExternalId = nil
}

func (c *CreateUserTokenRequest) GetName() *string {
	if c == nil {
		return nil
	}
	return c.Name
}

func (c *CreateUserTokenRequest) SetName(name string) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Name"] = true
	c.Name = &name
}

func (c *CreateUserTokenRequest) SetNameNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Name"] = true
	c.Name = nil
}
func (c CreateUserTokenRequest) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if c.touched["Email"] && c.Email == nil {
		data["email"] = nil
	} else if c.Email != nil {
		data["email"] = c.Email
	}

	if c.touched["Expiry"] && c.Expiry == nil {
		data["expiry"] = nil
	} else if c.Expiry != nil {
		data["expiry"] = c.Expiry
	}

	if c.touched["ExternalId"] && c.ExternalId == nil {
		data["external_id"] = nil
	} else if c.ExternalId != nil {
		data["external_id"] = c.ExternalId
	}

	if c.touched["Name"] && c.Name == nil {
		data["name"] = nil
	} else if c.Name != nil {
		data["name"] = c.Name
	}

	return json.Marshal(data)
}
