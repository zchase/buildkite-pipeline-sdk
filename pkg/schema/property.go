package schema

import (
	"fmt"

	"github.com/zchase/buildkite-pipeline-sdk/pkg/utils"
)

// new playground
type SchemaType string

const (
	schemaString  SchemaType = "string"
	schemaNumber  SchemaType = "number"
	schemaBoolean SchemaType = "boolean"
	schemaArray   SchemaType = "array"
	schemaObject  SchemaType = "object"
	schemaEnum    SchemaType = "enum"
)

// end playground

type PropertyName string

func (p PropertyName) TitleCase() string {
	return utils.SnakeCaseToTitleCase(string(p))
}

func (p PropertyName) CamelCase() string {
	return utils.SnakeCaseToCamelCase(string(p))
}

type ObjectArrayProperty struct {
	TypeRef   PropertyName `yaml:"$types"`
	Type      SchemaType   `yaml:"type"`
	Delimiter string       `yaml:"delimiter"`
	Values    []string     `yaml:"values"`

	Items struct {
		Type string `yaml:"type"`
	} `yaml:"items"`
}

type PropertyMap map[PropertyName]Property

func NewPropertyMap(propertyMap map[string]Property) map[PropertyName]Property {
	result := make(map[PropertyName]Property, len(propertyMap))
	for key, prop := range propertyMap {
		result[PropertyName(key)] = prop
	}
	return result
}

type Property struct {
	Description string                    `yaml:"description"`
	Type        SchemaType                `yaml:"type"`
	Items       ObjectArrayProperty       `yaml:"items"`
	Properties  map[PropertyName]Property `yaml:"properties"`
	Required    bool                      `yaml:"required"`

	TypeRef PropertyName `yaml:"$types"`
}

func (p Property) GoType() string {
	switch p.Type {
	case "string":
		return "string"
	case "number":
		return "int"
	case "boolean":
		return "bool"
	case "array":
		switch p.Items.Type {
		case "string":
			return "[]string"
		case "number":
			return "[]int"
		case "boolean":
			return "[]bool"
		default:
			return "[]interface{}"
		}
	case "object":
		switch p.Items.Type {
		case "string":
			return "map[string]string"
		case "number":
			return "map[string]int"
		case "boolean":
			return "map[string]bool"
		default:
			return "map[string]interface{}"
		}
	default:
		return "interface{}"
	}
}

func (p Property) TypeScriptType(name PropertyName, types PropertyMap) string {
	switch p.Type {
	case "string":
		return "string"
	case "number":
		return "number"
	case "boolean":
		return "boolean"
	case "array":
		switch p.Items.Type {
		case "string":
			return "string[]"
		case "number":
			return "number[]"
		case "boolean":
			return "boolean[]"
		default:
			return "any[]"
		}
	case "object":
		if p.Properties != nil {
			return string(name)
		}

		if p.Items.TypeRef != "" {
			fmt.Println(p)
			return fmt.Sprintf("Record<string, %s>", string(name))
		}

		switch p.Items.Type {
		case "string":
			return "Record<string, string>"
		case "number":
			return "Reacord<string, number>"
		case "boolean":
			return "Record<string, boolean>"
		default:
			return "Record<string, any>"
		}
	default:
		return "any"
	}
}
