package events

import (
	"encoding/json"
	"github.com/magicbell/magicbell-go-project-client/pkg/shared"
)

type EventCollection struct {
	Data  []Event       `json:"data,omitempty"`
	Links *shared.Links `json:"links,omitempty"`
}

func (e *EventCollection) GetData() []Event {
	if e == nil {
		return nil
	}
	return e.Data
}

func (e *EventCollection) SetData(data []Event) {
	e.Data = data
}

func (e *EventCollection) GetLinks() *shared.Links {
	if e == nil {
		return nil
	}
	return e.Links
}

func (e *EventCollection) SetLinks(links shared.Links) {
	e.Links = &links
}

func (e EventCollection) String() string {
	jsonData, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		return "error converting struct: EventCollection to string"
	}
	return string(jsonData)
}
