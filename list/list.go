package list

type List[T any] interface {
	AddBefore(value T)
	AddAfter(value T)
	DeleteLast() (value T, err error) 
	ToString() string
}
