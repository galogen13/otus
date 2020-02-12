package calendar

import "time"
import "github.com/chilts/sid"

type Calendar interface {
	(cal *calendar) AddEvent(event Event) (event Event, err error)
	(cal *calendar) DeleteEvent(id string) error
	(cal *calendar) EditEvent(id string, event Event) error
	(cal *calendar) GetEvent(id string) (event Event, err error)
	(cal *calendar) List() []Event

}

type Event struct{
	id string
	Start, Finish time.Time
	Name, Descr string
}

func NewEvent(start, finish time.Time, name, descr string) Event{
	return Event{id: sid.Id() , Start:start, Finish: finish, Name:name, Descr: descr}
}
