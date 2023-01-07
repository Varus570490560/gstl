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

func (i *Iterator[T]) Next() (*Iterator[T], bool) {
	if i.pNext == nil {
		return nil, false
	}
	return i.pNext, true
}

func (i *Iterator[T]) Previous() (*Iterator[T], bool) {
	if i.pPre == nil {
		return nil, false
	}
	return i.pPre, true
}
