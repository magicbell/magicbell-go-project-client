package integrations

import "encoding/json"

type ExpoConfig struct {
	Config *ExpoConfigPayload `json:"config,omitempty" required:"true"`
	Id     *string            `json:"id,omitempty" required:"true"`
	Name   *string            `json:"name,omitempty" required:"true"`
}

func (e *ExpoConfig) GetConfig() *ExpoConfigPayload {
	if e == nil {
		return nil
	}
	return e.Config
}

func (e *ExpoConfig) SetConfig(config ExpoConfigPayload) {
	e.Config = &config
}

func (e *ExpoConfig) GetId() *string {
	if e == nil {
		return nil
	}
	return e.Id
}

func (e *ExpoConfig) SetId(id string) {
	e.Id = &id
}

func (e *ExpoConfig) GetName() *string {
	if e == nil {
		return nil
	}
	return e.Name
}

func (e *ExpoConfig) SetName(name string) {
	e.Name = &name
}

func (e ExpoConfig) String() string {
	jsonData, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		return "error converting struct: ExpoConfig to string"
	}
	return string(jsonData)
}
