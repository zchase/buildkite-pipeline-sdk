package schema

type Step struct {
	Description    string                    `yaml:"description"`
	Properties     map[PropertyName]Property `yaml:"properties"`
	EmptyArgsValue string                    `yaml:"emptyArgsValue"`
}
