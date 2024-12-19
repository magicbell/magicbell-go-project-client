package channels

import (
	"encoding/json"
)

type DiscardResult struct {
	DiscardedAt *string `json:"discarded_at,omitempty"`
	Id          *string `json:"id,omitempty"`
	touched     map[string]bool
}

func (d *DiscardResult) GetDiscardedAt() *string {
	if d == nil {
		return nil
	}
	return d.DiscardedAt
}

func (d *DiscardResult) SetDiscardedAt(discardedAt string) {
	if d.touched == nil {
		d.touched = map[string]bool{}
	}
	d.touched["DiscardedAt"] = true
	d.DiscardedAt = &discardedAt
}

func (d *DiscardResult) SetDiscardedAtNil() {
	if d.touched == nil {
		d.touched = map[string]bool{}
	}
	d.touched["DiscardedAt"] = true
	d.DiscardedAt = nil
}

func (d *DiscardResult) GetId() *string {
	if d == nil {
		return nil
	}
	return d.Id
}

func (d *DiscardResult) SetId(id string) {
	if d.touched == nil {
		d.touched = map[string]bool{}
	}
	d.touched["Id"] = true
	d.Id = &id
}

func (d *DiscardResult) SetIdNil() {
	if d.touched == nil {
		d.touched = map[string]bool{}
	}
	d.touched["Id"] = true
	d.Id = nil
}
func (d DiscardResult) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if d.touched["DiscardedAt"] && d.DiscardedAt == nil {
		data["discarded_at"] = nil
	} else if d.DiscardedAt != nil {
		data["discarded_at"] = d.DiscardedAt
	}

	if d.touched["Id"] && d.Id == nil {
		data["id"] = nil
	} else if d.Id != nil {
		data["id"] = d.Id
	}

	return json.Marshal(data)
}
