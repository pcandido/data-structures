package list

import (
	"bytes"
	"fmt"
)

type circularLinkedListNode[T comparable] struct {
	value T
	next  *circularLinkedListNode[T]
}

type CircularLinkedList[T comparable] struct {
	tail *circularLinkedListNode[T]
}

// depracated
func (node *circularLinkedListNode[T]) isLast() bool {
	return node.next == nil
}

// Adds an element at the beginning of the list
// Complexity: O(1)
func (list *CircularLinkedList[T]) InsertStart(value T) {
	if list.tail == nil {
		node := &circularLinkedListNode[T]{value, nil}
		node.next = node
		list.tail = node
		return
	}

	list.tail.next = &circularLinkedListNode[T]{value, list.tail.next}
}

// Adds an element at the final of the list
// Complexity: O(1)
func (list *CircularLinkedList[T]) InsertEnd(value T) {
	if list.tail == nil {
		node := &circularLinkedListNode[T]{value, nil}
		node.next = node
		list.tail = node
		return
	}

	node := &circularLinkedListNode[T]{value, list.tail.next}
	list.tail.next = node
	list.tail = node
}

// Deletes and returns the first element of the list
// Complexity: O(1)
func (list *CircularLinkedList[T]) RemoveStart() (value T, err error) {
	if list.tail == nil {
		err = fmt.Errorf("could not delete from an empty list")
		return
	}

	toRemove := list.tail.next
	value = toRemove.value

	if toRemove==list.tail{
		list.tail=nil
		return
	}

	list.tail.next = toRemove.next
	return
}

// Deletes and returns the last element of the list
// Complexity: O(n)
func (list *CircularLinkedList[T]) RemoveEnd() (value T, err error) {
	if list.tail == nil {
		err = fmt.Errorf("could not delete from an empty list")
		return
	}

	value = list.tail.value

	if list.tail == list.tail.next {
		list.tail = nil
		return
	}

	beforeLast := list.tail
	for beforeLast.next != list.tail{
		beforeLast =beforeLast.next
	}

	beforeLast.next = list.tail.next
	list.tail = beforeLast
	return
}

// Find and element on the list and returns its index
// Complexity: O(n)
func (list CircularLinkedList[T]) Find(value T) (index int, err error) {
	if list.tail == nil {
		err = fmt.Errorf("element not found")
		return
	}
	
	currentNode := list.tail
	currentIndex := -1

	for  {
		currentNode=currentNode.next
		currentIndex = currentIndex + 1

		if currentNode.value == value {
			index = currentIndex
			return
		}

		if currentNode == list.tail {
			break
		}
	}

	err = fmt.Errorf("element not found")
	return
}

// Returns the value of the first element of the list
// Complexity: O(1)
func (list CircularLinkedList[T]) First() (value T, err error) {
	if list.tail == nil {
		err = fmt.Errorf("could not get first from an empty list")
		return
	}

	value = list.tail.next.value
	return
}

// Returns the value of the last element of the list
// Complexity: O(1)
func (list CircularLinkedList[T]) Last() (value T, err error) {
	if list.tail == nil {
		err = fmt.Errorf("could not get first from an empty list")
		return
	}

	value = list.tail.value
	return
}

// Generates a string that represents the whole list
// Complexity: O(n)
func (list CircularLinkedList[T]) ToString() string {
	var buffer bytes.Buffer

	if list.tail == nil {
		return "[  ]"
	}

	current := list.tail
	first := true

	for {
		current = current.next

		if first {
			first = false
		} else {
			buffer.WriteString(", ")
		}

		buffer.WriteString(fmt.Sprint(current.value))

		if current == list.tail {
			break
		}
	}

	return fmt.Sprintf("[ %s ]", buffer.String())
}
