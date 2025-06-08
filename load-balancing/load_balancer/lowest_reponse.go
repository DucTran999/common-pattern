package loadbalancer

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"patterns/utils"
	"sync"
	"time"
)

type lowestResponseAlg struct {
	backends   []*utils.SimpleHTTPServer
	proxyCache sync.Map
}

func NewLowestResponseAlg(targets []*utils.SimpleHTTPServer) (*lowestResponseAlg, error) {
	if len(targets) == 0 {
		return nil, ErrNoTargetServersFound
	}

	// Validate backend URLs
	for _, target := range targets {
		if target.GetUrl() == nil {
			return nil, ErrInvalidBackendUrl
		}
	}

	lr := &lowestResponseAlg{
		backends:   targets,
		proxyCache: sync.Map{},
	}

	return lr, nil
}

func (lb *lowestResponseAlg) ForwardRequest(w http.ResponseWriter, r *http.Request) {
	nextUrl := lb.getNextBackend()

	// Log the next URL to which the request will be forwarded
	log.Printf("[INFO] load balancer forwarding request to: %v\n", nextUrl.String())

	// Create a reverse proxy for the next backend
	proxy := lb.getOrCreateProxy(nextUrl)

	// Serve the request using the reverse proxy
	proxy.ServeHTTP(w, r)
}

func (lb *lowestResponseAlg) getOrCreateProxy(target *url.URL) *httputil.ReverseProxy {
	key := target.String()
	if proxy, ok := lb.proxyCache.Load(key); ok {
		return proxy.(*httputil.ReverseProxy)
	}

	proxy := httputil.NewSingleHostReverseProxy(target)
	lb.proxyCache.Store(key, proxy)

	return proxy
}

func (lb *lowestResponseAlg) getNextBackend() *url.URL {
	if len(lb.backends) == 1 {
		return lb.backends[0].GetUrl()
	}

	minLatency := lb.backends[0].Latency()
	backendIdx := 0
	backendLatency := []time.Duration{minLatency}

	for idx := 1; idx < len(lb.backends); idx++ {
		backend := lb.backends[idx]
		backendLatency = append(backendLatency, backend.Latency())

		if minLatency > backend.Latency() {
			minLatency = backend.Latency()
			backendIdx = idx
		}
	}

	log.Printf(
		"[INFO] backend latency: %v, select: %d, connection: %d \n",
		backendLatency, backendIdx, minLatency,
	)

	return lb.backends[backendIdx].GetUrl()
}
