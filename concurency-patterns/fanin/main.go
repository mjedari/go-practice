package main

import (
	"fanin/app"
	"fmt"
)

const NumberOfSensors = 3

func main() {
	sources := make([]<-chan string, 0)

	// just create 3 instance of identical sensors and collect them in a slice
	for i := 0; i < NumberOfSensors; i++ {
		// create an instance
		name := fmt.Sprintf("Sensor %v", i)
		sensor := app.NewSensor(name)

		// start reading data from sensor
		sensorOutputCh := sensor.Read()

		sources = append(sources, sensorOutputCh)
	}

	outputCh := app.Funnel(sources...)

	for {
		msg := <-outputCh
		fmt.Println("Out put is:", msg)
	}

}
