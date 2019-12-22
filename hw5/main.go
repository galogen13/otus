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

	first, _ := list.First()
	list.Remove(first)
	first, _ = list.First()
	list.Remove(first)
	last, _ := list.Last()
	list.Remove(last)

	first, _ = list.First()
	list.Remove(first)

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

// First - возвращает первый элемент списка и булево значение(Ложь - если список пустой, иначе Истина)#
func (list List) First() (item Item, isExists bool) {
	if list.first == nil {
		item = Item{}
	} else {
		item = *(list.first)
		isExists = true
	}
	return
}

// Last - возвращает последний элемент списка и булево значение(Ложь - если список пустой, иначе Истина)#
func (list List) Last() (item Item, isExists bool) {
	if list.last == nil {
		item = Item{}
	} else {
		item = *(list.last)
		isExists = true
	}
	return
}

// PushFront - добавляет в начало списка элемент #
func (list *List) PushFront(v interface{}) {
	var newFirst Item
	if list.first != nil {
		newFirst = Item{value: v, next: list.first}
		list.first.prev = &newFirst

	} else {
		newFirst = Item{value: v}
	}
	if list.Len() == 0 {
		list.last = &newFirst
	}
	list.first = &newFirst
	list.len++
}

// PushBack - добавляет в конец списка элемент #
func (list *List) PushBack(v interface{}) {
	var newLast Item
	if list.last != nil {
		newLast = Item{value: v, prev: list.last}
		list.last.next = &newLast
	} else {
		newLast = Item{value: v, prev: nil}
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
		{
			i.prev.next = i.next
			i.next.prev = i.prev
		}

	// Если удаляется крайний элемент, то предкрайний должен стать крайним
	case i.next == nil && i.prev != nil:
		{
			i.prev.next = nil
			list.last = i.prev
		}

	// Если удаляется первый элемент, то второй должен стать первым
	case i.next != nil && i.prev == nil:
		{
			i.next.prev = nil
			list.first = i.next
		}
	// Если это последний элемент в списке, очищаем первого и последнего в списке
	case i.next == nil && i.prev == nil:
		{
			list.first = nil
			list.last = nil
		}
	}

	list.len--
}
