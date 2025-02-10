package channels

import (
	"encoding/json"
	"github.com/magicbell/magicbell-go-project-client/pkg/shared"
)

type ArrayOfWebPushTokenResponses struct {
	Data    []WebPushTokenResponse `json:"data,omitempty"`
	Links   *shared.Links          `json:"links,omitempty"`
	touched map[string]bool
}

func (a *ArrayOfWebPushTokenResponses) GetData() []WebPushTokenResponse {
	if a == nil {
		return nil
	}
	return a.Data
}

func (a *ArrayOfWebPushTokenResponses) SetData(data []WebPushTokenResponse) {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["Data"] = true
	a.Data = data
}

func (a *ArrayOfWebPushTokenResponses) SetDataNil() {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["Data"] = true
	a.Data = nil
}

func (a *ArrayOfWebPushTokenResponses) GetLinks() *shared.Links {
	if a == nil {
		return nil
	}
	return a.Links
}

func (a *ArrayOfWebPushTokenResponses) SetLinks(links shared.Links) {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["Links"] = true
	a.Links = &links
}

func (a *ArrayOfWebPushTokenResponses) SetLinksNil() {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["Links"] = true
	a.Links = nil
}

func (a ArrayOfWebPushTokenResponses) MarshalJSON() ([]byte, error) {
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

func (a ArrayOfWebPushTokenResponses) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: ArrayOfWebPushTokenResponses to string"
	}
	return string(jsonData)
}
