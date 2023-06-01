package app

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type Sensor struct {
	Name string
}

func NewSensor(name string) *Sensor {
	return &Sensor{Name: name}
}

func (s *Sensor) Read() <-chan string {
	fmt.Println("start reading from sensor: ", s.Name)
	output := make(chan string) // be careful fo close the channel

	go func() {
		// make sure you close every channel you've made after finishing its job
		defer close(output)

		for {
			randomNumber := rand.Int()
			output <- fmt.Sprintf("%v: %v", s.Name, strconv.Itoa(randomNumber))
			time.Sleep(time.Millisecond * 100)
		}
	}()

	return output
}
