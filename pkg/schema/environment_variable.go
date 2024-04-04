package schema

import (
	"strings"

	"github.com/zchase/buildkite-pipeline-sdk/pkg/utils"
)

func NewEnvironmentVariableMap(envVars map[string]EnvironmentVariable) map[PropertyName]EnvironmentVariable {
	newMap := make(map[PropertyName]EnvironmentVariable, len(envVars))
	for key, envVar := range envVars {
		newMap[PropertyName(key)] = envVar
	}

	return newMap
}

type EnvironmentVariable struct {
	Type        SchemaType          `yaml:"type"`
	Description string              `yaml:"description"`
	Items       ObjectArrayProperty `yaml:"items"`

	// There are some dynamic environment variables created from tags.
	// To support those we allow users to supply args to the env var function.
	Dynamic bool
}

func (e EnvironmentVariable) TitleCase(name string) string {
	cleanName := strings.Replace(name, "BUILDKITE_", "", 1)
	return utils.SnakeCaseToTitleCase(cleanName)
}

func (e EnvironmentVariable) CamelCase(name string) string {
	cleanName := strings.Replace(name, "BUILDKITE_", "", 1)
	return utils.SnakeCaseToCamelCase(cleanName)
}
