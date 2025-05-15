package integrations

import (
	"encoding/json"
	"github.com/magicbell/magicbell-go/pkg/project-client/shared"
)

type GithubConfigCollection struct {
	Data  []GithubConfig `json:"data,omitempty"`
	Links *shared.Links  `json:"links,omitempty"`
}

func (g *GithubConfigCollection) GetData() []GithubConfig {
	if g == nil {
		return nil
	}
	return g.Data
}

func (g *GithubConfigCollection) SetData(data []GithubConfig) {
	g.Data = data
}

func (g *GithubConfigCollection) GetLinks() *shared.Links {
	if g == nil {
		return nil
	}
	return g.Links
}

func (g *GithubConfigCollection) SetLinks(links shared.Links) {
	g.Links = &links
}

func (g GithubConfigCollection) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GithubConfigCollection to string"
	}
	return string(jsonData)
}
