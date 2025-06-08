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

	appName := flag.String("app-name", "rr", "Load balance app to run")
	flag.Parse()

	switch *appName {
	case "rr":
		app.RunRoundRobinApp(logger)
	case "wrr": // weight round robin
		app.RunWeightRoundRobinApp(logger)
	case "sih": // source ip hash
		app.RunSourceIPhashApp(logger)
	case "lc": // least connection
		app.RunLeastConnectionApp(logger)
	case "ll": // lowest latency
		app.RunLowestLatencyApp(logger)
	case "rb": // resource base
		app.RunResourceBaseApp(logger)
	default:
		logger.Fatal().Msg("[ERROR] app not available")
	}
}
