package jwt

import (
	"encoding/json"
)

type CreateProjectTokenRequest struct {
	// The duration for which the token is valid (in seconds)
	Expiry *int64 `json:"expiry,omitempty" min:"1"`
	// The name of the token.
	Name    *string `json:"name,omitempty" required:"true" maxLength:"255"`
	touched map[string]bool
}

func (c *CreateProjectTokenRequest) GetExpiry() *int64 {
	if c == nil {
		return nil
	}
	return c.Expiry
}

func (c *CreateProjectTokenRequest) SetExpiry(expiry int64) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Expiry"] = true
	c.Expiry = &expiry
}

func (c *CreateProjectTokenRequest) SetExpiryNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Expiry"] = true
	c.Expiry = nil
}

func (c *CreateProjectTokenRequest) GetName() *string {
	if c == nil {
		return nil
	}
	return c.Name
}

func (c *CreateProjectTokenRequest) SetName(name string) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Name"] = true
	c.Name = &name
}

func (c *CreateProjectTokenRequest) SetNameNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Name"] = true
	c.Name = nil
}
func (c CreateProjectTokenRequest) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if c.touched["Expiry"] && c.Expiry == nil {
		data["expiry"] = nil
	} else if c.Expiry != nil {
		data["expiry"] = c.Expiry
	}

	if c.touched["Name"] && c.Name == nil {
		data["name"] = nil
	} else if c.Name != nil {
		data["name"] = c.Name
	}

	return json.Marshal(data)
}
