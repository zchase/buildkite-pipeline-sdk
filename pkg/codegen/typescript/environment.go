package tscodegen

import (
	"fmt"

	"github.com/zchase/buildkite-pipeline-sdk/pkg/schema"
	"github.com/zchase/buildkite-pipeline-sdk/pkg/utils"
)

func NewEnvironmentFile(envs map[schema.PropertyName]schema.EnvironmentVariable) string {
	file := NewFile()
	file.code = append(file.code, "class Environment {")

	for name, env := range envs {
		var returnStatement string
		returnType := env.Type
		switch env.Type {
		case "boolean":
			returnStatement = fmt.Sprintf("return Boolean(process.env.%s)", name)
		case "number":
			returnStatement = fmt.Sprintf("return Number(process.env.%s)", name)
		case "array":
			switch env.Items.Type {
			case "boolean":
				returnType = "boolean[]"
				returnStatement = fmt.Sprintf("return process.env.%s.split(\"%s\").map(v => Boolean(v))", name, env.Items.Delimiter)
			case "number":
				returnType = "number[]"
				returnStatement = fmt.Sprintf("return process.env.%s.split(\"%s\").map(v => Number(v))", name, env.Items.Delimiter)
			default:
				returnType = "string[]"
				returnStatement = fmt.Sprintf("return process.env.%s.split(\"%s\")", name, env.Items.Delimiter)
			}
		default:
			returnStatement = fmt.Sprintf("return process.env.%s", name)
		}

		block := utils.CodeBlock{
			utils.NewCodeComment(env.Description, 0),
			fmt.Sprintf("public %s(): %s {", env.CamelCase(string(name)), returnType),
			fmt.Sprintf("    %s", returnStatement),
			"}",
		}

		file.code = append(file.code, block.DisplayIndent(4))
	}

	file.code = append(file.code, "}")
	file.code = append(file.code, "\n\nexport default Environment;")

	return file.String()
}
