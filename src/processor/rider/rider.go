package rider

import (
	"config"
	"dbmanager"
	"encoding/json"
	"httprouter"
	"io"
	"lebangproto"
	"logger"
	"net/http"
	"processor/common"
	"sync"
	"time"

	"gopkg.in/mgo.v2/bson"
)

var ridermanager *RiderManager

type RiderManager struct {
	riderlocationmap map[string]*lebangproto.RiderLocation
	idmutex          sync.RWMutex
}

func Init(router *httprouter.Router) {
	ridermanager = &RiderManager{
		riderlocationmap: make(map[string]*lebangproto.RiderLocation),
	}
	// rider
	router.POST("/addrider", ridermanager.AddRider)
	router.POST("/getrider", ridermanager.GetRider)
	router.POST("/deleterider", ridermanager.DeleteRider)
	router.POST("/getriderlocation", ridermanager.GetRiderLocation)

	router.POST("/uploadriderlocation", ridermanager.UploadRiderLocation)

	router.POST("/modifyriderpassword", ridermanager.ModifyRiderPassword)

}

func (this *RiderManager) AddRider(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	this.idmutex.Lock()
	defer this.idmutex.Unlock()

	defer req.Body.Close()
	buf := make([]byte, req.ContentLength)
	common.GetBuffer(req, buf)

	var reqdata lebangproto.AddRiderReq
	if !common.Unmarshal(buf, &reqdata) {
		return
	}
	logger.PRINTLINE(reqdata.GetPhone())

	var rider lebangproto.Rider
	var response lebangproto.Response
	if dbmanager.GetMongo().Find(config.DB().DBName, config.DB().CollMap["rider"], bson.M{"phone": reqdata.GetPhone()}, nil, &rider) {
		if rider.State != int64(lebangproto.RiderState_RIDER_STATE_DIMISSION) {
			response.Errorcode = "骑手已经存在"
		} else {
			rider.State = int64(lebangproto.RiderState_RIDER_STATE_ONJOB)
			rider.Phone = reqdata.GetPhone()
			dbmanager.GetMongo().Update(config.DB().DBName, config.DB().CollMap["rider"], bson.M{"phone": reqdata.GetPhone()}, &rider)
		}
	} else {
		rider = lebangproto.Rider{Phone: reqdata.GetPhone(),
			Name:           reqdata.GetName(),
			Registertime:   time.Now().Unix() * 1000,
			Lastsignintime: time.Now().Unix() * 1000,
			Password:       "123456",
			State:          int64(lebangproto.RiderState_RIDER_STATE_ONJOB)}
		dbmanager.GetMongo().Insert(config.DB().DBName, config.DB().CollMap["rider"], rider)
	}

	sendbuf, err := json.Marshal(response)
	if err != nil {
		logger.PRINTLINE("Marshal response error: ", err)
		return
	}
	io.WriteString(w, string(sendbuf))
}

func (this *RiderManager) GetRider(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	this.idmutex.Lock()
	defer this.idmutex.Unlock()

	defer req.Body.Close()
	buf := make([]byte, req.ContentLength)
	common.GetBuffer(req, buf)

	var reqdata lebangproto.GetRiderReq
	if !common.Unmarshal(buf, &reqdata) {
		return
	}
	logger.PRINTLINE(reqdata.GetPhone())

	var response lebangproto.GetRiderRes
	if dbmanager.GetMongo().FindAll(config.DB().DBName, config.DB().CollMap["rider"],
		bson.M{"state": bson.M{"$lt": int64(lebangproto.RiderState_RIDER_STATE_DIMISSION)}}, "", nil, &response.Rider) {
		if len(response.Rider) == 0 {
			response.Errorcode = "no rider"
			logger.PRINTLINE("no rider")
		}
	}

	sendbuf, err := json.Marshal(response)
	if err != nil {
		logger.PRINTLINE("Marshal response error: ", err)
		return
	}
	io.WriteString(w, string(sendbuf))
}

func (this *RiderManager) DeleteRider(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	this.idmutex.Lock()
	defer this.idmutex.Unlock()

	defer req.Body.Close()
	buf := make([]byte, req.ContentLength)
	common.GetBuffer(req, buf)

	var reqdata lebangproto.DeleteRiderReq
	if !common.Unmarshal(buf, &reqdata) {
		return
	}
	logger.PRINTLINE(reqdata.GetPhone())

	var response lebangproto.GetRiderRes
	var rider lebangproto.Rider
	if dbmanager.GetMongo().Find(config.DB().DBName, config.DB().CollMap["rider"],
		bson.M{"phone": reqdata.GetPhone()}, nil, &rider) {
		rider.State = int64(lebangproto.RiderState_RIDER_STATE_DIMISSION)
		dbmanager.GetMongo().Update(config.DB().DBName, config.DB().CollMap["rider"], bson.M{"phone": reqdata.GetPhone()}, &rider)
	} else {
		response.Errorcode = "order not exist"
		logger.PRINTLINE("order not exist: ", reqdata.GetPhone())
	}

	sendbuf, err := json.Marshal(response)
	if err != nil {
		logger.PRINTLINE("Marshal response error: ", err)
		return
	}
	io.WriteString(w, string(sendbuf))
}

func (this *RiderManager) GetRiderLocation(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	this.idmutex.Lock()
	defer this.idmutex.Unlock()

	defer req.Body.Close()
	buf := make([]byte, req.ContentLength)
	common.GetBuffer(req, buf)

	var reqdata lebangproto.GetRiderLocationsReq
	if !common.Unmarshal(buf, &reqdata) {
		return
	}

	//	判断是否为管理员权限

	var response lebangproto.GetRiderLocationsRes
	//	判断骑手状态
	if len(this.riderlocationmap) > 0 {
		for _, location := range this.riderlocationmap {
			if time.Now().Unix()-location.GetTime() < 60 {
				response.Location = append(response.Location, location)
			}
		}
		if len(response.Location) <= 0 {
			response.Errorcode = "no rider"
			logger.PRINTLINE("no rider")
		}
	} else {
		response.Errorcode = "no rider"
		logger.PRINTLINE("no rider")
	}

	sendbuf, err := json.Marshal(response)
	if err != nil {
		logger.PRINTLINE("Marshal response error: ", err)
		return
	}
	io.WriteString(w, string(sendbuf))
}

func (this *RiderManager) UploadRiderLocation(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	this.idmutex.Lock()
	defer this.idmutex.Unlock()

	defer req.Body.Close()
	buf := make([]byte, req.ContentLength)
	common.GetBuffer(req, buf)

	var reqdata lebangproto.UploadRiderLocationReq
	if !common.Unmarshal(buf, &reqdata) {
		return
	}

	var response lebangproto.Response
	riderlocation, exist := this.riderlocationmap[reqdata.GetPhone()]
	if exist {
		riderlocation.Latitude = reqdata.GetLatitude()
		riderlocation.Longitude = reqdata.GetLongitude()
		riderlocation.Time = time.Now().Unix()
	} else {
		var rider lebangproto.Rider
		if dbmanager.GetMongo().Find(config.DB().DBName, config.DB().CollMap["rider"], bson.M{"phone": reqdata.GetPhone()}, nil, &rider) {
			newriderlocation := &lebangproto.RiderLocation{
				Phone:     rider.GetPhone(),
				Name:      rider.GetName(),
				Latitude:  reqdata.GetLatitude(),
				Longitude: reqdata.GetLongitude(),
				Time:      time.Now().Unix(),
			}

			this.riderlocationmap[reqdata.GetPhone()] = newriderlocation
		} else {
			response.Errorcode = "骑手不存在"
		}
	}
	//	for _, l := range this.riderlocationmap {
	//		logger.PRINTLINE(reqdata.GetPhone(), l.GetPhone(), l.GetLatitude(), l.GetLongitude(), l.GetTime())
	//	}

	sendbuf, err := json.Marshal(response)
	if err != nil {
		logger.PRINTLINE("Marshal response error: ", err)
		return
	}
	io.WriteString(w, string(sendbuf))
}

func (this *RiderManager) ModifyRiderPassword(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	this.idmutex.Lock()
	defer this.idmutex.Unlock()

	defer req.Body.Close()
	buf := make([]byte, req.ContentLength)
	common.GetBuffer(req, buf)

	var reqdata lebangproto.ModifyRiderPasswordReq
	if !common.Unmarshal(buf, &reqdata) {
		return
	}

	logger.PRINTLINE(reqdata)
	var rider lebangproto.Rider
	var response lebangproto.Response
	if dbmanager.GetMongo().Find(config.DB().DBName, config.DB().CollMap["rider"], bson.M{"phone": reqdata.GetPhone()}, nil, &rider) {
		if rider.GetPassword() != reqdata.GetOripassword() {
			response.Errorcode = "密码错误"
		} else {
			rider.Password = reqdata.GetNewpassword()
			dbmanager.GetMongo().Update(config.DB().DBName, config.DB().CollMap["rider"], bson.M{"phone": reqdata.GetPhone()}, &rider)
		}
	} else {
		response.Errorcode = "骑手不存在"
	}

	sendbuf, err := json.Marshal(response)
	if err != nil {
		logger.PRINTLINE("Marshal response error: ", err)
		return
	}
	io.WriteString(w, string(sendbuf))
}
