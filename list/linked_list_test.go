package list_test

import (
	"data_structures/list"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinkedList(t *testing.T) {

	t.Run("AddBefore", func(t *testing.T) {
		t.Run("should add elements at the beginning of the list", func(t *testing.T) {
			var list list.LinkedList[int]

			list.AddBefore(1)
			list.AddBefore(2)
			list.AddBefore(3)

			assert.Equal(t, "[ 3, 2, 1 ]", list.ToString())
		})
	})

	t.Run("AddAfter", func(t *testing.T) {
		t.Run("should add elements at the final of the list", func(t *testing.T) {
			var list list.LinkedList[int]

			list.AddAfter(1)
			list.AddAfter(2)
			list.AddAfter(3)

			assert.Equal(t, "[ 1, 2, 3 ]", list.ToString())
		})
	})

	t.Run("DeleteFirst", func(t *testing.T) {
		t.Run("should return an error if list is empty", func(t *testing.T) {
			var list list.LinkedList[int]
			_, err := list.DeleteLast()
			assert.NotNil(t, err)
		})

		t.Run("in an 1-element list", func(t *testing.T) {
			makeList := func() list.LinkedList[int] {
				var list list.LinkedList[int]
				list.AddAfter(5)
				return list
			}

			t.Run("should return the single element", func(t *testing.T) {
				list := makeList()

				val, err := list.DeleteFirst()

				assert.Nil(t, err)
				assert.Equal(t, 5, val)
			})

			t.Run("should empty the list after delete", func(t *testing.T) {
				list := makeList()

				list.DeleteFirst()

				assert.Equal(t, "[  ]", list.ToString())
			})
		})

		t.Run("in an 5-elements list", func(t *testing.T) {
			makeList := func() list.LinkedList[int] {
				var list list.LinkedList[int]
				list.AddAfter(1)
				list.AddAfter(2)
				list.AddAfter(3)
				list.AddAfter(4)
				list.AddAfter(5)
				return list
			}

			t.Run("should return the last element", func(t *testing.T) {
				list := makeList()

				val, err := list.DeleteFirst()
				assert.Nil(t, err)
				assert.Equal(t, 1, val)

				val, err = list.DeleteFirst()
				assert.Nil(t, err)
				assert.Equal(t, 2, val)
			})

			t.Run("should remove the element", func(t *testing.T) {
				list := makeList()

				list.DeleteFirst()
				assert.Equal(t, "[ 2, 3, 4, 5 ]", list.ToString())

				list.DeleteFirst()
				assert.Equal(t, "[ 3, 4, 5 ]", list.ToString())
			})
		})
	})

	t.Run("DeleteLast", func(t *testing.T) {
		t.Run("should return an error if list is empty", func(t *testing.T) {
			var list list.LinkedList[int]
			_, err := list.DeleteLast()
			assert.NotNil(t, err)
		})

		t.Run("in an 1-element list", func(t *testing.T) {
			makeList := func() list.LinkedList[int] {
				var list list.LinkedList[int]
				list.AddAfter(5)
				return list
			}

			t.Run("should return the single element", func(t *testing.T) {
				list := makeList()

				val, err := list.DeleteLast()

				assert.Nil(t, err)
				assert.Equal(t, 5, val)
			})

			t.Run("should empty the list after delete", func(t *testing.T) {
				list := makeList()

				list.DeleteLast()

				assert.Equal(t, "[  ]", list.ToString())
			})
		})

		t.Run("in an 5-elements list", func(t *testing.T) {
			makeList := func() list.LinkedList[int] {
				var list list.LinkedList[int]
				list.AddAfter(1)
				list.AddAfter(2)
				list.AddAfter(3)
				list.AddAfter(4)
				list.AddAfter(5)
				return list
			}

			t.Run("should return the last element", func(t *testing.T) {
				list := makeList()

				val, err := list.DeleteLast()
				assert.Nil(t, err)
				assert.Equal(t, 5, val)

				val, err = list.DeleteLast()
				assert.Nil(t, err)
				assert.Equal(t, 4, val)
			})

			t.Run("should remove the element", func(t *testing.T) {
				list := makeList()

				list.DeleteLast()
				assert.Equal(t, "[ 1, 2, 3, 4 ]", list.ToString())

				list.DeleteLast()
				assert.Equal(t, "[ 1, 2, 3 ]", list.ToString())
			})
		})
	})

}
