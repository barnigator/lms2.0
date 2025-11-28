package main

import "sync"

type SafeMap struct {
	m   map[string]interface{}
	mux sync.Mutex
}

func (s *SafeMap) Get(key string) interface{} {
	s.mux.Lock()
	value := s.m[key]
	s.mux.Unlock()
	return value
}

func (s *SafeMap) Set(key string, value interface{}) {
	s.mux.Lock()
	s.m[key] = value
	s.mux.Unlock()
}

func NewSafeMap() *SafeMap {
	m := make(map[string]interface{})
	var mu sync.Mutex
	return &SafeMap{m, mu}
}
