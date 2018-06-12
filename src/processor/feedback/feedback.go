package feedback

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
)

func Init(router *httprouter.Router) {
	// idcode and authentication
	router.POST("/feedback", FeedBack)

}

func FeedBack(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	defer req.Body.Close()

	buf := make([]byte, req.ContentLength)
	common.GetBuffer(req, buf)

	var reqdata lebangproto.FeedBack
	if !common.Unmarshal(buf, &reqdata) {
		return
	}
	logger.PRINTLINE(reqdata.GetPhone())

	var response lebangproto.Response
	if !dbmanager.GetMongo().Insert(config.DB().DBName, config.DB().CollMap["feedback"], reqdata) {
		response.Errorcode = "insert error"
	}

	sendbuf, err := json.Marshal(response)
	if err != nil {
		logger.PRINTLINE("Marshal response error: ", err)
		return
	}
	io.WriteString(w, string(sendbuf))
}
