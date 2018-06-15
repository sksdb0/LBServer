package mongo

import (
	"config"
	"errors"
	"fmt"
	"lebangproto"
	"logger"
	"strconv"
	"strings"
	"sync"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type MongoManager struct {
	session   []*mgo.Session
	index     int
	lockGuard sync.Mutex
}

func (this *MongoManager) updateErrandsMainClassification() {
	logger.LOGLINE("update errands main classification")
	if !this.IsCollExist(config.DB().DBName, config.DB().CollMap["errandsclassification"]) {
		this.Insert(config.DB().DBName, config.DB().CollMap["errandsclassification"],
			&lebangproto.ErrandsClassification{Classification: "main",
				Labels: config.DB().ErrandsClassification["labels"],
				Hint:   config.DB().ErrandsClassification["hint"]})
	} else {
		this.Update(config.DB().DBName, config.DB().CollMap["errandsclassification"],
			bson.M{"Classification": "main"},
			&lebangproto.ErrandsClassification{Classification: "main",
				Labels: config.DB().ErrandsClassification["labels"],
				Hint:   config.DB().ErrandsClassification["hint"]})
	}
}

func (this *MongoManager) updateErrandsSubClassification() {
	logger.LOGLINE("errands sub classification")
	errandsLabels := strings.Split(config.DB().ErrandsClassification["labels"], " ")
	for _, classification := range errandsLabels {
		if !this.IsExist(config.DB().DBName, config.DB().CollMap["errandssubclassification"], bson.M{"classification": classification}) {
			this.Insert(config.DB().DBName, config.DB().CollMap["errandssubclassification"],
				&lebangproto.ErrandsClassification{Classification: classification,
					Labels: config.DB().ErrandsSubClassification[classification]["labels"],
					Hint:   config.DB().ErrandsSubClassification[classification]["hint"]})
		} else {
			this.Update(config.DB().DBName, config.DB().CollMap["errandssubclassification"],
				bson.M{"classification": classification},
				&lebangproto.ErrandsClassification{Classification: classification,
					Labels: config.DB().ErrandsSubClassification[classification]["labels"],
					Hint:   config.DB().ErrandsSubClassification[classification]["hint"]})
		}
	}
}

func (this *MongoManager) updateClassificationView() {
	logger.LOGLINE("update classificationview")
	for name, types := range config.DB().ClassificationView {
		if !this.IsExist(config.DB().DBName, config.DB().CollMap["classificationview"], bson.M{"name": name}) {
			this.Insert(config.DB().DBName, config.DB().CollMap["classificationview"],
				&lebangproto.ClassificationView{Name: name, Typeids: types})
		} else {
			this.Update(config.DB().DBName, config.DB().CollMap["classificationview"],
				bson.M{"name": name},
				&lebangproto.ClassificationView{Name: name, Typeids: types})
		}
	}
}

func (this *MongoManager) updateClassification() {
	logger.LOGLINE("update classification")
	for name, typeidstr := range config.DB().Classification {
		if !this.IsExist(config.DB().DBName, config.DB().CollMap["classification"], bson.M{"name": name}) {
			typeid, _ := strconv.Atoi(typeidstr)
			this.Insert(config.DB().DBName, config.DB().CollMap["classification"],
				&lebangproto.Classification{Name: name, Typeid: int32(typeid)})
		} else {
			typeid, _ := strconv.Atoi(typeidstr)
			this.Update(config.DB().DBName, config.DB().CollMap["classification"],
				bson.M{"name": name},
				&lebangproto.Classification{Name: name, Typeid: int32(typeid)})
		}
	}
}

func (this *MongoManager) updateSubClassification() {
	logger.LOGLINE("update sub classification")
	for name, subinfostr := range config.DB().SubClassification {
		if !this.IsExist(config.DB().DBName, config.DB().CollMap["subclassification"], bson.M{"name": name}) {
			subinfo := strings.Split(subinfostr, " ")
			parenttypeid, _ := strconv.Atoi(subinfo[0])
			typeid, _ := strconv.Atoi(subinfo[1])
			this.Insert(config.DB().DBName, config.DB().CollMap["subclassification"],
				&lebangproto.SubClassification{Name: name,
					Typeid:       int32(typeid),
					Parenttypeid: int32(parenttypeid),
					Image:        subinfo[2]})
		} else {
			subinfo := strings.Split(subinfostr, " ")
			parenttypeid, _ := strconv.Atoi(subinfo[0])
			typeid, _ := strconv.Atoi(subinfo[1])
			this.Update(config.DB().DBName, config.DB().CollMap["subclassification"],
				bson.M{"name": name},
				&lebangproto.SubClassification{Name: name,
					Typeid:       int32(typeid),
					Parenttypeid: int32(parenttypeid),
					Image:        subinfo[2]})
		}
	}
}

func (this *MongoManager) Init(addrs string, port string) error {
	addr := fmt.Sprintf("%s:%s", addrs, port)

	session, err := mgo.Dial(addr)
	if err != nil {
		logger.PRINTLINE("open mongodb error:", err)
	} else {
		session.SetSocketTimeout(100 * time.Hour)
		session.SetMode(mgo.Monotonic, true)
		this.session = append(this.session, session)

		logger.LOGLINE("mongodb connect ", addrs, " success!")
	}

	// update errands main classification
	this.updateErrandsMainClassification()
	// update errands sub classification
	this.updateErrandsSubClassification()
	// update classificationview
	this.updateClassificationView()
	// update classification
	this.updateClassification()
	// update sub classification
	this.updateSubClassification()

	if len(this.session) > 0 {
		return nil
	}
	return errors.New("all mysql connect error!")
}

func (this *MongoManager) Close() {
	for _, node := range this.session {
		if node != nil {
			node.Close()
		}
	}
}

func (this *MongoManager) GetDB(name string) *mgo.Database {
	this.lockGuard.Lock()
	defer this.lockGuard.Unlock()

	return this.session[this.index].DB(name)
}

func NewMongoManager(addrs string, port string) *MongoManager {
	instance := MongoManager{
		session: nil,
		index:   0,
	}
	if instance.Init(addrs, port) == nil {
		return &instance
	} else {
		return nil
	}
}
