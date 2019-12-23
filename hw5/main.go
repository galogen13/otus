package main

import (
	"fmt"
)

func main() {

	list := List{}
	fmt.Println(list.Len())
	list.PushFront(1)
	list.PushFront(2)
	list.PushFront(3)
	list.PushBack(0)

	first := list.First()
	list.Remove(*first)
	first = list.First()
	list.Remove(*first)
	last := list.Last()
	list.Remove(*last)

	first = list.First()
	list.Remove(*first)

	fmt.Println(list.First())

}

// Item - элемент списка #
type Item struct {
	value interface{}
	next  *Item
	prev  *Item
}

// Value - возвращает значение элемента #
func (item Item) Value() interface{} {
	return item.value
}

// Next - возвращает указатель на следующий элемент #
func (item Item) Next() *Item {
	return item.next
}

// Prev - возвращает указатель на предыдущий элемент #
func (item Item) Prev() *Item {
	return item.prev
}

// List - двухсвязный список #
type List struct {
	len   int
	first *Item
	last  *Item
}

// Len - возвращает длину списка #
func (list List) Len() int {
	return list.len
}

// First - возвращает указатель на первый элемент списка#
func (list List) First() *Item {
	return list.first
}

// Last - возвращает указатель на последний элемент списка#
func (list List) Last() *Item {
	return list.last
}

// PushFront - добавляет в начало списка элемент #
func (list *List) PushFront(v interface{}) {
	newFirst := Item{value: v}
	if list.first != nil {
		newFirst.next = list.first
		list.first.prev = &newFirst
	}
	if list.Len() == 0 {
		list.last = &newFirst
	}
	list.first = &newFirst
	list.len++
}

// PushBack - добавляет в конец списка элемент #
func (list *List) PushBack(v interface{}) {
	newLast := Item{value: v}
	if list.last != nil {
		newLast.prev = list.last
		list.last.next = &newLast
	}

	if list.Len() == 0 {
		list.first = &newLast
	}
	list.last = &newLast
	list.len++
}

// Remove - удаляет элемент из списка #
func (list *List) Remove(i Item) {
	switch {
	// Если известен предудыщий и следующий - сошьем их
	case i.next != nil && i.prev != nil:

		i.prev.next = i.next
		i.next.prev = i.prev

	// Если удаляется крайний элемент, то предкрайний должен стать крайним
	case i.next == nil && i.prev != nil:

		i.prev.next = nil
		list.last = i.prev

	// Если удаляется первый элемент, то второй должен стать первым
	case i.next != nil && i.prev == nil:

		i.next.prev = nil
		list.first = i.next

	// Если это последний элемент в списке, очищаем первого и последнего в списке
	case i.next == nil && i.prev == nil:

		list.first = nil
		list.last = nil

	}

	list.len--
}
