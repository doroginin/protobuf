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

func WithRouter(r types.Router) StringsHTTPServerOption {
	return func(opts *_StringsHTTPServerOptions) {
		opts.rtr = r
	}
}

type _{{ $service.Name }}HTTPServerOptions struct{
	srv {{ $service.Name }}Server
	cdc types.Codec
	rtr types.Router
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
	if s.opts.rtr == nil {
		s.opts.cdc.WriteError(w, errors.New("router is not defined"))
		return		
	}
	name, err := s.opts.rtr.Route(r)
	if err != nil {
		s.opts.cdc.WriteError(w, err)
		return
	}
	if handler, ok := s.handlers[name]; ok {
		handler(w, r)
		return
	}
	s.opts.cdc.WriteError(w, errors.New("handler is not found for request"))
}

{{ end }}`))}