package list

type List[T any] interface {
	AddAfter(value T)
	DeleteLast() (value T, err error) 
	ToString() string
}
