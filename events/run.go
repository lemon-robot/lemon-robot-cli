package events

import (
	"lemon-robot-golang-commons/logger"
	"lemon-robot-golang-commons/utils/lrucmd"
)

func Run(params []string) {
	Build(params)
	// then run
	logger.Info("Start invoke lemon robot task...")
	lrucmd.OnlyConsoleDisplayExec("java -jar target/lemon_robot_task_pkg.jar")
	logger.Info("Lemon robot task execute complete!")
}
