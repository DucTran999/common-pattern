package loadbalancer

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"patterns/utils"
	"sync"
)

type resourceBaseLoadAlg struct {
	backends   []*utils.SimpleHTTPServer
	proxyCache sync.Map
}

func NewResourceBaseLoadAlg(targets []*utils.SimpleHTTPServer) (*resourceBaseLoadAlg, error) {
	if len(targets) == 0 {
		return nil, ErrNoTargetServersFound
	}

	// Validate backend URLs
	for _, target := range targets {
		if target.GetUrl() == nil {
			return nil, ErrInvalidBackendUrl
		}
	}

	rbl := &resourceBaseLoadAlg{
		backends:   targets,
		proxyCache: sync.Map{},
	}

	return rbl, nil
}

func (lb *resourceBaseLoadAlg) ForwardRequest(w http.ResponseWriter, r *http.Request) {
	nextUrl := lb.getNextBackend()

	// Log the next URL to which the request will be forwarded
	log.Printf("[INFO] load balancer forwarding request to: %v\n", nextUrl.String())

	// Create a reverse proxy for the next backend
	proxy := lb.getOrCreateProxy(nextUrl)

	// Serve the request using the reverse proxy
	proxy.ServeHTTP(w, r)
}

func (lb *resourceBaseLoadAlg) getOrCreateProxy(target *url.URL) *httputil.ReverseProxy {
	key := target.String()
	if proxy, ok := lb.proxyCache.Load(key); ok {
		return proxy.(*httputil.ReverseProxy)
	}

	proxy := httputil.NewSingleHostReverseProxy(target)
	lb.proxyCache.Store(key, proxy)

	return proxy
}

func (lc *resourceBaseLoadAlg) getNextBackend() *url.URL {
	// Only one backend server return it intermediately
	if len(lc.backends) == 1 {
		return lc.backends[0].GetUrl()
	}

	// Lookup the backends got lowest cpu load
	minCPULoad := lc.backends[0].GetCPULoad()
	backendIdx := 0
	backendCPUs := []float64{minCPULoad}

	for idx := 1; idx < len(lc.backends); idx++ {
		backend := lc.backends[idx]
		backendCPUs = append(backendCPUs, backend.GetCPULoad())

		if minCPULoad > lc.backends[idx].GetCPULoad() {
			minCPULoad = backend.GetCPULoad()
			backendIdx = idx
		}
	}

	log.Printf(
		"[INFO] backend connections: %v, select: %d, CPU load: %.2f \n",
		backendCPUs, backendIdx, minCPULoad,
	)

	return lc.backends[backendIdx].GetUrl()
}
