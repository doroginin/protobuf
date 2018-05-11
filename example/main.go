//go:generate protoc -I. -I/usr/local/include -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --gofast_out=plugins=grpc:. --go-http-server_out=verbose=true:. ./pb/strings.proto
package main

import (
	"net/http"
	"github.com/doroginin/protobuf/example/pb"
)

func main() {
	http.ListenAndServe(":8080", strings.NewStringsHTTPServer())
}
