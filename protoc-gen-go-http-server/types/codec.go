package types

import (
	"net/http"
)

type Codec interface {
	ReadRequest(r *http.Request, out interface{}) error
	WriteResponse(w http.ResponseWriter, resp interface{}) error
	WriteError(w http.ResponseWriter, err error) error
}