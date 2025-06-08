package utils

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
)

const (
	DefaultMaxConnection = 10
	DefaultMinConnection = 1
)

type SimpleHTTPServer struct {
	host string
	port int
	id   int

	weight     int
	connection int
	router     *mux.Router
	server     *http.Server
}

// Constructor function
func NewSimpleHTTPServer(host string, port int, id, weight int) *SimpleHTTPServer {
	return &SimpleHTTPServer{
		host:   host,
		port:   port,
		id:     id,
		weight: weight,
		router: mux.NewRouter(),
	}
}

func (s *SimpleHTTPServer) GetWeight() int {
	return s.weight
}

func (s *SimpleHTTPServer) GetConnection() int {
	return s.connection
}

func (s *SimpleHTTPServer) GetUrl() *url.URL {
	scheme := "http"

	if s.port == 443 {
		scheme = "https"
	}

	buildUrl := &url.URL{
		Scheme: scheme,
		Host:   fmt.Sprintf("%s:%d", s.host, s.port),
	}

	return buildUrl
}

// Start the server
func (s *SimpleHTTPServer) Start() error {
	s.routes()
	addr := fmt.Sprintf("%s:%d", s.host, s.port)
	s.server = &http.Server{
		Addr:              addr,
		Handler:           s.router,
		ReadHeaderTimeout: 500 * time.Millisecond,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	log.Info().Msgf("server running on http://%s , weight: %d", addr, s.weight)
	return s.server.ListenAndServe()
}

func (s *SimpleHTTPServer) Stop(ctx context.Context) error {
	defer func() {
		log.Info().Int("sever_id", s.id).Msg("shutdown")
	}()

	if s.server != nil {
		return s.server.Shutdown(ctx)
	}

	return nil
}

// Handler method
func (s *SimpleHTTPServer) reqHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	reqID := vars["req_id"]
	handleTime := time.Second * time.Duration(1/s.weight)
	time.Sleep(handleTime)

	// Simulate change the connection to this backend server
	s.connection = s.randomConnectionNumber(DefaultMinConnection, DefaultMaxConnection)

	if _, err := fmt.Fprintf(w, "Server %d, handle request %s!", s.id, reqID); err != nil {
		log.Error().Err(err).Msg("failed to write response")
	}
}

// Method to initialize routes
func (s *SimpleHTTPServer) routes() {
	s.router.HandleFunc("/req/{req_id}", s.reqHandler)
}

func (s *SimpleHTTPServer) randomConnectionNumber(min, max int) int {
	return rand.Intn(max-min+1) + min //nolint:gosec
}
