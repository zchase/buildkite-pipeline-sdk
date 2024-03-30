package codegen

import (
	"fmt"
	"os"
	"path"

	gocodegen "github.com/zchase/buildkite-pipeline-sdk/pkg/codegen/go"
	tscodegen "github.com/zchase/buildkite-pipeline-sdk/pkg/codegen/typescript"
	"github.com/zchase/buildkite-pipeline-sdk/pkg/schema"
)

const (
	SDK_FOLDER = "sdk"
)

type TypeScriptSDK struct{}

func (t TypeScriptSDK) FolderName() string {
	return "typescript"
}

func (TypeScriptSDK) Files(pipelineSchema schema.PipelineSchema) map[string]string {
	return map[string]string{
		"environment.ts": tscodegen.NewEnvironmentFile(pipelineSchema.Environment),
		"index.ts":       tscodegen.NewIndexFile(),
		"package.json":   tscodegen.NewPackageJSONFile(),
		"stepBuilder.ts": tscodegen.NewStepBuilderFile(pipelineSchema),
		"types.ts":       tscodegen.NewTypesFile(pipelineSchema.Types),
	}
}

type GoSDK struct{}

func (GoSDK) FolderName() string {
	return "go"
}

func (GoSDK) Files(pipelineSchema schema.PipelineSchema) map[string]string {
	return map[string]string{
		"root.go":         gocodegen.NewIndexFile(),
		"enviornment.go":  gocodegen.NewEnvironmentFile(pipelineSchema.Environment),
		"step_builder.go": gocodegen.NewStepBuilderFile(pipelineSchema),
		"types.go":        gocodegen.NewTypesFile(pipelineSchema.Types),
		"utils.go":        gocodegen.NewUtilsFiles(),
	}
}

var targets = []LanguageTarget{
	TypeScriptSDK{},
	GoSDK{},
}

func GenerateSDKs(pipelineSchema schema.PipelineSchema) error {
	err := os.Mkdir(SDK_FOLDER, os.ModePerm)
	if err != nil {
		return fmt.Errorf("creating sdk folder: %v", err)
	}

	for _, target := range targets {
		targetFolder := path.Join(SDK_FOLDER, target.FolderName())
		err = os.Mkdir(targetFolder, os.ModePerm)
		if err != nil {
			return err
		}

		files := target.Files(pipelineSchema)
		for name, contents := range files {
			out := path.Join(targetFolder, name)
			err = os.WriteFile(out, []byte(contents), os.ModePerm)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
