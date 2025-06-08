package loadbalancer

import (
	"net/http"
	"patterns/utils"
)

type AlgorithmImplementer interface {
	ForwardRequest(w http.ResponseWriter, r *http.Request)
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
	lb.algorithmImpl.ForwardRequest(w, r)
}

func (h *loadBalanceHandler) getAlgorithmImpl(alg Algorithm) (AlgorithmImplementer, error) {
	switch alg {
	case RoundRobin:
		return NewRoundRobinAlg(h.targets)
	case WeightedRoundRobin:
		return NewWeightedRoundRobinAlg(h.targets)
	case SourceIPHash:
		return NewSourceIPHashAlgorithm(h.targets)
	case LeastConnection:
		return NewLeastConnectionAlg(h.targets)
	case LowestLatency:
		return NewLowestLatencyAlg(h.targets)
	case ResourceBase:
		return NewResourceBaseLoadAlg(h.targets)
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
