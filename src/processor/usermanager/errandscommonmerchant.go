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

func UpdateErrandsCommonMerchant(phone string, merchant string) {
	if merchant == "就近购买" || merchant == "" {
		return
	}

	var merchantdata lebangproto.ErrandCommonMerchant

	if dbmanager.GetMongo().Find(config.DB().DBName, config.DB().CollMap["errandscommonmerchant"],
		bson.M{"phone": phone}, nil, &merchantdata) {
		mers := []string{merchant}
		for _, v := range merchantdata.Merchant {
			if merchant != v {
				mers = append(mers, v)
			}
			if len(mers) >= 10 {
				break
			}
		}
		merchantdata.Merchant = mers

		dbmanager.GetMongo().Update(config.DB().DBName, config.DB().CollMap["errandscommonmerchant"],
			bson.M{"phone": phone}, &merchantdata)
	} else {
		merchantdata = lebangproto.ErrandCommonMerchant{
			Phone:    phone,
			Merchant: []string{merchant},
		}

		dbmanager.GetMongo().Insert(config.DB().DBName, config.DB().CollMap["errandscommonmerchant"], &merchantdata)
	}
}

func GetErrandsCommonMerchant(w http.ResponseWriter, req *http.Request) {
	logger.PRINTLINE("GetErrandsCommonMerchant")

	defer req.Body.Close()
	buf := make([]byte, req.ContentLength)
	common.GetBuffer(req, buf)

	var reqdata lebangproto.GetErrandCommonMerchant
	if !common.Unmarshal(buf, &reqdata) {
		return
	}

	logger.PRINTLINE(reqdata)
	var response lebangproto.ResErrandCommonMerchant
	var merchant lebangproto.ErrandCommonMerchant
	if dbmanager.GetMongo().Find(config.DB().DBName, config.DB().CollMap["errandscommonmerchant"],
		bson.M{"phone": reqdata.GetPhone()}, nil, &merchant) {
		response.Merchant = merchant.GetMerchant()
	} else {
		response.Errorcode = "no errandscommonmerchant"
		logger.PRINTLINE("no errandscommonmerchant: ", reqdata.GetPhone())
	}

	sendbuf, err := json.Marshal(response)
	if err != nil {
		logger.PRINTLINE("Marshal response error: ", err)
		return
	}
	io.WriteString(w, string(sendbuf))
}
