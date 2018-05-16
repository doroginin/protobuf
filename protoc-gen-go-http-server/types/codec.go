package types

import (
	"net/http"
	"errors"
)

var ErrMethodNotFound = errors.New("method is not found")

type Codec interface {
	ReadRequest(r *http.Request) (req *http.Request, method string, data interface{}, err error)
	WriteResponse(w http.ResponseWriter, resp interface{}, err error) error
}