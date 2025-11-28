package main

import "sync"

type ConcurrentQueue struct {
	queue []interface{}
	mutex sync.Mutex
}

type Queue interface {
	Enqueue(element interface{}) // положить элемент в очередь
	Dequeue() interface{}        // забрать первый элемент из очереди
}

func (q *ConcurrentQueue) Enqueue(element interface{}) {
	q.mutex.Lock()
	q.queue = append(q.queue, element)
	q.mutex.Unlock()
}

func (q *ConcurrentQueue) Dequeue() interface{} {
	q.mutex.Lock()
	elem := q.queue[0]
	q.queue = q.queue[1:]
	q.mutex.Unlock()
	return elem
}
