package integrations

import (
	"encoding/json"
	"github.com/magicbell/magicbell-go-project-client/pkg/shared"
)

type ExpoConfigCollection struct {
	Data  []ExpoConfig  `json:"data,omitempty"`
	Links *shared.Links `json:"links,omitempty"`
}

func (e *ExpoConfigCollection) GetData() []ExpoConfig {
	if e == nil {
		return nil
	}
	return e.Data
}

func (e *ExpoConfigCollection) SetData(data []ExpoConfig) {
	e.Data = data
}

func (e *ExpoConfigCollection) GetLinks() *shared.Links {
	if e == nil {
		return nil
	}
	return e.Links
}

func (e *ExpoConfigCollection) SetLinks(links shared.Links) {
	e.Links = &links
}

func (e ExpoConfigCollection) String() string {
	jsonData, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		return "error converting struct: ExpoConfigCollection to string"
	}
	return string(jsonData)
}
