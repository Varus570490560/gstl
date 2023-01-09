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

func (l *List[T]) PushBackList(list *List[T]) {
	if list.length == 0 {
		return
	} else if l == list {
		list = l.Clone()
	}
	for iterator := list.pHead; iterator != nil; iterator = iterator.pNext {
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
		l.pEnd = newNode
	} else {
		l.pHead.pPre = newNode
		newNode.pNext = l.pHead
	}
	l.pHead = newNode
	l.nodeId++
	l.length++
	return
}

func (l *List[T]) PushFrontList(list *List[T]) {
	if list.length == 0 {
		return
	} else if l == list {
		list = list.Clone()
	}
	for iterator := list.pEnd; iterator != nil; iterator = iterator.pPre {
		l.PushFrontElement(iterator.element)
	}
}

func (l *List[T]) Clone() *List[T] {
	listClone := NewList[T]()
	listClone.PushBackList(l)
	return listClone
}

func (l *List[T]) Begin() (*Iterator[T], bool) {
	if l.length == 0 {
		return nil, false
	}
	return l.pHead, true
}

func (l *List[T]) End() (*Iterator[T], bool) {
	if l.length == 0 {
		return nil, false
	}
	return l.pEnd, true
}

func (l *List[T]) Empty() bool {
	return l.length == 0
}

// InsertFrontElement don't verify if the iterator in receiver list. It inserts the element front of the iterator.
func (l *List[T]) InsertFrontElement(position *Iterator[T], element T) {
	if position == l.pHead {
		l.PushFrontElement(element)
		return
	}
	newNode := &Iterator[T]{
		nodeId:  l.nodeId,
		element: element,
		pNext:   position,
		pPre:    position.pPre,
	}
	position.pPre = newNode
	newNode.pPre.pNext = newNode
	l.length++
	l.nodeId++
}

// InsertBackElement don't verify if the position iterator in receiver list. It inserts the element back of the iterator.
func (l *List[T]) InsertBackElement(position *Iterator[T], element T) {
	if position == l.pEnd {
		l.PushBackElement(element)
		return
	}
	newNode := &Iterator[T]{
		nodeId:  l.nodeId,
		element: element,
		pNext:   position.pNext,
		pPre:    position,
	}
	position.pNext = newNode
	newNode.pNext.pPre = newNode
	l.length++
	l.nodeId++
}

// InsertFrontList don't verify if the iterator in receiver list. It inserts the list front of the iterator.
func (l *List[T]) InsertFrontList(position *Iterator[T], list *List[T]) {
	if list.Empty() {
		return
	}
	if l == list {
		list = list.Clone()
	}
	for it, ok := list.End(); ok; it, ok = it.Previous() {
		list.InsertFrontElement(position, it.Value())
	}
}

// InsertBackList don't verify if the iterator in receiver list. It inserts the list back of the iterator.
func (l *List[T]) InsertBackList(position *Iterator[T], list *List[T]) {
	if list.Empty() {
		return
	}
	if l == list {
		list = list.Clone()
	}
	for it, ok := list.Begin(); ok; it, ok = it.Next() {
		list.InsertBackElement(position, it.Value())
	}
}

func (l *List[T]) EraseElement(iterator *Iterator[T]) {
	if iterator == l.pHead {
		l.PopFront()
		return
	} else if iterator == l.pEnd {
		l.PopBack()
		return
	}
	iterator.pPre.pNext = iterator.pNext
	iterator.pNext.pPre = iterator.pPre
	l.length--
	return
}

//// EraseElements remove the element between start iterator and end iterator, the start element will be removed but the end not.
//// Waring: You must insure start iterator in front of the end. For execution efficiency, EraseElements don't check it.
//func (l *List[T]) EraseElements(start *Iterator[T], end *Iterator[T]) {
//	if start == end {
//		return
//	}
//
//}
