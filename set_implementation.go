/*
 set Implementation in go
 you can define specific set by data type also
 you can use struct instead of bool in set
*/

package main

import "fmt"

type ISet interface {
	add(key interface{})
	remove(key interface{})
	exists(key interface{}) bool
	size() int
}
type MySet map[interface{}]bool

func NewSet() ISet {
	return &MySet{}
}

func (s MySet) add(key interface{}) {
	s[key] = true
}

func (s MySet) remove(key interface{}) {
	delete(s, key)
}

func (s MySet) exists(key interface{}) bool {
	return s[key]
}

func (s MySet) size() int {
	return len(s)
}

func main() {

	mySet := NewSet()
	mySet.add("salam")
	mySet.add("salam2")
	mySet.add("salam3")
	mySet.add(12)

	fmt.Println("Here is existence check for salam3", mySet.exists("salam3"))
	mySet.remove("salam3")
	fmt.Println("Here is existence check for salam3", mySet.exists("salam3"))
	fmt.Println("Here is existence check for 12", mySet.exists(12))
	mySet.remove(12)
	fmt.Println("Here is existence check for 12", mySet.exists(12))
}
