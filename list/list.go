package list

type List[T any] interface {
	AddAfter(value T)
	ToString() string
}
