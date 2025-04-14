package integrations

import (
	"encoding/json"
	"github.com/magicbell/magicbell-go-project-client/pkg/shared"
)

type ApnsConfigCollection struct {
	Data  []ApnsConfig  `json:"data,omitempty"`
	Links *shared.Links `json:"links,omitempty"`
}

func (a *ApnsConfigCollection) GetData() []ApnsConfig {
	if a == nil {
		return nil
	}
	return a.Data
}

func (a *ApnsConfigCollection) SetData(data []ApnsConfig) {
	a.Data = data
}

func (a *ApnsConfigCollection) GetLinks() *shared.Links {
	if a == nil {
		return nil
	}
	return a.Links
}

func (a *ApnsConfigCollection) SetLinks(links shared.Links) {
	a.Links = &links
}

func (a ApnsConfigCollection) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: ApnsConfigCollection to string"
	}
	return string(jsonData)
}
