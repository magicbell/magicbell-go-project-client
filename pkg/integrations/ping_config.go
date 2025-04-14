package integrations

import "encoding/json"

type PingConfig struct {
	Config *PingConfigPayload `json:"config,omitempty" required:"true"`
	Id     *string            `json:"id,omitempty" required:"true"`
	Name   *string            `json:"name,omitempty" required:"true"`
}

func (p *PingConfig) GetConfig() *PingConfigPayload {
	if p == nil {
		return nil
	}
	return p.Config
}

func (p *PingConfig) SetConfig(config PingConfigPayload) {
	p.Config = &config
}

func (p *PingConfig) GetId() *string {
	if p == nil {
		return nil
	}
	return p.Id
}

func (p *PingConfig) SetId(id string) {
	p.Id = &id
}

func (p *PingConfig) GetName() *string {
	if p == nil {
		return nil
	}
	return p.Name
}

func (p *PingConfig) SetName(name string) {
	p.Name = &name
}

func (p PingConfig) String() string {
	jsonData, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return "error converting struct: PingConfig to string"
	}
	return string(jsonData)
}
