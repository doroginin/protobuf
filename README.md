# Install

```bash
go get -u github.com/doroginin/protobuf/protoc-gen-go-http-server
```

# Usage
Available `protoc-gen-go-http-server` options:
- `verbose` - `bool`, show debug info, default `false`
- `impl` - `bool`, generate server implementation stub, default `true`
- `codec` - `bool`, generate codec for parsing http request, and write http response, default `true`
- `swagger` - `bool`, generate swagger documentation handler, default `true`

using: `protoc --go-http-server_out=verbose=true,impl=false,swagger=false,codec=false:. my.proto`

```bash
protoc -I. -I/usr/local/include  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --gofast_out=plugins=grpc:. --go-http-server_out=. strings.proto
```

# todo
 - swagger gen
 - middleware
 - grpc server gen