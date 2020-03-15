package localcalendar

import "github.com/galogen13/otus/golendarium/pkg/calendar"

// LocalCalendar .
type LocalCalendar struct {
	events map[string]calendar.Event
}

func NewLocalCalendar() *LocalCalendar {

	calend := LocalCalendar{}
	calend.events = make(map[string]calendar.Event)
	return &calend
}

// AddEvent .
func (cal *LocalCalendar) AddEvent(newEvent calendar.Event) error {
	err := cal.checkDateBusy(newEvent)
	if err != nil {
		return err
	}
	cal.events[newEvent.ID] = newEvent
	return nil
}

// DeleteEvent .
func (cal *LocalCalendar) DeleteEvent(id string) error {

	_, ok := cal.events[id]
	if ok {
		delete(cal.events, id)
		return nil
	}

	return calendar.ErrEventNotExist
}

// EditEvent .
func (cal *LocalCalendar) EditEvent(id string, event calendar.Event) error {

	err := cal.DeleteEvent(id)
	if err != nil {
		return err
	}

	event.ID = id
	err = cal.AddEvent(event)
	if err != nil {
		return err
	}

	return nil
}

// GetEvent .
func (cal LocalCalendar) GetEvent(id string) (*calendar.Event, error) {

	event, ok := cal.events[id]
	if ok {
		return &event, nil
	}

	return nil, calendar.ErrEventNotExist
}

// List .
func (cal LocalCalendar) List() (list []calendar.Event) {
	for _, event := range cal.events {
		list = append(list, event)
	}
	return
}

func (cal LocalCalendar) checkDateBusy(newEvent calendar.Event) error {
	for _, event := range cal.events {
		if newEvent.Start.After(event.Start) && event.Finish.After(newEvent.Start) || newEvent.Finish.After(event.Start) && event.Finish.After(newEvent.Finish) || newEvent.Start.After(event.Start) && event.Finish.After(newEvent.Finish) {
			return calendar.ErrDateBusy
		}
	}
	return nil
}
