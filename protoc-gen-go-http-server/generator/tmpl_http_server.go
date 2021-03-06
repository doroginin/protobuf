package generator

import (
	"strings"
	"text/template"
)

var ServerTemplate = &Template{
	FileName: "%s.pb.http.server.go",
	Body: template.Must(template.New(`file`).Funcs(template.FuncMap{
		`lower`: strings.ToLower,
	}).Parse(`
// Code generated by protoc-gen-go-http-server.
// source: {{ .Source }}

package {{ .Package }}

import (
	"errors"
	"net/http"

	"github.com/doroginin/protobuf/protoc-gen-go-http-server/types"
)

{{ range $sIdx, $service := .Services }}

type {{ $service.Name }}HTTPServer struct {
	handlers map[string]http.HandlerFunc
	opts _{{ $service.Name }}HTTPServerOptions
}

func WithImpl(s {{ $service.Name }}Server) {{ $service.Name }}HTTPServerOption {
	return func(opts *_{{ $service.Name }}HTTPServerOptions) {
		opts.srv = s 
	}
}

func WithCodec(c types.Codec) {{ $service.Name }}HTTPServerOption {
	return func(opts *_{{ $service.Name }}HTTPServerOptions) {
		opts.cdc = c
	}
}

func WithFallbackHandler(h http.Handler) {{ $service.Name }}HTTPServerOption {
	return func(opts *_StringsHTTPServerOptions) {
		opts.fallback = h
	}
}

type _{{ $service.Name }}HTTPServerOptions struct{
	srv {{ $service.Name }}Server
	cdc types.Codec
	fallback http.Handler
}

var default{{ $service.Name }}HTTPServerOptions _{{ $service.Name }}HTTPServerOptions

type {{ $service.Name }}HTTPServerOption func(*_{{ $service.Name }}HTTPServerOptions)

func New{{ $service.Name }}HTTPServer(opts... {{ $service.Name }}HTTPServerOption) *{{ $service.Name }}HTTPServer {
	s := &{{ $service.Name }}HTTPServer{opts: default{{ $service.Name }}HTTPServerOptions}
	for _, opt := range opts {
		opt(&s.opts)
	}
	s.handlers = map[string]http.HandlerFunc{ {{ range $hIdx, $handler := $service.Handlers }}
		"{{ $handler.Name }}": s.{{ $handler.Name }},
	{{ end }} }
	return s
}

func (s *{{ $service.Name }}HTTPServer) Handler(name string) (http.HandlerFunc, bool) {
	h, ok := s.handlers[name]
	return h, ok
}

func (s *{{ $service.Name }}HTTPServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if s.opts.cdc == nil {
		s.opts.cdc.WriteResponse(w, nil, errors.New("codec is not defined"))
		return		
	}
	r, method, _, err := s.opts.cdc.ReadRequest(r)
	if err == types.ErrMethodNotFound && s.opts.fallback != nil {
		s.opts.fallback.ServeHTTP(w, r)
		return
	}
	if err != nil {
		s.opts.cdc.WriteResponse(w, nil, err)
		return
	}
	if handler, ok := s.Handler(method); ok {
		handler(w, r)
		return
	}
	s.opts.cdc.WriteResponse(w, nil, errors.New("handler is not found for request"))
}

{{ end }}`))}