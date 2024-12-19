package channels

import (
	"encoding/json"
)

type MetadataExpoToken struct {
	Data     *ExpoToken     `json:"data,omitempty" required:"true"`
	Metadata *TokenMetadata `json:"metadata,omitempty" required:"true"`
	touched  map[string]bool
}

func (m *MetadataExpoToken) GetData() *ExpoToken {
	if m == nil {
		return nil
	}
	return m.Data
}

func (m *MetadataExpoToken) SetData(data ExpoToken) {
	if m.touched == nil {
		m.touched = map[string]bool{}
	}
	m.touched["Data"] = true
	m.Data = &data
}

func (m *MetadataExpoToken) SetDataNil() {
	if m.touched == nil {
		m.touched = map[string]bool{}
	}
	m.touched["Data"] = true
	m.Data = nil
}

func (m *MetadataExpoToken) GetMetadata() *TokenMetadata {
	if m == nil {
		return nil
	}
	return m.Metadata
}

func (m *MetadataExpoToken) SetMetadata(metadata TokenMetadata) {
	if m.touched == nil {
		m.touched = map[string]bool{}
	}
	m.touched["Metadata"] = true
	m.Metadata = &metadata
}

func (m *MetadataExpoToken) SetMetadataNil() {
	if m.touched == nil {
		m.touched = map[string]bool{}
	}
	m.touched["Metadata"] = true
	m.Metadata = nil
}
func (m MetadataExpoToken) MarshalJSON() ([]byte, error) {
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
