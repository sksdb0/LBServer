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
	"processor/usermanager"
	"time"

	"gopkg.in/mgo.v2/bson"
)

func Init() {
	lebangnet.RouteRegister("/neworder", NewOrder)
	lebangnet.RouteRegister("/getorder", GetOrder)
	lebangnet.RouteRegister("/cancelorder", CancelOrder)

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

	var reqdata lebangproto.Order
	if !common.Unmarshal(buf, &reqdata) {
		return
	}

	var userinfo lebangproto.UserInfo
	var response lebangproto.Response

	if dbmanager.GetMongo().Find(config.DB().DBName, config.DB().CollMap["user"], bson.M{"phone": reqdata.GetPhone()}, nil, &userinfo) {
		dbmanager.GetMongo().Insert(config.DB().DBName, config.DB().CollMap["order"], reqdata)
		userinfo.Ordertimes += 1

		usermanager.UpdateErrandsCommonMerchant(reqdata.GetPhone(), reqdata.GetMerchant())
		logger.PRINTLINE(time.Unix(reqdata.GetOrdertime()/1000, 0))
	} else {
		response.Errorcode = "user not exist"
		logger.PRINTLINE("user not exist: ", reqdata.GetPhone())
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

func CancelOrder(w http.ResponseWriter, req *http.Request) {
	logger.PRINTLINE("CancelOrder")

	defer req.Body.Close()
	buf := make([]byte, req.ContentLength)
	common.GetBuffer(req, buf)

	var reqdata lebangproto.CancelOrder
	if !common.Unmarshal(buf, &reqdata) {
		return
	}

	logger.PRINTLINE(reqdata)
	var response lebangproto.Response
	var order lebangproto.Order
	if dbmanager.GetMongo().Find(config.DB().DBName, config.DB().CollMap["order"], bson.M{"phone": reqdata.GetPhone()}, nil, &order) {
		dbmanager.GetMongo().Remove(config.DB().DBName, config.DB().CollMap["order"], order)
	} else {
		response.Errorcode = "order not exist"
		logger.PRINTLINE("order not exist: ", reqdata.GetPhone(), reqdata.GetOrdertime())
	}

	sendbuf, err := json.Marshal(response)
	if err != nil {
		logger.PRINTLINE("Marshal response error: ", err)
		return
	}
	io.WriteString(w, string(sendbuf))
}
