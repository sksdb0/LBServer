package ordermanager

import (
	"alisms"
	"config"
	"dbmanager"
	"encoding/json"
	"fmt"
	"httprouter"
	"io"
	"lebangproto"
	"logger"
	"net/http"
	"processor/common"
	"processor/usermanager"
	"sync"
	"time"
	"xinge"

	"gopkg.in/mgo.v2/bson"
)

var ordermanager *OrderManager

type OrderManager struct {
	idmutex sync.RWMutex
}

func Init(router *httprouter.Router) {
	ordermanager = &OrderManager{}
	router.POST("/neworder", ordermanager.NewOrder)
	router.POST("/getorder", ordermanager.GetOrder)
	router.POST("/getallorder", ordermanager.GetAllOrder)
	router.POST("/cancelorder", ordermanager.CancelOrder)
	router.POST("/modifyorder", ordermanager.ModifyOrder)

	router.POST("/", ordermanager.Span)
}

func (this *OrderManager) Span(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	this.idmutex.Lock()
	defer this.idmutex.Unlock()

	logger.PRINTLINE("aaaaaaaaaaa")
}

func (this *OrderManager) NewOrder(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
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
		now := time.Now()
		if now.Hour() > 21 || now.Hour() < 7 {
			response.Errorcode = "营业时间为早7点至晚10点"
			logger.PRINTLINE("营业时间为早7点至晚10点: ", reqdata.GetPhone())
		} else {
			reqdata.Ordertime = time.Now().Unix() * 1000
			dbmanager.GetMongo().Insert(config.DB().DBName, config.DB().CollMap["order"], reqdata)
			go func() {
				for _, v := range config.Instance().XingeToken {
					xingeres := xinge.PushTokenAndroid(config.Instance().XingeAccessId, config.Instance().XingeSecretKey,
						"新订单", reqdata.GetPhone(), v)
					logger.PRINTLINE(xingeres)
				}
			}()

			user.Ordertimes += 1
			usermanager.UpdateErrandsCommonMerchant(reqdata.GetPhone(), reqdata.GetMerchant())
			dbmanager.GetMongo().Update(config.DB().DBName, config.DB().CollMap["user"], bson.M{"phone": reqdata.GetPhone()}, &user)

			err := alisms.SendSms(config.Instance().AccessKeyID, config.Instance().AccessSecret, "15940647652",
				"乐帮跑腿", fmt.Sprintf("{code:%s}", "9999"), "SMS_135792492")
			if err != nil {
				logger.PRINTLINE("dysms.SendSms", err)
			}

			err = alisms.SendSms(config.Instance().AccessKeyID, config.Instance().AccessSecret, "13683330861",
				"乐帮跑腿", fmt.Sprintf("{code:%s}", "9999"), "SMS_135792492")
			if err != nil {
				logger.PRINTLINE("dysms.SendSms", err)
			}
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

func (this *OrderManager) GetOrder(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
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

func (this *OrderManager) GetAllOrder(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
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
	if reqdata.GetPhone() != config.Instance().Manager {
		response.Errorcode = "请申请管理员权限"
		logger.PRINTLINE("请申请管理员权限")
	} else {
		if dbmanager.GetMongo().FindAll(config.DB().DBName, config.DB().CollMap["order"], nil, "-ordertime", nil, &response.Order) {
			if len(response.Order) == 0 {
				response.Errorcode = "no order"
				logger.PRINTLINE("no order")
			}
		}
	}

	sendbuf, err := json.Marshal(response)
	if err != nil {
		logger.PRINTLINE("Marshal response error: ", err)
		return
	}
	io.WriteString(w, string(sendbuf))
}

func (this *OrderManager) ModifyOrder(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
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

func (this *OrderManager) CancelOrder(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
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
	if dbmanager.GetMongo().Find(config.DB().DBName, config.DB().CollMap["order"],
		bson.M{"phone": reqdata.GetPhone(), "ordertime": reqdata.GetOrdertime()}, nil, &order) {
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
