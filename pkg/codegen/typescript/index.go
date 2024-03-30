package tscodegen

var indexFile = `// This file is auto-generated please do not edit

import Environment from "./environment";
import StepBuilder from "./stepBuilder";

export const environment = new Environment();
export const stepBuilder = new StepBuilder();`

func NewIndexFile() string {
	return indexFile
}
