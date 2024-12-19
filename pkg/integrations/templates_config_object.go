package integrations

import (
	"encoding/json"
)

type TemplatesConfigObject struct {
	Config  any     `json:"config,omitempty" required:"true"`
	Id      *string `json:"id,omitempty" required:"true"`
	Name    *string `json:"name,omitempty" required:"true"`
	touched map[string]bool
}

func (t *TemplatesConfigObject) GetConfig() any {
	if t == nil {
		return nil
	}
	return t.Config
}

func (t *TemplatesConfigObject) SetConfig(config any) {
	if t.touched == nil {
		t.touched = map[string]bool{}
	}
	t.touched["Config"] = true
	t.Config = config
}

func (t *TemplatesConfigObject) SetConfigNil() {
	if t.touched == nil {
		t.touched = map[string]bool{}
	}
	t.touched["Config"] = true
	t.Config = nil
}

func (t *TemplatesConfigObject) GetId() *string {
	if t == nil {
		return nil
	}
	return t.Id
}

func (t *TemplatesConfigObject) SetId(id string) {
	if t.touched == nil {
		t.touched = map[string]bool{}
	}
	t.touched["Id"] = true
	t.Id = &id
}

func (t *TemplatesConfigObject) SetIdNil() {
	if t.touched == nil {
		t.touched = map[string]bool{}
	}
	t.touched["Id"] = true
	t.Id = nil
}

func (t *TemplatesConfigObject) GetName() *string {
	if t == nil {
		return nil
	}
	return t.Name
}

func (t *TemplatesConfigObject) SetName(name string) {
	if t.touched == nil {
		t.touched = map[string]bool{}
	}
	t.touched["Name"] = true
	t.Name = &name
}

func (t *TemplatesConfigObject) SetNameNil() {
	if t.touched == nil {
		t.touched = map[string]bool{}
	}
	t.touched["Name"] = true
	t.Name = nil
}
func (t TemplatesConfigObject) MarshalJSON() ([]byte, error) {
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
