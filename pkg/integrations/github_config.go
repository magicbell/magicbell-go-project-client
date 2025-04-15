package integrations

import "encoding/json"

type GithubConfig struct {
	Config *GithubConfigPayload `json:"config,omitempty" required:"true"`
	Id     *string              `json:"id,omitempty" required:"true"`
	Name   *string              `json:"name,omitempty" required:"true"`
}

func (g *GithubConfig) GetConfig() *GithubConfigPayload {
	if g == nil {
		return nil
	}
	return g.Config
}

func (g *GithubConfig) SetConfig(config GithubConfigPayload) {
	g.Config = &config
}

func (g *GithubConfig) GetId() *string {
	if g == nil {
		return nil
	}
	return g.Id
}

func (g *GithubConfig) SetId(id string) {
	g.Id = &id
}

func (g *GithubConfig) GetName() *string {
	if g == nil {
		return nil
	}
	return g.Name
}

func (g *GithubConfig) SetName(name string) {
	g.Name = &name
}

func (g GithubConfig) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GithubConfig to string"
	}
	return string(jsonData)
}
