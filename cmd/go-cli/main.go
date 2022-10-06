package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/nokamoto/grpc-cli-gen/internal/gen"
)

func main() {
	plugin := gen.NewPlugin()
	if err := plugin.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", filepath.Base(os.Args[0]), err)
		os.Exit(1)
	}
}
