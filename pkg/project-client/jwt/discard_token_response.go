package jwt

import "encoding/json"

type DiscardTokenResponse struct {
	DiscardedAt *string `json:"discarded_at,omitempty" required:"true"`
	TokenId     *string `json:"token_id,omitempty" required:"true"`
}

func (d *DiscardTokenResponse) GetDiscardedAt() *string {
	if d == nil {
		return nil
	}
	return d.DiscardedAt
}

func (d *DiscardTokenResponse) SetDiscardedAt(discardedAt string) {
	d.DiscardedAt = &discardedAt
}

func (d *DiscardTokenResponse) GetTokenId() *string {
	if d == nil {
		return nil
	}
	return d.TokenId
}

func (d *DiscardTokenResponse) SetTokenId(tokenId string) {
	d.TokenId = &tokenId
}

func (d DiscardTokenResponse) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DiscardTokenResponse to string"
	}
	return string(jsonData)
}
