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
	defer req.Body.Close()

	buf := make([]byte, req.ContentLength)
	common.GetBuffer(req, buf)

	var reqdata lebangproto.AddAddressReq
	if !common.Unmarshal(buf, &reqdata) {
		return
	}
	logger.PRINTLINE(reqdata.GetPhone())

	var useraddress lebangproto.UserAddress
	var response lebangproto.Response
	if dbmanager.GetMongo().Find(config.DB().DBName, config.DB().CollMap["address"], bson.M{"phone": reqdata.GetPhone()}, nil, &useraddress) {
		if len(useraddress.Address) == 0 {
			reqdata.Address.Isdefault = true
		}
		useraddress.Address = append(useraddress.Address, reqdata.GetAddress())
		dbmanager.GetMongo().Update(config.DB().DBName, config.DB().CollMap["address"], bson.M{"phone": reqdata.GetPhone()}, useraddress)
	} else {
		useraddress := lebangproto.UserAddress{Phone: reqdata.GetPhone()}
		reqdata.Address.Isdefault = true
		useraddress.Address = append(useraddress.Address, reqdata.GetAddress())
		dbmanager.GetMongo().Insert(config.DB().DBName, config.DB().CollMap["address"], useraddress)
	}

	sendbuf, err := json.Marshal(response)
	if err != nil {
		logger.PRINTLINE("Marshal response error: ", err)
		return
	}
	io.WriteString(w, string(sendbuf))
}

func ModifyAddress(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()

	buf := make([]byte, req.ContentLength)
	common.GetBuffer(req, buf)

	var reqdata lebangproto.ModifyAddressReq
	if !common.Unmarshal(buf, &reqdata) {
		return
	}
	logger.PRINTLINE(reqdata.GetPhone())

	var useraddress lebangproto.UserAddress
	var response lebangproto.Response
	if dbmanager.GetMongo().Find(config.DB().DBName, config.DB().CollMap["address"], bson.M{"phone": reqdata.GetPhone()}, nil, &useraddress) {
		useraddress.Address[reqdata.GetAddressnumber()] = reqdata.Address
		dbmanager.GetMongo().Update(config.DB().DBName, config.DB().CollMap["address"], bson.M{"phone": reqdata.GetPhone()}, useraddress)
	} else {
		response.Errorcode = "address not exist"
		logger.PRINTLINE("address not exist: ", reqdata.GetPhone())
	}

	sendbuf, err := json.Marshal(response)
	if err != nil {
		logger.PRINTLINE("Marshal response error: ", err)
		return
	}
	io.WriteString(w, string(sendbuf))
}

func DeleteAddress(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()

	buf := make([]byte, req.ContentLength)
	common.GetBuffer(req, buf)

	var reqdata lebangproto.DeleteAddressReq
	if !common.Unmarshal(buf, &reqdata) {
		return
	}
	logger.PRINTLINE(reqdata.GetPhone())

	var useraddress lebangproto.UserAddress
	var response lebangproto.Response
	if dbmanager.GetMongo().Find(config.DB().DBName, config.DB().CollMap["address"], bson.M{"phone": reqdata.GetPhone()}, nil, &useraddress) {
		useraddress.Address = append(useraddress.Address[:reqdata.GetAddressnumber()], useraddress.Address[reqdata.GetAddressnumber()+1:]...)
		dbmanager.GetMongo().Update(config.DB().DBName, config.DB().CollMap["address"], bson.M{"phone": reqdata.GetPhone()}, useraddress)
	} else {
		response.Errorcode = "address not exist"
		logger.PRINTLINE("address not exist: ", reqdata.GetPhone())
	}

	sendbuf, err := json.Marshal(response)
	if err != nil {
		logger.PRINTLINE("Marshal response error: ", err)
		return
	}
	io.WriteString(w, string(sendbuf))
}

func DefaultAddress(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	buf := make([]byte, req.ContentLength)
	common.GetBuffer(req, buf)

	var reqdata lebangproto.DefaultAddressReq
	if !common.Unmarshal(buf, &reqdata) {
		return
	}
	logger.PRINTLINE(reqdata.GetPhone())

	var useraddress lebangproto.UserAddress
	var response lebangproto.DefaultAddressRes
	if dbmanager.GetMongo().Find(config.DB().DBName, config.DB().CollMap["address"], bson.M{"phone": reqdata.GetPhone()}, nil, &useraddress) {
		for _, node := range useraddress.GetAddress() {
			if node.GetIsdefault() {
				response.Address = node
				break
			}
		}
		if response.Address == nil {
			response.Errorcode = "no default address"
			logger.PRINTLINE("no default address: ", reqdata.GetPhone())
		}
	} else {
		response.Errorcode = "no default address"
		logger.PRINTLINE("no default address: ", reqdata.GetPhone())
	}

	sendbuf, err := json.Marshal(response)
	if err != nil {
		logger.PRINTLINE("Marshal response error: ", err)
		return
	}
	io.WriteString(w, string(sendbuf))
}

func SetDefaultAddress(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	buf := make([]byte, req.ContentLength)
	common.GetBuffer(req, buf)

	var reqdata lebangproto.ModifyAddressReq
	if !common.Unmarshal(buf, &reqdata) {
		return
	}
	logger.PRINTLINE(reqdata.GetPhone())

	var useraddress lebangproto.UserAddress
	var response lebangproto.DefaultAddressRes
	if dbmanager.GetMongo().Find(config.DB().DBName, config.DB().CollMap["address"], bson.M{"phone": reqdata.GetPhone()}, nil, &useraddress) {
		for _, node := range useraddress.GetAddress() {
			if node.GetIsdefault() {
				node.Isdefault = false
			}
		}
		useraddress.Address[reqdata.GetAddressnumber()].Isdefault = true
		dbmanager.GetMongo().Update(config.DB().DBName, config.DB().CollMap["address"], bson.M{"phone": reqdata.GetPhone()}, useraddress)
	}

	sendbuf, err := json.Marshal(response)
	if err != nil {
		logger.PRINTLINE("Marshal response error: ", err)
		return
	}
	io.WriteString(w, string(sendbuf))
}

func GetAddress(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	buf := make([]byte, req.ContentLength)
	common.GetBuffer(req, buf)

	var reqdata lebangproto.GetAddressReq
	if !common.Unmarshal(buf, &reqdata) {
		return
	}
	logger.PRINTLINE(reqdata.GetPhone())

	var useraddress lebangproto.UserAddress
	var response lebangproto.GetAddressRes
	if dbmanager.GetMongo().Find(config.DB().DBName, config.DB().CollMap["address"], bson.M{"phone": reqdata.GetPhone()}, nil, &useraddress) {
		if len(useraddress.GetAddress()) == 0 {
			response.Errorcode = "no address"
			logger.PRINTLINE("no address: ", reqdata.GetPhone())
		} else {
			for _, node := range useraddress.GetAddress() {
				response.Address = append(response.Address, node)
			}
		}
	} else {
		response.Errorcode = "no address"
		logger.PRINTLINE("no address: ", reqdata.GetPhone())
	}

	sendbuf, err := json.Marshal(response)
	if err != nil {
		logger.PRINTLINE("Marshal response error: ", err)
		return
	}
	io.WriteString(w, string(sendbuf))
}
