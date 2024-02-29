package main

import (
	"github.com/dulguundd/logError-lib/logger"
	"os"
	"restAPIServer/app/driving"
)

func main() {
	logger.Info("Starting the application...")
	exePath, err := os.Executable()
	if err != nil {
		panic(err)
	}
	driving.Start(exePath)
}
