package integrations

import (
	"encoding/json"
	"github.com/magicbell/magicbell-go-project-client/pkg/shared"
)

type EventSourceConfigCollection struct {
	Data  []EventSourceConfig `json:"data,omitempty"`
	Links *shared.Links       `json:"links,omitempty"`
}

func (e *EventSourceConfigCollection) GetData() []EventSourceConfig {
	if e == nil {
		return nil
	}
	return e.Data
}

func (e *EventSourceConfigCollection) SetData(data []EventSourceConfig) {
	e.Data = data
}

func (e *EventSourceConfigCollection) GetLinks() *shared.Links {
	if e == nil {
		return nil
	}
	return e.Links
}

func (e *EventSourceConfigCollection) SetLinks(links shared.Links) {
	e.Links = &links
}

func (e EventSourceConfigCollection) String() string {
	jsonData, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		return "error converting struct: EventSourceConfigCollection to string"
	}
	return string(jsonData)
}
