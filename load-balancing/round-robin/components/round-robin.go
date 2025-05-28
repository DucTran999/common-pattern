package components

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync/atomic"
)

type loadBalancer struct {
	backends []*url.URL
	counter  uint64
}

func NewLoadBalancer(targets []string) (*loadBalancer, error) {
	var urls []*url.URL

	for _, target := range targets {
		u, err := url.Parse(target)

		if err != nil {
			return nil, fmt.Errorf("invalid backend URL: %w", err)
		}
		urls = append(urls, u)
	}

	return &loadBalancer{backends: urls}, nil
}

func (lb *loadBalancer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	target := lb.getNextBackend()
	log.Printf("Forwarding request to: %s\n", target)

	proxy := lb.getOrCreateProxy(target)
	proxy.ServeHTTP(w, r)
}

// Round robin
func (lb *loadBalancer) getNextBackend() *url.URL {
	idx := atomic.AddUint64(&lb.counter, 1)

	return lb.backends[idx%uint64(len(lb.backends))]
}

func (lb *loadBalancer) getOrCreateProxy(target *url.URL) *httputil.ReverseProxy {
	return httputil.NewSingleHostReverseProxy(target)
}
