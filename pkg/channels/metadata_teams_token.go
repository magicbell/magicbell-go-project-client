package channels

import (
	"encoding/json"
)

type MetadataTeamsToken struct {
	Data     *TeamsToken    `json:"data,omitempty" required:"true"`
	Metadata *TokenMetadata `json:"metadata,omitempty" required:"true"`
	touched  map[string]bool
}

func (m *MetadataTeamsToken) GetData() *TeamsToken {
	if m == nil {
		return nil
	}
	return m.Data
}

func (m *MetadataTeamsToken) SetData(data TeamsToken) {
	if m.touched == nil {
		m.touched = map[string]bool{}
	}
	m.touched["Data"] = true
	m.Data = &data
}

func (m *MetadataTeamsToken) SetDataNil() {
	if m.touched == nil {
		m.touched = map[string]bool{}
	}
	m.touched["Data"] = true
	m.Data = nil
}

func (m *MetadataTeamsToken) GetMetadata() *TokenMetadata {
	if m == nil {
		return nil
	}
	return m.Metadata
}

func (m *MetadataTeamsToken) SetMetadata(metadata TokenMetadata) {
	if m.touched == nil {
		m.touched = map[string]bool{}
	}
	m.touched["Metadata"] = true
	m.Metadata = &metadata
}

func (m *MetadataTeamsToken) SetMetadataNil() {
	if m.touched == nil {
		m.touched = map[string]bool{}
	}
	m.touched["Metadata"] = true
	m.Metadata = nil
}
func (m MetadataTeamsToken) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if m.touched["Data"] && m.Data == nil {
		data["data"] = nil
	} else if m.Data != nil {
		data["data"] = m.Data
	}

	if m.touched["Metadata"] && m.Metadata == nil {
		data["metadata"] = nil
	} else if m.Metadata != nil {
		data["metadata"] = m.Metadata
	}

	return json.Marshal(data)
}
