package main

import (
	"error-handling/src"
	"error-handling/src/errs"
	"fmt"
)

func main() {
	var vErr errs.ValidationError

	if err := src.DoSomeThing(); err != nil {
		vErr.Add(err)
	}

	if err := src.DoSomeThingTwo(); err != nil {
		vErr.Add(err)
	}

	if err := src.DoSomeThingWithName("Mahdi"); err != nil {
		vErr.Add(err)
	}

	fmt.Printf("All errs: %+v \n\n", vErr)
	for _, err := range vErr.Errors {
		fmt.Println(err)
	}
}
