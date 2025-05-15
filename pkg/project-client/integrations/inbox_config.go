package integrations

import "encoding/json"

type InboxConfig struct {
	Config *InboxConfigPayload `json:"config,omitempty" required:"true"`
	Id     *string             `json:"id,omitempty" required:"true"`
	Name   *string             `json:"name,omitempty" required:"true"`
}

func (i *InboxConfig) GetConfig() *InboxConfigPayload {
	if i == nil {
		return nil
	}
	return i.Config
}

func (i *InboxConfig) SetConfig(config InboxConfigPayload) {
	i.Config = &config
}

func (i *InboxConfig) GetId() *string {
	if i == nil {
		return nil
	}
	return i.Id
}

func (i *InboxConfig) SetId(id string) {
	i.Id = &id
}

func (i *InboxConfig) GetName() *string {
	if i == nil {
		return nil
	}
	return i.Name
}

func (i *InboxConfig) SetName(name string) {
	i.Name = &name
}

func (i InboxConfig) String() string {
	jsonData, err := json.MarshalIndent(i, "", "  ")
	if err != nil {
		return "error converting struct: InboxConfig to string"
	}
	return string(jsonData)
}
