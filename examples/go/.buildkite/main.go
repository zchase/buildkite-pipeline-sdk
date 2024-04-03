package main

import (
	"encoding/json"
	"fmt"
	"os"

	bk "github.com/zchase/buildkite-pipeline-sdk/sdk/go"
)

func run() error {
	// Create a new Buildkite Pipeline
	pipeline := bk.NewStepBuilder().
		AddCommandStep(&bk.CommandStepArgs{
			Commands: []string{
				"echo \"Hello World!\"",
			},
		})

	// Get the branch name of the current build
	branchName := bk.Environment.Branch()

	// Print out what branch we are on.
	if branchName == "main" {
		pipeline.AddCommandStep(&bk.CommandStepArgs{
			Commands: []string{
				`echo "I am on the main branch"`,
			},
		})
	} else {
		pipeline.AddCommandStep(&bk.CommandStepArgs{
			Commands: []string{
				fmt.Sprintf(`echo "I am on the %s branch"`, branchName),
			},
		})
	}

	str, err := json.Marshal(pipeline)
	if err != nil {
		return err
	}

	return os.WriteFile("pipeline.json", str, os.ModePerm)
}

func main() {
	err := run()
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		os.Exit(1)
	}
}
