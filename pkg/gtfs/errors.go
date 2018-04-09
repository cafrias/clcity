package gtfs

import (
	"bytes"
	"fmt"
)

// ErrValidation represents a validation error.
type ErrValidation struct {
	File   string
	Fields map[string]string
}

func (e *ErrValidation) Error() string {
	buf := bytes.NewBufferString(fmt.Sprintf("Validation Error on file '%s': \n", e.File))
	for key, value := range e.Fields {
		buf.WriteString(fmt.Sprintf("    - Invalid value '%s' for field '%s'\n", value, key))
	}
	return buf.String()
}
