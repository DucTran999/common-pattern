package observer

import (
	"slices"
	"sync"
)

type subscriber struct {
	name string
	news []string
	mu   sync.RWMutex
}

func NewSubscriber(name string) *subscriber {
	return &subscriber{name: name}
}

func (s *subscriber) UpdateNews(news string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.news = append(s.news, news)
}

func (s *subscriber) GetNews() []string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return slices.Clone(s.news) // Return a copy of the news slice
}
