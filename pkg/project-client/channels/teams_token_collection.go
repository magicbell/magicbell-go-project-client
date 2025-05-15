package channels

import (
	"encoding/json"
	"github.com/magicbell/magicbell-go/pkg/project-client/shared"
)

type TeamsTokenCollection struct {
	Data  []TeamsToken  `json:"data,omitempty"`
	Links *shared.Links `json:"links,omitempty"`
}

func (t *TeamsTokenCollection) GetData() []TeamsToken {
	if t == nil {
		return nil
	}
	return t.Data
}

func (t *TeamsTokenCollection) SetData(data []TeamsToken) {
	t.Data = data
}

func (t *TeamsTokenCollection) GetLinks() *shared.Links {
	if t == nil {
		return nil
	}
	return t.Links
}

func (t *TeamsTokenCollection) SetLinks(links shared.Links) {
	t.Links = &links
}

func (t TeamsTokenCollection) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TeamsTokenCollection to string"
	}
	return string(jsonData)
}
