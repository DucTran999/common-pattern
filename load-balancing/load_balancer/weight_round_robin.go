package loadbalancer

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"patterns/utils"
	"sort"
	"sync"
)

type weightedRoundRobin struct {
	backends      []*utils.SimpleHTTPServer
	currentWeight int
	currentIndex  int
	proxyCache    sync.Map
}

func NewWeightedRoundRobinAlg(targets []*utils.SimpleHTTPServer) (*weightedRoundRobin, error) {
	if len(targets) == 0 {
		return nil, ErrNoTargetServersFound
	}

	// Validate backend URLs
	for _, target := range targets {
		if target.GetUrl() == nil {
			return nil, ErrInvalidBackendUrl
		}
	}

	wrr := &weightedRoundRobin{
		backends:   targets,
		proxyCache: sync.Map{},
	}

	wrr.setupBackend()

	return wrr, nil
}

func (lb *weightedRoundRobin) ForwardRequest(w http.ResponseWriter, r *http.Request) {
	nextUrl := lb.getNextBackend()

	// Log the next URL to which the request will be forwarded
	log.Printf("[INFO] load balancer forwarding request to: %v\n", nextUrl.String())

	// Create a reverse proxy for the next backend
	proxy := lb.getOrCreateProxy(&nextUrl)

	// Serve the request using the reverse proxy
	proxy.ServeHTTP(w, r)
}

func (lb *weightedRoundRobin) getOrCreateProxy(target *url.URL) *httputil.ReverseProxy {
	key := target.String()
	if proxy, ok := lb.proxyCache.Load(key); ok {
		return proxy.(*httputil.ReverseProxy)
	}

	proxy := httputil.NewSingleHostReverseProxy(target)
	lb.proxyCache.Store(key, proxy)

	return proxy
}

func (lb *weightedRoundRobin) getNextBackend() url.URL {
	if lb.currentWeight == 0 {
		lb.currentIndex = lb.calculateNextIndex()
		lb.currentWeight = lb.backends[lb.currentIndex].Weight
	}

	lb.currentWeight--
	nextBackend := lb.backends[lb.currentIndex]
	return *nextBackend.GetUrl()
}

func (lb *weightedRoundRobin) setupBackend() {
	// Sort backends by weight in descending order
	sort.SliceStable(lb.backends, func(i, j int) bool {
		return lb.backends[i].Weight > lb.backends[j].Weight
	})

	// Initialize currentWeight and currentIndex
	if len(lb.backends) > 0 {
		lb.currentWeight = lb.backends[0].Weight
		lb.currentIndex = 0
	}
}

func (lb *weightedRoundRobin) calculateNextIndex() int {
	current := lb.currentIndex + 1
	if current > len(lb.backends)-1 {
		current = 0
	}

	return current
}
