package integrations

import (
	"encoding/json"
	"github.com/magicbell/magicbell-go-project-client/pkg/shared"
)

type IntegrationConfigCollection struct {
	Data  []IntegrationConfig `json:"data,omitempty"`
	Links *shared.Links       `json:"links,omitempty"`
}

func (i *IntegrationConfigCollection) GetData() []IntegrationConfig {
	if i == nil {
		return nil
	}
	return i.Data
}

func (i *IntegrationConfigCollection) SetData(data []IntegrationConfig) {
	i.Data = data
}

func (i *IntegrationConfigCollection) GetLinks() *shared.Links {
	if i == nil {
		return nil
	}
	return i.Links
}

func (i *IntegrationConfigCollection) SetLinks(links shared.Links) {
	i.Links = &links
}

func (i IntegrationConfigCollection) String() string {
	jsonData, err := json.MarshalIndent(i, "", "  ")
	if err != nil {
		return "error converting struct: IntegrationConfigCollection to string"
	}
	return string(jsonData)
}
