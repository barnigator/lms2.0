package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"sync"
)

type safeStudents struct {
	names []string
	marks []int
	mu    *sync.Mutex
	err   error
}

func (s *safeStudents) GetMark(index int) {
	url := fmt.Sprintf("http://localhost:8082/mark?name=%s", s.names[index])

	client := &http.Client{}

	req, err := http.NewRequest("Get", url, nil)
	if err != nil {
		s.err = err
	}

	resp, err := client.Do(req)
	if err != nil {
		s.err = err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		s.err = err
	}

	mark, err := strconv.Atoi(string(data))
	if err != nil {
		s.err = err
	}

	s.marks[index] = mark
}

func NewStudents(names []string) *safeStudents {
	marks := make([]int, len(names))
	return &safeStudents{names, marks, &sync.Mutex{}, nil}
}

func AverageMark(marks []int) int {
	sum := 0
	for _, mark := range marks {
		sum += mark
	}
	return int(sum / len(marks))
}

func CompareList(names []string) (map[string]string, error) {
	sts := NewStudents(names)

	wg := &sync.WaitGroup{}

	size := len(names)
	wg.Add(size)
	for i := 0; i < size; i++ {
		go func(index int) {
			sts.GetMark(index)
			wg.Done()
		}(i)
	}

	wg.Wait()

	if sts.err != nil {
		return nil, sts.err
	}

	avM := AverageMark(sts.marks)

	studentsRank := make(map[string]string, size)
	for i, st := range sts.names {
		switch {
		case sts.marks[i] > avM:
			studentsRank[st] = ">"
		case sts.marks[i] < avM:
			studentsRank[st] = "<"
		case sts.marks[i] == avM:
			studentsRank[st] = "="
		}
	}

	return studentsRank, nil
}
