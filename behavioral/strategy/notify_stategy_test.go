package strategy_test

import (
	"patterns/behavioral/strategy"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestChanelStrategyString(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		strategy strategy.ChannelStrategy
		expected string
	}{
		{"email strategy", 0, "Email"},
		{"sms strategy", 1, "SMS"},
		{"unknown", 99, "Unknown"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			actual := tt.strategy.String()

			assert.Equal(t, tt.expected, actual)
		})
	}
}

func Test_NotificationManager(t *testing.T) {
	t.Parallel()
	// Initialize the notification manager
	nm := strategy.NewNotificationManager()

	smsStrategy, err := strategy.NewSMSStrategy("sms.example.com", "+1234567890")
	if err != nil {
		t.Fatalf("Failed to create SMS strategy: %v", err)
	}
	emailStrategy, err := strategy.NewEmailStrategy("smtp.example.com", 587, "no-reply@gmail.com")
	if err != nil {
		t.Fatalf("Failed to create Email strategy: %v", err)
	}

	nm.RegisterNotifier(strategy.SMSStrategy, smsStrategy)
	nm.RegisterNotifier(strategy.EmailStrategy, emailStrategy)

	err = nm.SendNotification(strategy.SMSStrategy, "+0987654321", "Test SMS message")
	require.NoError(t, err)

	err = nm.SendNotification(strategy.EmailStrategy, "ab@gmail.com", "Test Email message")
	require.NoError(t, err)

	err = nm.SendNotification(99, "Invalid Channel", "This should fail")
	require.Error(t, err)
}
