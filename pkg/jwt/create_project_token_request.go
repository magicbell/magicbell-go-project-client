package jwt

import (
	"encoding/json"
	"github.com/magicbell/magicbell-go-project-client/pkg/util"
)

type CreateProjectTokenRequest struct {
	// The duration for which the token is valid (in seconds)
	Expiry *util.Nullable[int64] `json:"expiry,omitempty" min:"1"`
	// The name of the token.
	Name *string `json:"name,omitempty" required:"true" maxLength:"255"`
}

func (c *CreateProjectTokenRequest) GetExpiry() *util.Nullable[int64] {
	if c == nil {
		return nil
	}
	return c.Expiry
}

func (c *CreateProjectTokenRequest) SetExpiry(expiry util.Nullable[int64]) {
	c.Expiry = &expiry
}

func (c *CreateProjectTokenRequest) SetExpiryNull() {
	c.Expiry = &util.Nullable[int64]{IsNull: true}
}

func (c *CreateProjectTokenRequest) GetName() *string {
	if c == nil {
		return nil
	}
	return c.Name
}

func (c *CreateProjectTokenRequest) SetName(name string) {
	c.Name = &name
}

func (c CreateProjectTokenRequest) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CreateProjectTokenRequest to string"
	}
	return string(jsonData)
}
