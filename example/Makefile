go=/usr/local/go/bin/go

install:
	@${go} install -v ../protoc-gen-go-http-server

protoc: install
	@GOPATH=${GOPATH%%:*} ${go} generate

all: protoc