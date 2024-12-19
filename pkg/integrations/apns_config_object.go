package integrations

import (
	"encoding/json"
)

type ApnsConfigObject struct {
	Config  *ApnsConfig `json:"config,omitempty" required:"true"`
	Id      *string     `json:"id,omitempty" required:"true"`
	Name    *string     `json:"name,omitempty" required:"true"`
	touched map[string]bool
}

func (a *ApnsConfigObject) GetConfig() *ApnsConfig {
	if a == nil {
		return nil
	}
	return a.Config
}

func (a *ApnsConfigObject) SetConfig(config ApnsConfig) {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["Config"] = true
	a.Config = &config
}

func (a *ApnsConfigObject) SetConfigNil() {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["Config"] = true
	a.Config = nil
}

func (a *ApnsConfigObject) GetId() *string {
	if a == nil {
		return nil
	}
	return a.Id
}

func (a *ApnsConfigObject) SetId(id string) {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["Id"] = true
	a.Id = &id
}

func (a *ApnsConfigObject) SetIdNil() {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["Id"] = true
	a.Id = nil
}

func (a *ApnsConfigObject) GetName() *string {
	if a == nil {
		return nil
	}
	return a.Name
}

func (a *ApnsConfigObject) SetName(name string) {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["Name"] = true
	a.Name = &name
}

func (a *ApnsConfigObject) SetNameNil() {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["Name"] = true
	a.Name = nil
}
func (a ApnsConfigObject) MarshalJSON() ([]byte, error) {
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
