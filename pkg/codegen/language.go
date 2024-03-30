package codegen

import "github.com/zchase/buildkite-pipeline-sdk/pkg/schema"

type LanguageTarget interface {
	FolderName() string
	Files(pipelineSchema schema.PipelineSchema) map[string]string
}
