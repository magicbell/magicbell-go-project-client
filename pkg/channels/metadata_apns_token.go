package channels

import (
	"encoding/json"
)

type MetadataApnsToken struct {
	Data     *ApnsToken     `json:"data,omitempty" required:"true"`
	Metadata *TokenMetadata `json:"metadata,omitempty" required:"true"`
	touched  map[string]bool
}

func (m *MetadataApnsToken) GetData() *ApnsToken {
	if m == nil {
		return nil
	}
	return m.Data
}

func (m *MetadataApnsToken) SetData(data ApnsToken) {
	if m.touched == nil {
		m.touched = map[string]bool{}
	}
	m.touched["Data"] = true
	m.Data = &data
}

func (m *MetadataApnsToken) SetDataNil() {
	if m.touched == nil {
		m.touched = map[string]bool{}
	}
	m.touched["Data"] = true
	m.Data = nil
}

func (m *MetadataApnsToken) GetMetadata() *TokenMetadata {
	if m == nil {
		return nil
	}
	return m.Metadata
}

func (m *MetadataApnsToken) SetMetadata(metadata TokenMetadata) {
	if m.touched == nil {
		m.touched = map[string]bool{}
	}
	m.touched["Metadata"] = true
	m.Metadata = &metadata
}

func (m *MetadataApnsToken) SetMetadataNil() {
	if m.touched == nil {
		m.touched = map[string]bool{}
	}
	m.touched["Metadata"] = true
	m.Metadata = nil
}
func (m MetadataApnsToken) MarshalJSON() ([]byte, error) {
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
