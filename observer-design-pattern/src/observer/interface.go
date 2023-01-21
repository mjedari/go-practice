package observer

import (
	"observer/src/product"
)

type IObserver interface {
	Update(subject product.IProduct)
}
