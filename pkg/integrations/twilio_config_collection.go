package integrations

import (
	"encoding/json"
	"github.com/magicbell/magicbell-go-project-client/pkg/shared"
)

type TwilioConfigCollection struct {
	Data  []TwilioConfig `json:"data,omitempty"`
	Links *shared.Links  `json:"links,omitempty"`
}

func (t *TwilioConfigCollection) GetData() []TwilioConfig {
	if t == nil {
		return nil
	}
	return t.Data
}

func (t *TwilioConfigCollection) SetData(data []TwilioConfig) {
	t.Data = data
}

func (t *TwilioConfigCollection) GetLinks() *shared.Links {
	if t == nil {
		return nil
	}
	return t.Links
}

func (t *TwilioConfigCollection) SetLinks(links shared.Links) {
	t.Links = &links
}

func (t TwilioConfigCollection) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TwilioConfigCollection to string"
	}
	return string(jsonData)
}
