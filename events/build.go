package events

import (
	"lemon-robot-golang-commons/utils/cmd"
	"lemon-robot-golang-commons/utils/io"
	"lemon-robot-golang-commons/utils/logger"
	"os"
)

func Build(params []string) {
	buildPath, _ := os.Getwd()
	if !checkIsLRStandardProject(buildPath) {
		logger.Error("The current directory of the console is not a standard Lemon Robot Task project", nil)
	} else {
		// is standard task project
		logger.Info("Building lemon robot task project, please wait...")
		err := cmd.NoDisplayExec("mvn package")
		if err != nil {
			logger.Error("An error occurred during the build task", err)
			os.Exit(1)
		}
		logger.Info("Build success!")
	}
}

func checkIsLRStandardProject(path string) bool {
	return io.PathExists(path + "/" + "pom.xml")
}
