package observer

import (
	"slices"
	"sync"
)

type Subscriber interface {
	UpdateNews(newPost string)
}

type blogs struct {
	mutex     sync.Mutex
	observers []Subscriber
}

func NewBlogs() *blogs {
	return &blogs{
		mutex:     sync.Mutex{},
		observers: make([]Subscriber, 0),
	}
}

func (s *blogs) Subscribe(o Subscriber) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	// Check for duplicates
	if !slices.Contains(s.observers, o) {
		s.observers = append(s.observers, o)
	}
}

func (s *blogs) AddNews(newPost string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	for _, o := range s.observers {
		o.UpdateNews(newPost)
	}
}
