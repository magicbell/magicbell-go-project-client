package integrations

import (
	"encoding/json"
	"github.com/magicbell/magicbell-go/pkg/project-client/shared"
)

type StripeConfigCollection struct {
	Data  []StripeConfig `json:"data,omitempty"`
	Links *shared.Links  `json:"links,omitempty"`
}

func (s *StripeConfigCollection) GetData() []StripeConfig {
	if s == nil {
		return nil
	}
	return s.Data
}

func (s *StripeConfigCollection) SetData(data []StripeConfig) {
	s.Data = data
}

func (s *StripeConfigCollection) GetLinks() *shared.Links {
	if s == nil {
		return nil
	}
	return s.Links
}

func (s *StripeConfigCollection) SetLinks(links shared.Links) {
	s.Links = &links
}

func (s StripeConfigCollection) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: StripeConfigCollection to string"
	}
	return string(jsonData)
}
