package handlers

import (
	"errors"

	"github.com/magicbell/magicbell-go/pkg/project-client/internal/clients/rest/httptransport"
	"github.com/magicbell/magicbell-go/pkg/project-client/internal/contenttypes"
)

type UnmarshalHandler[T any] struct {
	nextHandler Handler[T]
}

func NewUnmarshalHandler[T any]() *UnmarshalHandler[T] {
	return &UnmarshalHandler[T]{
		nextHandler: nil,
	}
}

func (h *UnmarshalHandler[T]) Handle(request httptransport.Request) (*httptransport.Response[T], *httptransport.ErrorResponse[T]) {
	if h.nextHandler == nil {
		err := errors.New("Handler chain terminated without terminating handler")
		return nil, httptransport.NewErrorResponse[T](err, nil)
	}

	resp, handlerError := h.nextHandler.Handle(request)
	if handlerError != nil {
		return nil, handlerError
	}

	target := new(T)
	if request.ResponseContentType == httptransport.ContentTypeJson {
		err := contenttypes.FromJson(resp.Body, target)
		if err != nil {
			return nil, httptransport.NewErrorResponse[T](err, resp)
		}
	} else if request.ResponseContentType == httptransport.ContentTypeFormUrlEncoded {
		err := contenttypes.FromFormUrlEncoded(resp.Body, target)
		if err != nil {
			return nil, httptransport.NewErrorResponse[T](err, resp)
		}
	} else if request.ResponseContentType == httptransport.ContentTypeMultipartFormData {
		err := contenttypes.FromFormData(resp.Body, target)
		if err != nil {
			return nil, httptransport.NewErrorResponse[T](err, resp)
		}
	} else if request.ResponseContentType == httptransport.ContentTypeText {
		err := contenttypes.FromText[T](resp.Body, target)
		if err != nil {
			return nil, httptransport.NewErrorResponse[T](err, resp)
		}
	} else if request.ResponseContentType == httptransport.ContentTypeBinary {
		err := contenttypes.FromBinary(resp.Body, target)
		if err != nil {
			return nil, httptransport.NewErrorResponse[T](err, resp)
		}
	} else {
		err := contenttypes.FromBinary(resp.Body, target)
		if err != nil {
			return nil, httptransport.NewErrorResponse[T](err, resp)
		}
	}

	resp.Data = *target

	return resp, nil
}

func (h *UnmarshalHandler[T]) SetNext(handler Handler[T]) {
	h.nextHandler = handler
}
