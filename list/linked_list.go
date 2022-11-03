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

func (node *node[T]) isLast() bool {
	return node.next == nil
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

// Deletes and returns the las element of the list
// Complexity: O(n)
func (list *LinkedList[T]) DeleteLast() (value T, err error) {
	if list.head == nil {
		err = fmt.Errorf("could not pop from the list, it is empty")
		return
	}

	if list.head.isLast() {
		value = list.head.value
		list.head = nil
		return
	}

	current := list.head
	for !current.next.isLast() {
		current = current.next
	}

	value = current.next.value
	current.next = nil

	return
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
