package main

import (
	"flag"
	"log"
	"patterns/load-balancing/app"
)

func main() {
	appName := flag.String("app-name", "round-robin", "Load balance app to run")
	flag.Parse()

	switch *appName {
	case "round-robin":
		app.RunRoundRobinApp()
	default:
		log.Println("[ERROR] app not available")
	}
}
