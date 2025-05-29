package loadbalancer

import (
	"net/url"
	"patterns/utils"
	"sync/atomic"
)

type roundRobin struct {
	backends []*utils.SimpleHTTPServer
	counter  uint64
}

func NewRoundRobinAlg(targets []*utils.SimpleHTTPServer) (*roundRobin, error) {
	return &roundRobin{backends: targets}, nil
}

// Round robin
func (lb *roundRobin) GetNextBackend() url.URL {
	idx := atomic.AddUint64(&lb.counter, 1)

	next := lb.backends[idx%uint64(len(lb.backends))]

	return *next.GetUrl()
}
