package channels

import (
	"encoding/json"
)

type TeamsTokenResponse struct {
	CreatedAt   *string                    `json:"created_at,omitempty" required:"true"`
	DiscardedAt *string                    `json:"discarded_at,omitempty"`
	Id          *string                    `json:"id,omitempty" required:"true"`
	UpdatedAt   *string                    `json:"updated_at,omitempty"`
	Webhook     *TeamsTokenResponseWebhook `json:"webhook,omitempty"`
	touched     map[string]bool
}

func (t *TeamsTokenResponse) GetCreatedAt() *string {
	if t == nil {
		return nil
	}
	return t.CreatedAt
}

func (t *TeamsTokenResponse) SetCreatedAt(createdAt string) {
	if t.touched == nil {
		t.touched = map[string]bool{}
	}
	t.touched["CreatedAt"] = true
	t.CreatedAt = &createdAt
}

func (t *TeamsTokenResponse) SetCreatedAtNil() {
	if t.touched == nil {
		t.touched = map[string]bool{}
	}
	t.touched["CreatedAt"] = true
	t.CreatedAt = nil
}

func (t *TeamsTokenResponse) GetDiscardedAt() *string {
	if t == nil {
		return nil
	}
	return t.DiscardedAt
}

func (t *TeamsTokenResponse) SetDiscardedAt(discardedAt string) {
	if t.touched == nil {
		t.touched = map[string]bool{}
	}
	t.touched["DiscardedAt"] = true
	t.DiscardedAt = &discardedAt
}

func (t *TeamsTokenResponse) SetDiscardedAtNil() {
	if t.touched == nil {
		t.touched = map[string]bool{}
	}
	t.touched["DiscardedAt"] = true
	t.DiscardedAt = nil
}

func (t *TeamsTokenResponse) GetId() *string {
	if t == nil {
		return nil
	}
	return t.Id
}

func (t *TeamsTokenResponse) SetId(id string) {
	if t.touched == nil {
		t.touched = map[string]bool{}
	}
	t.touched["Id"] = true
	t.Id = &id
}

func (t *TeamsTokenResponse) SetIdNil() {
	if t.touched == nil {
		t.touched = map[string]bool{}
	}
	t.touched["Id"] = true
	t.Id = nil
}

func (t *TeamsTokenResponse) GetUpdatedAt() *string {
	if t == nil {
		return nil
	}
	return t.UpdatedAt
}

func (t *TeamsTokenResponse) SetUpdatedAt(updatedAt string) {
	if t.touched == nil {
		t.touched = map[string]bool{}
	}
	t.touched["UpdatedAt"] = true
	t.UpdatedAt = &updatedAt
}

func (t *TeamsTokenResponse) SetUpdatedAtNil() {
	if t.touched == nil {
		t.touched = map[string]bool{}
	}
	t.touched["UpdatedAt"] = true
	t.UpdatedAt = nil
}

func (t *TeamsTokenResponse) GetWebhook() *TeamsTokenResponseWebhook {
	if t == nil {
		return nil
	}
	return t.Webhook
}

func (t *TeamsTokenResponse) SetWebhook(webhook TeamsTokenResponseWebhook) {
	if t.touched == nil {
		t.touched = map[string]bool{}
	}
	t.touched["Webhook"] = true
	t.Webhook = &webhook
}

func (t *TeamsTokenResponse) SetWebhookNil() {
	if t.touched == nil {
		t.touched = map[string]bool{}
	}
	t.touched["Webhook"] = true
	t.Webhook = nil
}

func (t TeamsTokenResponse) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if t.touched["CreatedAt"] && t.CreatedAt == nil {
		data["created_at"] = nil
	} else if t.CreatedAt != nil {
		data["created_at"] = t.CreatedAt
	}

	if t.touched["DiscardedAt"] && t.DiscardedAt == nil {
		data["discarded_at"] = nil
	} else if t.DiscardedAt != nil {
		data["discarded_at"] = t.DiscardedAt
	}

	if t.touched["Id"] && t.Id == nil {
		data["id"] = nil
	} else if t.Id != nil {
		data["id"] = t.Id
	}

	if t.touched["UpdatedAt"] && t.UpdatedAt == nil {
		data["updated_at"] = nil
	} else if t.UpdatedAt != nil {
		data["updated_at"] = t.UpdatedAt
	}

	if t.touched["Webhook"] && t.Webhook == nil {
		data["webhook"] = nil
	} else if t.Webhook != nil {
		data["webhook"] = t.Webhook
	}

	return json.Marshal(data)
}

func (t TeamsTokenResponse) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TeamsTokenResponse to string"
	}
	return string(jsonData)
}

type TeamsTokenResponseWebhook struct {
	Url     *string `json:"url,omitempty"`
	touched map[string]bool
}

func (t *TeamsTokenResponseWebhook) GetUrl() *string {
	if t == nil {
		return nil
	}
	return t.Url
}

func (t *TeamsTokenResponseWebhook) SetUrl(url string) {
	if t.touched == nil {
		t.touched = map[string]bool{}
	}
	t.touched["Url"] = true
	t.Url = &url
}

func (t *TeamsTokenResponseWebhook) SetUrlNil() {
	if t.touched == nil {
		t.touched = map[string]bool{}
	}
	t.touched["Url"] = true
	t.Url = nil
}

func (t TeamsTokenResponseWebhook) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if t.touched["Url"] && t.Url == nil {
		data["url"] = nil
	} else if t.Url != nil {
		data["url"] = t.Url
	}

	return json.Marshal(data)
}

func (t TeamsTokenResponseWebhook) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TeamsTokenResponseWebhook to string"
	}
	return string(jsonData)
}
