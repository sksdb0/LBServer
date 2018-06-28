package classification

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
	"strconv"
	"strings"

	"gopkg.in/mgo.v2/bson"
)

func Init(router *httprouter.Router) {
	router.POST("/getsubclassification", GetSubClassification)
}

func GetSubClassification(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	defer req.Body.Close()
	buf := make([]byte, req.ContentLength)
	common.GetBuffer(req, buf)

	var reqdata lebangproto.GetSubClassificationViewReq
	if !common.Unmarshal(buf, &reqdata) {
		return
	}

	var response lebangproto.GetSubClassificationRes
	var classificationview lebangproto.ClassificationView
	if dbmanager.GetMongo().Find(config.DB().DBName, config.DB().CollMap["classificationview"],
		bson.M{"name": reqdata.GetName()}, nil, &classificationview) {

		typeids := strings.Split(classificationview.GetTypeids(), " ")
		for _, typeidstr := range typeids {
			typeid, _ := strconv.Atoi(typeidstr)
			classification := lebangproto.ClassificationRes{
				Name:           classificationview.Name,
				Typeid:         int32(typeid),
				Classification: make([]*lebangproto.SubClassification, 0),
			}

			if dbmanager.GetMongo().FindAll(config.DB().DBName, config.DB().CollMap["subclassification"],
				bson.M{"parenttypeid": typeid}, "typeid", nil, &classification.Classification) {
				response.Classification = append(response.Classification, &classification)
			} else {
				response.Errorcode = "no sub classificationview"
				logger.PRINTLINE(reqdata.GetName(), "no sub classificationview: ")
				break
			}
		}

	} else {
		response.Errorcode = "no classificationview"
		logger.PRINTLINE(reqdata.GetName(), "no classificationview: ")
	}

	sendbuf, err := json.Marshal(response)
	if err != nil {
		logger.PRINTLINE("Marshal response error: ", err)
		return
	}
	io.WriteString(w, string(sendbuf))
}
