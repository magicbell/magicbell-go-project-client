package channels

import (
	"encoding/json"
)

type MetadataFcmToken struct {
	Data     *FcmToken      `json:"data,omitempty" required:"true"`
	Metadata *TokenMetadata `json:"metadata,omitempty" required:"true"`
	touched  map[string]bool
}

func (m *MetadataFcmToken) GetData() *FcmToken {
	if m == nil {
		return nil
	}
	return m.Data
}

func (m *MetadataFcmToken) SetData(data FcmToken) {
	if m.touched == nil {
		m.touched = map[string]bool{}
	}
	m.touched["Data"] = true
	m.Data = &data
}

func (m *MetadataFcmToken) SetDataNil() {
	if m.touched == nil {
		m.touched = map[string]bool{}
	}
	m.touched["Data"] = true
	m.Data = nil
}

func (m *MetadataFcmToken) GetMetadata() *TokenMetadata {
	if m == nil {
		return nil
	}
	return m.Metadata
}

func (m *MetadataFcmToken) SetMetadata(metadata TokenMetadata) {
	if m.touched == nil {
		m.touched = map[string]bool{}
	}
	m.touched["Metadata"] = true
	m.Metadata = &metadata
}

func (m *MetadataFcmToken) SetMetadataNil() {
	if m.touched == nil {
		m.touched = map[string]bool{}
	}
	m.touched["Metadata"] = true
	m.Metadata = nil
}
func (m MetadataFcmToken) MarshalJSON() ([]byte, error) {
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
