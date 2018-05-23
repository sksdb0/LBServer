package db

import (
	"logger"

	"github.com/Unknwon/goconfig"
)

type DB struct {
	DB                  string
	DBName              string
	MongoDBLocalPort    string
	MongoDBLocalAddress string

	CollMap map[string]string
}

func (this *DB) Load() bool {
	dbcfg, err := goconfig.LoadConfigFile("src/config/db/db.cfg")
	if err != nil {
		logger.LOGLINE("读取配置文件失败[go.cfg]", err)
		return false
	}

	this.DB, _ = dbcfg.GetValue("config", "DB")

	this.DBName, _ = dbcfg.GetValue("mongodb", "DBName")
	this.MongoDBLocalPort, _ = dbcfg.GetValue("mongodb", "MongoDBLocalPort")
	this.MongoDBLocalAddress, _ = dbcfg.GetValue("mongodb", "MongoDBLocalAddress")

	this.CollMap, _ = dbcfg.GetSection("CollMap")
	return true
}

func NewDB() *DB {
	return &DB{
		CollMap: make(map[string]string),
	}
}
