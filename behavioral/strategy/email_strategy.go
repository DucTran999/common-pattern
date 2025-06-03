package strategy

import (
	"errors"
	"log"
	"strings"
)

var (
	ErrInvalidHost      = errors.New("invalid host name")
	ErrInvalidPort      = errors.New("invalid port number")
	ErrInvalidSender    = errors.New("invalid sender email address")
	ErrInvalidRecipient = errors.New("invalid recipient email address")
)

type emailStrategy struct {
	smtpHost string
	smtpPort int
	sender   string
}

func NewEmailStrategy(host string, port int, sender string) (*emailStrategy, error) {
	// Check if the host is empty or contains only whitespace
	if strings.Trim(host, " ") == "" {
		return nil, ErrInvalidHost
	}
	// Check if the port is within the valid range
	if port < 1 || port > 65535 {
		return nil, ErrInvalidPort
	}
	// A simple check for a valid email format
	if strings.Trim(sender, " ") == "" || !strings.Contains(sender, "@") {
		return nil, ErrInvalidSender
	}

	return &emailStrategy{
		smtpHost: host,
		smtpPort: port,
		sender:   sender,
	}, nil
}

func (e *emailStrategy) Send(to, message string) error {
	if strings.Trim(to, " ") == "" || !strings.Contains(to, "@") {
		return ErrInvalidRecipient
	}

	log.Printf(
		"Sending email to %s from %s via %s:%d with message: %s\n",
		to, e.sender, e.smtpHost, e.smtpPort, message,
	)

	return nil
}
