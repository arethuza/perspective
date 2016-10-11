package items

import (
	"encoding/json"
	"net/http"
)

type Item interface {
	TypeName() string
}

type HttpError struct {
	Code    int
	Message string
}

func (he HttpError) Error() string {
	return he.Message
}

type Action func(item Item) (ActionResult, *HttpError)

type ActionResult interface {
	SendResponse(w http.ResponseWriter) *HttpError
}

type JsonResult struct {
	value interface{}
}

func (jr JsonResult) SendResponse(w http.ResponseWriter) *HttpError {
	data, err := json.Marshal(jr.value)
	if err != nil {
		return &HttpError{Code: 500, Message: err.Error()}
	}
	_, err = w.Write(data)
	if err != nil {
		return &HttpError{Code: 500, Message: err.Error()}
	}
	return nil
}
