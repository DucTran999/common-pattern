package strategy_test

import (
	"patterns/behavioral/strategy"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_NewSmsStrategy(t *testing.T) {
	testCases := []struct {
		name        string
		gateway     string
		sender      string
		expectError error
	}{
		{
			name:        "Valid SMS Strategy",
			gateway:     "sms.example.com",
			sender:      "+1234567890",
			expectError: nil,
		},
		{
			name:        "Invalid SMS Gateway",
			gateway:     "",
			sender:      "+1234567890",
			expectError: strategy.ErrInvalidSmsHost,
		},
		{
			name:        "Invalid Phone Number",
			gateway:     "sms.example.com",
			sender:      "1234567890",
			expectError: strategy.ErrInvalidPhoneNumber,
		},
		{
			name:        "Empty SMS Gateway and Phone Number",
			gateway:     "",
			sender:      "",
			expectError: strategy.ErrInvalidSmsHost,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := strategy.NewSMSStrategy(tc.gateway, tc.sender)

			require.ErrorIs(t, err, tc.expectError)
		})
	}
}

func Test_SmsStrategy_SendNotification(t *testing.T) {
	smsStrategy, err := strategy.NewSMSStrategy("sms.example.com", "+1234567890")
	require.NoError(t, err)
	testCases := []struct {
		name        string
		to          string
		message     string
		expectError error
	}{
		{
			name:        "Valid SMS Notification",
			to:          "+1987654321",
			message:     "Hello, this is a test SMS!",
			expectError: nil,
		},
		{
			name:        "Invalid Phone Number",
			to:          "987654321",
			message:     "This should fail",
			expectError: strategy.ErrInvalidPhoneNumber,
		},
		{
			name:        "Empty Phone Number",
			to:          "",
			message:     "This should also fail",
			expectError: strategy.ErrInvalidPhoneNumber,
		},
		{
			name:        "Empty Message",
			to:          "+1987654321",
			message:     "",
			expectError: nil, // Empty message is allowed
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := smsStrategy.SendNotification(tc.to, tc.message)

			require.ErrorIs(t, err, tc.expectError)
		})
	}
}
