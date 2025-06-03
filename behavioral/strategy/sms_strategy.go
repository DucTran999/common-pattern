package strategy

import (
	"errors"
	"log"
	"strings"
)

var (
	ErrInvalidSmsHost     = errors.New("invalid sms host name")
	ErrInvalidPhoneNumber = errors.New("invalid phone number format")
)

type smsStrategy struct {
	smsGateway string
	sender     string
}

func NewSMSStrategy(gateway, sender string) (*smsStrategy, error) {
	if strings.TrimSpace(gateway) == "" {
		return nil, ErrInvalidSmsHost
	}
	// simple check for a valid phone number format
	if strings.TrimSpace(sender) == "" || !strings.HasPrefix(sender, "+") {
		return nil, ErrInvalidPhoneNumber
	}

	return &smsStrategy{
		smsGateway: gateway,
		sender:     sender,
	}, nil
}

func (s *smsStrategy) SendNotification(to, message string) error {
	if strings.TrimSpace(to) == "" || !strings.HasPrefix(to, "+") {
		return ErrInvalidPhoneNumber
	}

	// Simulate sending an SMS
	log.Printf(
		"Sending SMS to %s from %s via gateway %s with message: %s\n",
		to, s.sender, s.smsGateway, message,
	)

	return nil
}
