package utils

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
)

type SimpleHTTPServer struct {
	Host   string
	Port   int
	ID     int
	Weight int
	Router *mux.Router
	Server *http.Server
}

// Constructor function
func NewSimpleHTTPServer(host string, port int, id, weight int) *SimpleHTTPServer {
	return &SimpleHTTPServer{
		Host:   host,
		Port:   port,
		ID:     id,
		Weight: weight,
		Router: mux.NewRouter(),
	}
}

func (s *SimpleHTTPServer) GetWeight() int {
	return s.Weight
}

func (s *SimpleHTTPServer) GetUrl() *url.URL {
	scheme := "http"

	if s.Port == 443 {
		scheme = "https"
	}

	buildUrl := &url.URL{
		Scheme: scheme,
		Host:   fmt.Sprintf("%s:%d", s.Host, s.Port),
	}

	return buildUrl
}

// Start the server
func (s *SimpleHTTPServer) Start() error {
	s.routes()
	addr := fmt.Sprintf("%s:%d", s.Host, s.Port)
	s.Server = &http.Server{
		Addr:              addr,
		Handler:           s.Router,
		ReadHeaderTimeout: 500 * time.Millisecond,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	fmt.Printf("Server running on http://%s\n", addr)
	return s.Server.ListenAndServe()
}

func (s *SimpleHTTPServer) Stop(ctx context.Context) error {
	defer func() {
		log.Info().Int("sever_id", s.ID).Msg("shutdown")
	}()

	if s.Server != nil {
		return s.Server.Shutdown(ctx)
	}

	return nil
}

// Handler method
func (s *SimpleHTTPServer) reqHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	reqID := vars["req_id"]
	handleTime := time.Second * time.Duration(1/s.Weight)
	time.Sleep(time.Second * handleTime)

	if _, err := fmt.Fprintf(w, "Server %d, handle request %s!", s.ID, reqID); err != nil {
		log.Error().Err(err).Msg("failed to write response")
	}
}

// Method to initialize routes
func (s *SimpleHTTPServer) routes() {
	s.Router.HandleFunc("/req/{req_id}", s.reqHandler)
}
