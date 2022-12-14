package list

import (
	"bytes"
	"fmt"
)

type singlyLinkedListNode[T comparable] struct {
	value T
	next  *singlyLinkedListNode[T]
}

type SinglyLinkedList[T comparable] struct {
	head *singlyLinkedListNode[T]
}

func (node *singlyLinkedListNode[T]) isLast() bool {
	return node.next == nil
}

// Adds an element at the beginning of the list
// Complexity: O(1)
func (list *SinglyLinkedList[T]) InsertStart(value T) {
	list.head = &singlyLinkedListNode[T]{value, list.head}
}

// Adds an element at the final of the list
// Complexity: O(n)
func (list *SinglyLinkedList[T]) InsertEnd(value T) {
	newNode := &singlyLinkedListNode[T]{value, nil}

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

// Deletes and returns the first element of the list
// Complexity: O(1)
func (list *SinglyLinkedList[T]) RemoveStart() (value T, err error) {
	if list.head == nil {
		err = fmt.Errorf("could not delete from an empty list")
		return
	}

	value = list.head.value
	list.head = list.head.next

	return
}

// Deletes and returns the last element of the list
// Complexity: O(n)
func (list *SinglyLinkedList[T]) RemoveEnd() (value T, err error) {
	if list.head == nil {
		err = fmt.Errorf("could not delete from an empty list")
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

// Find and element on the list and returns its index
// Complexity: O(n)
func (list SinglyLinkedList[T]) Find(value T) (index int, err error) {
	currentNode := list.head
	currentIndex := 0

	for currentNode != nil {
		if currentNode.value == value {
			index = currentIndex
			return
		}

		currentIndex = currentIndex + 1
		currentNode = currentNode.next
	}

	err = fmt.Errorf("element not found")
	return
}

// Returns the value of the first element of the list
// Complexity: O(1)
func (list SinglyLinkedList[T]) First() (value T, err error) {
	if list.head == nil {
		err = fmt.Errorf("could not get first from an empty list")
		return
	}

	value = list.head.value
	return
}

// Returns the value of the last element of the list
// Complexity: O(n)
func (list SinglyLinkedList[T]) Last() (value T, err error) {
	if list.head == nil {
		err = fmt.Errorf("could not get first from an empty list")
		return
	}

	current := list.head
	for !current.isLast() {
		current = current.next
	}

	value = current.value
	return
}

// Generates a string that represents the whole list
// Complexity: O(n)
func (list SinglyLinkedList[T]) ToString() string {
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
