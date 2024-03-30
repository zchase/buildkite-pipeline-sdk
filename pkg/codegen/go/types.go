package gocodegen

import (
	"fmt"
	"strings"

	"github.com/zchase/buildkite-pipeline-sdk/pkg/schema"
	"github.com/zchase/buildkite-pipeline-sdk/pkg/utils"
)

func newStructProperties(props, types schema.PropertyMap) utils.CodeBlock {
	var properties []string
	for name, prop := range props {
		propType := prop.GoType()

		var ok bool
		if prop.TypeRef != "" {
			propType = prop.TypeRef.TitleCase()

			prop, ok = types[prop.TypeRef]
			if !ok {
				panic(fmt.Errorf("type not found in schema %s", name))
			}
		}

		properties = append(properties, fmt.Sprintf("%s\n%s\n",
			utils.NewCodeComment(prop.Description, 4),
			fmt.Sprintf("    %s %s `json:\"%s,omitempty\"`", name.TitleCase(), propType, name),
		))
	}

	return properties
}

func NewStructFromStep(name schema.PropertyName, step schema.Step, types schema.PropertyMap) utils.CodeBlock {
	properties := newStructProperties(step.Properties, types)

	return utils.CodeBlock{
		fmt.Sprintf("type %sArgs struct {", name.TitleCase()),
		strings.Join(properties, "\n"),
		"}",
	}
}

func newStructFromProperty(name schema.PropertyName, prop schema.Property, types schema.PropertyMap) utils.CodeBlock {
	properties := newStructProperties(prop.Properties, types)

	return utils.CodeBlock{
		fmt.Sprintf("type %s struct {", name.TitleCase()),
		strings.Join(properties, "\n"),
		"}",
	}
}

func newGoSimpleType(name, desc, typ string) string {
	return fmt.Sprintf("%s\ntype %s %s\n", utils.NewCodeComment(desc, 0), name, typ)
}

func newGoRecordType(name, desc, typ string) string {
	return fmt.Sprintf("%s\ntype %s map[string]%s\n", utils.NewCodeComment(desc, 0), name, typ)
}

func newGoStringEnum(name, desc string, values []string) string {
	consts := utils.CodeBlock{}
	for _, val := range values {
		consts = append(consts, fmt.Sprintf("    %s %s = \"%s\"", strings.ToUpper(val), name, val))
	}

	return utils.CodeBlock{
		utils.NewCodeComment(desc, 0),
		fmt.Sprintf("type %s string", name),
		"const (",
		consts.Display(),
		")",
	}.Display()
}

func NewTypesFile(types schema.PropertyMap) string {
	file := NewFile()

	for name, property := range types {
		switch property.Type {
		case "string":
			file.code = append(file.code, newGoSimpleType(name.TitleCase(), property.Description, "string"))
		case "boolean":
			file.code = append(file.code, newGoSimpleType(name.TitleCase(), property.Description, "bool"))
		case "number":
			file.code = append(file.code, newGoSimpleType(name.TitleCase(), property.Description, "int"))
		case "enum":
			switch property.Items.Type {
			case "string":
				file.code = append(file.code, newGoStringEnum(name.TitleCase(), property.Description, property.Items.Values))
			default:
				panic("enums for non strings is not supported currently")
			}
		case "object":
			if property.Properties != nil {
				file.code = append(file.code, newStructFromProperty(name, property, types).Display())
				break
			}

			if property.Items.TypeRef != "" {
				file.code = append(file.code, newGoRecordType(name.TitleCase(), property.Description, property.Items.TypeRef.TitleCase()))
				break
			}

			switch property.Items.Type {
			case "string":
				file.code = append(file.code, newGoRecordType(name.TitleCase(), property.Description, "string"))
			case "boolean":
				file.code = append(file.code, newGoRecordType(name.TitleCase(), property.Description, "bool"))
			case "number":
				file.code = append(file.code, newGoRecordType(name.TitleCase(), property.Description, "int"))
			}
		}
	}

	return file.String()
}
