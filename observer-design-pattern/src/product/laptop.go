package product

import "errors"

type Laptop struct {
	Name      string
	Count     uint
	observers []IObserver
}

func NewLaptop(name string, count uint) *Laptop {
	return &Laptop{Name: name, Count: count}
}

func (p *Laptop) Attach(observer IObserver) {
	p.observers = append(p.observers, observer)
}

func (p *Laptop) Notify() {
	for _, observer := range p.observers {
		observer.Update(p)
	}
}

func (p *Laptop) GetName() string {
	return p.Name
}

func (p *Laptop) GetRemaining() uint {
	return p.Count
}

func (p *Laptop) Sold(count uint) error {
	if count > p.Count {
		return errors.New("can not sell")
	}

	p.Count -= count
	p.Notify()
	return nil
}
