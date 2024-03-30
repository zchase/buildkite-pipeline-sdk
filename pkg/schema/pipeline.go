package schema

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

const (
	PIPELINE_SCHEMA_FILE = "schema.yml"
)

type PipelineSchema struct {
	Name        string                               `yaml:"name"`
	Types       map[PropertyName]Property            `yaml:"types"`
	Environment map[PropertyName]EnvironmentVariable `yaml:"environment"`
	Steps       map[PropertyName]Step                `yaml:"steps"`
}

func (p PipelineSchema) lookupType(name string) Property {
	prop, ok := p.Types[PropertyName(name)]
	if !ok {
		panic(fmt.Errorf("type not found in schema: %s", name))
	}
	return prop
}

func (p PipelineSchema) TypeScriptType(name PropertyName, prop Property) string {
	switch prop.Type {
	case "string":
		return "string"
	case "number":
		return "number"
	case "boolean":
		return "boolean"
	case "array":
		switch prop.Items.Type {
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
		if prop.Properties != nil {
			return string(name)
		}

		if prop.Items.TypeRef != "" {
			refProp := p.lookupType(string(prop.Items.TypeRef))
			vType := p.TypeScriptType(PropertyName(prop.Items.TypeRef), refProp)
			return fmt.Sprintf("Record<string, %s>", vType)
		}

		switch prop.Items.Type {
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

func ReadSchema() (PipelineSchema, error) {
	schemaBytes, err := os.ReadFile("./schema.yml")
	if err != nil {
		return PipelineSchema{}, err
	}

	schema := PipelineSchema{}
	err = yaml.Unmarshal(schemaBytes, &schema)
	if err != nil {
		return PipelineSchema{}, err
	}

	return schema, nil
}
