package app

import (
	"patterns/load-balancing/components"
	loadbalancer "patterns/load-balancing/load_balancer"

	"github.com/rs/zerolog"
)

func RunResourceBaseApp(logger zerolog.Logger) {
	// Initialize the backend builder and configure number of backend servers
	backendBuilder := components.NewBackendBuilder(logger)
	backendBuilder.SetNumberOfBackends(5)

	// Build the backend servers
	backends, err := backendBuilder.Build()
	if err != nil {
		logger.Fatal().Msgf("failed when build backends: %v", err)
	}

	// Create a new load balancer on localhost:8080 using the backends and using resource base algorithm
	lb, err := loadbalancer.NewLoadBalancer("localhost", 8080, backends, loadbalancer.ResourceBase)
	if err != nil {
		logger.Fatal().Msgf("failed to init loadbalancer: %v", err)
	}

	// Start the load balancer asynchronously
	if err := lb.Start(); err != nil {
		logger.Fatal().Msgf("failed to start load balancer: %v", err)
	}

	// Initialize a request sender component and start sending requests asynchronously
	rs := components.NewRequestSender(20)
	go rs.SendNow()

	// Wait for a graceful shutdown signal and stop the first backend cleanly
	components.GracefulShutdown(logger, backendBuilder.ShutdownAllBackends)
}
