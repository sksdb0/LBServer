package processor

import (
	"config"
	"httprouter"
	"logger"
	"net/http"
	"processor/classification"
	"processor/errandsclassification"
	"processor/feedback"
	"processor/ordermanager"
	"processor/picture"
	"processor/usermanager"
)

func Init() {
	router := httprouter.New()
	usermanager.Init(router)
	errandsclassification.Init(router)
	feedback.Init(router)
	ordermanager.Init(router)
	picture.Init(router)
	classification.Init(router)

	err := http.ListenAndServeTLS(config.Instance().HttpPort, config.Instance().CertFile, config.Instance().KeyFile, router)
	if err != nil {
		logger.PRINTLINE("ListenAndServe: ", err)
	}
}
