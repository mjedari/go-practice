package rooms

import "fmt"

type BaseRoom struct {
	Name     string
	Receiver chan string
}

func NewBaseRoom(name string) *BaseRoom {
	return &BaseRoom{Name: name, Receiver: make(chan string)}
}

func (f BaseRoom) Broadcast() {
	fmt.Println("broadcasting from base room...")
}

func (f FilteredRoom) GetName() string {
	return f.Name
}
