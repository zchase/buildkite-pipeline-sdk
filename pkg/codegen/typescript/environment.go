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
		envKey := fmt.Sprintf("process.env.%s", name)
		if env.Dynamic {
			envKey = "process.env[strs.join(\"_\").toUpperCase()]"
		}

		var returnStatement string
		returnType := env.Type
		switch env.Type {
		case "boolean":
			returnStatement = fmt.Sprintf("return Boolean(%s)", envKey)
		case "number":
			returnStatement = fmt.Sprintf("return Number(%s)", envKey)
		case "array":
			switch env.Items.Type {
			case "boolean":
				returnType = "boolean[]"
				returnStatement = fmt.Sprintf("return %s.split(\"%s\").map(v => Boolean(v))", envKey, env.Items.Delimiter)
			case "number":
				returnType = "number[]"
				returnStatement = fmt.Sprintf("return %s.split(\"%s\").map(v => Number(v))", envKey, env.Items.Delimiter)
			default:
				returnType = "string[]"
				returnStatement = fmt.Sprintf("return %s.split(\"%s\")", envKey, env.Items.Delimiter)
			}
		default:
			returnStatement = fmt.Sprintf("return %s", envKey)
		}

		dynamicArgs := ""
		if env.Dynamic {
			dynamicArgs = "...strs: string[]"
		}

		block := utils.CodeBlock{
			utils.NewCodeComment(env.Description, 0),
			fmt.Sprintf("public %s(%s): %s {", env.CamelCase(string(name)), dynamicArgs, returnType),
			fmt.Sprintf("    %s;", returnStatement),
			"}",
		}

		file.code = append(file.code, block.DisplayIndent(4))
	}

	file.code = append(file.code, "}")
	file.code = append(file.code, "\n\nexport default Environment;")

	return file.String()
}
