package main

import (
	"os"
	"restgo/src/routes"
)

func main() {
	var appPort = os.Getenv("appPort")

	if appPort == "" {
		appPort = "9090"
	}

	routes.SetupServer(appPort)
}
