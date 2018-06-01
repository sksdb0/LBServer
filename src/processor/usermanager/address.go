package usermanager

import (
	"config"
	"dbmanager"
	"encoding/json"
	"io"
	"lebangproto"
	"logger"
	"net/http"
	"processor/common"

	"gopkg.in/mgo.v2/bson"
)

func AddAddress(w http.ResponseWriter, req *http.Request) {
	logger.PRINTLINE("AddAddress")
	defer req.Body.Close()

	buf := make([]byte, req.ContentLength)
	common.GetBuffer(req, buf)

	var newinfo lebangproto.UserInfo
	if !common.Unmarshal(buf, &newinfo) {
		return
	}

	var userinfo lebangproto.UserInfo
	var response lebangproto.Response
	if dbmanager.GetMongo().Find(config.DB().DBName, config.DB().CollMap["user"], bson.M{"phone": newinfo.GetPhone()}, nil, &userinfo) {
		// 第一个为默认地址
		if len(userinfo.GetAddress()) == 0 {
			newinfo.Address[0].Isdefault = true
		}
		userinfo.Address = append(userinfo.Address, newinfo.Address[0])
		dbmanager.GetMongo().Update(config.DB().DBName, config.DB().CollMap["user"], bson.M{"phone": newinfo.GetPhone()}, userinfo)
	} else {
		response.Errorcode = "user not exist"
		logger.PRINTLINE("user not exist: ", newinfo.GetPhone())
	}

	sendbuf, err := json.Marshal(response)
	if err != nil {
		logger.PRINTLINE("Marshal response error: ", err)
		return
	}
	io.WriteString(w, string(sendbuf))
}

func ModifyAddress(w http.ResponseWriter, req *http.Request) {
	logger.PRINTLINE("ModifyAddress")
	defer req.Body.Close()

	buf := make([]byte, req.ContentLength)
	common.GetBuffer(req, buf)

	var modifyaddr lebangproto.ModifyAddress
	if !common.Unmarshal(buf, &modifyaddr) {
		return
	}

	logger.PRINTLINE(modifyaddr)
	var user lebangproto.UserInfo
	var response lebangproto.Response

	if dbmanager.GetMongo().Find(config.DB().DBName, config.DB().CollMap["user"], bson.M{"phone": modifyaddr.GetPhone()}, nil, &user) {
		user.Address[modifyaddr.GetAddressnumber()] = modifyaddr.Address
		dbmanager.GetMongo().Update(config.DB().DBName, config.DB().CollMap["user"], bson.M{"phone": modifyaddr.GetPhone()}, user)
	} else {
		response.Errorcode = "user not exist"
		logger.PRINTLINE("user not exist: ", modifyaddr.GetPhone())
	}

	sendbuf, err := json.Marshal(response)
	if err != nil {
		logger.PRINTLINE("Marshal response error: ", err)
		return
	}
	io.WriteString(w, string(sendbuf))
}

func DeleteAddress(w http.ResponseWriter, req *http.Request) {
	logger.PRINTLINE("DeleteAddress")
	defer req.Body.Close()

	buf := make([]byte, req.ContentLength)
	common.GetBuffer(req, buf)

	var deleteaddr lebangproto.DeleteAddress
	if !common.Unmarshal(buf, &deleteaddr) {
		return
	}

	logger.PRINTLINE(deleteaddr)
	var userinfo lebangproto.UserInfo
	var response lebangproto.Response

	if dbmanager.GetMongo().Find(config.DB().DBName, config.DB().CollMap["user"], bson.M{"phone": deleteaddr.GetPhone()}, nil, &userinfo) {
		userinfo.Address = append(userinfo.Address[:deleteaddr.GetAddressnumber()], userinfo.Address[deleteaddr.GetAddressnumber()+1:]...)
		dbmanager.GetMongo().Update(config.DB().DBName, config.DB().CollMap["user"], bson.M{"phone": deleteaddr.GetPhone()}, userinfo)
	} else {
		response.Errorcode = "user not exist"
		logger.PRINTLINE("user not exist: ", deleteaddr.GetPhone())
	}

	sendbuf, err := json.Marshal(response)
	if err != nil {
		logger.PRINTLINE("Marshal response error: ", err)
		return
	}
	io.WriteString(w, string(sendbuf))
}

func DefaultAddress(w http.ResponseWriter, req *http.Request) {
	logger.PRINTLINE("DefaultAddress")

	defer req.Body.Close()
	buf := make([]byte, req.ContentLength)
	common.GetBuffer(req, buf)

	var userinfo lebangproto.UserInfo
	if !common.Unmarshal(buf, &userinfo) {
		return
	}

	logger.PRINTLINE(userinfo)
	var response lebangproto.DefaultAddressRes

	if dbmanager.GetMongo().Find(config.DB().DBName, config.DB().CollMap["user"], bson.M{"phone": userinfo.GetPhone()}, nil, &userinfo) {
		for _, node := range userinfo.GetAddress() {
			if node.GetIsdefault() {
				response.Address = node
			}
		}
		if response.Address == nil {
			response.Errorcode = "no default address"
			logger.PRINTLINE("no default address: ", userinfo.GetPhone())
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

func SetDefaultAddress(w http.ResponseWriter, req *http.Request) {
	logger.PRINTLINE("SetDefaultAddress")

	defer req.Body.Close()
	buf := make([]byte, req.ContentLength)
	common.GetBuffer(req, buf)

	var addrinfo lebangproto.ModifyAddress
	if !common.Unmarshal(buf, &addrinfo) {
		return
	}

	logger.PRINTLINE(addrinfo)
	var userinfo lebangproto.UserInfo
	var response lebangproto.DefaultAddressRes

	if dbmanager.GetMongo().Find(config.DB().DBName, config.DB().CollMap["user"], bson.M{"phone": addrinfo.GetPhone()}, nil, &userinfo) {
		for _, node := range userinfo.GetAddress() {
			if node.GetIsdefault() {
				node.Isdefault = false
			}
		}
		userinfo.Address[addrinfo.GetAddressnumber()].Isdefault = true
		dbmanager.GetMongo().Update(config.DB().DBName, config.DB().CollMap["user"], bson.M{"phone": addrinfo.GetPhone()}, userinfo)
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

func GetAddress(w http.ResponseWriter, req *http.Request) {
	logger.PRINTLINE("GetAddress")

	defer req.Body.Close()
	buf := make([]byte, req.ContentLength)
	common.GetBuffer(req, buf)

	var userinfo lebangproto.UserInfo
	if !common.Unmarshal(buf, &userinfo) {
		return
	}

	logger.PRINTLINE(userinfo)
	var response lebangproto.GetAddressRes

	if dbmanager.GetMongo().Find(config.DB().DBName, config.DB().CollMap["user"], bson.M{"phone": userinfo.GetPhone()}, nil, &userinfo) {
		if len(userinfo.GetAddress()) == 0 {
			response.Errorcode = "no address"
			logger.PRINTLINE("no address: ", userinfo.GetPhone())
		} else {
			for _, node := range userinfo.GetAddress() {
				response.Address = append(response.Address, node)
			}
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
