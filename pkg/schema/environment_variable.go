package schema

import (
	"strings"

	"github.com/zchase/buildkite-pipeline-sdk/pkg/utils"
)

type EnvironmentVariable struct {
	Type        string `yaml:"type"`
	Description string `yaml:"description"`
	Items       struct {
		Type      string `yaml:"type"`
		Delimiter string `yaml:"delimiter"`
	} `yaml:"items"`
}

func (e EnvironmentVariable) TitleCase(name string) string {
	cleanName := strings.Replace(name, "BUILDKITE_", "", 1)
	return utils.SnakeCaseToTitleCase(cleanName)
}

func (e EnvironmentVariable) CamelCase(name string) string {
	cleanName := strings.Replace(name, "BUILDKITE_", "", 1)
	return utils.SnakeCaseToCamelCase(cleanName)
}
