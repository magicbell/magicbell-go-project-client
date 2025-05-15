package jwt

import (
	"encoding/json"
	"github.com/magicbell/magicbell-go/pkg/project-client/util"
)

type CreateUserTokenRequest struct {
	// The user's email.
	Email *util.Nullable[string] `json:"email,omitempty" maxLength:"255"`
	// The duration for which the token is valid (in seconds)
	Expiry *util.Nullable[int64] `json:"expiry,omitempty" min:"1"`
	// A unique string that MagicBell can utilize to identify the user uniquely. We recommend setting this attribute to the ID of the user in your database. Provide the external id if the user's email is unavailable.
	ExternalId *util.Nullable[string] `json:"external_id,omitempty" maxLength:"255"`
	// The name of the token.
	Name *util.Nullable[string] `json:"name,omitempty" maxLength:"255"`
}

func (c *CreateUserTokenRequest) GetEmail() *util.Nullable[string] {
	if c == nil {
		return nil
	}
	return c.Email
}

func (c *CreateUserTokenRequest) SetEmail(email util.Nullable[string]) {
	c.Email = &email
}

func (c *CreateUserTokenRequest) SetEmailNull() {
	c.Email = &util.Nullable[string]{IsNull: true}
}

func (c *CreateUserTokenRequest) GetExpiry() *util.Nullable[int64] {
	if c == nil {
		return nil
	}
	return c.Expiry
}

func (c *CreateUserTokenRequest) SetExpiry(expiry util.Nullable[int64]) {
	c.Expiry = &expiry
}

func (c *CreateUserTokenRequest) SetExpiryNull() {
	c.Expiry = &util.Nullable[int64]{IsNull: true}
}

func (c *CreateUserTokenRequest) GetExternalId() *util.Nullable[string] {
	if c == nil {
		return nil
	}
	return c.ExternalId
}

func (c *CreateUserTokenRequest) SetExternalId(externalId util.Nullable[string]) {
	c.ExternalId = &externalId
}

func (c *CreateUserTokenRequest) SetExternalIdNull() {
	c.ExternalId = &util.Nullable[string]{IsNull: true}
}

func (c *CreateUserTokenRequest) GetName() *util.Nullable[string] {
	if c == nil {
		return nil
	}
	return c.Name
}

func (c *CreateUserTokenRequest) SetName(name util.Nullable[string]) {
	c.Name = &name
}

func (c *CreateUserTokenRequest) SetNameNull() {
	c.Name = &util.Nullable[string]{IsNull: true}
}

func (c CreateUserTokenRequest) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CreateUserTokenRequest to string"
	}
	return string(jsonData)
}
