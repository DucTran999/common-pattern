package utils

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
)

type SimpleHTTPServer struct {
	Host   string
	Port   int
	ID     int
	Router *mux.Router
	Server *http.Server
}

// Constructor function
func NewSimpleHTTPServer(host string, port int, ID int) *SimpleHTTPServer {
	return &SimpleHTTPServer{
		Host:   host,
		Port:   port,
		ID:     ID,
		Router: mux.NewRouter(),
	}
}

// Handler method
func (s *SimpleHTTPServer) reqHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	reqID := vars["req_id"]
	time.Sleep(time.Second * 4)
	fmt.Fprintf(w, "Server %d, handle request %s!", s.ID, reqID)
}

// Method to initialize routes
func (s *SimpleHTTPServer) routes() {
	s.Router.HandleFunc("/req/{req_id}", s.reqHandler)
}

// Start the server
func (s *SimpleHTTPServer) Start() error {
	s.routes()
	addr := fmt.Sprintf("%s:%d", s.Host, s.Port)
	s.Server = &http.Server{
		Addr:    addr,
		Handler: s.Router,
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
