package list

type Iterator[T any] struct {
	element T
	nodeId  int
	pNext   *Iterator[T]
	pPre    *Iterator[T]
}

func (i *Iterator[T]) Value() T {
	return i.element
}
