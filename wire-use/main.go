package main

import (
	"fmt"
	"wire-use/wiring"
)

func main() {
	userService, err := wiring.InitializeServices()
	if err != nil {
		fmt.Println("Error initializing services:", err)
		return
	}

	fmt.Println("User service:", userService)
}
