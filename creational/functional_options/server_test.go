package functionaloptions_test

import (
	functionaloptions "patterns/creational/functional_options"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_NewServer(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name         string
		options      []functionaloptions.Option
		expectedHost string
		expectedPort int
		expectedTLS  bool
		wantErr      error
	}{
		{
			name:         "default options",
			options:      []functionaloptions.Option{},
			expectedHost: "localhost",
			expectedPort: 8080,
			expectedTLS:  false,
			wantErr:      nil,
		},
		{
			name: "valid host and port",
			options: []functionaloptions.Option{
				functionaloptions.WithHost("10.20.20.10"),
				functionaloptions.WithPort(80),
				functionaloptions.WithTLS(true),
			},
			expectedHost: "10.20.20.10",
			expectedPort: 80,
			expectedTLS:  true,
			wantErr:      nil,
		},
		{
			name: "invalid port",
			options: []functionaloptions.Option{
				functionaloptions.WithPort(70000), // Invalid port
			},
			wantErr: functionaloptions.ErrInvalidPort,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			server, err := functionaloptions.NewHTTPServer(tt.options...)
			require.ErrorIs(t, err, tt.wantErr)
			if err != nil {
				return // Skip further checks if there is an error
			}

			host, port, tls := server.GetServerConfig()
			require.Equal(t, tt.expectedHost, host)
			require.Equal(t, tt.expectedPort, port)
			require.Equal(t, tt.expectedTLS, tls)
		})
	}
}
