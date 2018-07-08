package config

import (
	"config/db"
	"logger"

	"github.com/Unknwon/goconfig"
)

var _instance *config

type config struct {
	HttpPort string
	CertFile string
	KeyFile  string
	Manager  string

	AccessKeyID  string
	AccessSecret string

	XingeAccessId  int64
	XingeSecretKey string

	DB *db.DB
}

func Init() bool {
	_instance = newconfig()
	return _instance.load()
}

func Instance() *config {
	return _instance
}

func DB() *db.DB {
	return _instance.DB
}

func (this *config) load() bool {
	cfg, err := goconfig.LoadConfigFile("src/config/server.cfg")
	if err != nil {
		logger.LOGLINE("读取配置文件失败[server.cfg]", err)
		return false
	}

	this.DB.Load()

	this.HttpPort, _ = cfg.GetValue("http", "Port")
	this.CertFile, _ = cfg.GetValue("http", "CertFile")
	this.KeyFile, _ = cfg.GetValue("http", "KeyFile")
	this.Manager, _ = cfg.GetValue("http", "Manager")

	this.AccessKeyID, _ = cfg.GetValue("aliyun", "AccessKeyID")
	this.AccessSecret, _ = cfg.GetValue("aliyun", "AccessSecret")

	this.XingeAccessId, _ = cfg.Int64("xinge", "XingeAccessId")
	this.XingeSecretKey, _ = cfg.GetValue("xinge", "XingeSecretKey")

	return true
}

func newconfig() *config {
	return &config{
		DB: db.NewDB(),
	}
}
