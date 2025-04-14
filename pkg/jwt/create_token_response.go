package jwt

import (
	"encoding/json"
	"github.com/magicbell/magicbell-go-project-client/pkg/util"
)

type CreateTokenResponse struct {
	CreatedAt *string                `json:"created_at,omitempty" required:"true"`
	ExpiresAt *util.Nullable[string] `json:"expires_at,omitempty"`
	Token     *string                `json:"token,omitempty" required:"true"`
	TokenId   *string                `json:"token_id,omitempty" required:"true"`
}

func (c *CreateTokenResponse) GetCreatedAt() *string {
	if c == nil {
		return nil
	}
	return c.CreatedAt
}

func (c *CreateTokenResponse) SetCreatedAt(createdAt string) {
	c.CreatedAt = &createdAt
}

func (c *CreateTokenResponse) GetExpiresAt() *util.Nullable[string] {
	if c == nil {
		return nil
	}
	return c.ExpiresAt
}

func (c *CreateTokenResponse) SetExpiresAt(expiresAt util.Nullable[string]) {
	c.ExpiresAt = &expiresAt
}

func (c *CreateTokenResponse) SetExpiresAtNull() {
	c.ExpiresAt = &util.Nullable[string]{IsNull: true}
}

func (c *CreateTokenResponse) GetToken() *string {
	if c == nil {
		return nil
	}
	return c.Token
}

func (c *CreateTokenResponse) SetToken(token string) {
	c.Token = &token
}

func (c *CreateTokenResponse) GetTokenId() *string {
	if c == nil {
		return nil
	}
	return c.TokenId
}

func (c *CreateTokenResponse) SetTokenId(tokenId string) {
	c.TokenId = &tokenId
}

func (c CreateTokenResponse) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CreateTokenResponse to string"
	}
	return string(jsonData)
}
