package integrations

import (
	"encoding/json"
	"github.com/magicbell/magicbell-go-project-client/pkg/shared"
)

type ArrayOfApnsConfigObjects struct {
	Data    []ApnsConfigObject `json:"data,omitempty"`
	Links   *shared.Links      `json:"links,omitempty"`
	touched map[string]bool
}

func (a *ArrayOfApnsConfigObjects) GetData() []ApnsConfigObject {
	if a == nil {
		return nil
	}
	return a.Data
}

func (a *ArrayOfApnsConfigObjects) SetData(data []ApnsConfigObject) {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["Data"] = true
	a.Data = data
}

func (a *ArrayOfApnsConfigObjects) SetDataNil() {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["Data"] = true
	a.Data = nil
}

func (a *ArrayOfApnsConfigObjects) GetLinks() *shared.Links {
	if a == nil {
		return nil
	}
	return a.Links
}

func (a *ArrayOfApnsConfigObjects) SetLinks(links shared.Links) {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["Links"] = true
	a.Links = &links
}

func (a *ArrayOfApnsConfigObjects) SetLinksNil() {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["Links"] = true
	a.Links = nil
}
func (a ArrayOfApnsConfigObjects) MarshalJSON() ([]byte, error) {
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
