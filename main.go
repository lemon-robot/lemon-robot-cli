package main

import (
	"fmt"
	"lemon-robot-cli/events"
	"lemon-robot-golang-commons/utils/logger"
	"os"
	"reflect"
	//"reflect"
)

var eventsMapping = map[string]interface{}{
	"build":   events.Build,
	"run":     events.Run,
	"compile": events.Compile,
	"clean":   events.Clean}

func main() {
	printBaseInfo()
	if len(os.Args) <= 1 {
		logger.Error("Please tell me what you want to do, for example:", nil)
		logger.Error("lr build", nil)
		logger.Error("lr run", nil)
		os.Exit(-1)
	}
	eventName := os.Args[1]
	_, containEvent := eventsMapping[eventName]
	if containEvent {
		loc, _ := os.Getwd()
		logger.Info("Current path: " + loc)
		logger.Info("Invoke event: " + os.Args[1])
		event := reflect.ValueOf(eventsMapping[os.Args[1]])
		paramsCount := len(os.Args) - 2
		params := make([]string, paramsCount)
		for i := 0; i < paramsCount; i++ {
			params[i] = os.Args[i+2]
		}
		event.Call([]reflect.Value{reflect.ValueOf(params)})
	} else {
		logger.Error("The event not support: "+eventName, nil)
	}

}

func printBaseInfo() {
	fmt.Println(`
 _                               ______       _                
| |                             (_____ \     | |           _   
| |      _____ ____   ___  ____  _____) )___ | |__   ___ _| |_ 
| |     | ___ |    \ / _ \|  _ \|  __  // _ \|  _ \ / _ (_   _)
| |_____| ____| | | | |_| | | | | |  \ \ |_| | |_) ) |_| || |_ 
|_______)_____)_|_|_|\___/|_| |_|_|   |_\___/|____/ \___/  \__)

	`)
	fmt.Printf("-- %c[0;0;32mLemon Robot Cli - Version: 1.0.0 - https://www.lemonit.cn%c[0m --\n\n", 0x1B, 0x1B)
}
