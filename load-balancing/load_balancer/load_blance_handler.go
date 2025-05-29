package loadbalancer

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"patterns/utils"
)

type AlgorithmImplementer interface {
	GetNextBackend() url.URL
}

type loadBalanceHandler struct {
	targets       []*utils.SimpleHTTPServer
	algorithmImpl AlgorithmImplementer
}

func NewLoadBalancerHandler(
	alg Algorithm, targets []*utils.SimpleHTTPServer,
) (*loadBalanceHandler, error) {
	hdl := &loadBalanceHandler{
		targets: targets,
	}

	algorithmImpl, err := hdl.getAlgorithmImpl(alg)
	if err != nil {
		return nil, err
	}
	hdl.algorithmImpl = algorithmImpl

	if err = hdl.validateConfig(); err != nil {
		return nil, err
	}

	return hdl, nil
}

func (lb *loadBalanceHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	nextUrl := lb.algorithmImpl.GetNextBackend()
	log.Printf("Forwarding request to: %v\n", nextUrl.String())

	proxy := lb.getOrCreateProxy(&nextUrl)
	proxy.ServeHTTP(w, r)
}

func (lb *loadBalanceHandler) getOrCreateProxy(target *url.URL) *httputil.ReverseProxy {
	return httputil.NewSingleHostReverseProxy(target)
}

func (h *loadBalanceHandler) getAlgorithmImpl(alg Algorithm) (AlgorithmImplementer, error) {
	switch alg {
	case RoundRobin:
		return NewRoundRobinAlg(h.targets)
	case WeightedRoundRobin:
		return nil, nil
	default:
		return nil, ErrUnsupportedAlg
	}
}

func (lb *loadBalanceHandler) validateConfig() error {
	if len(lb.targets) == 0 {
		return ErrNoTargetServersFound
	}

	for _, s := range lb.targets {
		if s == nil {
			return ErrNoTargetServersFound
		}
	}

	return nil
}
