package main

type Calendar1 struct {
	events []Event
}

func (cal *Calendar1) AddEvent(event Event) error{
	// проверка на существование
	cal.events = append(cal.events, event)
	return
}
func (cal *Calendar1) DeleteEvent(id string) error{
	
		
	for i :=0 ; i < len(cal.events); i++{
		if cal.events[i].id = id{
			cal.events = append(cal.events[0:i], cal.events[i+1:])

		}
	}
}
	// (cal *Calendar) EditEvent(id string, event Event) error
	// (cal *Calendar) GetEvent(id string) (event Event, err error)
	// (cal *Calendar) List() []Event

//var Store []Event
