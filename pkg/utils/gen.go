package utils

import (
	"fmt"
	"strings"
)

// A CodeBlock contains lines of generated code.
type CodeBlock []string

func (c CodeBlock) Display() string {
	return strings.Join(c, "\n")
}

func (c CodeBlock) DisplayIndent(spaces int) string {
	indent := strings.Repeat(" ", spaces)
	return fmt.Sprintf("%s%s",
		indent,
		strings.Join(c, fmt.Sprintf("\n%s", indent)),
	)
}
