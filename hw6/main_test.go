package main

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

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

	err := Run(tasks, 5, 5)
	if err == nil {
		t.Fatalf("ERR is empty")
	}

	tasks = tasks[:0]

	fmt.Println("next")

	for i := 0; i < 10; i++ {
		f = OkFunc()
		tasks = append(tasks, f)
	}

	err = Run(tasks, 2, 3)
	if err != nil {
		t.Fatalf("ERR is not empty")
	}

}

func FailFunc() func() error {

	return func() error {
		time.Sleep(time.Second)
		fmt.Println("FAIL")
		return errors.New("Fail")
	}

}

func OkFunc() func() error {

	return func() error {
		var err error
		time.Sleep(time.Second)
		fmt.Println("OK")
		return err
	}

}
