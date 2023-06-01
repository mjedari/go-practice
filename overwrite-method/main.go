package main

import (
	"fmt"
	"overwrite/rooms"
)

type IRoom interface {
	Broadcast()
}

func main() {

	newRoom := rooms.NewBaseRoom("this")

	newFilteredRoom := rooms.NewFilteredRoom(newRoom)

	newFilteredRoom.Broadcast()

	fmt.Println(newFilteredRoom.GetName())

}
