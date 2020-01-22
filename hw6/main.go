package main

import (
	"errors"
	_ "fmt"
)

// Run .
func Run(task []func() error, N int, M int) (err error) {

	startCh := make(chan interface{})
	resultCh := make(chan error, len(task))

	i := 0
	taskCount := len(task)
	maxGoroutines := min(N, taskCount)

	tasksCh := make(chan func() error, maxGoroutines)

	// Подготавливаем к запуску первые maxGoroutines функций
	for i < maxGoroutines {
		f := task[i]
		tasksCh <- f
		go func() {
			<-startCh
			execTask(tasksCh, resultCh)
		}()

		i++
	}
	close(startCh)

	finishedTasks := 0

	for {
		select {

		case res := <-resultCh:
			if res != nil {
				M--
			}
			if M == 0 {
				close(tasksCh)
				err = errors.New("Too much errors")
				return err
			}
			finishedTasks++
			if finishedTasks == taskCount {
				return
			}
		default:
		}
		if i < taskCount {
			select {
			case tasksCh <- task[i]:
				go func() {
					execTask(tasksCh, resultCh)
				}()
				i++
			default:
			}
		}

	}

}

func execTask(tasksCh <-chan func() error, resultCh chan<- error) {
	f, ok := <-tasksCh
	if ok {
		res := f()
		resultCh <- res
	}
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
