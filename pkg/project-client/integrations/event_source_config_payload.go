package integrations

import "encoding/json"

type EventSourceConfigPayload struct {
	Source *string `json:"source,omitempty" required:"true"`
}

func (e *EventSourceConfigPayload) GetSource() *string {
	if e == nil {
		return nil
	}
	return e.Source
}

func (e *EventSourceConfigPayload) SetSource(source string) {
	e.Source = &source
}

func (e EventSourceConfigPayload) String() string {
	jsonData, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		return "error converting struct: EventSourceConfigPayload to string"
	}
	return string(jsonData)
}
