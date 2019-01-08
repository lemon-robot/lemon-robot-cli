package events

import (
	"lemon-robot-golang-commons/utils/io"
	"lemon-robot-golang-commons/utils/logger"
	"os"
)

func Clean(params []string) {
	CleanParams()
	CleanGroovy()
}

func CleanParams() {
	loc, _ := os.Getwd()
	destFullPath := loc + DestParamsPath
	err := os.RemoveAll(destFullPath)
	if err != nil {
		logger.Error("Clean parameters dest folder error", err)
		os.Exit(1)
	}
	logger.Info("Clean parameters dest folder complete!")
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
