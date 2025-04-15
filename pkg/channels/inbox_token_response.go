package channels

import (
	"encoding/json"
	"github.com/magicbell/magicbell-go-project-client/pkg/util"
)

type InboxTokenResponse struct {
	ConnectionId *util.Nullable[string] `json:"connection_id,omitempty"`
	CreatedAt    *string                `json:"created_at,omitempty" required:"true"`
	DiscardedAt  *util.Nullable[string] `json:"discarded_at,omitempty"`
	Id           *string                `json:"id,omitempty" required:"true"`
	Token        *string                `json:"token,omitempty" required:"true" minLength:"64"`
	UpdatedAt    *util.Nullable[string] `json:"updated_at,omitempty"`
}

func (i *InboxTokenResponse) GetConnectionId() *util.Nullable[string] {
	if i == nil {
		return nil
	}
	return i.ConnectionId
}

func (i *InboxTokenResponse) SetConnectionId(connectionId util.Nullable[string]) {
	i.ConnectionId = &connectionId
}

func (i *InboxTokenResponse) SetConnectionIdNull() {
	i.ConnectionId = &util.Nullable[string]{IsNull: true}
}

func (i *InboxTokenResponse) GetCreatedAt() *string {
	if i == nil {
		return nil
	}
	return i.CreatedAt
}

func (i *InboxTokenResponse) SetCreatedAt(createdAt string) {
	i.CreatedAt = &createdAt
}

func (i *InboxTokenResponse) GetDiscardedAt() *util.Nullable[string] {
	if i == nil {
		return nil
	}
	return i.DiscardedAt
}

func (i *InboxTokenResponse) SetDiscardedAt(discardedAt util.Nullable[string]) {
	i.DiscardedAt = &discardedAt
}

func (i *InboxTokenResponse) SetDiscardedAtNull() {
	i.DiscardedAt = &util.Nullable[string]{IsNull: true}
}

func (i *InboxTokenResponse) GetId() *string {
	if i == nil {
		return nil
	}
	return i.Id
}

func (i *InboxTokenResponse) SetId(id string) {
	i.Id = &id
}

func (i *InboxTokenResponse) GetToken() *string {
	if i == nil {
		return nil
	}
	return i.Token
}

func (i *InboxTokenResponse) SetToken(token string) {
	i.Token = &token
}

func (i *InboxTokenResponse) GetUpdatedAt() *util.Nullable[string] {
	if i == nil {
		return nil
	}
	return i.UpdatedAt
}

func (i *InboxTokenResponse) SetUpdatedAt(updatedAt util.Nullable[string]) {
	i.UpdatedAt = &updatedAt
}

func (i *InboxTokenResponse) SetUpdatedAtNull() {
	i.UpdatedAt = &util.Nullable[string]{IsNull: true}
}

func (i InboxTokenResponse) String() string {
	jsonData, err := json.MarshalIndent(i, "", "  ")
	if err != nil {
		return "error converting struct: InboxTokenResponse to string"
	}
	return string(jsonData)
}
