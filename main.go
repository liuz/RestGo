package main

import (
	"os"
	"restgo/src/routes"
)

func main() {
	var appPort = os.Getenv("appPort")

	if appPort == "" {
		appPort = "1111"
	}

	routes.SetupServer(appPort)
}
