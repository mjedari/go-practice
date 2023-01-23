package src

import "errors"

func DoSomeThing() error {
	return errors.New("some thing went wrong")
}

func DoSomeThingTwo() error {
	return errors.New("some thing went wrong 2")
}

func DoSomeThingWithName(name Name) error {
	if err := name.Validate(); err != nil {
		return err
	}
	return nil
}
