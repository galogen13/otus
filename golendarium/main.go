package main

import (
	"fmt"

	"time"

	cal "github.com/galogen13/otus/golendarium/cmd/calendar"
	loccal "github.com/galogen13/otus/golendarium/cmd/localcalendar"
)

func main() {

	var calend cal.Calendar
	calend = &loccal.LocalCalendar{}
	event, err := cal.NewEvent(
		time.Date(2020, 03, 1, 20, 34, 58, 0, time.UTC),
		time.Date(2020, 03, 1, 21, 56, 0, 0, time.UTC),
		"Mew event",
		"new event descr")
	if err != nil {
		panic(err)
	}
	calend.AddEvent(event)
	err = calend.DeleteEvent("dflksdflksdmfkl")
	fmt.Print(err)
	fmt.Println(calend.List())
	ev, _ := calend.GetEvent(event.ID)
	fmt.Println(ev)

}
