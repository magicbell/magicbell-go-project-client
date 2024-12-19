package integrations

import (
	"encoding/json"
)

type MailgunConfigObject struct {
	Config  *MailgunConfig `json:"config,omitempty" required:"true"`
	Id      *string        `json:"id,omitempty" required:"true"`
	Name    *string        `json:"name,omitempty" required:"true"`
	touched map[string]bool
}

func (m *MailgunConfigObject) GetConfig() *MailgunConfig {
	if m == nil {
		return nil
	}
	return m.Config
}

func (m *MailgunConfigObject) SetConfig(config MailgunConfig) {
	if m.touched == nil {
		m.touched = map[string]bool{}
	}
	m.touched["Config"] = true
	m.Config = &config
}

func (m *MailgunConfigObject) SetConfigNil() {
	if m.touched == nil {
		m.touched = map[string]bool{}
	}
	m.touched["Config"] = true
	m.Config = nil
}

func (m *MailgunConfigObject) GetId() *string {
	if m == nil {
		return nil
	}
	return m.Id
}

func (m *MailgunConfigObject) SetId(id string) {
	if m.touched == nil {
		m.touched = map[string]bool{}
	}
	m.touched["Id"] = true
	m.Id = &id
}

func (m *MailgunConfigObject) SetIdNil() {
	if m.touched == nil {
		m.touched = map[string]bool{}
	}
	m.touched["Id"] = true
	m.Id = nil
}

func (m *MailgunConfigObject) GetName() *string {
	if m == nil {
		return nil
	}
	return m.Name
}

func (m *MailgunConfigObject) SetName(name string) {
	if m.touched == nil {
		m.touched = map[string]bool{}
	}
	m.touched["Name"] = true
	m.Name = &name
}

func (m *MailgunConfigObject) SetNameNil() {
	if m.touched == nil {
		m.touched = map[string]bool{}
	}
	m.touched["Name"] = true
	m.Name = nil
}
func (m MailgunConfigObject) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if m.touched["Config"] && m.Config == nil {
		data["config"] = nil
	} else if m.Config != nil {
		data["config"] = m.Config
	}

	if m.touched["Id"] && m.Id == nil {
		data["id"] = nil
	} else if m.Id != nil {
		data["id"] = m.Id
	}

	if m.touched["Name"] && m.Name == nil {
		data["name"] = nil
	} else if m.Name != nil {
		data["name"] = m.Name
	}

	return json.Marshal(data)
}
