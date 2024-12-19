package integrations

import (
	"encoding/json"
)

type GithubConfigObject struct {
	Config  *GithubConfig `json:"config,omitempty" required:"true"`
	Id      *string       `json:"id,omitempty" required:"true"`
	Name    *string       `json:"name,omitempty" required:"true"`
	touched map[string]bool
}

func (g *GithubConfigObject) GetConfig() *GithubConfig {
	if g == nil {
		return nil
	}
	return g.Config
}

func (g *GithubConfigObject) SetConfig(config GithubConfig) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Config"] = true
	g.Config = &config
}

func (g *GithubConfigObject) SetConfigNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Config"] = true
	g.Config = nil
}

func (g *GithubConfigObject) GetId() *string {
	if g == nil {
		return nil
	}
	return g.Id
}

func (g *GithubConfigObject) SetId(id string) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Id"] = true
	g.Id = &id
}

func (g *GithubConfigObject) SetIdNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Id"] = true
	g.Id = nil
}

func (g *GithubConfigObject) GetName() *string {
	if g == nil {
		return nil
	}
	return g.Name
}

func (g *GithubConfigObject) SetName(name string) {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Name"] = true
	g.Name = &name
}

func (g *GithubConfigObject) SetNameNil() {
	if g.touched == nil {
		g.touched = map[string]bool{}
	}
	g.touched["Name"] = true
	g.Name = nil
}
func (g GithubConfigObject) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if g.touched["Config"] && g.Config == nil {
		data["config"] = nil
	} else if g.Config != nil {
		data["config"] = g.Config
	}

	if g.touched["Id"] && g.Id == nil {
		data["id"] = nil
	} else if g.Id != nil {
		data["id"] = g.Id
	}

	if g.touched["Name"] && g.Name == nil {
		data["name"] = nil
	} else if g.Name != nil {
		data["name"] = g.Name
	}

	return json.Marshal(data)
}
