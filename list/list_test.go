package list

import (
	"testing"
)

func TestList_InsertElementBack(t *testing.T) {
	list := NewList[int]()
	list.PushBackList(list)
	list.PushBackElement(1)
	list.PushBackElement(2)
	list.PushBackElement(3)
	list.PushBackElement(4)
	list.PushBackElement(5)
	list.PushBackElement(6)
	begin, _ := list.Begin()
	begin, _ = begin.Next()
	begin, _ = begin.Next()
	begin, _ = begin.Next()

	begin, _ = begin.Next()

	begin, _ = begin.Next()

	end, _ := list.End()
	end, _ = end.Previous()
	end, _ = end.Previous()
	list.EraseElements(begin, end)
	for it, ok := list.Begin(); ok; it, ok = it.Next() {
		t.Log(it.Value())
	}
	t.Log(list.Size())
}
