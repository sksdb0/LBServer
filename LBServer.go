package main

import (
	"config"
	"dbmanager"
	"lebangnet"
	"logger"
	"processor/ordermanager"
	"processor/usermanager"
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

	lebangnet.Init()

	usermanager.Init()
	ordermanager.Init()

	lebangnet.Run()
}
