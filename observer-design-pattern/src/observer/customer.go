package observer

import (
	"fmt"
	"observer/src/product"
)

type Customer struct {
	Name    string
	Country string
}

func NewCustomer(name string, country string) *Customer {
	return &Customer{Name: name, Country: country}
}

func (c Customer) Update(product product.IProduct) {
	fmt.Println("Diling **** Diling *****")
	fmt.Printf("Customer: \"%v\" received message for updating product: \"%v\" count remining \"%v\" \n",
		c.Name, product.GetName(), product.GetRemaining())
	fmt.Println("")
}
