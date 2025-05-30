package main

import (
	"context"
	"os"
	"os/signal"
	"patterns/load-balancing/components"
	loadbalancer "patterns/load-balancing/load_balancer"
	"syscall"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	// Initialize zerolog with ConsoleWriter for pretty terminal output
	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()

	// Initialize the backend builder and configure number of backend servers
	backendBuilder := components.NewBackendBuilder(logger)
	backendBuilder.SetNumberOfBackends(5)

	// Build the backend servers
	backends, err := backendBuilder.Build()
	if err != nil {
		log.Fatal().Msgf("failed when build backends: %v", err)
	}

	// Create a new load balancer on localhost:8080 using the backends and round-robin algorithm
	lb, err := loadbalancer.NewLoadBalancer("localhost", 8080, backends, loadbalancer.RoundRobin)
	if err != nil {
		log.Fatal().Msgf("failed to init loadbalancer: %v", err)
	}

	// Start the load balancer asynchronously
	if err := lb.Start(); err != nil {
		log.Error().Msgf("failed to start load balancer: %v", err)
	}

	// Initialize a request sender component and start sending requests asynchronously
	rs := components.NewRequestSender(20)
	go rs.SendNow()

	// Wait for a graceful shutdown signal and stop the first backend cleanly
	GracefulShutdown(backendBuilder.ShutdownAllBackends)
}

// GracefulShutdown handles OS signals and performs a graceful shutdown of the server.
func GracefulShutdown(shutdownTasks ...func(ctx context.Context) error) {
	const shutdownTimeout = 5 * time.Second

	// Listen for SIGINT or SIGTERM
	shutdownCtx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	<-shutdownCtx.Done()
	log.Info().Msg("shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	cleanExit := true
	for _, task := range shutdownTasks {
		if err := task(ctx); err != nil {
			log.Warn().Err(err).Msg("shutdown task error")
			cleanExit = false
		}
	}

	if cleanExit {
		log.Info().Msg("server shut down cleanly")
	} else {
		log.Warn().Msg("server encountered errors during shutdown")
	}
}
