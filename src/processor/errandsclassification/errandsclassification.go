package errandsclassification

import (
	"config"
	"dbmanager"
	"encoding/json"
	"httprouter"
	"io"
	"lebangproto"
	"logger"
	"net/http"
	"processor/common"

	"gopkg.in/mgo.v2/bson"
)

func Init(router *httprouter.Router) {
	router.POST("/geterrandssubclassification", GetErrandsSubClassification)
	router.POST("/geterrandsclassification", GetErrandsClassification)
}

func GetErrandsClassification(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	defer req.Body.Close()
	buf := make([]byte, req.ContentLength)
	common.GetBuffer(req, buf)
	logger.PRINTLINE("GetErrandsClassification")

	var response lebangproto.GetErrandsClassificationRes
	if dbmanager.GetMongo().Find(config.DB().DBName, config.DB().CollMap["errandsclassification"],
		bson.M{"classification": "main"}, nil, &response.Classification) {
	} else {
		response.Errorcode = "no errandsclassification"
		logger.PRINTLINE("no errandsclassification: ")
	}

	sendbuf, err := json.Marshal(response)
	if err != nil {
		logger.PRINTLINE("Marshal response error: ", err)
		return
	}
	io.WriteString(w, string(sendbuf))
}

func GetErrandsSubClassification(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	defer req.Body.Close()
	buf := make([]byte, req.ContentLength)
	common.GetBuffer(req, buf)

	var reqdata lebangproto.GetErrandsClassificationReq
	if !common.Unmarshal(buf, &reqdata) {
		return
	}
	logger.PRINTLINE(reqdata.GetTypename())

	var response lebangproto.GetErrandsClassificationRes
	if dbmanager.GetMongo().Find(config.DB().DBName, config.DB().CollMap["errandssubclassification"],
		bson.M{"classification": reqdata.GetTypename()}, nil, &response.Classification) {
	} else {
		response.Errorcode = "no errandssubclassification"
		logger.PRINTLINE(reqdata.GetTypename(), "no errandssubclassification: ")
	}

	sendbuf, err := json.Marshal(response)
	if err != nil {
		logger.PRINTLINE("Marshal response error: ", err)
		return
	}
	io.WriteString(w, string(sendbuf))
}
