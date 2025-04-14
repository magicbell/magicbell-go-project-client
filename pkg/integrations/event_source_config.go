package integrations

import "encoding/json"

type EventSourceConfig struct {
	Config *EventSourceConfigPayload `json:"config,omitempty" required:"true"`
	Id     *string                   `json:"id,omitempty" required:"true"`
	Name   *string                   `json:"name,omitempty" required:"true"`
}

func (e *EventSourceConfig) GetConfig() *EventSourceConfigPayload {
	if e == nil {
		return nil
	}
	return e.Config
}

func (e *EventSourceConfig) SetConfig(config EventSourceConfigPayload) {
	e.Config = &config
}

func (e *EventSourceConfig) GetId() *string {
	if e == nil {
		return nil
	}
	return e.Id
}

func (e *EventSourceConfig) SetId(id string) {
	e.Id = &id
}

func (e *EventSourceConfig) GetName() *string {
	if e == nil {
		return nil
	}
	return e.Name
}

func (e *EventSourceConfig) SetName(name string) {
	e.Name = &name
}

func (e EventSourceConfig) String() string {
	jsonData, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		return "error converting struct: EventSourceConfig to string"
	}
	return string(jsonData)
}
