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
	"sync"

	"gopkg.in/mgo.v2/bson"
)

var ordermanager *OrderManager

type OrderManager struct {
	idmutex sync.RWMutex
}

func Init() {
	ordermanager = &OrderManager{}
	lebangnet.RouteRegister("/neworder", ordermanager.NewOrder)
	lebangnet.RouteRegister("/getorder", ordermanager.GetOrder)
	lebangnet.RouteRegister("/getallorder", ordermanager.GetAllOrder)
	lebangnet.RouteRegister("/cancelorder", ordermanager.CancelOrder)
	lebangnet.RouteRegister("/modifyorder", ordermanager.ModifyOrder)

	lebangnet.RouteRegister("/", ordermanager.Span)
}

func (this *OrderManager) Span(w http.ResponseWriter, req *http.Request) {
	this.idmutex.Lock()
	defer this.idmutex.Unlock()

	logger.PRINTLINE("aaaaaaaaaaa")
}

func (this *OrderManager) NewOrder(w http.ResponseWriter, req *http.Request) {
	this.idmutex.Lock()
	defer this.idmutex.Unlock()

	defer req.Body.Close()
	buf := make([]byte, req.ContentLength)
	common.GetBuffer(req, buf)

	var reqdata lebangproto.Order
	if !common.Unmarshal(buf, &reqdata) {
		return
	}
	logger.PRINTLINE(reqdata.GetPhone())

	var user lebangproto.User
	var response lebangproto.Response
	if dbmanager.GetMongo().Find(config.DB().DBName, config.DB().CollMap["user"], bson.M{"phone": reqdata.GetPhone()}, nil, &user) {
		dbmanager.GetMongo().Insert(config.DB().DBName, config.DB().CollMap["order"], reqdata)
		user.Ordertimes += 1

		usermanager.UpdateErrandsCommonMerchant(reqdata.GetPhone(), reqdata.GetMerchant())

		dbmanager.GetMongo().Update(config.DB().DBName, config.DB().CollMap["user"], bson.M{"phone": reqdata.GetPhone()}, &user)
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

func (this *OrderManager) GetOrder(w http.ResponseWriter, req *http.Request) {
	this.idmutex.Lock()
	defer this.idmutex.Unlock()

	defer req.Body.Close()
	buf := make([]byte, req.ContentLength)
	common.GetBuffer(req, buf)

	var reqdata lebangproto.GetOrderReq
	if !common.Unmarshal(buf, &reqdata) {
		return
	}
	logger.PRINTLINE(reqdata.GetPhone())

	var response lebangproto.GetOrderRes
	if dbmanager.GetMongo().FindAll(config.DB().DBName, config.DB().CollMap["order"],
		bson.M{"phone": reqdata.GetPhone()}, "-ordertime", nil, &response.Order) {
		if len(response.Order) == 0 {
			response.Errorcode = "no order"
			logger.PRINTLINE("no order: ", reqdata.GetPhone())
		}
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

func (this *OrderManager) GetAllOrder(w http.ResponseWriter, req *http.Request) {
	this.idmutex.Lock()
	defer this.idmutex.Unlock()

	logger.PRINTLINE("GetAllOrder")
	defer req.Body.Close()
	buf := make([]byte, req.ContentLength)
	common.GetBuffer(req, buf)

	var reqdata lebangproto.GetOrderReq
	if !common.Unmarshal(buf, &reqdata) {
		return
	}

	var response lebangproto.GetOrderRes
	if dbmanager.GetMongo().FindAll(config.DB().DBName, config.DB().CollMap["order"], nil, "-ordertime", nil, &response.Order) {
		if len(response.Order) == 0 {
			response.Errorcode = "no order"
			logger.PRINTLINE("no order")
		}
	}

	sendbuf, err := json.Marshal(response)
	if err != nil {
		logger.PRINTLINE("Marshal response error: ", err)
		return
	}
	io.WriteString(w, string(sendbuf))
}

func (this *OrderManager) ModifyOrder(w http.ResponseWriter, req *http.Request) {
	this.idmutex.Lock()
	defer this.idmutex.Unlock()

	defer req.Body.Close()
	buf := make([]byte, req.ContentLength)
	common.GetBuffer(req, buf)

	var reqdata lebangproto.ModifyOrderReq
	if !common.Unmarshal(buf, &reqdata) {
		return
	}
	logger.PRINTLINE(reqdata.GetPhone())

	var response lebangproto.Response
	if dbmanager.GetMongo().IsExist(config.DB().DBName, config.DB().CollMap["order"],
		bson.M{"phone": reqdata.GetPhone(), "ordertime": reqdata.GetOrder().GetOrdertime()}) {

		dbmanager.GetMongo().Update(config.DB().DBName, config.DB().CollMap["order"],
			bson.M{"phone": reqdata.GetPhone(), "ordertime": reqdata.GetOrder().GetOrdertime()}, reqdata.Order)
	} else {
		response.Errorcode = "order not exist"
		logger.PRINTLINE("order not exist: ", reqdata.GetPhone(), reqdata.GetOrder())
	}

	sendbuf, err := json.Marshal(response)
	if err != nil {
		logger.PRINTLINE("Marshal response error: ", err)
		return
	}
	io.WriteString(w, string(sendbuf))
}

func (this *OrderManager) CancelOrder(w http.ResponseWriter, req *http.Request) {
	this.idmutex.Lock()
	defer this.idmutex.Unlock()

	defer req.Body.Close()
	buf := make([]byte, req.ContentLength)
	common.GetBuffer(req, buf)

	var reqdata lebangproto.CancelOrderReq
	if !common.Unmarshal(buf, &reqdata) {
		return
	}
	logger.PRINTLINE(reqdata.GetPhone())

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
