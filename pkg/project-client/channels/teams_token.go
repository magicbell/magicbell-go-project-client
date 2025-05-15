package channels

import (
	"encoding/json"
	"github.com/magicbell/magicbell-go/pkg/project-client/util"
)

type TeamsToken struct {
	CreatedAt   *string                `json:"created_at,omitempty" required:"true"`
	DiscardedAt *util.Nullable[string] `json:"discarded_at,omitempty"`
	Id          *string                `json:"id,omitempty" required:"true"`
	UpdatedAt   *util.Nullable[string] `json:"updated_at,omitempty"`
	Webhook     *TeamsTokenWebhook     `json:"webhook,omitempty"`
}

func (t *TeamsToken) GetCreatedAt() *string {
	if t == nil {
		return nil
	}
	return t.CreatedAt
}

func (t *TeamsToken) SetCreatedAt(createdAt string) {
	t.CreatedAt = &createdAt
}

func (t *TeamsToken) GetDiscardedAt() *util.Nullable[string] {
	if t == nil {
		return nil
	}
	return t.DiscardedAt
}

func (t *TeamsToken) SetDiscardedAt(discardedAt util.Nullable[string]) {
	t.DiscardedAt = &discardedAt
}

func (t *TeamsToken) SetDiscardedAtNull() {
	t.DiscardedAt = &util.Nullable[string]{IsNull: true}
}

func (t *TeamsToken) GetId() *string {
	if t == nil {
		return nil
	}
	return t.Id
}

func (t *TeamsToken) SetId(id string) {
	t.Id = &id
}

func (t *TeamsToken) GetUpdatedAt() *util.Nullable[string] {
	if t == nil {
		return nil
	}
	return t.UpdatedAt
}

func (t *TeamsToken) SetUpdatedAt(updatedAt util.Nullable[string]) {
	t.UpdatedAt = &updatedAt
}

func (t *TeamsToken) SetUpdatedAtNull() {
	t.UpdatedAt = &util.Nullable[string]{IsNull: true}
}

func (t *TeamsToken) GetWebhook() *TeamsTokenWebhook {
	if t == nil {
		return nil
	}
	return t.Webhook
}

func (t *TeamsToken) SetWebhook(webhook TeamsTokenWebhook) {
	t.Webhook = &webhook
}

func (t TeamsToken) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TeamsToken to string"
	}
	return string(jsonData)
}

type TeamsTokenWebhook struct {
	Url *string `json:"url,omitempty"`
}

func (t *TeamsTokenWebhook) GetUrl() *string {
	if t == nil {
		return nil
	}
	return t.Url
}

func (t *TeamsTokenWebhook) SetUrl(url string) {
	t.Url = &url
}

func (t TeamsTokenWebhook) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TeamsTokenWebhook to string"
	}
	return string(jsonData)
}
