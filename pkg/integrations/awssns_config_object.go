package integrations

import (
	"encoding/json"
)

type AwssnsConfigObject struct {
	Config  *AwssnsConfig `json:"config,omitempty" required:"true"`
	Id      *string       `json:"id,omitempty" required:"true"`
	Name    *string       `json:"name,omitempty" required:"true"`
	touched map[string]bool
}

func (a *AwssnsConfigObject) GetConfig() *AwssnsConfig {
	if a == nil {
		return nil
	}
	return a.Config
}

func (a *AwssnsConfigObject) SetConfig(config AwssnsConfig) {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["Config"] = true
	a.Config = &config
}

func (a *AwssnsConfigObject) SetConfigNil() {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["Config"] = true
	a.Config = nil
}

func (a *AwssnsConfigObject) GetId() *string {
	if a == nil {
		return nil
	}
	return a.Id
}

func (a *AwssnsConfigObject) SetId(id string) {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["Id"] = true
	a.Id = &id
}

func (a *AwssnsConfigObject) SetIdNil() {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["Id"] = true
	a.Id = nil
}

func (a *AwssnsConfigObject) GetName() *string {
	if a == nil {
		return nil
	}
	return a.Name
}

func (a *AwssnsConfigObject) SetName(name string) {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["Name"] = true
	a.Name = &name
}

func (a *AwssnsConfigObject) SetNameNil() {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["Name"] = true
	a.Name = nil
}
func (a AwssnsConfigObject) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if a.touched["Config"] && a.Config == nil {
		data["config"] = nil
	} else if a.Config != nil {
		data["config"] = a.Config
	}

	if a.touched["Id"] && a.Id == nil {
		data["id"] = nil
	} else if a.Id != nil {
		data["id"] = a.Id
	}

	if a.touched["Name"] && a.Name == nil {
		data["name"] = nil
	} else if a.Name != nil {
		data["name"] = a.Name
	}

	return json.Marshal(data)
}
