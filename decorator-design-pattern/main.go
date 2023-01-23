package main

import (
	"decorator/src/services"
	"fmt"
	"github.com/fatih/color"
)

type myDecoratedFunction struct {
	Id      uint
	Name    string
	Friends []string
}

func (m myDecoratedFunction) Run() {
	color.New(color.Italic, color.FgMagenta, color.Bold).Println("Hi, I'm the original function!")
	fmt.Printf("You can pass data to this function by populating its struct."+
		"Now you can see the function name which is: %v \nNow these are %v's frinds: \n", m.Name, m.Name)

	for _, name := range m.Friends {
		fmt.Println(name)
	}
	print("Thank you from function Id: ")
	color.New(color.Italic, color.FgMagenta, color.Bold).Println(m.Id)
	println("bye...")

}

func main() {
	fmt.Println("Program started...")
	myFunc := myDecoratedFunction{
		Id:      112,
		Name:    "Elen",
		Friends: []string{"Peter", "David"},
	}

	logger := services.NewLogger()
	metrics := services.NewMetrics()

	decoratedWithLogger := logger.Decorate(myFunc)
	decoratedWithMetrics := metrics.Decorate(decoratedWithLogger)

	decoratedWithMetrics.Run()

}
