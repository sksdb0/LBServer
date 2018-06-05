package usermanager

import (
	//	"alisms"
	"config"
	"dbmanager"
	"encoding/json"
	"fmt"
	"io"
	"lebangnet"
	"lebangproto"
	"logger"
	"math/rand"
	"net/http"
	"processor/common"
	"time"

	"gopkg.in/mgo.v2/bson"
)

func Init() {
	// idcode and authentication
	lebangnet.RouteRegister("/login", Login)
	lebangnet.RouteRegister("/getidcode", GetIDCode)
	lebangnet.RouteRegister("/authentication", Authentication)

	// merchant
	lebangnet.RouteRegister("/geterrandscommonmerchant", GetErrandsCommonMerchant)

	lebangnet.RouteRegister("/getaddress", GetAddress)
	lebangnet.RouteRegister("/addaddress", AddAddress)
	lebangnet.RouteRegister("/modifyaddress", ModifyAddress)
	lebangnet.RouteRegister("/deleteaddress", DeleteAddress)
	lebangnet.RouteRegister("/setdefaultaddress", SetDefaultAddress)
	lebangnet.RouteRegister("/defaultaddress", DefaultAddress)
}

func Authentication(w http.ResponseWriter, req *http.Request) {
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
		if idcode.GetCode() != reqdata.GetCode() {
			response.Errorcode = "验证码错误"
			logger.PRINTLINE("authentication error: ", idcode.GetPhone(), idcode.GetCode(), reqdata.GetCode())
		} else if time.Unix(reqdata.GetTime()/1000, 0).Sub(time.Unix(idcode.GetTime()/1000, 0)).Seconds() > 90 {
			durationsecond := time.Unix(reqdata.GetTime()/1000, 0).Sub(time.Unix(idcode.GetTime()/1000, 0)).Seconds()
			logger.PRINTLINE(durationsecond)
			response.Errorcode = "验证码超时"
			logger.PRINTLINE("authentication error: ", idcode.GetPhone(), idcode.GetCode(), reqdata.GetCode())
		} else {
			var userdata lebangproto.User
			if dbmanager.GetMongo().Find(config.DB().DBName, config.DB().CollMap["user"],
				bson.M{"phone": reqdata.GetPhone()}, nil, &userdata) {
				logger.PRINTLINE("update")
				userdata.Lastsignintime = reqdata.GetTime()
				dbmanager.GetMongo().Update(config.DB().DBName, config.DB().CollMap["user"], bson.M{"phone": reqdata.GetPhone()}, userdata)
			} else {
				logger.PRINTLINE("insert")
				userdata := lebangproto.User{
					Phone:          reqdata.GetPhone(),
					Registertime:   reqdata.GetTime(),
					Lastsignintime: reqdata.GetTime(),
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

func GetIDCode(w http.ResponseWriter, req *http.Request) {
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
			Code:  fmt.Sprintf("%06d", rand.New(rand.NewSource(time.Now().UnixNano())).Intn(999999)),
			Time:  time.Now().Unix() * 1000,
		}
		logger.PRINTLINE("exist", idcodeinfo.GetCode())
		dbmanager.GetMongo().Update(config.DB().DBName, config.DB().CollMap["idcode"], bson.M{"phone": reqdata.GetPhone()}, &idcodeinfo)
	} else {
		idcodeinfo = lebangproto.IDCode{
			Phone: reqdata.GetPhone(),
			Code:  fmt.Sprintf("%06d", rand.New(rand.NewSource(time.Now().UnixNano())).Intn(999999)),
			Time:  time.Now().Unix() * 1000,
		}
		logger.PRINTLINE("not exist", idcodeinfo.GetCode())
		dbmanager.GetMongo().Insert(config.DB().DBName, config.DB().CollMap["idcode"], &idcodeinfo)
	}

	//	err := alisms.SendSms(config.Instance().AccessKeyID, config.Instance().AccessSecret, reqdata.GetPhone(),
	//		"LeBang", fmt.Sprintf("{code:%s}", idcodeinfo.Code), "SMS_135792492")
	//	if err != nil {
	//		logger.PRINTLINE("dysms.SendSms", err)
	//	}

	sendbuf, err := json.Marshal(response)
	if err != nil {
		logger.PRINTLINE("Marshal response error: ", err)
		return
	}

	io.WriteString(w, string(sendbuf))
}

func Login(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	buf := make([]byte, req.ContentLength)
	common.GetBuffer(req, buf)

	var reqdata lebangproto.LoginReq
	if !common.Unmarshal(buf, &reqdata) {
		return
	}
	logger.PRINTLINE(reqdata.GetPhone())

	var response lebangproto.Response
	if reqdata.GetPhone() != "123456" {
		response.Errorcode = "user not exist"
	}

	sendbuf, err := json.Marshal(response)
	if err != nil {
		logger.PRINTLINE("Marshal response error: ", err)
		return
	}

	io.WriteString(w, string(sendbuf))
}
