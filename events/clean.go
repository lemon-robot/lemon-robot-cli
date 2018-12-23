package events

import (
	"lemon-robot-cli/logger"
	"lemon-robot-cli/utils/io"
	"os"
)

func Clean(params []string) {
	CleanParams()
	CleanGroovy()
}

func CleanParams() {

}

func CleanGroovy() {
	loc, _ := os.Getwd()
	cleanPath := loc + DestGroovyPath
	if io.PathExists(cleanPath) {
		_ = os.Remove(cleanPath)
		logger.Info("Clean file : " + cleanPath)
	}
	logger.Info("Clean task complete!")
}
