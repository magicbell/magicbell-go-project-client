package integrations

import (
	"encoding/json"
	"github.com/magicbell/magicbell-go/pkg/project-client/shared"
)

type AwssnsConfigCollection struct {
	Data  []AwssnsConfig `json:"data,omitempty"`
	Links *shared.Links  `json:"links,omitempty"`
}

func (a *AwssnsConfigCollection) GetData() []AwssnsConfig {
	if a == nil {
		return nil
	}
	return a.Data
}

func (a *AwssnsConfigCollection) SetData(data []AwssnsConfig) {
	a.Data = data
}

func (a *AwssnsConfigCollection) GetLinks() *shared.Links {
	if a == nil {
		return nil
	}
	return a.Links
}

func (a *AwssnsConfigCollection) SetLinks(links shared.Links) {
	a.Links = &links
}

func (a AwssnsConfigCollection) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AwssnsConfigCollection to string"
	}
	return string(jsonData)
}
