package observer

import (
	"fmt"
	"observer/src/product"
)

type SupplyManager struct {
	Name  string
	Phone string
}

func NewSupplyManager(name string, phone string) *SupplyManager {
	return &SupplyManager{Name: name, Phone: phone}
}

func (s SupplyManager) Update(product product.IProduct) {
	fmt.Println("Diling **** Diling *****")
	fmt.Printf("supply manager: \"%v\" received message for updating product: \"%v\" count remining \"%v\" \n",
		s.Name, product.GetName(), product.GetRemaining())
	fmt.Println("")
}
