package lebangnet

import (
	"config"
	"logger"
	"net/http"
)

var instance *Net

type Net struct {
}

func Init() {
	logger.PRINTLINE("Net Init")

}

func RouteRegister(partten string, f http.HandlerFunc) {
	http.HandleFunc(partten, f)
}

func Run() {
	err := http.ListenAndServeTLS(config.Instance().HttpPort, config.Instance().CertFile, config.Instance().KeyFile, nil)
	if err != nil {
		logger.PRINTLINE("ListenAndServe: ", err)
	}
}
