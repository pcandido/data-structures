package main

import (
	"data_structures/list"
	"fmt"
)

func main() {
	var x list.List[int] = &list.SinglyLinkedList[int]{}

	x.InsertEnd(1)
	x.InsertEnd(2)

	fmt.Println(x.ToString())
}
