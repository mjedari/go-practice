package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Starting program...")
	numbers := []uint{1, 20, 30, 43, 6, 6757, 2, 3, 35, 657, 78, 90}
	var convertedNumbers []int

	for _, number := range numbers {
		convertedNumbers = append(convertedNumbers, int(number))
	}
	sort.Ints(convertedNumbers)
	ok := numbers[2]
	fmt.Println("sorted lists is: ", numbers, ok)
}
