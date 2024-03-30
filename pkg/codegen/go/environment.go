package gocodegen

import (
	"fmt"

	"github.com/zchase/buildkite-pipeline-sdk/pkg/schema"
	"github.com/zchase/buildkite-pipeline-sdk/pkg/utils"
)

var environmentFile = `// This file is auto-generated please do not edit

package buildkite

import (
	"os"
)

type environment struct{}`

func NewEnvironmentFile(envs map[schema.PropertyName]schema.EnvironmentVariable) string {
	file := utils.CodeBlock{environmentFile}

	methods := utils.CodeBlock{}
	for name, env := range envs {
		methodBody := utils.CodeBlock{
			fmt.Sprintf(`str := os.Getenv("%s")`, name),
		}

		var returnType string
		switch env.Type {
		case "boolean":
			returnType = "bool"
			methodBody = append(methodBody, "return ParseStringToBool(str)")
		case "number":
			returnType = "int"
			methodBody = append(methodBody, "return ParseStringToInt(str)")
		case "arrary":
			switch env.Items.Type {
			case "boolean":
				returnType = "bool"
				methodBody = append(methodBody, fmt.Sprintf("return ParseStringSliceToBoolSlice(strings.Split(str, \"%s\"))", env.Items.Delimiter))
			case "number":
				returnType = "int"
				methodBody = append(methodBody, fmt.Sprintf("return ParseStringSliceToIntSlice(strings.Split(str, \"%s\"))", env.Items.Delimiter))
			default:
				returnType = "string"
				methodBody = append(methodBody, fmt.Sprintf("return strings.Split(str, \"%s\")", env.Items.Delimiter))
			}
		default:
			returnType = "string"
			methodBody = append(methodBody, "return str")
		}

		methods = append(methods, fmt.Sprintf("%s\n%s\n%s\n%s\n",
			utils.NewCodeComment(env.Description, 0),
			fmt.Sprintf("func (e environment) %s() %s {", env.TitleCase(string(name)), returnType),
			methodBody.DisplayIndent(4),
			"}",
		))
	}

	file = append(file, methods.Display())

	return file.Display()
}
