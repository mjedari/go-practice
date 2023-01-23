package errs

import (
	"strings"
)

type ValidationError struct {
	Errors []error
}

func (v ValidationError) Error() string {
	// we should compact all errs into one string to pass through code
	var errors []string
	for _, err := range v.Errors {
		errors = append(errors, err.Error())
	}
	return strings.Join(errors, "; ")
}

func (v *ValidationError) Add(err error) {
	v.Errors = append(v.Errors, err)
}
