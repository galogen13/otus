package calendar

import (
	"errors"
	"time"

	"github.com/chilts/sid"
)

// ErrDateBusy .
var ErrDateBusy = errors.New("Date busy")

// ErrEmptyName .
var ErrEmptyName = errors.New("Empty name")

// ErrEmptyName .
var ErrEmptyStartTime = errors.New("Empty start time")

// ErrEmptyFinishTime .
var ErrEmptyFinishTime = errors.New("Empty finish time")

// Calendar .
type Calendar interface {
	AddEvent(event Event) error
	DeleteEvent(id string) error
	EditEvent(id string, event Event) error
	GetEvent(id string) (event Event, err error)
	List() []Event
}

// Event .
type Event struct {
	ID            string
	Start, Finish time.Time
	Name, Descr   string
}

// NewEvent .
func NewEvent(start, finish time.Time, name, descr string) Event {
	return Event{ID: sid.Id(), Start: start, Finish: finish, Name: name, Descr: descr}
}
