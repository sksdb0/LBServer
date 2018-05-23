package dbmanager

import (
	"config"
	"dbmanager/mongo"
	"logger"
)

var instance *DBManager

type DBManager struct {
	mongo *mongo.MongoManager
}

func Init() bool {
	logger.PRINTLINE("DB Init")
	instance = new(DBManager)

	instance.mongo = mongo.NewMongoManager(config.DB().MongoDBLocalAddress, config.DB().MongoDBLocalPort)
	if instance.mongo == nil {
		logger.PRINTLINE("mongodb connect error")
	}

	logger.PRINTLINE("DB Init Success")
	return true
}

func Close() {
	if instance.mongo != nil {
		instance.mongo.Close()
	}
}

func GetMongo() *mongo.MongoManager {
	if instance == nil {
		instance = new(DBManager)
	}
	return instance.mongo
}
