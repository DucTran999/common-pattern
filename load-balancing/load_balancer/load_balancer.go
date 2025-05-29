package loadbalancer

import (
	"net/http"
	"patterns/utils"
	"time"

	"github.com/rs/zerolog/log"
)

type Algorithm int

func (a Algorithm) String() string {
	switch a {
	case RoundRobin:
		return "Round Robin"
	case WeightedRoundRobin:
		return "Weighted Round Robin"
	default:
		return ""
	}
}

const (
	RoundRobin Algorithm = iota
	WeightedRoundRobin
)

type LoadBalancer interface {
	Start() error
}

type loadBalancer struct {
	handler http.Handler
	server  *http.Server
}

func NewLoadBalancer(
	targets []*utils.SimpleHTTPServer, alg Algorithm,
) (*loadBalancer, error) {

	hdl, err := NewLoadBalancerHandler(alg, targets)
	if err != nil {
		return nil, err
	}

	lb := &loadBalancer{
		server: &http.Server{
			Addr:         ":8080",
			Handler:      hdl,
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 10 * time.Second,
			IdleTimeout:  60 * time.Second,
		},
	}

	return lb, nil
}

func (lb *loadBalancer) Start() error {
	log.Info().Msg("Load Balancer running on :8080")

	if err := lb.server.ListenAndServe(); err != nil {
		log.Fatal().Str("err", err.Error()).Msg("failed to start load balancer")
		return err
	}

	return nil
}
