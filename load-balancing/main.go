package main

import (
	"flag"
	"os"

	"patterns/load-balancing/app"

	"github.com/rs/zerolog"
)

func main() {
	// Initialize zerolog with ConsoleWriter for pretty terminal output
	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()

	appName := flag.String("app-name", "round-robin", "Load balance app to run")
	flag.Parse()

	switch *appName {
	case "round-robin":
		app.RunRoundRobinApp(logger)
	default:
		logger.Fatal().Msg("[ERROR] app not available")
	}
}
