package queue

import (
	"errors"
	"sync"
)

var ErrQueueEmpty = errors.New("queue empty")

type Queue interface {
	Len() int
	Enqueue(val int)
	Dequeue() (val int, err error)
}

type queue struct {
	list  []int
	mutex sync.Mutex
}

func NewQueue(cap int) Queue {
	return &queue{
		list:  make([]int, 0, cap),
		mutex: sync.Mutex{},
	}
}

func (q *queue) Len() int {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	return len(q.list)
}

func (q *queue) Enqueue(val int) {
	q.mutex.Lock()
	q.list = append(q.list, val)
	q.mutex.Unlock()
}

func (q *queue) Dequeue() (int, error) {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	if len(q.list) == 0 {
		return 0, ErrQueueEmpty
	}

	val := q.list[0]
	q.list = q.list[1:]

	return val, nil
}
