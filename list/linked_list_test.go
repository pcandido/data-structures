package list_test

import (
	"data_structures/list"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinkedList(t *testing.T) {

	t.Run("Push", func(t *testing.T) {
		t.Run("should add elements at the final of list", func(t *testing.T) {
			var list list.LinkedList[int]

			list.Push(1)
			list.Push(2)
			list.Push(3)

			assert.Equal(t, "[ 1, 2, 3 ]", list.ToString())
		})
	})

	t.Run("Pop", func(t *testing.T) {
		t.Run("should return an error if pop from an empty list", func(t *testing.T) {
			var list list.LinkedList[int]
			_, err := list.Pop()
			assert.NotNil(t, err)
		})

		t.Run("in an 1-element list", func(t *testing.T) {
			makeList := func() list.LinkedList[int] {
				var list list.LinkedList[int]
				list.Push(5)
				return list
			}

			t.Run("should return the single element", func(t *testing.T) {
				list := makeList()

				val, err := list.Pop()

				assert.Nil(t, err)
				assert.Equal(t, 5, val)
			})

			t.Run("should empty the list after pop", func(t *testing.T) {
				list := makeList()

				list.Pop()

				assert.Equal(t, "[  ]", list.ToString())
			})
		})

		t.Run("in an 5-elements list", func(t *testing.T) {
			makeList := func() list.LinkedList[int] {
				var list list.LinkedList[int]
				list.Push(1)
				list.Push(2)
				list.Push(3)
				list.Push(4)
				list.Push(5)
				return list
			}

			t.Run("should return the last element", func(t *testing.T) {
				list := makeList()

				val, err := list.Pop()
				assert.Nil(t, err)
				assert.Equal(t, 5, val)

				val, err = list.Pop()
				assert.Nil(t, err)
				assert.Equal(t, 4, val)
			})

			t.Run("should remove the popped element", func(t *testing.T) {
				list := makeList()

				list.Pop()
				assert.Equal(t, "[ 1, 2, 3, 4 ]", list.ToString())

				list.Pop()
				assert.Equal(t, "[ 1, 2, 3 ]", list.ToString())
			})
		})
	})

}
