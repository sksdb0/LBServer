package classificationmanager

import (
	"config"
	"dbmanager"
	"encoding/json"
	"io"
	"lebangnet"
	"lebangproto"
	"logger"
	"net/http"
	"processor/common"

	"gopkg.in/mgo.v2/bson"
)

func Init() {
	lebangnet.RouteRegister("/getsubclassification", GetSubClassification)
}

func GetSubClassification(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	buf := make([]byte, req.ContentLength)
	common.GetBuffer(req, buf)

	var reqdata lebangproto.GetClassificationReq
	if !common.Unmarshal(buf, &reqdata) {
		return
	}
	logger.PRINTLINE(reqdata.GetTypename())

	var response lebangproto.GetClassificationRes
	if dbmanager.GetMongo().Find(config.DB().DBName, config.DB().CollMap["subclassification"],
		bson.M{"classification": reqdata.GetTypename()}, nil, &response.Classification) {
	} else {
		response.Errorcode = "no subclassification"
		logger.PRINTLINE(reqdata.GetTypename(), "no classification: ")
	}

	sendbuf, err := json.Marshal(response)
	if err != nil {
		logger.PRINTLINE("Marshal response error: ", err)
		return
	}
	io.WriteString(w, string(sendbuf))
}
