package integrations

import (
	"encoding/json"
)

type InboxConfigObject struct {
	Config  *InboxConfig `json:"config,omitempty" required:"true"`
	Id      *string      `json:"id,omitempty" required:"true"`
	Name    *string      `json:"name,omitempty" required:"true"`
	touched map[string]bool
}

func (i *InboxConfigObject) GetConfig() *InboxConfig {
	if i == nil {
		return nil
	}
	return i.Config
}

func (i *InboxConfigObject) SetConfig(config InboxConfig) {
	if i.touched == nil {
		i.touched = map[string]bool{}
	}
	i.touched["Config"] = true
	i.Config = &config
}

func (i *InboxConfigObject) SetConfigNil() {
	if i.touched == nil {
		i.touched = map[string]bool{}
	}
	i.touched["Config"] = true
	i.Config = nil
}

func (i *InboxConfigObject) GetId() *string {
	if i == nil {
		return nil
	}
	return i.Id
}

func (i *InboxConfigObject) SetId(id string) {
	if i.touched == nil {
		i.touched = map[string]bool{}
	}
	i.touched["Id"] = true
	i.Id = &id
}

func (i *InboxConfigObject) SetIdNil() {
	if i.touched == nil {
		i.touched = map[string]bool{}
	}
	i.touched["Id"] = true
	i.Id = nil
}

func (i *InboxConfigObject) GetName() *string {
	if i == nil {
		return nil
	}
	return i.Name
}

func (i *InboxConfigObject) SetName(name string) {
	if i.touched == nil {
		i.touched = map[string]bool{}
	}
	i.touched["Name"] = true
	i.Name = &name
}

func (i *InboxConfigObject) SetNameNil() {
	if i.touched == nil {
		i.touched = map[string]bool{}
	}
	i.touched["Name"] = true
	i.Name = nil
}
func (i InboxConfigObject) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if i.touched["Config"] && i.Config == nil {
		data["config"] = nil
	} else if i.Config != nil {
		data["config"] = i.Config
	}

	if i.touched["Id"] && i.Id == nil {
		data["id"] = nil
	} else if i.Id != nil {
		data["id"] = i.Id
	}

	if i.touched["Name"] && i.Name == nil {
		data["name"] = nil
	} else if i.Name != nil {
		data["name"] = i.Name
	}

	return json.Marshal(data)
}
