package shared

import (
	"encoding/json"

	"github.com/magicbell/magicbell-go/pkg/user-client/internal/clients/rest/httptransport"
)

type UserClientResponse[T any] struct {
	Data     T
	Metadata UserClientResponseMetadata
}

type UserClientResponseMetadata struct {
	Headers    map[string]string
	StatusCode int
}

func NewUserClientResponse[T any](resp *httptransport.Response[T]) *UserClientResponse[T] {
	return &UserClientResponse[T]{
		Data: resp.Data,
		Metadata: UserClientResponseMetadata{
			StatusCode: resp.StatusCode,
			Headers:    resp.Headers,
		},
	}
}

func (r *UserClientResponse[T]) GetData() T {
	return r.Data
}

func (r UserClientResponse[T]) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: UserClientResponse to string"
	}
	return string(jsonData)
}
