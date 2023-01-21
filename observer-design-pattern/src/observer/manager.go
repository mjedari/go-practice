package observer

import (
	"fmt"
	"observer/src/product"
)

type Manager struct {
	Name string
	Year string
}

func NewManager(name string, year string) *Manager {
	return &Manager{Name: name, Year: year}
}

func (m Manager) Update(product product.IProduct) {
	fmt.Println("Diling **** Diling *****")
	fmt.Printf("manager: \"%v\" received message for updating product: \"%v\" count remining \"%v\" \n",
		m.Name, product.GetName(), product.GetRemaining())
	fmt.Println("")
}
