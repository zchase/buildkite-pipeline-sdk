# Buildkite Pipeline SDK

This is an experimental approach to generating SDKs for defining [Buildkite Dynamic Pipelines](https://buildkite.com/docs/pipelines/defining-steps#dynamic-pipelines) from a YAML schema.

The purpose of this repo is to demonstrate the technique and provide simple examples. You should not use the code in this repo for any sort of production use.

## Usage

The SDKs are intended to be used locally but you __could__ use the Go SDK on any machine since the repo is public. The TypeScript isn't published to NPM so it can only be used after building/install locally.

### Build and Install

To build and install the SDKs, you just need to run the build and install script:

```
$ ./scripts/build_and_install.sh
```

### Try it out

To try out the SDKs you can create a new project and use the files in the `examples` directory as a guide.

## TypeScript Example

```typescript
import * as bk from "buildkite-pipline-sdk";

let pipeline = bk.stepBuilder
    .addCommandStep({
        commands: [ "echo \"Hello World!\"" ],
    });

const branchName = bk.environment.branch();
if (branchName === "main") {
    pipeline = pipeline.addCommandStep({
        commands: [ `echo "I am on the main branch"` ],
    })
} else {
    pipeline = pipeline.addCommandStep({
        commands: [ `echo "This isn't the main its the ${branchName} branch"` ],
    })
}

pipeline.write();
```

## Go Example

```go
package main

import (
	"fmt"

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

	return pipeline.Print()
}

func main() {
	err := run()
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		os.Exit(1)
	}
}
```