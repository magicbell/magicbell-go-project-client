package integrations

import (
	"encoding/json"
)

type ExpoConfigObject struct {
	Config  *ExpoConfig `json:"config,omitempty" required:"true"`
	Id      *string     `json:"id,omitempty" required:"true"`
	Name    *string     `json:"name,omitempty" required:"true"`
	touched map[string]bool
}

func (e *ExpoConfigObject) GetConfig() *ExpoConfig {
	if e == nil {
		return nil
	}
	return e.Config
}

func (e *ExpoConfigObject) SetConfig(config ExpoConfig) {
	if e.touched == nil {
		e.touched = map[string]bool{}
	}
	e.touched["Config"] = true
	e.Config = &config
}

func (e *ExpoConfigObject) SetConfigNil() {
	if e.touched == nil {
		e.touched = map[string]bool{}
	}
	e.touched["Config"] = true
	e.Config = nil
}

func (e *ExpoConfigObject) GetId() *string {
	if e == nil {
		return nil
	}
	return e.Id
}

func (e *ExpoConfigObject) SetId(id string) {
	if e.touched == nil {
		e.touched = map[string]bool{}
	}
	e.touched["Id"] = true
	e.Id = &id
}

func (e *ExpoConfigObject) SetIdNil() {
	if e.touched == nil {
		e.touched = map[string]bool{}
	}
	e.touched["Id"] = true
	e.Id = nil
}

func (e *ExpoConfigObject) GetName() *string {
	if e == nil {
		return nil
	}
	return e.Name
}

func (e *ExpoConfigObject) SetName(name string) {
	if e.touched == nil {
		e.touched = map[string]bool{}
	}
	e.touched["Name"] = true
	e.Name = &name
}

func (e *ExpoConfigObject) SetNameNil() {
	if e.touched == nil {
		e.touched = map[string]bool{}
	}
	e.touched["Name"] = true
	e.Name = nil
}
func (e ExpoConfigObject) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if e.touched["Config"] && e.Config == nil {
		data["config"] = nil
	} else if e.Config != nil {
		data["config"] = e.Config
	}

	if e.touched["Id"] && e.Id == nil {
		data["id"] = nil
	} else if e.Id != nil {
		data["id"] = e.Id
	}

	if e.touched["Name"] && e.Name == nil {
		data["name"] = nil
	} else if e.Name != nil {
		data["name"] = e.Name
	}

	return json.Marshal(data)
}
