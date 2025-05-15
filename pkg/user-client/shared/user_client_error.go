package shared

import (
	"github.com/magicbell/magicbell-go/pkg/user-client/internal/clients/rest/httptransport"
)

type UserClientError struct {
	Err      error
	Body     []byte
	Metadata UserClientErrorMetadata
}

type UserClientErrorMetadata struct {
	Headers    map[string]string
	StatusCode int
}

func NewUserClientError[T any](transportError *httptransport.ErrorResponse[T]) *UserClientError {
	return &UserClientError{
		Err:  transportError.GetError(),
		Body: transportError.GetBody(),
		Metadata: UserClientErrorMetadata{
			StatusCode: transportError.GetStatusCode(),
			Headers:    transportError.GetHeaders(),
		},
	}
}

func (e *UserClientError) Error() string {
	return e.Err.Error()
}
