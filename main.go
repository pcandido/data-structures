package main

import (
	"data_structures/list"
	"fmt"
)

func main() {
	var x list.List[int] = &list.LinkedList[int]{}

	x.AddAfter(1)
	x.AddAfter(2)

	fmt.Println(x.ToString())
}
