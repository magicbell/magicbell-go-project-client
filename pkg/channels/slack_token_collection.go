package channels

import (
	"encoding/json"
	"github.com/magicbell/magicbell-go-project-client/pkg/shared"
)

type SlackTokenCollection struct {
	Data  []SlackToken  `json:"data,omitempty"`
	Links *shared.Links `json:"links,omitempty"`
}

func (s *SlackTokenCollection) GetData() []SlackToken {
	if s == nil {
		return nil
	}
	return s.Data
}

func (s *SlackTokenCollection) SetData(data []SlackToken) {
	s.Data = data
}

func (s *SlackTokenCollection) GetLinks() *shared.Links {
	if s == nil {
		return nil
	}
	return s.Links
}

func (s *SlackTokenCollection) SetLinks(links shared.Links) {
	s.Links = &links
}

func (s SlackTokenCollection) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SlackTokenCollection to string"
	}
	return string(jsonData)
}
