package events

import (
	"lemon-robot-cli/logger"
	"os"
)

func Build() {
	buildPath, _ := os.Getwd()
	logger.Info("Build path: " + buildPath)
}
