package main

import (
	"flag"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/plugin"
	"github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway/descriptor"
	"github.com/doroginin/protobuf/protoc-gen-go-http-server/generator"
	"fmt"
)

var (
	verboseMode = flag.Bool("verbose", false, "Show debug information")
	withImpl = flag.Bool("impl", true, "Generate simple implementations for proto Services")
	withSwagger = flag.Bool("swagger", true, "Generate swagger docs")
	withRouter = flag.Bool("router", true, "Generate http router for handlers")
)

func parseReq(r io.Reader) (*plugin_go.CodeGeneratorRequest, error) {
	input, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, fmt.Errorf("failed to read code generator request: %v", err)
	}
	req := new(plugin_go.CodeGeneratorRequest)
	if err = proto.Unmarshal(input, req); err != nil {
		return nil, fmt.Errorf("failed to unmarshal code generator request: %v", err)
	}
	return req, nil
}

func main() {
	flag.Parse()
	var (
		reg      = descriptor.NewRegistry()
		req, err = parseReq(os.Stdin)
	)
	if err != nil {
		log.Panic(err)
	}
	processParameters(req, reg)
	if !*verboseMode {
		log.SetOutput(ioutil.Discard)
	}

	g := generator.New(reg)

	if err := reg.Load(req); err != nil {
		emitError(err)
		return
	}

	var targets []*descriptor.File
	for _, target := range req.FileToGenerate {
		f, err := reg.LookupFile(target)
		if err != nil {
			log.Panic(err)
		}
		targets = append(targets, f)
	}

	out, err := g.GenerateServer(targets)
	if err != nil {
		emitError(err)
		return
	}
	emitFiles(out)

	out, err = g.GenerateHandlers(targets)
	if err != nil {
		emitError(err)
		return
	}
	emitFiles(out)

	if *withImpl {
		out, err = g.GenerateImpl(targets)
		if err != nil {
			emitError(err)
			return
		}
		emitFiles(out)
	}

	if *withRouter {
		out, err = g.GenerateRouter(targets)
		if err != nil {
			emitError(err)
			return
		}
		emitFiles(out)
	}
}

func emitFiles(out []*plugin_go.CodeGeneratorResponse_File) {
	emitResp(&plugin_go.CodeGeneratorResponse{File: out})
}

func emitError(err error) {
	emitResp(&plugin_go.CodeGeneratorResponse{Error: proto.String(err.Error())})
}

func emitResp(resp *plugin_go.CodeGeneratorResponse) {
	buf, err := proto.Marshal(resp)
	if err != nil {
		log.Panic(err)
	}
	if _, err := os.Stdout.Write(buf); err != nil {
		log.Panic(err)
	}
}

func processParameters(req *plugin_go.CodeGeneratorRequest, reg *descriptor.Registry) {
	if req.Parameter != nil {
		for _, p := range strings.Split(req.GetParameter(), ",") {
			spec := strings.SplitN(p, "=", 2)
			if len(spec) == 1 {
				if err := flag.CommandLine.Set(spec[0], ""); err != nil {
					log.Panicf("Cannot set flag %s", p)
				}
				continue
			}
			name, value := spec[0], spec[1]
			if strings.HasPrefix(name, "M") {
				reg.AddPkgMap(name[1:], value)
				continue
			}
			if err := flag.CommandLine.Set(name, value); err != nil {
				log.Panicf("Cannot set flag %s", p)
			}
		}
	}
}
