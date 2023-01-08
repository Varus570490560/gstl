package list

import "testing"

func TestList_InsertElementBack(t *testing.T) {
	list := NewList[int]()
	list.PushBackList(list)
	list.PushBackElement(1)
	list.PushBackElement(2)
	list.PushBackElement(3)
	for it, ok := list.Begin(); ok; it, ok = it.Next() {
		t.Log(it.Value())
	}
	it2, _ := list.End()
	it2, _ = it2.Previous()
	it2, _ = it2.Previous()
	list.InsertElementBack(it2, 1000)
	for it, ok := list.Begin(); ok; it, ok = it.Next() {
		t.Log(it.Value())
	}
}
