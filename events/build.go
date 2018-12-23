package events

import (
	"lemon-robot-cli/logger"
	"os"
)

func Build(params []string) {
	buildPath, _ := os.Getwd()
	logger.Info("Build path: " + buildPath)
}
