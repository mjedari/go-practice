package main

import (
	"log"
	"observer/src/observer"
	"observer/src/product"
)

func main() {
	// create a laptop
	laptop := product.NewLaptop("Mac Book M1 pro", 9)
	monitor := product.NewMonitor("Dell", 24, 5)

	// create an observer
	supplyManagerObserver := observer.NewSupplyManager("Edd Das", "983 434 6565")
	customerObserver := observer.NewCustomer("Emily Ban", "983 564 3212")

	// attach observer to laptop and monitor
	laptop.Attach(supplyManagerObserver)
	laptop.Attach(customerObserver)
	monitor.Attach(supplyManagerObserver)

	// sell 5 Mac Book pro
	if err := laptop.Sold(5); err != nil {
		log.Fatal(err)
	}

	// sell 2 Dell monitor
	if err := monitor.Sold(5); err != nil {
		log.Fatal(err)
	}

}
