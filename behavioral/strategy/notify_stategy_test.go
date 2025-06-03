package strategy_test

import (
	"patterns/behavioral/strategy"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_NotificationManager(t *testing.T) {
	// Initialize the notification manager
	nm, err := strategy.NewNotificationManager()
	if err != nil {
		t.Fatalf("Failed to create notification manager: %v", err)
	}

	smsStrategy, err := strategy.NewSMSStrategy("sms.example.com", "+1234567890")
	if err != nil {
		t.Fatalf("Failed to create SMS strategy: %v", err)
	}
	emailStrategy, err := strategy.NewEmailStrategy("smtp.example.com", 587, "no-reply@gmail.com")
	if err != nil {
		t.Fatalf("Failed to create Email strategy: %v", err)
	}

	nm.RegisterNotifier("sms", smsStrategy)
	nm.RegisterNotifier("email", emailStrategy)

	err = nm.NotifierMap["sms"].SendNotification("+1987654321", "Test SMS message")
	require.NoError(t, err)

	err = nm.NotifierMap["email"].SendNotification("hello@gmail.com", "Test Email message")
	require.NoError(t, err)
}
