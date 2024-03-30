package utils

import (
	"fmt"
	"strings"
)

func NewCodeComment(str string, indent int) string {
	parts := strings.Split(str, "\n")
	spaces := strings.Repeat(" ", indent)

	var result []string
	for i, part := range parts {
		// If the last line is empty don't include it.
		if i == (len(parts)-1) && part == "" {
			continue
		}

		result = append(result, fmt.Sprintf("%s// %s", spaces, part))
	}

	return strings.Join(result, "\n")
}
