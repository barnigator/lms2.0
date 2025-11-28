package main

import "sync"

type Counter struct {
	value int
	mu    sync.RWMutex
}

type Сount interface {
	Increment()    // увеличение счётчика на единицу
	GetValue() int // получение текущего значения
}

func (c *Counter) Increment() {
	c.mu.Lock()
	c.value++
	c.mu.Unlock()
}

func (c *Counter) GetValue() int {
	c.mu.Lock()
	res := c.value
	c.mu.Unlock()
	return res
}
