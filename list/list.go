package list

/*
 List is a doubly linked list that does not guarantee concurrency safety.
*/

type List[T any] struct {
	pHead  *Iterator[T]
	pEnd   *Iterator[T]
	length int
	nodeId int
}

type Iterator[T any] struct {
	element T
	nodeId  int
	pNext   *Iterator[T]
	pPre    *Iterator[T]
}

func NewList[T any]() *List[T] {
	return &List[T]{
		pHead:  nil,
		pEnd:   nil,
		length: 0,
		nodeId: 0,
	}
}

func (l *List[T]) PushBackElement(element T) {
	nodeId := l.nodeId
	newNode := &Iterator[T]{
		nodeId:  nodeId,
		element: element,
		pNext:   nil,
		pPre:    nil,
	}
	if l.length == 0 {
		l.pHead = newNode
	} else {
		l.pEnd.pNext = newNode
		newNode.pPre = l.pEnd
	}
	l.pEnd = newNode
	l.nodeId++
	l.length++
	return
}

func (l *List[T]) PushBackList(list List[T]) {
	if list.length == 0 {
		return
	}
	for iterator := list.pHead; iterator.pNext != nil; iterator = iterator.pNext {
		l.PushBackElement(iterator.element)
	}
}

func (l *List[T]) Size() int {
	return l.length
}

func (l *List[T]) PopBack() {
	if l.length == 0 {
		panic("List pop with length is 0.")
	} else if l.length == 1 {
		l.pHead = nil
		l.pEnd = nil
		l.length = 0
		return
	} else {
		l.pEnd = l.pEnd.pPre
		l.length--
	}
}

func (l *List[T]) PopFront() {
	if l.length == 0 {
		panic("List pop with length is 0.")
	} else if l.length == 1 {
		l.pHead = nil
		l.pEnd = nil
		l.length = 0
		return
	} else {
		l.pHead = l.pHead.pNext
		l.length--
	}
}

func (l *List[T]) PushFrontElement(element T) {
	nodeId := l.nodeId
	newNode := &Iterator[T]{
		nodeId:  nodeId,
		element: element,
		pNext:   nil,
		pPre:    nil,
	}
	if l.length == 0 {
		l.pHead = newNode
	} else {
		l.pHead.pPre = newNode
		newNode.pNext = l.pHead
	}
	l.pHead = newNode
	l.nodeId++
	l.length++
	return
}

func (l *List[T]) PushFrontList(list List[T]) {
	if list.length == 0 {
		return
	}
	for iterator := list.pEnd; iterator.pPre != nil; iterator = iterator.pPre {
		l.PushFrontElement(iterator.element)
	}
}
