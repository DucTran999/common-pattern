package functionaloptions

type server struct {
	host string
	port int
	tls  bool
}

func NewHTTPServer(options ...Option) (*server, error) {
	s := &server{
		host: "localhost", // default host
		port: 8080,        // default port
		tls:  false,       // default TLS setting
	}

	// Apply each option to the server instance
	for _, opt := range options {
		if err := opt(s); err != nil {
			return nil, err
		}
	}

	return s, nil
}

func (s *server) GetServerConfig() (string, int, bool) {
	return s.host, s.port, s.tls
}
