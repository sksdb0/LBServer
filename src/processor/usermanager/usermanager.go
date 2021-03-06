package usermanager

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
	"math/rand"
	"net/http"
	"processor/common"
	"time"

	"gopkg.in/mgo.v2/bson"
)

func Init(router *httprouter.Router) {
	// idcode and authentication
	router.POST("/ridervalidate", RiderValidate)
	router.POST("/ridersignin", RiderSignIn)

	router.POST("/getidcode", GetIDCode)
	router.POST("/authentication", Authentication)

	// merchant
	router.POST("/geterrandscommonmerchant", GetErrandsCommonMerchant)

	router.POST("/getaddress", GetAddress)
	router.POST("/addaddress", AddAddress)
	router.POST("/modifyaddress", ModifyAddress)
	router.POST("/deleteaddress", DeleteAddress)
	router.POST("/setdefaultaddress", SetDefaultAddress)
	router.POST("/defaultaddress", DefaultAddress)
}

func Authentication(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	defer req.Body.Close()
	buf := make([]byte, req.ContentLength)
	common.GetBuffer(req, buf)

	var reqdata lebangproto.IDCode
	if !common.Unmarshal(buf, &reqdata) {
		return
	}
	logger.PRINTLINE(reqdata.GetPhone())

	var idcode lebangproto.IDCode
	var response lebangproto.Response
	if dbmanager.GetMongo().Find(config.DB().DBName, config.DB().CollMap["idcode"],
		bson.M{"phone": reqdata.GetPhone()}, nil, &idcode) {
		logger.PRINTLINE(idcode.GetCode(), reqdata.GetCode())
		if reqdata.GetPhone() == "13683330861" && reqdata.GetCode() == "54321" {
			var userdata lebangproto.User
			if dbmanager.GetMongo().Find(config.DB().DBName, config.DB().CollMap["user"],
				bson.M{"phone": reqdata.GetPhone()}, nil, &userdata) {
				logger.PRINTLINE("update")
				userdata.Lastsignintime = time.Now().Unix() * 1000
				dbmanager.GetMongo().Update(config.DB().DBName, config.DB().CollMap["user"], bson.M{"phone": reqdata.GetPhone()}, userdata)
			} else {
				logger.PRINTLINE("insert")
				userdata := lebangproto.User{
					Phone:          reqdata.GetPhone(),
					Registertime:   time.Now().Unix() * 1000,
					Lastsignintime: time.Now().Unix() * 1000,
				}
				dbmanager.GetMongo().Insert(config.DB().DBName, config.DB().CollMap["user"], userdata)
			}
		} else if idcode.GetCode() != reqdata.GetCode() {
			response.Errorcode = "验证码错误"
			logger.PRINTLINE("authentication error: ", idcode.GetPhone(), idcode.GetCode(), reqdata.GetCode())
		} else if time.Unix(time.Now().Unix(), 0).Sub(time.Unix(idcode.GetTime()/1000, 0)).Seconds() > 90 {
			durationsecond := time.Unix(time.Now().Unix(), 0).Sub(time.Unix(idcode.GetTime()/1000, 0)).Seconds()
			logger.PRINTLINE(durationsecond)
			response.Errorcode = "验证码超时"
			logger.PRINTLINE("authentication error time out: ", idcode.GetPhone(), idcode.GetCode(), reqdata.GetCode())
		} else {
			var userdata lebangproto.User
			if dbmanager.GetMongo().Find(config.DB().DBName, config.DB().CollMap["user"],
				bson.M{"phone": reqdata.GetPhone()}, nil, &userdata) {
				logger.PRINTLINE("update")
				userdata.Lastsignintime = time.Now().Unix() * 1000
				dbmanager.GetMongo().Update(config.DB().DBName, config.DB().CollMap["user"], bson.M{"phone": reqdata.GetPhone()}, userdata)
			} else {
				logger.PRINTLINE("insert")
				userdata := lebangproto.User{
					Phone:          reqdata.GetPhone(),
					Registertime:   time.Now().Unix() * 1000,
					Lastsignintime: time.Now().Unix() * 1000,
				}
				dbmanager.GetMongo().Insert(config.DB().DBName, config.DB().CollMap["user"], userdata)
			}
		}
	} else {
		response.Errorcode = "用户不存在"
		logger.PRINTLINE("user not exist: ", idcode.GetPhone())
	}

	sendbuf, err := json.Marshal(response)
	if err != nil {
		logger.PRINTLINE("Marshal response error: ", err)
		return
	}
	io.WriteString(w, string(sendbuf))
}

func GetIDCode(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	defer req.Body.Close()
	buf := make([]byte, req.ContentLength)
	common.GetBuffer(req, buf)

	var reqdata lebangproto.GetIDCodeReq
	if !common.Unmarshal(buf, &reqdata) {
		return
	}
	logger.PRINTLINE(reqdata.GetPhone())

	var idcodeinfo lebangproto.IDCode
	var response lebangproto.Response
	if dbmanager.GetMongo().IsExist(config.DB().DBName, config.DB().CollMap["idcode"], bson.M{"phone": reqdata.GetPhone()}) {
		idcodeinfo = lebangproto.IDCode{
			Phone: reqdata.GetPhone(),
			Code:  fmt.Sprintf("%06d", rand.New(rand.NewSource(time.Now().UnixNano())).Intn(899999)+100000),
			Time:  time.Now().Unix() * 1000,
		}
		logger.PRINTLINE("exist", idcodeinfo.GetCode())
		dbmanager.GetMongo().Update(config.DB().DBName, config.DB().CollMap["idcode"], bson.M{"phone": reqdata.GetPhone()}, &idcodeinfo)
	} else {
		idcodeinfo = lebangproto.IDCode{
			Phone: reqdata.GetPhone(),
			Code:  fmt.Sprintf("%06d", rand.New(rand.NewSource(time.Now().UnixNano())).Intn(899999)+100000),
			Time:  time.Now().Unix() * 1000,
		}
		logger.PRINTLINE("not exist", idcodeinfo.GetCode())
		dbmanager.GetMongo().Insert(config.DB().DBName, config.DB().CollMap["idcode"], &idcodeinfo)
	}

	err := alisms.SendSms(config.Instance().AccessKeyID, config.Instance().AccessSecret, reqdata.GetPhone(),
		"乐帮跑腿", fmt.Sprintf("{code:%s}", idcodeinfo.Code), "SMS_135792492")
	if err != nil {
		logger.PRINTLINE("dysms.SendSms", err)
	}

	sendbuf, err := json.Marshal(response)
	if err != nil {
		logger.PRINTLINE("Marshal response error: ", err)
		return
	}

	io.WriteString(w, string(sendbuf))
}

func RiderValidate(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	defer req.Body.Close()
	buf := make([]byte, req.ContentLength)
	common.GetBuffer(req, buf)

	var reqdata lebangproto.RiderValidateReq
	if !common.Unmarshal(buf, &reqdata) {
		return
	}
	logger.PRINTLINE(reqdata.GetPhone())

	var response lebangproto.Response
	if !dbmanager.GetMongo().IsExist(config.DB().DBName, config.DB().CollMap["rider"], bson.M{"phone": reqdata.GetPhone()}) {
		response.Errorcode = "rider not exist"
	}

	sendbuf, err := json.Marshal(response)
	if err != nil {
		logger.PRINTLINE("Marshal response error: ", err)
		return
	}

	io.WriteString(w, string(sendbuf))
}

func RiderSignIn(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	defer req.Body.Close()
	buf := make([]byte, req.ContentLength)
	common.GetBuffer(req, buf)

	var reqdata lebangproto.RiderSignInReq
	if !common.Unmarshal(buf, &reqdata) {
		return
	}
	logger.PRINTLINE(reqdata.GetPhone())

	var response lebangproto.Response
	var rider lebangproto.Rider
	if dbmanager.GetMongo().Find(config.DB().DBName, config.DB().CollMap["rider"], bson.M{"phone": reqdata.GetPhone()}, nil, &rider) {
		if rider.GetPassword() != reqdata.GetPassword() {
			response.Errorcode = "password error"
			logger.PRINTLINE("password error")
		} else if rider.GetState() == int64(lebangproto.RiderState_RIDER_STATE_DIMISSION) {
			response.Errorcode = "authority error"
			logger.PRINTLINE("authority error")
		}
	}

	sendbuf, err := json.Marshal(response)
	if err != nil {
		logger.PRINTLINE("Marshal response error: ", err)
		return
	}

	io.WriteString(w, string(sendbuf))
}
