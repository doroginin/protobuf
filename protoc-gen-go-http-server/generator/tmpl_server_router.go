package generator

import (
	"strings"
	"text/template"
)

var ServerRouterTemplate = &Template{
	FileName: "%s.pb.http.server.router.go",
	Body: template.Must(template.New(`file`).Funcs(template.FuncMap{
		`ToLower`: strings.ToLower,
	}).Parse(`
// Code generated by protoc-gen-go-http-server.
// source: {{ .Source }}

package {{ .Package }}

import (
	"strings"
	"errors"
	"net/http"
)

func init() { {{ range $si, $service := .Services }}
	default{{ $service.Name }}HTTPServerOptions.rtr = &{{ $service.Name }}Router{}
{{ end }} }

{{ range $si, $service := .Services }}
type {{ $service.Name }}Router struct{
}

func (r *{{ $service.Name }}Router) Route(req *http.Request) (string, error) {
	path := strings.Split(strings.Trim(req.URL.Path, "/"), "/")
	{{ range $hi, $handler := $service.Handlers }}
		{{ if $handler.Bindings }}
   			{{ range $bId, $binding := $handler.Bindings }}
				// method: {{ $binding.HTTPMethod }}
			{{ end }}
		{{ else }}
			if len(path) == 2 && path[0] == "{{ $service.Name }}" && path[1] == "{{ $handler.Name }}" {
				if req.Method != "POST" {
					return "", errors.New("excepted POST method")
				}
				return path[1], nil
			}
		{{ end }}
	{{ end }}
	return "", errors.New("route is not found")
}

{{ end }}
`))}
