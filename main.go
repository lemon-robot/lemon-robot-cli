package main

import (
	"fmt"
	"lemon-robot-cli/events"
	"lemon-robot-cli/logger"
	"os"
)

var eventsMapping = map[string]func(){
	"build":          events.Build,
	"compile.groovy": events.CompileGroovy}

func main() {
	printLogo()
	if len(os.Args) <= 1 {
		logger.Error("Please tell me what you want to do, for example:")
		logger.Error("lr build")
		logger.Error("lr run")
		os.Exit(-1)
	}
	eventName := os.Args[1]
	_, containEvent := eventsMapping[eventName]
	if containEvent {
		loc, _ := os.Getwd()
		logger.Info("Current path: " + loc)
		logger.Info("Invoke event: " + os.Args[1])
		eventsMapping[os.Args[1]]()
	} else {
		logger.Error("The event not support: " + eventName)
	}

}

func printLogo() {
	fmt.Println(`
 _                               ______       _                
| |                             (_____ \     | |           _   
| |      _____ ____   ___  ____  _____) )___ | |__   ___ _| |_ 
| |     | ___ |    \ / _ \|  _ \|  __  // _ \|  _ \ / _ (_   _)
| |_____| ____| | | | |_| | | | | |  \ \ |_| | |_) ) |_| || |_ 
|_______)_____)_|_|_|\___/|_| |_|_|   |_\___/|____/ \___/  \__)

	`)
	fmt.Printf("🍋 %c[0;0;32mLemon Robot Cli - Version: 1.0.0 - https://www.lemonit.cn%c[0m 🍋\n\n", 0x1B, 0x1B)
}
