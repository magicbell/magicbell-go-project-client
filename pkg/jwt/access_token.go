package jwt

import (
	"encoding/json"
)

type AccessToken struct {
	CreatedAt *string `json:"created_at,omitempty" required:"true"`
	ExpiresAt *string `json:"expires_at,omitempty"`
	Token     *string `json:"token,omitempty" required:"true"`
	TokenId   *string `json:"token_id,omitempty" required:"true"`
	touched   map[string]bool
}

func (a *AccessToken) GetCreatedAt() *string {
	if a == nil {
		return nil
	}
	return a.CreatedAt
}

func (a *AccessToken) SetCreatedAt(createdAt string) {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["CreatedAt"] = true
	a.CreatedAt = &createdAt
}

func (a *AccessToken) SetCreatedAtNil() {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["CreatedAt"] = true
	a.CreatedAt = nil
}

func (a *AccessToken) GetExpiresAt() *string {
	if a == nil {
		return nil
	}
	return a.ExpiresAt
}

func (a *AccessToken) SetExpiresAt(expiresAt string) {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["ExpiresAt"] = true
	a.ExpiresAt = &expiresAt
}

func (a *AccessToken) SetExpiresAtNil() {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["ExpiresAt"] = true
	a.ExpiresAt = nil
}

func (a *AccessToken) GetToken() *string {
	if a == nil {
		return nil
	}
	return a.Token
}

func (a *AccessToken) SetToken(token string) {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["Token"] = true
	a.Token = &token
}

func (a *AccessToken) SetTokenNil() {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["Token"] = true
	a.Token = nil
}

func (a *AccessToken) GetTokenId() *string {
	if a == nil {
		return nil
	}
	return a.TokenId
}

func (a *AccessToken) SetTokenId(tokenId string) {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["TokenId"] = true
	a.TokenId = &tokenId
}

func (a *AccessToken) SetTokenIdNil() {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["TokenId"] = true
	a.TokenId = nil
}
func (a AccessToken) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if a.touched["CreatedAt"] && a.CreatedAt == nil {
		data["created_at"] = nil
	} else if a.CreatedAt != nil {
		data["created_at"] = a.CreatedAt
	}

	if a.touched["ExpiresAt"] && a.ExpiresAt == nil {
		data["expires_at"] = nil
	} else if a.ExpiresAt != nil {
		data["expires_at"] = a.ExpiresAt
	}

	if a.touched["Token"] && a.Token == nil {
		data["token"] = nil
	} else if a.Token != nil {
		data["token"] = a.Token
	}

	if a.touched["TokenId"] && a.TokenId == nil {
		data["token_id"] = nil
	} else if a.TokenId != nil {
		data["token_id"] = a.TokenId
	}

	return json.Marshal(data)
}
