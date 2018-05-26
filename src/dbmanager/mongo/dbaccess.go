package mongo

import (
	"config"
	"errors"
	"fmt"
	"lebangproto"
	"logger"
	"model"
	"sync"
	"time"

	"gopkg.in/mgo.v2"
)

type MongoManager struct {
	session   []*mgo.Session
	index     int
	lockGuard sync.Mutex
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

	if !this.IsCollExist(config.DB().DBName, config.DB().CollMap["classification"]) {
		for classificationName, index := range config.DB().Classification {
			this.Insert(config.DB().DBName, config.DB().CollMap["classification"],
				&model.Classification{Name: classificationName})

			this.Insert(config.DB().DBName, config.DB().CollMap["subclassification"],
				&lebangproto.Classification{Classification: classificationName,
					Labels: config.DB().SubClassification[classificationName]["labels"],
					Hint:   config.DB().SubClassification[classificationName]["hint"]})
		}

	}

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
