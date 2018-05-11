package types

import "net/http"

type Router interface {
	Route(r *http.Request) (handler string, err error)
}