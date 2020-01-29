package main

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

var CompletedTasks int = 0 

func TestRun(t *testing.T) {

	var tasks []func() error

	var f func() error

	for i := 0; i < 10; i++ {
		f = FailFunc()
		tasks = append(tasks, f)
	}

	for i := 0; i < 5; i++ {
		f = OkFunc()
		tasks = append(tasks, f)
	}

	CompletedTasks = 0
	err := Run(tasks, 5, 5)
	if err == nil {
		t.Fatalf("ERR is empty")
	}
	if CompletedTasks >= 5+5{
		t.Fatalf("Too much completed tasks: %d", CompletedTasks)	
	}

	tasks = tasks[:0]

	for i := 0; i < 10; i++ {
		f = OkFunc()
		tasks = append(tasks, f)
	}

	CompletedTasks = 0
	err = Run(tasks, 2, 3)
	if err != nil {
		t.Fatalf("ERR is not empty")
	}
	if CompletedTasks < len(tasks){
		t.Fatalf("Completed tasks: %d; exp completed tasks: %d", CompletedTasks, len(tasks))	
	}

	CompletedTasks = 0
	err = Run(tasks, 0, 0)
	if err != nil {
		t.Fatalf("ERR is not empty")
	}

}

func FailFunc() func() error {

	return func() error {
		time.Sleep(time.Second)
		fmt.Println("FAIL")
		CompletedTasks++
		return errors.New("Fail")
	}

}

func OkFunc() func() error {

	return func() error {
		var err error
		time.Sleep(time.Second)
		fmt.Println("OK")
		CompletedTasks++
		return err
	}

}
