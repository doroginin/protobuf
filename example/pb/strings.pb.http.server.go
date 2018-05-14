// Code generated by protoc-gen-go-http-server.
// source: pb/strings.proto

package strings

import (
	"errors"
	"net/http"

	"github.com/doroginin/protobuf/protoc-gen-go-http-server/types"
)

type StringsHTTPServer struct {
	handlers map[string]http.HandlerFunc
	opts     _StringsHTTPServerOptions
}

func WithImpl(s StringsServer) StringsHTTPServerOption {
	return func(opts *_StringsHTTPServerOptions) {
		opts.srv = s
	}
}

func WithCodec(c types.Codec) StringsHTTPServerOption {
	return func(opts *_StringsHTTPServerOptions) {
		opts.cdc = c
	}
}

type _StringsHTTPServerOptions struct {
	srv StringsServer
	cdc types.Codec
}

var defaultStringsHTTPServerOptions _StringsHTTPServerOptions

type StringsHTTPServerOption func(*_StringsHTTPServerOptions)

func NewStringsHTTPServer(opts ...StringsHTTPServerOption) *StringsHTTPServer {
	s := &StringsHTTPServer{opts: defaultStringsHTTPServerOptions}
	for _, opt := range opts {
		opt(&s.opts)
	}
	s.handlers = map[string]http.HandlerFunc{
		"ToUpper": s.ToUpper,

		"ToLower": s.ToLower,
	}
	return s
}

func (s *StringsHTTPServer) Handler(name string) (http.HandlerFunc, bool) {
	h, ok := s.handlers[name]
	return h, ok
}

func (s *StringsHTTPServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if s.opts.cdc == nil {
		s.opts.cdc.WriteResponse(w, nil, errors.New("Codec is not defined"))
		return
	}
	r, method, _, err := s.opts.cdc.ReadRequest(r)
	if err != nil {
		s.opts.cdc.WriteResponse(w, nil, err)
		return
	}
	if handler, ok := s.handlers[method]; ok {
		handler(w, r)
		return
	}
	s.opts.cdc.WriteResponse(w, nil, errors.New("handler is not found for request"))
}
