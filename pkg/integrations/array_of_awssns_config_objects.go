package integrations

import (
	"encoding/json"
	"github.com/magicbell/magicbell-go-project-client/pkg/shared"
)

type ArrayOfAwssnsConfigObjects struct {
	Data    []AwssnsConfigObject `json:"data,omitempty"`
	Links   *shared.Links        `json:"links,omitempty"`
	touched map[string]bool
}

func (a *ArrayOfAwssnsConfigObjects) GetData() []AwssnsConfigObject {
	if a == nil {
		return nil
	}
	return a.Data
}

func (a *ArrayOfAwssnsConfigObjects) SetData(data []AwssnsConfigObject) {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["Data"] = true
	a.Data = data
}

func (a *ArrayOfAwssnsConfigObjects) SetDataNil() {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["Data"] = true
	a.Data = nil
}

func (a *ArrayOfAwssnsConfigObjects) GetLinks() *shared.Links {
	if a == nil {
		return nil
	}
	return a.Links
}

func (a *ArrayOfAwssnsConfigObjects) SetLinks(links shared.Links) {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["Links"] = true
	a.Links = &links
}

func (a *ArrayOfAwssnsConfigObjects) SetLinksNil() {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["Links"] = true
	a.Links = nil
}
func (a ArrayOfAwssnsConfigObjects) MarshalJSON() ([]byte, error) {
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
