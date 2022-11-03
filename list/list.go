package list

type List[T any] interface {
	AddBefore(value T)
	AddAfter(value T)
	DeleteFirst() (value T, err error) 
	DeleteLast() (value T, err error) 
	Find(value T) (index int, err error) 
	ToString() string
}
