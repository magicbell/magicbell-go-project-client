package integrations

import (
	"encoding/json"
)

type FcmConfigObject struct {
	Config  *FcmConfig `json:"config,omitempty" required:"true"`
	Id      *string    `json:"id,omitempty" required:"true"`
	Name    *string    `json:"name,omitempty" required:"true"`
	touched map[string]bool
}

func (f *FcmConfigObject) GetConfig() *FcmConfig {
	if f == nil {
		return nil
	}
	return f.Config
}

func (f *FcmConfigObject) SetConfig(config FcmConfig) {
	if f.touched == nil {
		f.touched = map[string]bool{}
	}
	f.touched["Config"] = true
	f.Config = &config
}

func (f *FcmConfigObject) SetConfigNil() {
	if f.touched == nil {
		f.touched = map[string]bool{}
	}
	f.touched["Config"] = true
	f.Config = nil
}

func (f *FcmConfigObject) GetId() *string {
	if f == nil {
		return nil
	}
	return f.Id
}

func (f *FcmConfigObject) SetId(id string) {
	if f.touched == nil {
		f.touched = map[string]bool{}
	}
	f.touched["Id"] = true
	f.Id = &id
}

func (f *FcmConfigObject) SetIdNil() {
	if f.touched == nil {
		f.touched = map[string]bool{}
	}
	f.touched["Id"] = true
	f.Id = nil
}

func (f *FcmConfigObject) GetName() *string {
	if f == nil {
		return nil
	}
	return f.Name
}

func (f *FcmConfigObject) SetName(name string) {
	if f.touched == nil {
		f.touched = map[string]bool{}
	}
	f.touched["Name"] = true
	f.Name = &name
}

func (f *FcmConfigObject) SetNameNil() {
	if f.touched == nil {
		f.touched = map[string]bool{}
	}
	f.touched["Name"] = true
	f.Name = nil
}
func (f FcmConfigObject) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if f.touched["Config"] && f.Config == nil {
		data["config"] = nil
	} else if f.Config != nil {
		data["config"] = f.Config
	}

	if f.touched["Id"] && f.Id == nil {
		data["id"] = nil
	} else if f.Id != nil {
		data["id"] = f.Id
	}

	if f.touched["Name"] && f.Name == nil {
		data["name"] = nil
	} else if f.Name != nil {
		data["name"] = f.Name
	}

	return json.Marshal(data)
}
