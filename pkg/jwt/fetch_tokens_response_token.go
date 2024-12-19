package jwt

import (
	"encoding/json"
)

type FetchTokensResponseToken struct {
	CreatedAt   *string `json:"created_at,omitempty" required:"true"`
	DiscardedAt *string `json:"discarded_at,omitempty"`
	ExpiresAt   *string `json:"expires_at,omitempty"`
	Id          *string `json:"id,omitempty"`
	Name        *string `json:"name,omitempty"`
	touched     map[string]bool
}

func (f *FetchTokensResponseToken) GetCreatedAt() *string {
	if f == nil {
		return nil
	}
	return f.CreatedAt
}

func (f *FetchTokensResponseToken) SetCreatedAt(createdAt string) {
	if f.touched == nil {
		f.touched = map[string]bool{}
	}
	f.touched["CreatedAt"] = true
	f.CreatedAt = &createdAt
}

func (f *FetchTokensResponseToken) SetCreatedAtNil() {
	if f.touched == nil {
		f.touched = map[string]bool{}
	}
	f.touched["CreatedAt"] = true
	f.CreatedAt = nil
}

func (f *FetchTokensResponseToken) GetDiscardedAt() *string {
	if f == nil {
		return nil
	}
	return f.DiscardedAt
}

func (f *FetchTokensResponseToken) SetDiscardedAt(discardedAt string) {
	if f.touched == nil {
		f.touched = map[string]bool{}
	}
	f.touched["DiscardedAt"] = true
	f.DiscardedAt = &discardedAt
}

func (f *FetchTokensResponseToken) SetDiscardedAtNil() {
	if f.touched == nil {
		f.touched = map[string]bool{}
	}
	f.touched["DiscardedAt"] = true
	f.DiscardedAt = nil
}

func (f *FetchTokensResponseToken) GetExpiresAt() *string {
	if f == nil {
		return nil
	}
	return f.ExpiresAt
}

func (f *FetchTokensResponseToken) SetExpiresAt(expiresAt string) {
	if f.touched == nil {
		f.touched = map[string]bool{}
	}
	f.touched["ExpiresAt"] = true
	f.ExpiresAt = &expiresAt
}

func (f *FetchTokensResponseToken) SetExpiresAtNil() {
	if f.touched == nil {
		f.touched = map[string]bool{}
	}
	f.touched["ExpiresAt"] = true
	f.ExpiresAt = nil
}

func (f *FetchTokensResponseToken) GetId() *string {
	if f == nil {
		return nil
	}
	return f.Id
}

func (f *FetchTokensResponseToken) SetId(id string) {
	if f.touched == nil {
		f.touched = map[string]bool{}
	}
	f.touched["Id"] = true
	f.Id = &id
}

func (f *FetchTokensResponseToken) SetIdNil() {
	if f.touched == nil {
		f.touched = map[string]bool{}
	}
	f.touched["Id"] = true
	f.Id = nil
}

func (f *FetchTokensResponseToken) GetName() *string {
	if f == nil {
		return nil
	}
	return f.Name
}

func (f *FetchTokensResponseToken) SetName(name string) {
	if f.touched == nil {
		f.touched = map[string]bool{}
	}
	f.touched["Name"] = true
	f.Name = &name
}

func (f *FetchTokensResponseToken) SetNameNil() {
	if f.touched == nil {
		f.touched = map[string]bool{}
	}
	f.touched["Name"] = true
	f.Name = nil
}
func (f FetchTokensResponseToken) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if f.touched["CreatedAt"] && f.CreatedAt == nil {
		data["created_at"] = nil
	} else if f.CreatedAt != nil {
		data["created_at"] = f.CreatedAt
	}

	if f.touched["DiscardedAt"] && f.DiscardedAt == nil {
		data["discarded_at"] = nil
	} else if f.DiscardedAt != nil {
		data["discarded_at"] = f.DiscardedAt
	}

	if f.touched["ExpiresAt"] && f.ExpiresAt == nil {
		data["expires_at"] = nil
	} else if f.ExpiresAt != nil {
		data["expires_at"] = f.ExpiresAt
	}

	if f.touched["Id"] && f.Id == nil {
		data["id"] = nil
	} else if f.Id != nil {
		data["id"] = f.Id
	}

	if f.touched["Name"] && f.Name == nil {
		data["name"] = nil
	} else if f.Name != nil {
		data["name"] = f.Name
	}

	return json.Marshal(data)
}
