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

// ErrEmptyStartTime .
var ErrEmptyStartTime = errors.New("Empty start time")

// ErrEmptyFinishTime .
var ErrEmptyFinishTime = errors.New("Empty finish time")

// ErrEventNotExist .
var ErrEventNotExist = errors.New("Event not exist")

// ErrBadStartFinish .
var ErrBadStartFinish = errors.New("Start date is greater than finish date")

// Calendar .
type Calendar interface {
	AddEvent(event Event) error
	DeleteEvent(id string) error
	EditEvent(id string, event Event) error
	GetEvent(id string) (event *Event, err error)
	List() []Event
}

// Event .
type Event struct {
	ID            string
	Start, Finish time.Time
	Name, Descr   string
}

// NewEvent .
func NewEvent(start, finish time.Time, name, descr string) (*Event, error) {

	if name == "" {
		return nil, ErrEmptyName
	}

	if start.IsZero() {
		return nil, ErrEmptyStartTime
	}

	if finish.IsZero() {
		return nil, ErrEmptyFinishTime
	}

	if start.After(finish) || start.Equal(finish) {
		return nil, ErrBadStartFinish
	}

	return &Event{ID: sid.Id(), Start: start, Finish: finish, Name: name, Descr: descr}, nil
}
