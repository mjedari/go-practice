package src

import (
	"error-handling/src/errs"
	"errors"
	"reflect"
)

type ITypeValidator interface {
	Validate() error
}

type Name string

func (c Name) Validate() error {
	if len([]rune(c)) < 10 {
		return errors.New("name is not valid")
	}
	return nil
}

type Create struct {
	Name Name
}

func ValidateCommand(cmd interface{}) error {
	fields := reflect.TypeOf(cmd)
	values := reflect.ValueOf(cmd)
	var vError errs.ValidationError
	for i := 0; i < fields.NumField(); i++ {
		if validator, ok := values.Field(i).Interface().(ITypeValidator); ok {
			err := validator.Validate()
			if err != nil {
				vError.Errors = append(vError.Errors, err)
			}
		}
	}
	return vError
}
