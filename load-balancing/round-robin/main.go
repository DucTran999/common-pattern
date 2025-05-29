package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os/signal"
	loadbalancer "patterns/load-balancing/load_balancer"
	"patterns/utils"
	"syscall"
	"time"

	"github.com/rs/zerolog/log"
)

func main() {
	s1 := utils.NewSimpleHTTPServer("localhost", 11111, 1, 2)
	go func() {
		if err := s1.Start(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Error().Str("err", err.Error()).Msg("Server 1 stopped")
		}
	}()

	s2 := utils.NewSimpleHTTPServer("localhost", 11112, 2, 2)
	go func() {
		if err := s2.Start(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Error().Str("err", err.Error()).Msg("Server 2 stopped")
		}
	}()

	s3 := utils.NewSimpleHTTPServer("localhost", 11113, 3, 2)
	go func() {
		if err := s3.Start(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Error().Str("err", err.Error()).Msg("Server 3 stopped")
		}
	}()

	time.Sleep(time.Second * 3)
	servers := []*utils.SimpleHTTPServer{s1, s2, s3}
	StartLoadBalancer(servers)

	go AutoSendRequest()

	GracefulShutdown(s1.Stop, s2.Stop, s3.Stop)
}

func AutoSendRequest() {
	for i := range 20 {
		c := http.Client{}
		go func() {
			endpoint := fmt.Sprintf("http://localhost:8080/req/%d", i)
			resp, err := c.Get(endpoint)
			if err != nil {
				log.Error().Str("err", err.Error()).Msg("make request error")
				return
			}

			defer func() {
				if err := resp.Body.Close(); err != nil {
					log.Error().Err(err).Msg("failed to close response body")
				}
			}()

			body, _ := io.ReadAll(resp.Body)
			fmt.Println(string(body))
		}()

		time.Sleep(time.Second * 1)
	}
}

func StartLoadBalancer(targets []*utils.SimpleHTTPServer) {
	lb, err := loadbalancer.NewLoadBalancer("localhost", 8080, targets, loadbalancer.RoundRobin)
	if err != nil {
		log.Fatal().Str("err", err.Error()).Msg("failed to init loadbalancer")
	}

	go func() {
		if err := lb.Start(); err != nil {
			log.Error().Str("error", err.Error()).Msg("err start load balancer")
		}
	}()
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
