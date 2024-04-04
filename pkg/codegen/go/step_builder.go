package gocodegen

import (
	"fmt"

	"github.com/zchase/buildkite-pipeline-sdk/pkg/schema"
)

func newStepBuilderMethod(name schema.PropertyName, step schema.Step) string {
	return fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n",
		fmt.Sprintf("func (s *stepBuilder) Add%s(step *%sArgs) *stepBuilder {", name.TitleCase(), name.TitleCase()),
		"    if step == nil {",
		fmt.Sprintf("        s.Steps = append(s.Steps, \"%s\")", step.EmptyArgsValue),
		"        return s",
		"    }",
		"    s.Steps = append(s.Steps, step)",
		"    return s",
		"}",
	)
}

func NewStepBuilderFile(pipelineSchema schema.PipelineSchema) string {
	file := NewFile()
	file.imports.AddImports("encoding/json", "os")
	file.code = append(file.code, fmt.Sprintf(`type stepBuilder struct {
	Steps []interface{} %s
}

func NewStepBuilder() *stepBuilder {
	return &stepBuilder{}
}`, "`json:\"steps\"`"))

	for name, step := range pipelineSchema.Steps {
		goStruct := NewStructFromStep(name, step, pipelineSchema.Types)
		file.code = append(file.code, goStruct.Display())
		file.code = append(file.code, newStepBuilderMethod(name, step))
	}

	file.code = append(file.code, fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n",
		"func (s *stepBuilder) Print() error {",
		`	jsonBytes, err := json.MarshalIndent(s, "", "\t")`,
		"	if err != nil {",
		"	    return err",
		"	}",
		"",
		"    return os.WriteFile(\"pipeline.json\", jsonBytes, os.ModePerm)",
		"}",
	))

	return file.String()
}
