package list

type List[T any] interface {
	Push(value T)
	Pop() (value T, err error) 
	ToString() string
}
