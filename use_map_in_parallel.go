package main

import (
	"errors"
	"fmt"
	"github.com/sourcegraph/conc/iter"
)

func main() {
	myMap := map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
		"D": 4,
		"E": 5,
		"F": 0,
	}

	var values []int
	for _, value := range myMap {
		values = append(values, value)
	}

	doubledSlice := iter.Map(values, func(v *int) int {
		return (*v) * (*v)
	})
	fmt.Println("Doubled Map result: ", doubledSlice)

	tripledSliceWithErr, err := iter.MapErr(values, func(v *int) (int, error) {
		if *v <= 0 {
			return 0, errors.New("inputs are not valid")
		}
		return (*v) * (*v) * (*v), nil
	})

	if err != nil {
		panic(err)
	}

	fmt.Println("Tripled Map result: ", tripledSliceWithErr)
}
