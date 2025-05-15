package integrations

import (
	"encoding/json"
	"github.com/magicbell/magicbell-go/pkg/project-client/shared"
)

type FcmConfigCollection struct {
	Data  []FcmConfig   `json:"data,omitempty"`
	Links *shared.Links `json:"links,omitempty"`
}

func (f *FcmConfigCollection) GetData() []FcmConfig {
	if f == nil {
		return nil
	}
	return f.Data
}

func (f *FcmConfigCollection) SetData(data []FcmConfig) {
	f.Data = data
}

func (f *FcmConfigCollection) GetLinks() *shared.Links {
	if f == nil {
		return nil
	}
	return f.Links
}

func (f *FcmConfigCollection) SetLinks(links shared.Links) {
	f.Links = &links
}

func (f FcmConfigCollection) String() string {
	jsonData, err := json.MarshalIndent(f, "", "  ")
	if err != nil {
		return "error converting struct: FcmConfigCollection to string"
	}
	return string(jsonData)
}
