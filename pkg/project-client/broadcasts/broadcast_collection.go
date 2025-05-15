package broadcasts

import (
	"encoding/json"
	"github.com/magicbell/magicbell-go/pkg/project-client/shared"
)

type BroadcastCollection struct {
	Data  []Broadcast   `json:"data,omitempty"`
	Links *shared.Links `json:"links,omitempty"`
}

func (b *BroadcastCollection) GetData() []Broadcast {
	if b == nil {
		return nil
	}
	return b.Data
}

func (b *BroadcastCollection) SetData(data []Broadcast) {
	b.Data = data
}

func (b *BroadcastCollection) GetLinks() *shared.Links {
	if b == nil {
		return nil
	}
	return b.Links
}

func (b *BroadcastCollection) SetLinks(links shared.Links) {
	b.Links = &links
}

func (b BroadcastCollection) String() string {
	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "error converting struct: BroadcastCollection to string"
	}
	return string(jsonData)
}
