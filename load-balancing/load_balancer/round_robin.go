package loadbalancer

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"patterns/utils"
	"sync/atomic"
)

type roundRobin struct {
	backends []*utils.SimpleHTTPServer
	counter  uint64
}

func NewRoundRobinAlg(targets []*utils.SimpleHTTPServer) (*roundRobin, error) {
	if len(targets) == 0 {
		return nil, ErrNoTargetServersFound
	}

	// Validate backend URLs
	for _, target := range targets {
		if target.GetUrl() == nil {
			return nil, ErrInvalidBackendUrl
		}
	}

	return &roundRobin{
		backends: targets,
	}, nil
}

func (lb *roundRobin) ForwardRequest(w http.ResponseWriter, r *http.Request) {
	nextUrl := lb.getNextBackend()

	// Log the next URL to which the request will be forwarded
	log.Printf("[INFO] load balancer forwarding request to: %v\n", nextUrl.String())

	// Create a reverse proxy for the next backend
	proxy := lb.getOrCreateProxy(&nextUrl)

	// Serve the request using the reverse proxy
	proxy.ServeHTTP(w, r)
}

func (lb *roundRobin) getOrCreateProxy(target *url.URL) *httputil.ReverseProxy {
	return httputil.NewSingleHostReverseProxy(target)
}

func (lb *roundRobin) getNextBackend() url.URL {
	idx := atomic.AddUint64(&lb.counter, 1)

	next := lb.backends[idx%uint64(len(lb.backends))]

	return *next.GetUrl()
}
