package generator

import (
	"bytes"
	"errors"
	"fmt"
	"go/format"
	"log"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/plugin"
	"github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway/descriptor"
	options "google.golang.org/genproto/googleapis/api/annotations"
)

var (
	errNoTargetService = errors.New("no target service defined in the file")
)

type Generator struct {
	reg *descriptor.Registry
}

func New(reg *descriptor.Registry) *Generator {
	return &Generator{
		reg: reg,
	}
}

func (g *Generator) GenerateServer(targets []*descriptor.File) (files []*plugin_go.CodeGeneratorResponse_File, err error) {
	ff, err := g.buildFiles(targets, ServerTemplate)
	if err != nil {
		return nil, err
	}
	files = append(files, ff...)
	return files, nil
}

func (g *Generator) GenerateHandlers(targets []*descriptor.File) (files []*plugin_go.CodeGeneratorResponse_File, err error) {
	ff, err := g.buildFiles(targets, HandlerTemplate)
	if err != nil {
		return nil, err
	}
	files = append(files, ff...)
	return files, nil
}

func (g *Generator) GenerateSwagger(targets []*descriptor.File, fields... TemplateField) (files []*plugin_go.CodeGeneratorResponse_File, err error) {
	ff, err := g.buildFiles(targets, SwaggerTemplate, fields...)
	if err != nil {
		return nil, err
	}
	files = append(files, ff...)
	return files, nil
}

func (g *Generator) GenerateImpl(targets []*descriptor.File) (files []*plugin_go.CodeGeneratorResponse_File, err error) {
	ff, err := g.buildFiles(targets, ServerImplTemplate)
	if err != nil {
		return nil, err
	}
	files = append(files, ff...)
	return files, nil
}

func (g *Generator) GenerateCodec(targets []*descriptor.File) (files []*plugin_go.CodeGeneratorResponse_File, err error) {
	ff, err := g.buildFiles(targets, CodecTemplate)
	if err != nil {
		return nil, err
	}
	files = append(files, ff...)
	return files, nil
}

func (g *Generator) buildFiles(targets []*descriptor.File, tmpl *Template, fields... TemplateField) (files []*plugin_go.CodeGeneratorResponse_File, err error) {
	for _, file := range targets {
		log.Printf("Processing %s -> %s", file.GetName(), fmt.Sprintf(tmpl.FileName, file.GetName()))

		code, err := g.generateFrom(file, tmpl.Body, fields...)
		if err == errNoTargetService {
			log.Printf("%s: %v", file.GetName(), err)
			continue
		}
		if err != nil {
			return nil, err
		}

		formatted, err := format.Source([]byte(code))
		if err != nil {
			log.Printf("%v: %s", err, code)
			return nil, err
		}

		var (
			name   = file.GetName()
			ext    = filepath.Ext(name)
			base   = strings.TrimSuffix(name, ext)
			output = fmt.Sprintf(tmpl.FileName, base)
		)
		files = append(files, &plugin_go.CodeGeneratorResponse_File{
			Name:    proto.String(output),
			Content: proto.String(string(formatted)),
		})
	}

	return files, nil
}

func (g *Generator) generateFrom(file *descriptor.File, t *template.Template, fields... TemplateField) (string, error) {
	pkgSeen := make(map[string]bool)
	var imports []descriptor.GoPackage
	tFileInfo := &templateFileInfo{
		Source:  file.GetName(),
		Package: file.GoPkg.Name,
		Fields:  make(map[string]interface{}),
	}
	for _, f := range fields {
		if f.FileName == tFileInfo.Source {
			tFileInfo.Fields[f.Key] = f.Value
		}
	}
	for _, svc := range file.Services {
		tService := &templateService{Name: svc.GetName()}
		tFileInfo.Services = append(tFileInfo.Services, tService)

		for _, m := range svc.Methods {
			if m.GetServerStreaming() || m.GetClientStreaming() {
				continue
			}

			tService.Handlers = append(tService.Handlers, &templateHandler{
				Name:     m.GetName(),
				In:       m.GetInputType()[1:],
				Out:      m.GetOutputType()[1:],
				Bindings: m.Bindings,
			})

			g.markSeen(file, m, pkgSeen, imports)
		}
	}

	buf := bytes.NewBuffer([]byte{})
	t.Execute(buf, tFileInfo)

	return buf.String(), nil
}

func (g *Generator) markSeen(file *descriptor.File, m *descriptor.Method, pkgSeen map[string]bool, imports []descriptor.GoPackage) {
	pkg := m.RequestType.File.GoPkg
	if m.Options == nil || !proto.HasExtension(m.Options, options.E_Http) ||
		pkg == file.GoPkg || pkgSeen[pkg.Path] {
		return
	}
	pkgSeen[pkg.Path] = true
	imports = append(imports, pkg)
}
