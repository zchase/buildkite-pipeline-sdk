package main

import (
	"fmt"
	"os"

	"github.com/zchase/buildkite-pipeline-sdk/pkg/codegen"
	"github.com/zchase/buildkite-pipeline-sdk/pkg/schema"
)

func run() error {
	err := codegen.GenerateSDKs(schema.Schema)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	err := run()
	if err != nil {
		fmt.Printf("ERROR: %v", err)
		os.Exit(1)
	}
}
