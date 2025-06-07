package loadbalancer

import (
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"patterns/utils"
)

type sourceIPHash struct {
	backends []*utils.SimpleHTTPServer
	ipHash   []*url.URL
}

func NewSourceIPHashAlgorithm(targets []*utils.SimpleHTTPServer) (*sourceIPHash, error) {
	if len(targets) == 0 {
		return nil, ErrNoTargetServersFound
	}

	// Validate backend URLs
	for _, target := range targets {
		if target.GetUrl() == nil {
			return nil, ErrInvalidBackendUrl
		}
	}

	sih := &sourceIPHash{backends: targets}
	sih.addBackendToIPHash()

	return sih, nil
}

func (lb *sourceIPHash) ForwardRequest(w http.ResponseWriter, r *http.Request) {
	ip := lb.getClientIP(r)

	nextUrl := lb.getNextBackend(ip)

	// Log the next URL to which the request will be forwarded
	log.Printf(
		"[INFO] source ip: %s -> load balancer forwarding request to: %v\n",
		ip, nextUrl.String(),
	)

	// Create a reverse proxy for the next backend
	proxy := lb.getOrCreateProxy(&nextUrl)

	// Serve the request using the reverse proxy
	proxy.ServeHTTP(w, r)
}

func (lb *sourceIPHash) getClientIP(r *http.Request) string {
	if ip := r.Header.Get("X-Forwarded-For"); ip != "" {
		return ip
	}

	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		log.Printf("[ERROR] failed to get client ip")
	}

	return host
}

func (lb *sourceIPHash) addBackendToIPHash() {
	for _, backend := range lb.backends {
		lb.ipHash = append(lb.ipHash, backend.GetUrl())
	}
}

func (lb *sourceIPHash) getNextBackend(sourceIP string) url.URL {
	idx := lb.simpleHash(sourceIP, len(lb.backends))
	return *lb.ipHash[idx]
}

func (lb *sourceIPHash) simpleHash(s string, buckets int) int {
	hash := 0
	for _, char := range s {
		hash += int(char) // Add ASCII value (or rune value for Unicode)
	}
	return hash % buckets // Modulus to fit bucket range
}

func (lb *sourceIPHash) getOrCreateProxy(target *url.URL) *httputil.ReverseProxy {
	return httputil.NewSingleHostReverseProxy(target)
}
