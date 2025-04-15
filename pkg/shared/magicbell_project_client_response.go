package shared

import (
	"encoding/json"

	"github.com/magicbell/magicbell-go-project-client/internal/clients/rest/httptransport"
)

type MagicbellProjectClientResponse[T any] struct {
	Data     T
	Metadata MagicbellProjectClientResponseMetadata
}

type MagicbellProjectClientResponseMetadata struct {
	Headers    map[string]string
	StatusCode int
}

func NewMagicbellProjectClientResponse[T any](resp *httptransport.Response[T]) *MagicbellProjectClientResponse[T] {
	return &MagicbellProjectClientResponse[T]{
		Data: resp.Data,
		Metadata: MagicbellProjectClientResponseMetadata{
			StatusCode: resp.StatusCode,
			Headers:    resp.Headers,
		},
	}
}

func (r *MagicbellProjectClientResponse[T]) GetData() T {
	return r.Data
}

func (r MagicbellProjectClientResponse[T]) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: MagicbellProjectClientResponse to string"
	}
	return string(jsonData)
}
