package db

import (
	"logger"
	"strings"

	"github.com/Unknwon/goconfig"
)

type DB struct {
	DB                  string
	DBName              string
	MongoDBLocalPort    string
	MongoDBLocalAddress string

	CollMap map[string]string

	Classification     map[string]string
	SubClassification  map[string]string
	ClassificationView map[string]string

	ErrandsClassification    map[string]string
	ErrandsSubClassification map[string]map[string]string
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

	this.ErrandsClassification, _ = dbcfg.GetSection("ErrandsClassification")
	errandsLabels := strings.Split(this.ErrandsClassification["labels"], " ")
	for _, classification := range errandsLabels {
		this.ErrandsSubClassification[classification], _ = dbcfg.GetSection(classification)
	}

	this.Classification, _ = dbcfg.GetSection("Classification")
	this.SubClassification, _ = dbcfg.GetSection("SubClassification")
	this.ClassificationView, _ = dbcfg.GetSection("ClassificationView")
	return true
}

func NewDB() *DB {
	return &DB{
		CollMap:                  make(map[string]string),
		Classification:           make(map[string]string),
		SubClassification:        make(map[string]string),
		ErrandsClassification:    make(map[string]string),
		ErrandsSubClassification: make(map[string]map[string]string),
	}
}
