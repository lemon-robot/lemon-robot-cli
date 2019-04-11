package events

import (
	"lemon-robot-golang-commons/logger"
	"lemon-robot-golang-commons/utils/lrucmd"
	"lemon-robot-golang-commons/utils/lruio"
	"os"
)

func Build(params []string) {
	buildPath, _ := os.Getwd()
	if !checkIsLRStandardProject(buildPath) {
		logger.Error("The current directory of the console is not a standard Lemon Robot Task project", nil)
	} else {
		// is standard task project
		logger.Info("Building lemon robot task project, please wait...")
		err := lrucmd.NoDisplayExec("mvn package")
		if err != nil {
			logger.Error("An error occurred during the build task", err)
			os.Exit(1)
		}
		logger.Info("Build success!")
	}
}

func checkIsLRStandardProject(path string) bool {
	return lruio.PathExists(path + "/" + "pom.xml")
}
