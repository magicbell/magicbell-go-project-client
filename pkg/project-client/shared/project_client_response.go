package shared

import (
	"encoding/json"

	"github.com/magicbell/magicbell-go/pkg/project-client/internal/clients/rest/httptransport"
)

type ProjectClientResponse[T any] struct {
	Data     T
	Metadata ProjectClientResponseMetadata
}

type ProjectClientResponseMetadata struct {
	Headers    map[string]string
	StatusCode int
}

func NewProjectClientResponse[T any](resp *httptransport.Response[T]) *ProjectClientResponse[T] {
	return &ProjectClientResponse[T]{
		Data: resp.Data,
		Metadata: ProjectClientResponseMetadata{
			StatusCode: resp.StatusCode,
			Headers:    resp.Headers,
		},
	}
}

func (r *ProjectClientResponse[T]) GetData() T {
	return r.Data
}

func (r ProjectClientResponse[T]) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: ProjectClientResponse to string"
	}
	return string(jsonData)
}
