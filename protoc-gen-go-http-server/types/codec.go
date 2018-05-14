package types

import (
	"net/http"
)

type Codec interface {
	ReadRequest(r *http.Request) (req *http.Request, method string, data interface{}, err error)
	WriteResponse(w http.ResponseWriter, resp interface{}, err error) error
}