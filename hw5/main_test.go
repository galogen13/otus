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
	firstItem := list.First()
	list.Remove(*firstItem)
	listLen := list.Len()
	expectedLen := 3
	if listLen != expectedLen {
		t.Fatalf("Expected LEN %d Result len %d", expectedLen, listLen)
	}

}

func TestFirst(t *testing.T) {

	list := List{}
	first := list.First()
	if first != nil {
		t.Fatalf("EMPTY LIST Expected first %#v Result first %p", nil, first)
	}

	list.PushFront(2)
	list.PushFront(1)
	first = list.First()
	if first == nil {
		t.Fatalf("NOT EMPTY LIST Expected first %p Result first %p", list.first, first)
	}

	expectedValue := 1
	if first.Value() != expectedValue {
		t.Fatalf("Expected value %d Result value %d", expectedValue, first.Value())
	}
}

func TestLast(t *testing.T) {

	list := List{}
	last := list.Last()
	if last != nil {
		t.Fatalf("EMPTY LIST Expected last %#v Result last %p", nil, last)
	}

	list.PushFront(2)
	list.PushFront(1)
	last = list.Last()
	if last == nil {
		t.Fatalf("NOT EMPTY LIST Expected last %p Result last %p", list.last, last)
	}

	expectedValue := 2
	if last.Value() != expectedValue {
		t.Fatalf("Expected value %d Result value %d", expectedValue, last.Value())
	}
}

func TestPushFront(t *testing.T) {

	list := List{}

	list.PushFront(2)
	list.PushFront(1)

	expLen := 2
	if list.Len() != expLen {
		t.Fatalf("Expected Len %d Result len %d", expLen, list.Len())
	}

	item := list.First()
	expValue := 1
	if item.Value() != expValue {
		t.Fatalf("Expected value %d Result value %d", expValue, item.Value())
	}
}

func TestPushBack(t *testing.T) {

	list := List{}

	list.PushBack(2)
	list.PushBack(1)

	expLen := 2
	if list.Len() != expLen {
		t.Fatalf("Expected Len %d Result len %d", expLen, list.Len())
	}

	item := list.Last()
	expValue := 1
	if item.Value() != expValue {
		t.Fatalf("Expected value %d Result value %d", expValue, item.Value())
	}
}

func TestRemove(t *testing.T) {

	list := List{}

	list.PushFront(1)
	list.PushFront(2)
	list.PushFront(3)
	item := list.First()
	list.Remove(*item)

	expectedLen := 2
	if list.Len() != expectedLen {
		t.Fatalf("Expected Len %d Result len %d", expectedLen, list.Len())
	}

	expectedValue := 2
	item = list.First()
	if item.Value() != expectedValue {
		t.Fatalf("Expected value %d Result value %d", expectedValue, item.Value())
	}
}

func TestValue(t *testing.T) {

	list := List{}

	expectedValue := 1
	list.PushFront(expectedValue)
	item := list.First()

	if item.Value() != expectedValue {
		t.Fatalf("Expected value %d Result value %d", expectedValue, item.Value())
	}
}

func TestNext(t *testing.T) {

	list := List{}

	list.PushFront(1)
	list.PushFront(2)
	first := list.First()

	if list.last != first.Next() {
		t.Fatalf("Expected value %p Result value %p", list.last, first.Next())
	}

}

func TestPrev(t *testing.T) {

	list := List{}

	list.PushFront(1)
	list.PushFront(2)
	last := list.Last()

	if list.first != last.Prev() {
		t.Fatalf("Expected value %p Result value %p", list.first, last.Prev())
	}

}
