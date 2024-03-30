package tscodegen

import (
	"fmt"

	"github.com/zchase/buildkite-pipeline-sdk/pkg/utils"
)

// Imports
type ImportBlock struct {
	imports map[string]utils.CodeBlock
}

func (i *ImportBlock) AddImports(pkg string, names ...string) {
	for _, name := range names {
		i.imports[pkg] = append(i.imports[pkg], fmt.Sprintf("    %s,", name))
	}
}

func (i *ImportBlock) String() string {
	block := utils.CodeBlock{}

	for pkg, imports := range i.imports {
		block = append(block, fmt.Sprintf("%s\n%s\n%s\n",
			"import {",
			imports.Display(),
			fmt.Sprintf("} from \"%s\";", pkg),
		))
	}

	return block.Display()
}

// File
type File struct {
	header  string
	imports *ImportBlock
	code    utils.CodeBlock
}

func (f *File) String() string {
	file := utils.CodeBlock{
		f.header,
		f.imports.String(),
		f.code.Display(),
	}

	return file.Display()
}

func NewFile() *File {
	return &File{
		header: "// This file is auto-generated please do not edit",
		imports: &ImportBlock{
			imports: make(map[string]utils.CodeBlock),
		},
	}
}
