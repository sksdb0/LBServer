package ordermanager

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
	"time"

	"gopkg.in/mgo.v2/bson"
)

func Init() {
	lebangnet.RouteRegister("/neworder", NewOrder)
	lebangnet.RouteRegister("/getorder", GetOrder)
	lebangnet.RouteRegister("/", Span)
}

func Span(w http.ResponseWriter, req *http.Request) {
	logger.PRINTLINE("aaaaaaaaaaa")
}

func NewOrder(w http.ResponseWriter, req *http.Request) {
	logger.PRINTLINE("DefaultAddress")

	defer req.Body.Close()
	buf := make([]byte, req.ContentLength)
	common.GetBuffer(req, buf)

	var neworder lebangproto.Order
	if !common.Unmarshal(buf, &neworder) {
		return
	}

	var userinfo lebangproto.UserInfo
	var response lebangproto.Response

	if dbmanager.GetMongo().Find(config.DB().DBName, config.DB().CollMap["user"], bson.M{"phone": neworder.GetPhone()}, nil, &userinfo) {
		dbmanager.GetMongo().Insert(config.DB().DBName, config.DB().CollMap["order"], neworder)
		userinfo.Ordertimes += 1

		logger.PRINTLINE(time.Unix(neworder.GetOrdertime()/1000, 0))
	} else {
		response.Errorcode = "user not exist"
		logger.PRINTLINE("user not exist: ", neworder.GetPhone())
	}

	sendbuf, err := json.Marshal(response)
	if err != nil {
		logger.PRINTLINE("Marshal response error: ", err)
		return
	}
	io.WriteString(w, string(sendbuf))
}

func GetOrder(w http.ResponseWriter, req *http.Request) {
	logger.PRINTLINE("GetOrder")

	defer req.Body.Close()
	buf := make([]byte, req.ContentLength)
	common.GetBuffer(req, buf)

	var userinfo lebangproto.UserInfo
	if !common.Unmarshal(buf, &userinfo) {
		return
	}

	logger.PRINTLINE(userinfo)
	var response lebangproto.GetOrderRes

	if dbmanager.GetMongo().FindAll(config.DB().DBName, config.DB().CollMap["order"],
		bson.M{"phone": userinfo.GetPhone()}, "", nil, &response.Order) {
		if len(response.Order) == 0 {
			response.Errorcode = "no order"
			logger.PRINTLINE("no order: ", userinfo.GetPhone())
		}
	} else {
		response.Errorcode = "user not exist"
		logger.PRINTLINE("user not exist: ", userinfo.GetPhone())
	}

	sendbuf, err := json.Marshal(response)
	if err != nil {
		logger.PRINTLINE("Marshal response error: ", err)
		return
	}
	io.WriteString(w, string(sendbuf))
}
