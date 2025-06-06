package functionaloptions

import (
	"errors"
	"fmt"
)

type Option func(*server) error

var (
	ErrInvalidPort = errors.New("invalid port number")
)

func WithHost(host string) Option {
	return func(s *server) error {
		s.host = host
		return nil
	}
}

func WithPort(port int) Option {
	return func(s *server) error {
		if port <= 0 || port > 65535 {
			return fmt.Errorf("failed to set port: %d, %w", port, ErrInvalidPort)
		}

		s.port = port
		return nil
	}
}

func WithTLS(tls bool) Option {
	return func(s *server) error {
		s.tls = tls
		return nil
	}
}
