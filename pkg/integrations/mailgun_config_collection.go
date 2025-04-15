package integrations

import (
	"encoding/json"
	"github.com/magicbell/magicbell-go-project-client/pkg/shared"
)

type MailgunConfigCollection struct {
	Data  []MailgunConfig `json:"data,omitempty"`
	Links *shared.Links   `json:"links,omitempty"`
}

func (m *MailgunConfigCollection) GetData() []MailgunConfig {
	if m == nil {
		return nil
	}
	return m.Data
}

func (m *MailgunConfigCollection) SetData(data []MailgunConfig) {
	m.Data = data
}

func (m *MailgunConfigCollection) GetLinks() *shared.Links {
	if m == nil {
		return nil
	}
	return m.Links
}

func (m *MailgunConfigCollection) SetLinks(links shared.Links) {
	m.Links = &links
}

func (m MailgunConfigCollection) String() string {
	jsonData, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return "error converting struct: MailgunConfigCollection to string"
	}
	return string(jsonData)
}
