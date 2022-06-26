package tag

import (
	"strings"
)

type validationError struct {
	errorMessages []string
}

func (v *validationError) Error() string {
	return strings.Join(v.errorMessages, "; ")
}

func (v *validationError) append(str string) {
	v.errorMessages = append(v.errorMessages, str)
}

const (
	invalidTagErr = `invalid tag "%s", expected one of: %s`
)
