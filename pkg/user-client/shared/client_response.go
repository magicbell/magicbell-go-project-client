package shared

import (
	"encoding/json"

	"github.com/magicbell/magicbell-go/pkg/user-client/internal/clients/rest/httptransport"
)

type ClientResponse[T any] struct {
	Data     T
	Metadata ClientResponseMetadata
}

type ClientResponseMetadata struct {
	Headers    map[string]string
	StatusCode int
}

func NewClientResponse[T any](resp *httptransport.Response[T]) *ClientResponse[T] {
	return &ClientResponse[T]{
		Data: resp.Data,
		Metadata: ClientResponseMetadata{
			StatusCode: resp.StatusCode,
			Headers:    resp.Headers,
		},
	}
}

func (r *ClientResponse[T]) GetData() T {
	return r.Data
}

func (r ClientResponse[T]) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: ClientResponse to string"
	}
	return string(jsonData)
}
