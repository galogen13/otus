package calendar

import (
	"testing"
	"time"
)

func TestLen(t *testing.T) {

	_, err := NewEvent(
		time.Date(2020, 03, 1, 20, 34, 58, 0, time.UTC),
		time.Date(2020, 03, 1, 21, 56, 0, 0, time.UTC),
		"Mew event",
		"new event descr")
	if err != nil {
		t.Fatalf("EXPECTED nil error RESULT %s", err)
	}

	_, err = NewEvent(
		time.Date(2020, 03, 2, 20, 34, 58, 0, time.UTC),
		time.Date(2020, 03, 1, 21, 56, 0, 0, time.UTC),
		"Mew event",
		"new event descr")
	if err != ErrBadStartFinish {
		t.Fatalf("EXPECTED \"%s\" error RESULT %s", ErrBadStartFinish, err)
	}

	_, err = NewEvent(
		time.Date(2020, 03, 1, 20, 34, 58, 0, time.UTC),
		time.Date(2020, 03, 1, 20, 34, 58, 0, time.UTC),
		"Mew event",
		"new event descr")
	if err != ErrBadStartFinish {
		t.Fatalf("EXPECTED \"%s\" error RESULT %s", ErrBadStartFinish, err)
	}

	_, err = NewEvent(
		time.Date(2020, 03, 1, 20, 34, 58, 0, time.UTC),
		time.Date(2020, 03, 1, 21, 56, 0, 0, time.UTC),
		"",
		"new event descr")
	if err != ErrEmptyName {
		t.Fatalf("EXPECTED \"%s\" error RESULT %s", ErrEmptyName, err)
	}

	_, err = NewEvent(
		time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC),
		time.Date(2020, 03, 1, 21, 56, 0, 0, time.UTC),
		"New event",
		"new event descr")
	if err != ErrEmptyStartTime {
		t.Fatalf("EXPECTED \"%s\" error RESULT %s", ErrEmptyStartTime, err)
	}

	_, err = NewEvent(
		time.Date(2020, 03, 1, 21, 56, 0, 0, time.UTC),
		time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC),
		"New event",
		"new event descr")
	if err != ErrEmptyFinishTime {
		t.Fatalf("EXPECTED \"%s\" error RESULT %s", ErrEmptyFinishTime, err)
	}

}
