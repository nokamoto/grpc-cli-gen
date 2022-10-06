package gen

import (
	"io"
	"io/ioutil"
	"os"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/pluginpb"
)

// ref. https://github.com/protocolbuffers/protobuf-go/blob/7a48e2b66218ba306ed801e42126b374aefce255/compiler/protogen/protogen.go#L1
type Plugin struct {
	in  io.Reader
	out io.Writer
}

// NewPlugin returns a plugin which reads a CodeGeneratorRequest message from os.Stdin
// and writes a CodeGeneratorResponse message to os.Stdout.
func NewPlugin() *Plugin {
	return &Plugin{
		in:  os.Stdin,
		out: os.Stdout,
	}
}

// Run executes a function as a protoc plugin.
func (p *Plugin) Run() error {
	in, err := ioutil.ReadAll(p.in)
	if err != nil {
		return err
	}
	req := &pluginpb.CodeGeneratorRequest{}
	if err := proto.Unmarshal(in, req); err != nil {
		return err
	}
	res := p.response(req)
	out, err := proto.Marshal(res)
	if err != nil {
		return err
	}
	if _, err := p.out.Write(out); err != nil {
		return err
	}
	return nil
}

func (p *Plugin) response(req *pluginpb.CodeGeneratorRequest) *pluginpb.CodeGeneratorResponse {
	return nil
}
