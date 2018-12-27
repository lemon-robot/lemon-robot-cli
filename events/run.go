package events

import (
	"lemon-robot-cli/logger"
	"lemon-robot-cli/utils/cmd"
)

func Run(params []string) {
	Build(params)
	// then run
	logger.Info("Start invoke lemon robot task...")
	cmd.OnlyConsoleDisplayExec("java -jar target/lemon_robot_task_pkg.jar")
	logger.Info("Lemon robot task execute complete!")
}
