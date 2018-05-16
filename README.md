# Install

```bash
go get -u github.com/doroginin/protobuf/protoc-gen-go-http-server
```

# Usage
1. Write your proto file `strings/strings.proto`:
```proto
service Strings {
	message String {
		string s = 1;
	}
	rpc ToUpper (String) returns (String) {
		option (google.api.http) = {
			get: "/strings/upper/{string}"
		};
	}
}
```
2. Run code generation
```bash
protoc -I. -I/usr/local/include  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --gofast_out=plugins=grpc:. --go-http-server_out=. strings.proto
```
3. Write main file and run
```go
package main

import (
	"net/http"
	"strings"
)

func main() {
	http.ListenAndServe(":8080", strings.NewStringsHTTPServer())
}
```
4. Implement business logic in `strings.pb.server.impl.go`.
Replace `return &String{}, nil` with `return &String{S: strings.ToUpper(req.S)}, nil` for example and rerun app.
Check url: `http://localhost:8080/strings/upper/test` and you will get result:
```json
{
	"s": "TEST"
}
```
Profit
5. Add swagger if you want:
```go
func main () {
		swg := http.NewServeMux()
    	swg.Handle("/docs/swagger.json", strings.SwaggerJSONHandler)
    	swg.Handle("/docs/", http.StripPrefix("/docs", strings.SwaggerUIHandler))
		http.ListenAndServe(":8080", strings.NewStringsHTTPServer(strings.WithFallbackHandler(swg)))
}
```
and check `http://localhost:8080/docs`

# More options
Available `protoc-gen-go-http-server` options:
- `verbose` - `bool`, show debug info, default `false`
- `impl` - `bool`, generate server implementation stub, default `true`
- `codec` - `bool`, generate codec for parsing http request, and write http response, default `true`
- `swagger` - `bool`, generate swagger documentation handler, default `true`

using: `protoc --go-http-server_out=verbose=true,impl=false,swagger=false,codec=false:. my.proto`

# TODO
 - swagger gen
 - middleware
 - grpc server gen