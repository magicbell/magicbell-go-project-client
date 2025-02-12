package channels

import (
	"encoding/json"
	"github.com/magicbell/magicbell-go-project-client/pkg/shared"
)

type ArrayOfExpoTokenResponses struct {
	Data    []ExpoTokenResponse `json:"data,omitempty"`
	Links   *shared.Links       `json:"links,omitempty"`
	touched map[string]bool
}

func (a *ArrayOfExpoTokenResponses) GetData() []ExpoTokenResponse {
	if a == nil {
		return nil
	}
	return a.Data
}

func (a *ArrayOfExpoTokenResponses) SetData(data []ExpoTokenResponse) {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["Data"] = true
	a.Data = data
}

func (a *ArrayOfExpoTokenResponses) SetDataNil() {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["Data"] = true
	a.Data = nil
}

func (a *ArrayOfExpoTokenResponses) GetLinks() *shared.Links {
	if a == nil {
		return nil
	}
	return a.Links
}

func (a *ArrayOfExpoTokenResponses) SetLinks(links shared.Links) {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["Links"] = true
	a.Links = &links
}

func (a *ArrayOfExpoTokenResponses) SetLinksNil() {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["Links"] = true
	a.Links = nil
}

func (a ArrayOfExpoTokenResponses) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if a.touched["Data"] && a.Data == nil {
		data["data"] = nil
	} else if a.Data != nil {
		data["data"] = a.Data
	}

	if a.touched["Links"] && a.Links == nil {
		data["links"] = nil
	} else if a.Links != nil {
		data["links"] = a.Links
	}

	return json.Marshal(data)
}

func (a ArrayOfExpoTokenResponses) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: ArrayOfExpoTokenResponses to string"
	}
	return string(jsonData)
}
