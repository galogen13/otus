package main

import "time"
import "github.com/chilts/sid"

type Calendar interface {
	AddEvent(event Event) error
	DeleteEvent(id string) error
	EditEvent(id string, event Event) error
	GetEvent(id string) (event Event, err error)
	List() []Event
}

type Event struct {
	id            string
	Start, Finish time.Time
	Name, Descr   string
}

func NewEvent(start, finish time.Time, name, descr string) Event {
	return Event{id: sid.Id(), Start: start, Finish: finish, Name: name, Descr: descr}
}
