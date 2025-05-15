package integrations

import (
	"encoding/json"
	"github.com/magicbell/magicbell-go/pkg/project-client/shared"
)

type SendgridConfigCollection struct {
	Data  []SendgridConfig `json:"data,omitempty"`
	Links *shared.Links    `json:"links,omitempty"`
}

func (s *SendgridConfigCollection) GetData() []SendgridConfig {
	if s == nil {
		return nil
	}
	return s.Data
}

func (s *SendgridConfigCollection) SetData(data []SendgridConfig) {
	s.Data = data
}

func (s *SendgridConfigCollection) GetLinks() *shared.Links {
	if s == nil {
		return nil
	}
	return s.Links
}

func (s *SendgridConfigCollection) SetLinks(links shared.Links) {
	s.Links = &links
}

func (s SendgridConfigCollection) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SendgridConfigCollection to string"
	}
	return string(jsonData)
}
