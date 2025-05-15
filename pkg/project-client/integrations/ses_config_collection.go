package integrations

import (
	"encoding/json"
	"github.com/magicbell/magicbell-go/pkg/project-client/shared"
)

type SesConfigCollection struct {
	Data  []SesConfig   `json:"data,omitempty"`
	Links *shared.Links `json:"links,omitempty"`
}

func (s *SesConfigCollection) GetData() []SesConfig {
	if s == nil {
		return nil
	}
	return s.Data
}

func (s *SesConfigCollection) SetData(data []SesConfig) {
	s.Data = data
}

func (s *SesConfigCollection) GetLinks() *shared.Links {
	if s == nil {
		return nil
	}
	return s.Links
}

func (s *SesConfigCollection) SetLinks(links shared.Links) {
	s.Links = &links
}

func (s SesConfigCollection) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SesConfigCollection to string"
	}
	return string(jsonData)
}
