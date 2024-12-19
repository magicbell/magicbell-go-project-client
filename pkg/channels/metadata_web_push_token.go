package channels

import (
	"encoding/json"
)

type MetadataWebPushToken struct {
	Data     *WebPushToken  `json:"data,omitempty" required:"true"`
	Metadata *TokenMetadata `json:"metadata,omitempty" required:"true"`
	touched  map[string]bool
}

func (m *MetadataWebPushToken) GetData() *WebPushToken {
	if m == nil {
		return nil
	}
	return m.Data
}

func (m *MetadataWebPushToken) SetData(data WebPushToken) {
	if m.touched == nil {
		m.touched = map[string]bool{}
	}
	m.touched["Data"] = true
	m.Data = &data
}

func (m *MetadataWebPushToken) SetDataNil() {
	if m.touched == nil {
		m.touched = map[string]bool{}
	}
	m.touched["Data"] = true
	m.Data = nil
}

func (m *MetadataWebPushToken) GetMetadata() *TokenMetadata {
	if m == nil {
		return nil
	}
	return m.Metadata
}

func (m *MetadataWebPushToken) SetMetadata(metadata TokenMetadata) {
	if m.touched == nil {
		m.touched = map[string]bool{}
	}
	m.touched["Metadata"] = true
	m.Metadata = &metadata
}

func (m *MetadataWebPushToken) SetMetadataNil() {
	if m.touched == nil {
		m.touched = map[string]bool{}
	}
	m.touched["Metadata"] = true
	m.Metadata = nil
}
func (m MetadataWebPushToken) MarshalJSON() ([]byte, error) {
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
