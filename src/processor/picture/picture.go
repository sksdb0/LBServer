package picture

import (
	"config"
	"dbmanager"
	"fmt"
	"httprouter"
	"logger"
	"net/http"
	"sync"
)

var instance *Picture

type Picture struct {
	idmutex sync.RWMutex
}

func Init(router *httprouter.Router) {
	instance = &Picture{}
	router.GET("/picture/:name", instance.GetPicture)
}

func (this *Picture) GetPicture(w http.ResponseWriter, req *http.Request, param httprouter.Params) {
	this.idmutex.Lock()
	defer this.idmutex.Unlock()

	filename := fmt.Sprintf("%s.jpg", param.ByName("name"))
	result, data := dbmanager.GetMongo().SearchFile(config.DB().DBName, config.DB().CollMap["picture"], filename)
	if result {
		w.Write(data)
	} else {
		logger.PRINTLINE("faile", filename)
	}
}
