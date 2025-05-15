package integrations

import (
	"encoding/json"
	"github.com/magicbell/magicbell-go/pkg/project-client/shared"
)

type PingConfigCollection struct {
	Data  []PingConfig  `json:"data,omitempty"`
	Links *shared.Links `json:"links,omitempty"`
}

func (p *PingConfigCollection) GetData() []PingConfig {
	if p == nil {
		return nil
	}
	return p.Data
}

func (p *PingConfigCollection) SetData(data []PingConfig) {
	p.Data = data
}

func (p *PingConfigCollection) GetLinks() *shared.Links {
	if p == nil {
		return nil
	}
	return p.Links
}

func (p *PingConfigCollection) SetLinks(links shared.Links) {
	p.Links = &links
}

func (p PingConfigCollection) String() string {
	jsonData, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return "error converting struct: PingConfigCollection to string"
	}
	return string(jsonData)
}
