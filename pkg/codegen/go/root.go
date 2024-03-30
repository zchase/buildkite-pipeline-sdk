package gocodegen

var indexFile = `package buildkite

var Environment = environment{}
var StepBuilder = &stepBuilder{}`

func NewIndexFile() string {
	return indexFile
}
