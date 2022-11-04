package list_test

import (
	"data_structures/list"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSinglyLinkedList(t *testing.T) {

	t.Run("InsertStart", func(t *testing.T) {
		t.Run("should add elements at the beginning of the list", func(t *testing.T) {
			var list list.SinglyLinkedList[int]

			list.InsertStart(1)
			list.InsertStart(2)
			list.InsertStart(3)

			assert.Equal(t, "[ 3, 2, 1 ]", list.ToString())
		})
	})

	t.Run("InsertEnd", func(t *testing.T) {
		t.Run("should add elements at the final of the list", func(t *testing.T) {
			var list list.SinglyLinkedList[int]

			list.InsertEnd(1)
			list.InsertEnd(2)
			list.InsertEnd(3)

			assert.Equal(t, "[ 1, 2, 3 ]", list.ToString())
		})
	})

	t.Run("RemoveStart", func(t *testing.T) {
		t.Run("should return an error if list is empty", func(t *testing.T) {
			var list list.SinglyLinkedList[int]
			_, err := list.RemoveEnd()
			assert.NotNil(t, err)
		})

		t.Run("in an 1-element list", func(t *testing.T) {
			makeList := func() list.SinglyLinkedList[int] {
				var list list.SinglyLinkedList[int]
				list.InsertEnd(5)
				return list
			}

			t.Run("should return the single element", func(t *testing.T) {
				list := makeList()

				val, err := list.RemoveStart()

				assert.Nil(t, err)
				assert.Equal(t, 5, val)
			})

			t.Run("should empty the list after delete", func(t *testing.T) {
				list := makeList()

				list.RemoveStart()

				assert.Equal(t, "[  ]", list.ToString())
			})
		})

		t.Run("in an 5-elements list", func(t *testing.T) {
			makeList := func() list.SinglyLinkedList[int] {
				var list list.SinglyLinkedList[int]
				list.InsertEnd(1)
				list.InsertEnd(2)
				list.InsertEnd(3)
				list.InsertEnd(4)
				list.InsertEnd(5)
				return list
			}

			t.Run("should return the last element", func(t *testing.T) {
				list := makeList()

				val, err := list.RemoveStart()
				assert.Nil(t, err)
				assert.Equal(t, 1, val)

				val, err = list.RemoveStart()
				assert.Nil(t, err)
				assert.Equal(t, 2, val)
			})

			t.Run("should remove the element", func(t *testing.T) {
				list := makeList()

				list.RemoveStart()
				assert.Equal(t, "[ 2, 3, 4, 5 ]", list.ToString())

				list.RemoveStart()
				assert.Equal(t, "[ 3, 4, 5 ]", list.ToString())
			})
		})
	})

	t.Run("RemoveEnd", func(t *testing.T) {
		t.Run("should return an error if list is empty", func(t *testing.T) {
			var list list.SinglyLinkedList[int]
			_, err := list.RemoveEnd()
			assert.NotNil(t, err)
		})

		t.Run("in an 1-element list", func(t *testing.T) {
			makeList := func() list.SinglyLinkedList[int] {
				var list list.SinglyLinkedList[int]
				list.InsertEnd(5)
				return list
			}

			t.Run("should return the single element", func(t *testing.T) {
				list := makeList()

				val, err := list.RemoveEnd()

				assert.Nil(t, err)
				assert.Equal(t, 5, val)
			})

			t.Run("should empty the list after delete", func(t *testing.T) {
				list := makeList()

				list.RemoveEnd()

				assert.Equal(t, "[  ]", list.ToString())
			})
		})

		t.Run("in an 5-elements list", func(t *testing.T) {
			makeList := func() list.SinglyLinkedList[int] {
				var list list.SinglyLinkedList[int]
				list.InsertEnd(1)
				list.InsertEnd(2)
				list.InsertEnd(3)
				list.InsertEnd(4)
				list.InsertEnd(5)
				return list
			}

			t.Run("should return the last element", func(t *testing.T) {
				list := makeList()

				val, err := list.RemoveEnd()
				assert.Nil(t, err)
				assert.Equal(t, 5, val)

				val, err = list.RemoveEnd()
				assert.Nil(t, err)
				assert.Equal(t, 4, val)
			})

			t.Run("should remove the element", func(t *testing.T) {
				list := makeList()

				list.RemoveEnd()
				assert.Equal(t, "[ 1, 2, 3, 4 ]", list.ToString())

				list.RemoveEnd()
				assert.Equal(t, "[ 1, 2, 3 ]", list.ToString())
			})
		})
	})

	t.Run("Find", func(t *testing.T) {
		t.Run("should return an error if the list is empty", func(t *testing.T) {
			var list list.SinglyLinkedList[int]
			_, err := list.Find(3)
			assert.NotNil(t, err)
		})

		t.Run("should return the index of found element", func(t *testing.T) {
			var list list.SinglyLinkedList[int]
			list.InsertEnd(1)
			list.InsertEnd(2)
			list.InsertEnd(3)
			list.InsertEnd(4)

			index, err := list.Find(3)

			assert.Nil(t, err)
			assert.Equal(t, 2, index)
		})

		t.Run("should return an error if element does not exist on the list", func(t *testing.T) {
			var list list.SinglyLinkedList[int]
			list.InsertEnd(1)
			list.InsertEnd(2)
			list.InsertEnd(3)
			list.InsertEnd(4)

			_, err := list.Find(5)

			assert.NotNil(t, err)
		})
	})

	t.Run("First", func(t *testing.T) {
		t.Run("should return an error if list is empty", func(t *testing.T) {
			var list list.SinglyLinkedList[int]
			_, err := list.First()
			assert.NotNil(t, err)
		})

		t.Run("should return the first element of the list", func(t *testing.T) {
			var list list.SinglyLinkedList[int]
			list.InsertEnd(1)
			list.InsertEnd(2)
			list.InsertEnd(3)

			value, err := list.First()

			assert.Nil(t, err)
			assert.Equal(t, 1, value)
		})
	})

	t.Run("Last", func(t *testing.T) {
		t.Run("should return an error if list is empty", func(t *testing.T) {
			var list list.SinglyLinkedList[int]
			_, err := list.Last()
			assert.NotNil(t, err)
		})

		t.Run("should return the last element of the list", func(t *testing.T) {
			var list list.SinglyLinkedList[int]
			list.InsertEnd(1)
			list.InsertEnd(2)
			list.InsertEnd(3)

			value, err := list.Last()

			assert.Nil(t, err)
			assert.Equal(t, 3, value)
		})
	})
}
