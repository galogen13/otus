package main

import (
	"fmt"
	"os"

	"time"

	cal "github.com/galogen13/otus/golendarium/pkg/calendar"
	loccal "github.com/galogen13/otus/golendarium/pkg/localcalendar"
)

func main() {

	calend := loccal.NewLocalCalendar()
	event, err := cal.NewEvent(
		time.Date(2020, 03, 1, 20, 34, 58, 0, time.UTC),
		time.Date(2020, 03, 1, 21, 56, 0, 0, time.UTC),
		"Mew event",
		"new event descr")
	if err != nil {
		fmt.Fprint(os.Stderr, err)
	}
	calend.AddEvent(*event)
	err = calend.DeleteEvent("dflksdflksdmfkl")
	fmt.Println(err)
	fmt.Println(calend.List())
	ev, err := calend.GetEvent(event.ID)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
	}
	fmt.Println(ev)

}
