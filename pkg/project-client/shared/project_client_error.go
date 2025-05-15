package shared

import (
	"github.com/magicbell/magicbell-go/pkg/project-client/internal/clients/rest/httptransport"
)

type ProjectClientError struct {
	Err      error
	Body     []byte
	Metadata ProjectClientErrorMetadata
}

type ProjectClientErrorMetadata struct {
	Headers    map[string]string
	StatusCode int
}

func NewProjectClientError[T any](transportError *httptransport.ErrorResponse[T]) *ProjectClientError {
	return &ProjectClientError{
		Err:  transportError.GetError(),
		Body: transportError.GetBody(),
		Metadata: ProjectClientErrorMetadata{
			StatusCode: transportError.GetStatusCode(),
			Headers:    transportError.GetHeaders(),
		},
	}
}

func (e *ProjectClientError) Error() string {
	return e.Err.Error()
}
