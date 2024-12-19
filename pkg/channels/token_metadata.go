package channels

import (
	"encoding/json"
)

type TokenMetadata struct {
	CreatedAt   *string `json:"created_at,omitempty" required:"true"`
	DiscardedAt *string `json:"discarded_at,omitempty"`
	Id          *string `json:"id,omitempty" required:"true"`
	UpdatedAt   *string `json:"updated_at,omitempty"`
	touched     map[string]bool
}

func (t *TokenMetadata) GetCreatedAt() *string {
	if t == nil {
		return nil
	}
	return t.CreatedAt
}

func (t *TokenMetadata) SetCreatedAt(createdAt string) {
	if t.touched == nil {
		t.touched = map[string]bool{}
	}
	t.touched["CreatedAt"] = true
	t.CreatedAt = &createdAt
}

func (t *TokenMetadata) SetCreatedAtNil() {
	if t.touched == nil {
		t.touched = map[string]bool{}
	}
	t.touched["CreatedAt"] = true
	t.CreatedAt = nil
}

func (t *TokenMetadata) GetDiscardedAt() *string {
	if t == nil {
		return nil
	}
	return t.DiscardedAt
}

func (t *TokenMetadata) SetDiscardedAt(discardedAt string) {
	if t.touched == nil {
		t.touched = map[string]bool{}
	}
	t.touched["DiscardedAt"] = true
	t.DiscardedAt = &discardedAt
}

func (t *TokenMetadata) SetDiscardedAtNil() {
	if t.touched == nil {
		t.touched = map[string]bool{}
	}
	t.touched["DiscardedAt"] = true
	t.DiscardedAt = nil
}

func (t *TokenMetadata) GetId() *string {
	if t == nil {
		return nil
	}
	return t.Id
}

func (t *TokenMetadata) SetId(id string) {
	if t.touched == nil {
		t.touched = map[string]bool{}
	}
	t.touched["Id"] = true
	t.Id = &id
}

func (t *TokenMetadata) SetIdNil() {
	if t.touched == nil {
		t.touched = map[string]bool{}
	}
	t.touched["Id"] = true
	t.Id = nil
}

func (t *TokenMetadata) GetUpdatedAt() *string {
	if t == nil {
		return nil
	}
	return t.UpdatedAt
}

func (t *TokenMetadata) SetUpdatedAt(updatedAt string) {
	if t.touched == nil {
		t.touched = map[string]bool{}
	}
	t.touched["UpdatedAt"] = true
	t.UpdatedAt = &updatedAt
}

func (t *TokenMetadata) SetUpdatedAtNil() {
	if t.touched == nil {
		t.touched = map[string]bool{}
	}
	t.touched["UpdatedAt"] = true
	t.UpdatedAt = nil
}
func (t TokenMetadata) MarshalJSON() ([]byte, error) {
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

	return json.Marshal(data)
}
