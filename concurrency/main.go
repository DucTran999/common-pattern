//go:build ignore

package main

import (
	"flag"
	"os"
	"patterns/concurrency/app"

	"github.com/rs/zerolog"
)

func main() {
	// Initialize zerolog with ConsoleWriter for pretty terminal output
	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()

	appName := flag.String("app-name", "worker-pool", "Load concurrency app to run")
	flag.Parse()

	switch *appName {
	case "worker-pool":
		app.ReadCsvWithWorkerPool()
	case "generator":
		app.SecretConversationApp()
	default:
		logger.Fatal().Msg("[ERROR] app not available")
	}

}
