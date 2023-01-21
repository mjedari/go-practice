package product

import "errors"

type Monitor struct {
	Name      string
	Size      uint
	Count     uint
	observers []IObserver
}

func NewMonitor(name string, size uint, count uint) *Monitor {
	return &Monitor{Name: name, Size: size, Count: count}
}

func (m *Monitor) Attach(observer IObserver) {
	m.observers = append(m.observers, observer)

}

func (m *Monitor) Notify() {
	for _, observer := range m.observers {
		observer.Update(m)
	}
}

func (m *Monitor) Sold(count uint) error {
	if count > m.Count {
		return errors.New("can not sell monitor")
	}

	m.Count -= count
	m.Notify()
	return nil
}

func (m *Monitor) GetName() string {
	return m.Name
}

func (m *Monitor) GetRemaining() uint {
	return m.Count
}
