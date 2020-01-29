package main

import (
	"errors"
	_ "fmt"
	"sync"
)

// Run .
func Run(tasks []func() error, N int, M int) (err error) {

	L := len(tasks)
	if M == 0 {
		M = L
	}

	if N == 0 {
		N = L
	}

	wg := sync.WaitGroup{}
	wg.Add(N)
	taskCh := make(chan func() error)
	errCh := make(chan error, L)

	for i := 0; i < N; i++ {
		go func() {
			worker(taskCh, errCh)
			wg.Done()
		}()
	}

	i := 0
	for M > 0 && i < L {
		select {
		case <-errCh:
			M--
		default:
		}
		if M <= 0 {
			break
		}
		select {
		case <-errCh:
			M--
		case taskCh <- tasks[i]:
			i++
		}
	}

	close(taskCh)
	wg.Wait()
	close(errCh)
	if M <= 0 {
		err = errors.New("Too much errors")
	}
	return
}

func worker(taskCh chan func() error, errCh chan error) {
	for task := range taskCh {
		err := task()
		if err != nil {
			errCh <- err
		}
	}
}
