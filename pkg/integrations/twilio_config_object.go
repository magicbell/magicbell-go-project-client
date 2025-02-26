package integrations

import (
	"encoding/json"
)

type TwilioConfigObject struct {
	Config  *TwilioConfig `json:"config,omitempty" required:"true"`
	Id      *string       `json:"id,omitempty" required:"true"`
	Name    *string       `json:"name,omitempty" required:"true"`
	touched map[string]bool
}

func (t *TwilioConfigObject) GetConfig() *TwilioConfig {
	if t == nil {
		return nil
	}
	return t.Config
}

func (t *TwilioConfigObject) SetConfig(config TwilioConfig) {
	if t.touched == nil {
		t.touched = map[string]bool{}
	}
	t.touched["Config"] = true
	t.Config = &config
}

func (t *TwilioConfigObject) SetConfigNil() {
	if t.touched == nil {
		t.touched = map[string]bool{}
	}
	t.touched["Config"] = true
	t.Config = nil
}

func (t *TwilioConfigObject) GetId() *string {
	if t == nil {
		return nil
	}
	return t.Id
}

func (t *TwilioConfigObject) SetId(id string) {
	if t.touched == nil {
		t.touched = map[string]bool{}
	}
	t.touched["Id"] = true
	t.Id = &id
}

func (t *TwilioConfigObject) SetIdNil() {
	if t.touched == nil {
		t.touched = map[string]bool{}
	}
	t.touched["Id"] = true
	t.Id = nil
}

func (t *TwilioConfigObject) GetName() *string {
	if t == nil {
		return nil
	}
	return t.Name
}

func (t *TwilioConfigObject) SetName(name string) {
	if t.touched == nil {
		t.touched = map[string]bool{}
	}
	t.touched["Name"] = true
	t.Name = &name
}

func (t *TwilioConfigObject) SetNameNil() {
	if t.touched == nil {
		t.touched = map[string]bool{}
	}
	t.touched["Name"] = true
	t.Name = nil
}
func (t TwilioConfigObject) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if t.touched["Config"] && t.Config == nil {
		data["config"] = nil
	} else if t.Config != nil {
		data["config"] = t.Config
	}

	if t.touched["Id"] && t.Id == nil {
		data["id"] = nil
	} else if t.Id != nil {
		data["id"] = t.Id
	}

	if t.touched["Name"] && t.Name == nil {
		data["name"] = nil
	} else if t.Name != nil {
		data["name"] = t.Name
	}

	return json.Marshal(data)
}
