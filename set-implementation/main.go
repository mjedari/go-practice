/*
 set Implementation in go
 you can define specific set by data type also
 you can use struct instead of bool in set
*/

package main

import (
	"fmt"
	"set-implementation/src/set"
)

func main() {

	mySet := set.NewSet()
	mySet.Add("salam")
	mySet.Add("salam2")
	mySet.Add("salam3")
	mySet.Add(12)

	fmt.Println("Here is existence check for salam3", mySet.Exists("salam3"))
	mySet.Remove("salam3")
	fmt.Println("Here is existence check for salam3", mySet.Exists("salam3"))
	fmt.Println("Here is existence check for 12", mySet.Exists(12))
	mySet.Remove(12)
	fmt.Println("Here is existence check for 12", mySet.Exists(12))
}
