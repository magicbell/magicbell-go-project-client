package integrations

import (
	"encoding/json"
	"github.com/magicbell/magicbell-go/pkg/project-client/shared"
)

type TemplatesConfigCollection struct {
	Data  []TemplatesConfig `json:"data,omitempty"`
	Links *shared.Links     `json:"links,omitempty"`
}

func (t *TemplatesConfigCollection) GetData() []TemplatesConfig {
	if t == nil {
		return nil
	}
	return t.Data
}

func (t *TemplatesConfigCollection) SetData(data []TemplatesConfig) {
	t.Data = data
}

func (t *TemplatesConfigCollection) GetLinks() *shared.Links {
	if t == nil {
		return nil
	}
	return t.Links
}

func (t *TemplatesConfigCollection) SetLinks(links shared.Links) {
	t.Links = &links
}

func (t TemplatesConfigCollection) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TemplatesConfigCollection to string"
	}
	return string(jsonData)
}
