package generator

import (
	"github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway/descriptor"
	"text/template"
)

type Template struct {
	FileName string
	Body *template.Template
}

type templateFileInfo struct {
	Source   string
	Package  string
	Services []*templateService
}

type templateService struct {
	Name     string
	Handlers []*templateHandler
}

type templateHandler struct {
	Name     string
	Service  string
	In       string
	Out      string
	Bindings []*descriptor.Binding
}
