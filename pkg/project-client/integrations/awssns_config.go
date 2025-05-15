package integrations

import "encoding/json"

type AwssnsConfig struct {
	Config *AwssnsConfigPayload `json:"config,omitempty" required:"true"`
	Id     *string              `json:"id,omitempty" required:"true"`
	Name   *string              `json:"name,omitempty" required:"true"`
}

func (a *AwssnsConfig) GetConfig() *AwssnsConfigPayload {
	if a == nil {
		return nil
	}
	return a.Config
}

func (a *AwssnsConfig) SetConfig(config AwssnsConfigPayload) {
	a.Config = &config
}

func (a *AwssnsConfig) GetId() *string {
	if a == nil {
		return nil
	}
	return a.Id
}

func (a *AwssnsConfig) SetId(id string) {
	a.Id = &id
}

func (a *AwssnsConfig) GetName() *string {
	if a == nil {
		return nil
	}
	return a.Name
}

func (a *AwssnsConfig) SetName(name string) {
	a.Name = &name
}

func (a AwssnsConfig) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AwssnsConfig to string"
	}
	return string(jsonData)
}
