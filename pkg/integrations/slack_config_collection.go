package integrations

import (
	"encoding/json"
	"github.com/magicbell/magicbell-go-project-client/pkg/shared"
)

type SlackConfigCollection struct {
	Data  []SlackConfig `json:"data,omitempty"`
	Links *shared.Links `json:"links,omitempty"`
}

func (s *SlackConfigCollection) GetData() []SlackConfig {
	if s == nil {
		return nil
	}
	return s.Data
}

func (s *SlackConfigCollection) SetData(data []SlackConfig) {
	s.Data = data
}

func (s *SlackConfigCollection) GetLinks() *shared.Links {
	if s == nil {
		return nil
	}
	return s.Links
}

func (s *SlackConfigCollection) SetLinks(links shared.Links) {
	s.Links = &links
}

func (s SlackConfigCollection) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SlackConfigCollection to string"
	}
	return string(jsonData)
}
