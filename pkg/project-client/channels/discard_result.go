package channels

import "encoding/json"

type DiscardResult struct {
	DiscardedAt *string `json:"discarded_at,omitempty"`
	Id          *string `json:"id,omitempty"`
}

func (d *DiscardResult) GetDiscardedAt() *string {
	if d == nil {
		return nil
	}
	return d.DiscardedAt
}

func (d *DiscardResult) SetDiscardedAt(discardedAt string) {
	d.DiscardedAt = &discardedAt
}

func (d *DiscardResult) GetId() *string {
	if d == nil {
		return nil
	}
	return d.Id
}

func (d *DiscardResult) SetId(id string) {
	d.Id = &id
}

func (d DiscardResult) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DiscardResult to string"
	}
	return string(jsonData)
}
