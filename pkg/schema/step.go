package schema

type Step struct {
	Description    string                    `yaml:"description"`
	Properties     map[PropertyName]Property `yaml:"properties"`
	EmptyArgsValue string                    `yaml:"emptyArgsValue"`
}

func NewStepMap(stepMap map[string]Step) map[PropertyName]Step {
	result := make(map[PropertyName]Step, len(stepMap))
	for key, step := range stepMap {
		result[PropertyName(key)] = step
	}
	return result
}
