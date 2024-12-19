package integrations

import (
	"encoding/json"
)

type PingConfigObject struct {
	Config  *PingConfig `json:"config,omitempty" required:"true"`
	Id      *string     `json:"id,omitempty" required:"true"`
	Name    *string     `json:"name,omitempty" required:"true"`
	touched map[string]bool
}

func (p *PingConfigObject) GetConfig() *PingConfig {
	if p == nil {
		return nil
	}
	return p.Config
}

func (p *PingConfigObject) SetConfig(config PingConfig) {
	if p.touched == nil {
		p.touched = map[string]bool{}
	}
	p.touched["Config"] = true
	p.Config = &config
}

func (p *PingConfigObject) SetConfigNil() {
	if p.touched == nil {
		p.touched = map[string]bool{}
	}
	p.touched["Config"] = true
	p.Config = nil
}

func (p *PingConfigObject) GetId() *string {
	if p == nil {
		return nil
	}
	return p.Id
}

func (p *PingConfigObject) SetId(id string) {
	if p.touched == nil {
		p.touched = map[string]bool{}
	}
	p.touched["Id"] = true
	p.Id = &id
}

func (p *PingConfigObject) SetIdNil() {
	if p.touched == nil {
		p.touched = map[string]bool{}
	}
	p.touched["Id"] = true
	p.Id = nil
}

func (p *PingConfigObject) GetName() *string {
	if p == nil {
		return nil
	}
	return p.Name
}

func (p *PingConfigObject) SetName(name string) {
	if p.touched == nil {
		p.touched = map[string]bool{}
	}
	p.touched["Name"] = true
	p.Name = &name
}

func (p *PingConfigObject) SetNameNil() {
	if p.touched == nil {
		p.touched = map[string]bool{}
	}
	p.touched["Name"] = true
	p.Name = nil
}
func (p PingConfigObject) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if p.touched["Config"] && p.Config == nil {
		data["config"] = nil
	} else if p.Config != nil {
		data["config"] = p.Config
	}

	if p.touched["Id"] && p.Id == nil {
		data["id"] = nil
	} else if p.Id != nil {
		data["id"] = p.Id
	}

	if p.touched["Name"] && p.Name == nil {
		data["name"] = nil
	} else if p.Name != nil {
		data["name"] = p.Name
	}

	return json.Marshal(data)
}
