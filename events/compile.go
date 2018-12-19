package events

import (
	"io/ioutil"
	"lemon-robot-cli/logger"
	"lemon-robot-cli/utils/io"
	"os"
	"strings"
)

var groovyScriptDest = ""

func CompileGroovy() {
	//遍历当前工程的src/main/groovy文件夹，把所有groovy文件拼接
	logger.Info("Prepare to build groovy scripts")
	loc, _ := os.Getwd()
	if !isStandardPath(loc) {
		logger.Error("The current directory is not a standard Lemon Robot Task project!", nil)
		os.Exit(1)
	}
	// reset
	groovyScriptDest = ""
	logger.Info("Prepare to read your scripts")
	readGroovyScripts(loc)
	logger.Info("Prepare to write dest script file")
	writeGroovyScriptDest(loc+"/src/main/resources/__lemon_robot_task.groovy", groovyScriptDest)
}

/**
判断是否为标准的LemonRobot任务目录
 */
func isStandardPath(path string) bool {
	return io.PathExists(path + "/pom.xml")
}

func readGroovyScripts(groovyPath string) {
	rd, _ := ioutil.ReadDir(groovyPath)
	for _, fi := range rd {
		filePath := groovyPath + "/" + fi.Name()
		if fi.IsDir() {
			readGroovyScripts(filePath)
		} else {
			if strings.HasSuffix(fi.Name(), ".groovy") {
				logger.Info("Find groovy script file: " + filePath)
				file, _ := os.Open(filePath)
				scriptContent, _ := ioutil.ReadAll(file)
				groovyScriptDest += string(scriptContent) + "\n\n"
			}
		}
	}
}

func writeGroovyScriptDest(destPath string, script string) {
	if io.PathExists(destPath) {
		errRemove := os.Remove(destPath)
		if errRemove != nil {
			logger.Error("Can not remove exists dest file: "+destPath, errRemove)
			os.Exit(1)
		}
	}
	_, errCreate := os.Create(destPath)
	if errCreate != nil {
		logger.Error("Can not create the groovy dest file: "+destPath, errCreate)
		os.Exit(2)
	}
	errWrite := ioutil.WriteFile(destPath, []byte(script), 0600)
	if errWrite != nil {
		logger.Error("Can not write the groovy dest file: "+destPath, errWrite)
		os.Exit(3)
	}
	logger.Info("Write dest file complete! : " + destPath)
}
