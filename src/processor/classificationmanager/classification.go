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
	logger.PRINTLINE("GetSubClassification")

	defer req.Body.Close()
	buf := make([]byte, req.ContentLength)
	common.GetBuffer(req, buf)

	var classitype lebangproto.ClassificationType
	if !common.Unmarshal(buf, &classitype) {
		return
	}

	var response lebangproto.GetClassificationRes

	if dbmanager.GetMongo().Find(config.DB().DBName, config.DB().CollMap["subclassification"],
		bson.M{"classification": classitype.GetTypename()}, nil, &response.Classification) {
	} else {
		response.Errorcode = "no classification"
		logger.PRINTLINE("no classification: ", classitype.GetTypename())
	}

	sendbuf, err := json.Marshal(response)
	if err != nil {
		logger.PRINTLINE("Marshal response error: ", err)
		return
	}
	io.WriteString(w, string(sendbuf))
}
