package rooms

import (
	"fmt"
)

type FilteredRoom struct {
	*BaseRoom
}

func NewFilteredRoom(baseRoom *BaseRoom) *FilteredRoom {
	return &FilteredRoom{BaseRoom: baseRoom}
}

func (f FilteredRoom) Broadcast() {
	fmt.Println("broadcasting...")
}
