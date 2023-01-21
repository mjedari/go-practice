package main

import (
	"fmt"
	"github.com/sourcegraph/conc"
	"time"
)

func main() {
	p := conc.PanicCatcher{}

	go func() {
		p.Try(func() {
			time.Sleep(time.Second * 5)
			panic("this is a panic error")
		})
	}()

	time.Sleep(2 * time.Second)
	err := p.Recovered()
	fmt.Println(err.Value)
}
