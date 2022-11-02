package list

import (
	"bytes"
	"fmt"
)

type node[T any] struct {
	value T
	next  *node[T]
}

type LinkedList[T any] struct {
	head *node[T]
}

// Adds an element at the final of the list
// Complexity: O(n)
func (list *LinkedList[T]) AddAfter(value T) {
	newNode := &node[T]{value, nil}

	if list.head == nil {
		list.head = newNode
		return
	}

	current := list.head
	for current.next != nil {
		current = current.next
	}

	current.next = newNode
}

// Generates a string that represents the whole list
// Complexity: O(n)
func (list LinkedList[T]) ToString() string {
	var buffer bytes.Buffer

	head := list.head
	first := true

	for head != nil {
		if first {
			first = false
		} else {
			buffer.WriteString(", ")
		}

		buffer.WriteString(fmt.Sprint(head.value))
		head = head.next
	}

	return fmt.Sprintf("[ %s ]", buffer.String())
}
