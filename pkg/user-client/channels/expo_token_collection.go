package channels

import (
	"encoding/json"
	"github.com/magicbell/magicbell-go/pkg/user-client/shared"
)

type ExpoTokenCollection struct {
	Data  []ExpoToken   `json:"data,omitempty"`
	Links *shared.Links `json:"links,omitempty"`
}

func (e *ExpoTokenCollection) GetData() []ExpoToken {
	if e == nil {
		return nil
	}
	return e.Data
}

func (e *ExpoTokenCollection) SetData(data []ExpoToken) {
	e.Data = data
}

func (e *ExpoTokenCollection) GetLinks() *shared.Links {
	if e == nil {
		return nil
	}
	return e.Links
}

func (e *ExpoTokenCollection) SetLinks(links shared.Links) {
	e.Links = &links
}

func (e ExpoTokenCollection) String() string {
	jsonData, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		return "error converting struct: ExpoTokenCollection to string"
	}
	return string(jsonData)
}
