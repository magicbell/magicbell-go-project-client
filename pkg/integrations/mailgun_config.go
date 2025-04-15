package integrations

import "encoding/json"

type MailgunConfig struct {
	Config *MailgunConfigPayload `json:"config,omitempty" required:"true"`
	Id     *string               `json:"id,omitempty" required:"true"`
	Name   *string               `json:"name,omitempty" required:"true"`
}

func (m *MailgunConfig) GetConfig() *MailgunConfigPayload {
	if m == nil {
		return nil
	}
	return m.Config
}

func (m *MailgunConfig) SetConfig(config MailgunConfigPayload) {
	m.Config = &config
}

func (m *MailgunConfig) GetId() *string {
	if m == nil {
		return nil
	}
	return m.Id
}

func (m *MailgunConfig) SetId(id string) {
	m.Id = &id
}

func (m *MailgunConfig) GetName() *string {
	if m == nil {
		return nil
	}
	return m.Name
}

func (m *MailgunConfig) SetName(name string) {
	m.Name = &name
}

func (m MailgunConfig) String() string {
	jsonData, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return "error converting struct: MailgunConfig to string"
	}
	return string(jsonData)
}
