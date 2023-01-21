package main

import (
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
	}

	var keys []string
	for key, _ := range myMap {
		keys = append(keys, key)
	}

	iter.ForEach(keys, func(s *string) {
		fmt.Println("This is iteration over a map concurrently", *s)
	})

}
