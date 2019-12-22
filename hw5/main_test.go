package main

import (
	"testing"
)

func TestLen(t *testing.T) {

	list := List{}
	list.PushFront(1)
	list.PushFront(2)
	list.PushFront(3)
	list.PushBack(0)
	firstItem, _ := list.First()
	list.Remove(firstItem)
	listLen := list.Len()
	expectedLen := 3
	if listLen != expectedLen {
		t.Fatalf("Expected LEN %d Result len %d", expectedLen, listLen)
	}

}

func TestFirst(t *testing.T) {

	list := List{}
	_, ok := list.First()
	if ok {
		t.Fatalf("EMPTY LIST Expected ERR %t Result ERR %t", false, ok)
	}

	list.PushFront(2)
	list.PushFront(1)
	first, ok := list.First()
	if !ok {
		t.Fatalf("NOT EMPTY LIST Expected ERR %t Result ERR %t", true, ok)
	}

	if first.Value() != 1 {
		t.Fatalf("Expected value %d Result value %d", 1, first.Value())
	}
}

func TestLast(t *testing.T) {

	list := List{}
	_, ok := list.Last()
	if ok {
		t.Fatalf("EMPTY LIST Expected ERR %t Result ERR %t", false, ok)
	}

	list.PushFront(2)
	list.PushFront(1)
	last, ok := list.Last()
	if !ok {
		t.Fatalf("NOT EMPTY LIST Expected ERR %t Result ERR %t", true, ok)
	}

	if last.Value() != 2 {
		t.Fatalf("Expected value %d Result value %d", 1, last.Value())
	}
}

func TestPushFront(t *testing.T) {

	list := List{}

	list.PushFront(2)
	list.PushFront(1)

	if list.Len() != 2 {
		t.Fatalf("Expected Len %d Result len %d", 2, list.Len())
	}

	item, _ := list.First()
	if item.Value() != 1 {
		t.Fatalf("Expected value %d Result value %d", 1, item.Value())
	}
}

func TestPushBack(t *testing.T) {

	list := List{}

	list.PushBack(2)
	list.PushBack(1)

	if list.Len() != 2 {
		t.Fatalf("Expected Len %d Result len %d", 2, list.Len())
	}

	item, _ := list.Last()
	if item.Value() != 1 {
		t.Fatalf("Expected value %d Result value %d", 1, item.Value())
	}
}

func TestRemove(t *testing.T) {

	list := List{}

	list.PushFront(1)
	list.PushFront(2)
	list.PushFront(3)
	item, _ := list.First()
	list.Remove(item)

	expectedLen := 2
	if list.Len() != expectedLen {
		t.Fatalf("Expected Len %d Result len %d", expectedLen, list.Len())
	}

	expectedValue := 2
	item, _ = list.First()
	if item.Value() != expectedValue {
		t.Fatalf("Expected value %d Result value %d", expectedValue, item.Value())
	}
}

func TestValue(t *testing.T) {

	list := List{}

	expectedValue := 1
	list.PushFront(expectedValue)
	item, _ := list.First()

	if item.Value() != expectedValue {
		t.Fatalf("Expected value %d Result value %d", expectedValue, item.Value())
	}
}

func TestNext(t *testing.T) {

	list := List{}

	list.PushFront(1)
	list.PushFront(2)
	first, _ := list.First()

	if list.last != first.Next() {
		t.Fatalf("Expected value %p Result value %p", list.last, first.Next())
	}

}

func TestPrev(t *testing.T) {

	list := List{}

	list.PushFront(1)
	list.PushFront(2)
	last, _ := list.Last()

	if list.first != last.Prev() {
		t.Fatalf("Expected value %p Result value %p", list.first, last.Prev())
	}

}
