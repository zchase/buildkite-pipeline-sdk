package tscodegen

import (
	"fmt"

	"github.com/zchase/buildkite-pipeline-sdk/pkg/schema"
	"github.com/zchase/buildkite-pipeline-sdk/pkg/utils"
)

var stepBuilderCode = `%s

class StepBuilder {
	private steps: any[] = [];

	public write() {
		console.log(JSON.stringify({ steps: this.steps }, null, 4));
	}`

func NewStepBuilderFile(pipelineSchema schema.PipelineSchema) string {
	file := NewFile()
	file.code = append(file.code, stepBuilderCode)
	interfaces := utils.CodeBlock{}

	for name, step := range pipelineSchema.Steps {
		stepInterface, typeImports := NewInterfaceFromStep(name, step, pipelineSchema.Types)
		interfaces = append(interfaces, fmt.Sprintf("%s\n", stepInterface))
		file.imports.AddImports("./types", typeImports...)

		optionalArgsMarker := ""
		if step.EmptyArgsValue != "" {
			optionalArgsMarker = "?"
		}

		file.code = append(file.code, utils.CodeBlock{
			utils.NewCodeComment(step.Description, 0),
			fmt.Sprintf("public add%s(args%s: %sArgs): this {", name.TitleCase(), optionalArgsMarker, name.TitleCase()),
			"    if (args === undefined) {",
			fmt.Sprintf("        this.steps.push(\"%s\")", step.EmptyArgsValue),
			"    }",
			"    this.steps.push({ ...args });",
			"    return this;",
			"}",
		}.DisplayIndent(4))
	}

	file.code[0] = fmt.Sprintf(file.code[0], interfaces.Display())

	return fmt.Sprintf("%s\n}\nexport default StepBuilder;", file.String())
}
