package integrations

import (
	"encoding/json"
	"github.com/magicbell/magicbell-go-project-client/pkg/shared"
)

type InboxConfigCollection struct {
	Data  []InboxConfig `json:"data,omitempty"`
	Links *shared.Links `json:"links,omitempty"`
}

func (i *InboxConfigCollection) GetData() []InboxConfig {
	if i == nil {
		return nil
	}
	return i.Data
}

func (i *InboxConfigCollection) SetData(data []InboxConfig) {
	i.Data = data
}

func (i *InboxConfigCollection) GetLinks() *shared.Links {
	if i == nil {
		return nil
	}
	return i.Links
}

func (i *InboxConfigCollection) SetLinks(links shared.Links) {
	i.Links = &links
}

func (i InboxConfigCollection) String() string {
	jsonData, err := json.MarshalIndent(i, "", "  ")
	if err != nil {
		return "error converting struct: InboxConfigCollection to string"
	}
	return string(jsonData)
}
