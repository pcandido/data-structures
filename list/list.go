package list

type List[T any] interface {
	InsertStart(value T)
	InsertEnd(value T)
	RemoveStart() (value T, err error)
	RemoveEnd() (value T, err error)
	Find(value T) (index int, err error)
	First() (value T, err error)
	ToString() string
}
