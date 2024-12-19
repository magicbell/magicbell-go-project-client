package shared

import (
	"github.com/magicbell/magicbell-go-project-client/internal/clients/rest/httptransport"
)

type MagicbellProjectClientError struct {
	Err      error
	Body     []byte
	Metadata MagicbellProjectClientErrorMetadata
}

type MagicbellProjectClientErrorMetadata struct {
	Headers    map[string]string
	StatusCode int
}

func NewMagicbellProjectClientError[T any](transportError *httptransport.ErrorResponse[T]) *MagicbellProjectClientError {
	return &MagicbellProjectClientError{
		Err:  transportError.GetError(),
		Body: transportError.GetBody(),
		Metadata: MagicbellProjectClientErrorMetadata{
			StatusCode: transportError.GetStatusCode(),
			Headers:    transportError.GetHeaders(),
		},
	}
}

func (e *MagicbellProjectClientError) Error() string {
	return e.Err.Error()
}
