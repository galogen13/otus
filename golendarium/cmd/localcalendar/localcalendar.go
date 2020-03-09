package main

import cal "github.com/galogen13/otus/golendarium/calendar"

// Calendar1 .
type Calendar1 struct {
	events []cal.Event
}

// AddEvent .
func (cal *Calendar1) AddEvent(event cal.Event) error {
	// проверка на существование
	cal.events = append(cal.events, event)
	return nil
}

// DeleteEvent .
func (cal *Calendar1) DeleteEvent(id string) error {

	for i := 0; i < len(cal.events); i++ {
		if cal.events[i].ID == id {
			cal.events = append(cal.events[0:i], cal.events[i+1:]...)
			return nil
		}
	}

	return nil
}

// (cal *Calendar) EditEvent(id string, event Event) error
// (cal *Calendar) GetEvent(id string) (event Event, err error)
// (cal *Calendar) List() []Event

//var Store []Event
