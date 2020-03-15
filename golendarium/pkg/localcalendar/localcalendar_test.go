package localcalendar

import (
	"github.com/galogen13/otus/golendarium/pkg/calendar"
	"reflect"
	"testing"
	"time"
)

func TestLen(t *testing.T) {

	cal := NewLocalCalendar()

	event1, _ := calendar.NewEvent(
		time.Date(2020, 03, 1, 20, 34, 58, 0, time.UTC),
		time.Date(2020, 03, 1, 21, 56, 0, 0, time.UTC),
		"Mew event 1",
		"new event descr 1")

	err := cal.AddEvent(*event1)
	if err != nil {
		t.Fatalf("EXPECTED nil error RESULT %s", err)
	}
	if len(cal.List()) != 1 {
		t.Fatalf("EXPECTED list len = 1 RESULT %d", len(cal.List()))
	}

	event2, _ := calendar.NewEvent(
		time.Date(2020, 03, 1, 21, 34, 58, 0, time.UTC),
		time.Date(2020, 03, 1, 22, 56, 0, 0, time.UTC),
		"Mew event 2",
		"new event descr 2")

	err = cal.AddEvent(*event2)
	if err != calendar.ErrDateBusy {
		t.Fatalf("EXPECTED %s RESULT %s", calendar.ErrDateBusy, err)
	}
	if len(cal.List()) != 1 {
		t.Fatalf("EXPECTED list len = 1 RESULT %d", len(cal.List()))
	}

	event3, _ := calendar.NewEvent(
		time.Date(2020, 3, 2, 21, 34, 58, 0, time.UTC),
		time.Date(2020, 3, 2, 22, 56, 0, 0, time.UTC),
		"Mew event 3",
		"new event descr 3")

	_ = cal.AddEvent(*event3)
	findEvent, err := cal.GetEvent(event1.ID)
	if !reflect.DeepEqual(event1, findEvent) {
		t.Fatalf("GetEvent EXPECTED %s RESULT %s", event1, findEvent)
	}

	err = cal.DeleteEvent("12312312312312312321")
	if err != calendar.ErrEventNotExist {
		t.Fatalf("EXPECTED %s RESULT %s", calendar.ErrEventNotExist, err)
	}

	err = cal.DeleteEvent(event1.ID)
	if err != nil {
		t.Fatalf("EXPECTED nil error RESULT %s", err)
	}
	if len(cal.List()) != 1 {
		t.Fatalf("EXPECTED list len = 1 RESULT %d", len(cal.List()))
	}

	event3Edited, _ := calendar.NewEvent(
		time.Date(2020, 3, 4, 21, 34, 58, 0, time.UTC),
		time.Date(2020, 3, 4, 22, 56, 0, 0, time.UTC),
		"Mew event 3 edited",
		"new event descr 3 edited")
	event3Edited.ID = event3.ID

	err = cal.EditEvent("123123123123123123", *event3Edited)
	if err != calendar.ErrEventNotExist {
		t.Fatalf("EXPECTED %s RESULT %s", calendar.ErrEventNotExist, err)
	}

	err = cal.EditEvent(event3.ID, *event3Edited)
	if err != nil {
		t.Fatalf("EXPECTED nil error RESULT %s", err)
	}
	event3Find, _ := cal.GetEvent(event3.ID)
	if !reflect.DeepEqual(*event3Find, *event3Edited) {
		t.Fatalf("EXPECTED %s RESULT %s", event3Edited, event3Find)
	}

}
