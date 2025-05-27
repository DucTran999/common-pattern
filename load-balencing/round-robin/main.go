package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os/signal"
	"patterns/load-balencing/round-robin/components"
	"patterns/utils"
	"syscall"
	"time"

	"github.com/rs/zerolog/log"
)

func main() {
	s1 := utils.NewSimpleHTTPServer("localhost", 11111, 1)
	go func() {
		if err := s1.Start(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Error().Str("err", err.Error()).Msg("Server 1 stopped")
		}
	}()

	s2 := utils.NewSimpleHTTPServer("localhost", 11112, 2)
	go func() {
		if err := s2.Start(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Error().Str("err", err.Error()).Msg("Server 2 stopped")
		}
	}()

	s3 := utils.NewSimpleHTTPServer("localhost", 11113, 3)
	go func() {
		if err := s3.Start(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Error().Str("err", err.Error()).Msg("Server 3 stopped")
		}
	}()

	time.Sleep(time.Second * 3)
	StartLoadBalancer()

	go AutoSendRequest()

	GracefulShutdown(s1.Stop, s2.Stop, s3.Stop)
}

func AutoSendRequest() {
	for i := range 20 {
		go func() {
			c := http.Client{}
			endpoint := fmt.Sprintf("http://localhost:8080/req/%d", i)
			resp, err := c.Get(endpoint)
			if err != nil {
				log.Error().Str("err", err.Error()).Msg("make request error")
			}
			defer func() { _ = resp.Body.Close() }()

			body, _ := io.ReadAll(resp.Body)
			fmt.Println(string(body))
		}()

		time.Sleep(time.Second * 1)
	}
}

func StartLoadBalancer() {
	targets := []string{
		"http://localhost:11111",
		"http://localhost:11112",
		"http://localhost:11113",
	}

	lb, err := components.NewLoadBalancer(targets)
	if err != nil {
		log.Fatal().Str("err", err.Error()).Msg("failed to init loadbalancer")
	}

	go func() {
		log.Info().Msg("Load Balancer running on :8080")
		server := &http.Server{
			Addr:         ":8080",
			Handler:      lb,
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 10 * time.Second,
			IdleTimeout:  60 * time.Second,
		}

		if err := server.ListenAndServe(); err != nil {
			log.Fatal().Str("err", err.Error()).Msg("failed to start load balancer")
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
