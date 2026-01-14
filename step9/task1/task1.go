package main

import (
	"sync"
	"testing"
)

func TestSortIntegers(t *testing.T) {
	norm := true
	wg := &sync.WaitGroup{}
	numbers := []int{5, 3, 9, 6}
	SortIntegers(numbers)

	for i := 1; i < len(numbers); i++ {
		wg.Add(1)
		go func(flag *bool, i int) {
			if numbers[i] < numbers[i-1] {
				*flag = false
			}
			wg.Done()
		}(&norm, i)
	}

	wg.Wait()

	if !norm {
		t.Errorf("result: %t", norm)
	}
}
