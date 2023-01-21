package product

type IObserver interface {
	Update(subject IProduct)
}

type IProduct interface {
	Attach(observer IObserver)
	Notify()
	Sold(count uint) error

	GetName() string
	GetRemaining() uint
}
