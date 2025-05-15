package integrations

import "encoding/json"

type IntegrationConfig struct {
	Config any     `json:"config,omitempty" required:"true"`
	Id     *string `json:"id,omitempty" required:"true"`
	Name   *string `json:"name,omitempty" required:"true"`
}

func (i *IntegrationConfig) GetConfig() any {
	if i == nil {
		return nil
	}
	return i.Config
}

func (i *IntegrationConfig) SetConfig(config any) {
	i.Config = &config
}

func (i *IntegrationConfig) GetId() *string {
	if i == nil {
		return nil
	}
	return i.Id
}

func (i *IntegrationConfig) SetId(id string) {
	i.Id = &id
}

func (i *IntegrationConfig) GetName() *string {
	if i == nil {
		return nil
	}
	return i.Name
}

func (i *IntegrationConfig) SetName(name string) {
	i.Name = &name
}

func (i IntegrationConfig) String() string {
	jsonData, err := json.MarshalIndent(i, "", "  ")
	if err != nil {
		return "error converting struct: IntegrationConfig to string"
	}
	return string(jsonData)
}
