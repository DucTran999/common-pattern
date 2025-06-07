package loadbalancer

import (
	"net/url"
	"patterns/utils"
	"sort"
)

type weightedRoundRobin struct {
	backends      []*utils.SimpleHTTPServer
	currentWeight int
	currentIndex  int
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
		backends: targets,
	}

	wrr.sortBackendsByWeight()

	return wrr, nil
}

func (lb *weightedRoundRobin) GetNextBackend() url.URL {
	if lb.currentWeight == 0 {
		lb.currentIndex = lb.calculateNextIndex()
		lb.currentWeight = lb.backends[lb.currentIndex].Weight
	}

	lb.currentWeight--
	nextBackend := lb.backends[lb.currentIndex]
	return *nextBackend.GetUrl()
}

func (lb *weightedRoundRobin) sortBackendsByWeight() {
	// Sort backends by weight in descending order
	sort.SliceStable(lb.backends, func(i, j int) bool {
		return lb.backends[i].Weight < lb.backends[j].Weight
	})
}

func (lb *weightedRoundRobin) calculateNextIndex() int {
	current := lb.currentIndex + 1
	if current > len(lb.backends)-1 {
		current = 0
	}

	return current
}
