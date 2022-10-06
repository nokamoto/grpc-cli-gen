package gen

import (
	"bytes"
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/testing/protocmp"
	"google.golang.org/protobuf/types/pluginpb"
)

func TestPlugin_Run(t *testing.T) {
	table := []struct {
		name     string
		in       *pluginpb.CodeGeneratorRequest
		expected *pluginpb.CodeGeneratorResponse
		err      error
	}{
		{
			name:     "todo",
			in:       &pluginpb.CodeGeneratorRequest{},
			expected: &pluginpb.CodeGeneratorResponse{},
		},
	}

	for _, c := range table {
		t.Run(c.name, func(t *testing.T) {
			in, err := proto.Marshal(c.in)
			if err != nil {
				t.Fatal(err)
			}
			var out bytes.Buffer
			sut := &Plugin{
				in:  bytes.NewReader(in),
				out: &out,
			}

			err = sut.Run()
			if !errors.Is(err, c.err) {
				t.Errorf("expected %v but actual %v", err, c.err)
			}

			var res pluginpb.CodeGeneratorResponse
			err = proto.Unmarshal(out.Bytes(), &res)
			if err != nil {
				t.Fatal(err)
			}

			if diff := cmp.Diff(c.expected, &res, protocmp.Transform()); diff != "" {
				t.Error(diff)
			}
		})
	}
}
