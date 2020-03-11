package localcalendar

import "github.com/galogen13/otus/golendarium/cmd/calendar"

// Calendar1 .
type LocalCalendar struct {
	events []calendar.Event
}

// AddEvent .
func (cal *LocalCalendar) AddEvent(newEvent calendar.Event) error {
	err := checkDateBusy(cal.events, newEvent)
	if err != nil {
		return err
	}
	cal.events = append(cal.events, newEvent)
	return nil
}

// DeleteEvent .
func (cal *LocalCalendar) DeleteEvent(id string) error {

	for i := 0; i < len(cal.events); i++ {
		if cal.events[i].ID == id {
			cal.events = append(cal.events[0:i], cal.events[i+1:]...)
			return nil
		}
	}

	return calendar.ErrEventNotExist
}

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

func (cal LocalCalendar) GetEvent(id string) (event calendar.Event, err error) {
	for _, event := range cal.events {
		if event.ID == id {
			return event, nil
		}
	}
	return calendar.Event{}, calendar.ErrEventNotExist
}

func (cal LocalCalendar) List() []calendar.Event {
	return cal.events
}

func checkDateBusy(events []calendar.Event, newEvent calendar.Event) error {
	for _, event := range events {
		if newEvent.Start.After(event.Start) && event.Finish.After(newEvent.Start) || newEvent.Finish.After(event.Start) && event.Finish.After(newEvent.Finish) || newEvent.Start.After(event.Start) && event.Finish.After(newEvent.Finish) {
			return calendar.ErrDateBusy
		}
	}
	return nil
}
