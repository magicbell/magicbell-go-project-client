package integrations

import "encoding/json"

type ApnsConfig struct {
	Config *ApnsConfigPayload `json:"config,omitempty" required:"true"`
	Id     *string            `json:"id,omitempty" required:"true"`
	Name   *string            `json:"name,omitempty" required:"true"`
}

func (a *ApnsConfig) GetConfig() *ApnsConfigPayload {
	if a == nil {
		return nil
	}
	return a.Config
}

func (a *ApnsConfig) SetConfig(config ApnsConfigPayload) {
	a.Config = &config
}

func (a *ApnsConfig) GetId() *string {
	if a == nil {
		return nil
	}
	return a.Id
}

func (a *ApnsConfig) SetId(id string) {
	a.Id = &id
}

func (a *ApnsConfig) GetName() *string {
	if a == nil {
		return nil
	}
	return a.Name
}

func (a *ApnsConfig) SetName(name string) {
	a.Name = &name
}

func (a ApnsConfig) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: ApnsConfig to string"
	}
	return string(jsonData)
}
