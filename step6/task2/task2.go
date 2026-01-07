package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"sync"
)

type safeNames struct {
	names []string
	mu    *sync.Mutex
}

func New(names []string) *safeNames {
	return &safeNames{names, &sync.Mutex{}}
}

func (s *safeNames) GetMark(index int) (int, error) {
	s.mu.Lock()
	url := fmt.Sprintf("http://localhost:8082/mark?name=%s", s.names[index])
	s.mu.Unlock()

	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return 0, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	num, err := strconv.Atoi(string(data))
	if err != nil {
		return 0, err
	}
	return num, nil
}

func Average(names []string) (int, error) {
	snames := New(names)

	// wg := &sync.WaitGroup{}
	ch := make(chan int, len(names))
	errChan := make(chan error)

	// wg.Add(len(names))
	for i := 0; i < len(names); i++ {
		go func(ch chan int, errCh chan error, i int) {
			mark, err := snames.GetMark(i)
			if err != nil {
				errCh <- err
				return
			}

			ch <- mark
		}(ch, errChan, i)
	}
	// wg.Wait()

	sum := 0
	for i := 0; i < len(names); i++ {
		select {
		case num := <-ch:
			sum += num
		case err := <-errChan:
			close(errChan)
			close(ch)
			return 0, err

		}
	}
	close(errChan)
	close(ch)

	return int(sum / len(names)), nil
}
