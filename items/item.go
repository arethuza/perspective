package items

import "net/http"

type Item interface {
	TypeName() string
}

type HttpError struct {
	Code    int
	message string
}

func (he HttpError) Error() string {
	return he.message
}

type Action func(item Item) (*ActionResult, *HttpError)

type ActionResult interface {
	SendResponse(w http.ResponseWriter)
}
