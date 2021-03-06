package main

import (
	"config"
	"dbmanager"
	"logger"
	"os"
	"processor"
)

func main() {
	// init logger
	if !logger.GetInstance().Init("LBServer.log") {
		return
	}
	defer logger.GetInstance().Close()

	if !config.Init() {
		logger.PRINTLINE("config init error")
		return
	}

	if !dbmanager.Init() {
		logger.PRINTLINE("dbmanager init error")
		return
	}
	defer dbmanager.Close()

	if len(os.Args) > 1 {
		if os.Args[1] == "init" {
			dbmanager.GetMongo().InitFile()
		}
	}

	processor.Init()
}
