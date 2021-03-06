//go:generate protoc -I. -I/usr/local/include -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --gofast_out=plugins=grpc:. --go-http-server_out=verbose=true,impl=false:. ./pb/strings.proto
package main

import (
	"net/http"
	"github.com/doroginin/protobuf/example/pb"
)

func main() {
	swg := http.NewServeMux()
	swg.Handle("/docs/swagger.json", strings.SwaggerJSONHandler)
	swg.Handle("/docs/", http.StripPrefix("/docs", strings.SwaggerUIHandler))
	http.ListenAndServe(":8080", strings.NewStringsHTTPServer(strings.WithFallbackHandler(swg)))
}
