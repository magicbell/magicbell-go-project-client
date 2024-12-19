package jwt

import (
	"encoding/json"
)

type DiscardTokenResponse struct {
	DiscardedAt *string `json:"discarded_at,omitempty" required:"true"`
	TokenId     *string `json:"token_id,omitempty" required:"true"`
	touched     map[string]bool
}

func (d *DiscardTokenResponse) GetDiscardedAt() *string {
	if d == nil {
		return nil
	}
	return d.DiscardedAt
}

func (d *DiscardTokenResponse) SetDiscardedAt(discardedAt string) {
	if d.touched == nil {
		d.touched = map[string]bool{}
	}
	d.touched["DiscardedAt"] = true
	d.DiscardedAt = &discardedAt
}

func (d *DiscardTokenResponse) SetDiscardedAtNil() {
	if d.touched == nil {
		d.touched = map[string]bool{}
	}
	d.touched["DiscardedAt"] = true
	d.DiscardedAt = nil
}

func (d *DiscardTokenResponse) GetTokenId() *string {
	if d == nil {
		return nil
	}
	return d.TokenId
}

func (d *DiscardTokenResponse) SetTokenId(tokenId string) {
	if d.touched == nil {
		d.touched = map[string]bool{}
	}
	d.touched["TokenId"] = true
	d.TokenId = &tokenId
}

func (d *DiscardTokenResponse) SetTokenIdNil() {
	if d.touched == nil {
		d.touched = map[string]bool{}
	}
	d.touched["TokenId"] = true
	d.TokenId = nil
}
func (d DiscardTokenResponse) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if d.touched["DiscardedAt"] && d.DiscardedAt == nil {
		data["discarded_at"] = nil
	} else if d.DiscardedAt != nil {
		data["discarded_at"] = d.DiscardedAt
	}

	if d.touched["TokenId"] && d.TokenId == nil {
		data["token_id"] = nil
	} else if d.TokenId != nil {
		data["token_id"] = d.TokenId
	}

	return json.Marshal(data)
}
