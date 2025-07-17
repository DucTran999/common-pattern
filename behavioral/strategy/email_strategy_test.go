package strategy_test

import (
	"patterns/behavioral/strategy"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_NewEmailStrategy(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name        string
		host        string
		port        int
		sender      string
		expectError error
	}{
		{
			name:        "Valid Email Strategy",
			host:        "smtp.example.com",
			port:        587,
			sender:      "daniel@gmail.com",
			expectError: nil,
		},
		{
			name:        "Invalid Host",
			host:        "",
			port:        587,
			expectError: strategy.ErrInvalidHost,
		},
		{
			name:        "Invalid Port",
			host:        "smtp.example.com",
			port:        70000,
			expectError: strategy.ErrInvalidPort,
		},
		{
			name:        "Invalid Sender Email",
			host:        "smtp.example.com",
			port:        587,
			sender:      "danielgmail.com",
			expectError: strategy.ErrInvalidSender,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			_, err := strategy.NewEmailStrategy(tc.host, tc.port, tc.sender)

			require.ErrorIs(t, err, tc.expectError)
		})
	}
}

func Test_EmailStrategy_SendNotification(t *testing.T) {
	t.Parallel()
	emailStrategy, err := strategy.NewEmailStrategy("smtp.example.com", 587, "no-reply@gmail.com")
	require.NoError(t, err)

	testCases := []struct {
		name        string
		to          string
		message     string
		expectError error
	}{
		{
			name:        "Valid Email Notification",
			to:          "daniel@gmail.com",
			message:     "Hello, this is a test email!",
			expectError: nil,
		},
		{
			name:        "Invalid Recipient Email",
			to:          "danielgmail.com",
			message:     "This should fail due to invalid email format",
			expectError: strategy.ErrInvalidRecipient,
		},
		{
			name:        "Empty Recipient Email",
			to:          "",
			message:     "This should fail due to empty recipient",
			expectError: strategy.ErrInvalidRecipient,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			err := emailStrategy.Send(tc.to, tc.message)

			require.ErrorIs(t, err, tc.expectError)
		})
	}
}
