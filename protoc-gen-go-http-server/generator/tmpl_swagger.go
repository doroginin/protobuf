package generator

import (
	"text/template"
	"strings"
)

var SwaggerTemplate = &Template{
	FileName: "%s.pb.swagger.go",
	Body: template.Must(template.New(`file`).Funcs(template.FuncMap{
		"escape": func(s string) string { return strings.Replace(s, "`", "` + \"``\" + `", -1) },
	}).Parse(`
// Code generated by protoc-gen-go-http-server.
// source: {{ .Source }}

package {{ .Package }}

import (
	"net/http"

	"github.com/doroginin/protobuf/protoc-gen-go-http-server/swagger"
)

{{ range $sIdx, $service := .Services }}

var SwaggerJSONHandler = http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/javascript")
	w.WriteHeader(http.StatusOK)
	w.Write(_swaggerJSON)
})

var SwaggerUIHandler = swaggerui.NewHTTPHandler() 

{{ end }}

var _swaggerJSON = []byte(` + "`" + `{{ escape (.Fields.Swagger) }}` + `` + "`)" + `
`))}