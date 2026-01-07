package main

import (
	"fmt"
	"io"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"sync"
)

type safeStudents struct {
	students []string
	marks    []int
	mu       *sync.Mutex
	err      error
}

func (s *safeStudents) GetMark(index int) {
	s.mu.Lock()
	url := fmt.Sprintf("http://localhost:8082/mark?name=%s", s.students[index])
	s.mu.Unlock()

	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
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

	num, err := strconv.Atoi(string(data))
	if err != nil {
		s.err = err
	}

	s.marks[index] = num
}

func NewStudents(names []string) *safeStudents {
	marks := make([]int, len(names))

	sort.Slice(names, func(i, j int) bool {
		return names[i] < names[j]
	})

	return &safeStudents{names, marks, &sync.Mutex{}, nil}
}

func AverageMark(marks []int) int {
	sum := 0
	for _, mark := range marks {
		sum += mark
	}
	return int(sum / len(marks))
}

func BestStudents(names []string) (string, error) {
	wg := &sync.WaitGroup{}

	students := NewStudents(names)

	size := len(names)

	wg.Add(size)
	for i := 0; i < size; i++ {
		go func(index int) {
			students.GetMark(index)
			wg.Done()
		}(i)
	}

	wg.Wait()

	if students.err != nil {
		return "", students.err
	}

	avMark := AverageMark(students.marks)

	champList := make([]string, 0, len(names))
	for i, student := range students.students {
		if students.marks[i] > avMark {
			champList = append(champList, student)
		}
	}

	return strings.Join(champList, ","), nil
}
