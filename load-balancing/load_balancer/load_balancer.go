package loadbalancer

import (
	"fmt"
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
	port   int
	host   string
	server *http.Server
}

func NewLoadBalancer(
	host string, port int, targets []*utils.SimpleHTTPServer, alg Algorithm,
) (*loadBalancer, error) {

	hdl, err := NewLoadBalancerHandler(alg, targets)
	if err != nil {
		return nil, err
	}

	lb := &loadBalancer{
		host: host,
		port: port,
		server: &http.Server{
			Addr:         fmt.Sprintf("%s:%d", host, port),
			Handler:      hdl,
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 10 * time.Second,
			IdleTimeout:  60 * time.Second,
		},
	}

	return lb, nil
}

func (lb *loadBalancer) Start() error {
	log.Info().Msgf("Load Balancer running on %s:%d", lb.host, lb.port)

	if err := lb.server.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
