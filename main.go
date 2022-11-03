package main

import (
	"data_structures/list"
	"fmt"
)

func main() {
	var x list.List[int] = &list.LinkedList[int]{}

	x.Push(1)
	x.Push(2)

	fmt.Println(x.ToString())
}
