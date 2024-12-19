package jwt

import (
	"encoding/json"
	"github.com/magicbell/magicbell-go-project-client/pkg/shared"
)

type ArrayOfFetchTokensResponseTokens struct {
	Data    []FetchTokensResponseToken `json:"data,omitempty"`
	Links   *shared.Links              `json:"links,omitempty"`
	touched map[string]bool
}

func (a *ArrayOfFetchTokensResponseTokens) GetData() []FetchTokensResponseToken {
	if a == nil {
		return nil
	}
	return a.Data
}

func (a *ArrayOfFetchTokensResponseTokens) SetData(data []FetchTokensResponseToken) {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["Data"] = true
	a.Data = data
}

func (a *ArrayOfFetchTokensResponseTokens) SetDataNil() {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["Data"] = true
	a.Data = nil
}

func (a *ArrayOfFetchTokensResponseTokens) GetLinks() *shared.Links {
	if a == nil {
		return nil
	}
	return a.Links
}

func (a *ArrayOfFetchTokensResponseTokens) SetLinks(links shared.Links) {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["Links"] = true
	a.Links = &links
}

func (a *ArrayOfFetchTokensResponseTokens) SetLinksNil() {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["Links"] = true
	a.Links = nil
}
func (a ArrayOfFetchTokensResponseTokens) MarshalJSON() ([]byte, error) {
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
