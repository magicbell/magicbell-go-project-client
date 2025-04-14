package jwt

import (
	"encoding/json"
	"github.com/magicbell/magicbell-go-project-client/pkg/util"
)

type AccessToken struct {
	CreatedAt   *string                `json:"created_at,omitempty" required:"true"`
	DiscardedAt *util.Nullable[string] `json:"discarded_at,omitempty"`
	ExpiresAt   *string                `json:"expires_at,omitempty"`
	Id          *string                `json:"id,omitempty"`
	Name        *util.Nullable[string] `json:"name,omitempty"`
}

func (a *AccessToken) GetCreatedAt() *string {
	if a == nil {
		return nil
	}
	return a.CreatedAt
}

func (a *AccessToken) SetCreatedAt(createdAt string) {
	a.CreatedAt = &createdAt
}

func (a *AccessToken) GetDiscardedAt() *util.Nullable[string] {
	if a == nil {
		return nil
	}
	return a.DiscardedAt
}

func (a *AccessToken) SetDiscardedAt(discardedAt util.Nullable[string]) {
	a.DiscardedAt = &discardedAt
}

func (a *AccessToken) SetDiscardedAtNull() {
	a.DiscardedAt = &util.Nullable[string]{IsNull: true}
}

func (a *AccessToken) GetExpiresAt() *string {
	if a == nil {
		return nil
	}
	return a.ExpiresAt
}

func (a *AccessToken) SetExpiresAt(expiresAt string) {
	a.ExpiresAt = &expiresAt
}

func (a *AccessToken) GetId() *string {
	if a == nil {
		return nil
	}
	return a.Id
}

func (a *AccessToken) SetId(id string) {
	a.Id = &id
}

func (a *AccessToken) GetName() *util.Nullable[string] {
	if a == nil {
		return nil
	}
	return a.Name
}

func (a *AccessToken) SetName(name util.Nullable[string]) {
	a.Name = &name
}

func (a *AccessToken) SetNameNull() {
	a.Name = &util.Nullable[string]{IsNull: true}
}

func (a AccessToken) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AccessToken to string"
	}
	return string(jsonData)
}
