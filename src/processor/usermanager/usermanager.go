package usermanager

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
	lebangnet.RouteRegister("/signin", SignIn)
	lebangnet.RouteRegister("/signup", SignUp)
	lebangnet.RouteRegister("/getaddress", GetAddress)
	lebangnet.RouteRegister("/addaddress", AddAddress)
	lebangnet.RouteRegister("/modifyaddress", ModifyAddress)
	lebangnet.RouteRegister("/deleteaddress", DeleteAddress)
	lebangnet.RouteRegister("/setdefaultaddress", SetDefaultAddress)
	lebangnet.RouteRegister("/defaultaddress", DefaultAddress)
}

func SignIn(w http.ResponseWriter, req *http.Request) {
	logger.PRINTLINE("SignIn")

	defer req.Body.Close()
	buf := make([]byte, req.ContentLength)
	common.GetBuffer(req, buf)

	var signin lebangproto.SignIn
	if !common.Unmarshal(buf, &signin) {
		return
	}

	var userinfo lebangproto.UserInfo
	var response lebangproto.SignInRes

	if dbmanager.GetMongo().Find(config.DB().DBName, config.DB().CollMap["user"], bson.M{"phone": signin.GetPhone()}, nil, &userinfo) {
		if userinfo.GetPassword() == signin.GetPassword() {
			response.Phone = userinfo.GetPhone()
			logger.PRINTLINE("user signin: ", signin.GetPhone())
		} else {
			response.Errorcode = "password or username error"
			logger.PRINTLINE("password or username error: ", signin.GetPhone())
		}

	} else {
		response.Errorcode = "user not exist"
		logger.PRINTLINE("user not exist: ", signin.GetPhone())
	}

	sendbuf, err := json.Marshal(response)
	if err != nil {
		logger.PRINTLINE("Marshal response error: ", err)
		return
	}
	io.WriteString(w, string(sendbuf))
}

func SignUp(w http.ResponseWriter, req *http.Request) {
	logger.PRINTLINE("SignUp")

	defer req.Body.Close()
	buf := make([]byte, req.ContentLength)
	common.GetBuffer(req, buf)

	var signup lebangproto.SignUp
	if !common.Unmarshal(buf, &signup) {
		return
	}

	var response lebangproto.SignInRes
	if !dbmanager.GetMongo().IsExist(config.DB().DBName, config.DB().CollMap["user"], bson.M{"phone": signup.GetPhone()}) {
		userinfo := lebangproto.UserInfo{
			Phone:    signup.GetPhone(),
			Password: "",
		}
		dbmanager.GetMongo().Insert(config.DB().DBName, config.DB().CollMap["user"], userinfo)
		response.Phone = signup.GetPhone()
		logger.PRINTLINE("user signup: ", signup.GetPhone())
	} else {
		response.Errorcode = "user exist"
		logger.PRINTLINE("user exist: ", signup.GetPhone())
	}

	sendbuf, err := json.Marshal(response)
	if err != nil {
		logger.PRINTLINE("Marshal response error: ", err)
		return
	}
	io.WriteString(w, string(sendbuf))
}
