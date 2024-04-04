package tscodegen

import (
	"fmt"
	"strings"

	"github.com/zchase/buildkite-pipeline-sdk/pkg/schema"
	"github.com/zchase/buildkite-pipeline-sdk/pkg/utils"
)

func NewInterfaceFromStep(name schema.PropertyName, step schema.Step, types schema.PropertyMap) (string, []string) {
	var typeImports []string
	tsInterface := utils.CodeBlock{
		fmt.Sprintf("interface %sArgs {", name.TitleCase()),
	}

	var properties []string
	for name, prop := range step.Properties {
		propType := prop.TypeScriptType(name, types)

		var ok bool
		if prop.TypeRef != "" {
			propType = prop.TypeRef.TitleCase()
			typeImports = append(typeImports, prop.TypeRef.TitleCase())

			prop, ok = types[schema.PropertyName(prop.TypeRef)]
			if !ok {
				panic(fmt.Errorf("type not found in schema %s", name))
			}
		}

		optionalMarker := ""
		if !prop.Required {
			optionalMarker = "?"
		}

		properties = append(properties, utils.CodeBlock{
			utils.NewCodeComment(prop.Description, 4),
			fmt.Sprintf("    %s%s: %s;\n", name.CamelCase(), optionalMarker, propType),
		}.Display())
	}

	tsInterface = append(tsInterface, properties...)
	tsInterface = append(tsInterface, "}")

	return tsInterface.Display(), typeImports
}

func newInterfaceFromSchemaType(name schema.PropertyName, prop schema.Property, types schema.PropertyMap) string {
	tsInterface := utils.CodeBlock{
		fmt.Sprintf("export interface %s {", name.TitleCase()),
	}

	var properties []string
	for n, p := range prop.Properties {
		propType := p.TypeScriptType(n, types)

		var ok bool
		if prop.TypeRef != "" {
			propType = prop.TypeRef.TitleCase()

			prop, ok = types[schema.PropertyName(prop.TypeRef)]
			if !ok {
				panic(fmt.Errorf("type not found in schema %s", name))
			}
		}

		optionalMarker := ""
		if !p.Required {
			optionalMarker = "?"
		}

		properties = append(properties, utils.CodeBlock{
			utils.NewCodeComment(p.Description, 0),
			fmt.Sprintf("%s%s: %s;", n.CamelCase(), optionalMarker, propType),
		}.DisplayIndent(4))
	}

	tsInterface = append(tsInterface, properties...)
	tsInterface = append(tsInterface, "}")

	return tsInterface.Display()
}

func newTypeScriptSimpleType(name, desc, typ string) string {
	return fmt.Sprintf("%s\nexport type %s = %s;\n", utils.NewCodeComment(desc, 0), name, typ)
}

func newTypeScriptRecordType(name, desc, typ string) string {
	return fmt.Sprintf("%s\nexport type %s = Record<string, %s>\n", utils.NewCodeComment(desc, 0), name, typ)
}

func newTypeScriptStringEnum(name, desc string, values []string) string {
	enumProps := utils.CodeBlock{}
	for _, val := range values {
		enumProps = append(enumProps, fmt.Sprintf("    %s = \"%s\",", strings.ToUpper(val), val))
	}

	return utils.CodeBlock{
		utils.NewCodeComment(desc, 0),
		fmt.Sprintf("export enum %s {", name),
		enumProps.Display(),
		"}\n",
	}.Display()
}

func NewTypesFile(types schema.PropertyMap) string {
	file := NewFile()

	for name, property := range types {
		switch property.Type {
		case "string":
			file.code = append(file.code, newTypeScriptSimpleType(name.TitleCase(), property.Description, "string"))
		case "boolean":
			file.code = append(file.code, newTypeScriptSimpleType(name.TitleCase(), property.Description, "boolean"))
		case "number":
			file.code = append(file.code, newTypeScriptSimpleType(name.TitleCase(), property.Description, "number"))
		case "enum":
			switch property.Items.Type {
			case "string":
				file.code = append(file.code, newTypeScriptStringEnum(name.TitleCase(), property.Description, property.Items.Values))
			default:
				panic("enums for non strings is not supported currently")
			}
		case "object":
			if property.Properties != nil {
				file.code = append(file.code, newInterfaceFromSchemaType(name, property, types))
				break
			}

			if property.Items.TypeRef != "" {
				file.code = append(file.code, newTypeScriptRecordType(name.TitleCase(), property.Description, property.Items.TypeRef.TitleCase()))
				break
			}

			switch property.Items.Type {
			case "string":
				file.code = append(file.code, newTypeScriptRecordType(name.TitleCase(), property.Description, "string"))
			case "boolean":
				file.code = append(file.code, newTypeScriptRecordType(name.TitleCase(), property.Description, "boolean"))
			case "number":
				file.code = append(file.code, newTypeScriptRecordType(name.TitleCase(), property.Description, "number"))
			}
		}
	}

	return file.String()
}
