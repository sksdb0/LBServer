package mongo

import (
	"config"
	"io"
	"logger"
	"os"
)

func (this *MongoManager) SearchFile(db string, coll string, fielname string) (bool, []byte) {
	file, err := this.GetDB(db).GridFS(coll).Open(fielname)
	if err != nil {
		logger.LOGLINE(err)
		return false, nil
	}

	data := make([]byte, file.Size())
	_, err = file.Read(data)
	if err != nil {
		logger.LOGLINE(err)
		return false, nil
	}
	return true, data
}

func (this *MongoManager) InsertFile(db string, coll string, filepath string, fielname string) bool {
	file, err := this.GetDB(db).GridFS(coll).Create(fielname)
	if err != nil {
		logger.LOGLINE(err)
		return false
	}

	filedata, err := os.OpenFile(filepath, os.O_RDWR, 0666)
	if err != nil {
		logger.LOGLINE(err)
		return false
	}

	_, err = io.Copy(file, filedata)
	if err != nil {
		logger.LOGLINE(err)
		return false
	}

	err = file.Close()
	if err != nil {
		logger.LOGLINE(err)
		return false
	}

	err = filedata.Close()
	if err != nil {
		logger.LOGLINE(err)
		return false
	}
	return true
}

func (this *MongoManager) RemoveFile(db string, coll string, fielname string) bool {
	err := this.GetDB(db).GridFS(coll).Remove(fielname)
	if err != nil {
		logger.LOGLINE(err)
		return false
	}
	return true
}

func (this *MongoManager) InitFile() {
	this.InsertFile(config.DB().DBName, config.DB().CollMap["picture"], "picture/fruit/classification_fruit_apple.jpg", "classification_fruit_apple.jpg")
	this.InsertFile(config.DB().DBName, config.DB().CollMap["picture"], "picture/fruit/classification_fruit_cherry.jpg", "classification_fruit_cherry.jpg")
	this.InsertFile(config.DB().DBName, config.DB().CollMap["picture"], "picture/fruit/classification_fruit_grape.jpg", "classification_fruit_grape.jpg")
	this.InsertFile(config.DB().DBName, config.DB().CollMap["picture"], "picture/fruit/classification_fruit_kiwi.jpg", "classification_fruit_kiwi.jpg")
	this.InsertFile(config.DB().DBName, config.DB().CollMap["picture"], "picture/fruit/classification_fruit_mango.jpg", "classification_fruit_mango.jpg")
	this.InsertFile(config.DB().DBName, config.DB().CollMap["picture"], "picture/fruit/classification_fruit_melons.jpg", "classification_fruit_melons.jpg")
	this.InsertFile(config.DB().DBName, config.DB().CollMap["picture"], "picture/fruit/classification_fruit_orange.jpg", "classification_fruit_orange.jpg")
	this.InsertFile(config.DB().DBName, config.DB().CollMap["picture"], "picture/fruit/classification_fruit_pear.jpg", "classification_fruit_pear.jpg")
	this.InsertFile(config.DB().DBName, config.DB().CollMap["picture"], "picture/fruit/classification_fruit_more.jpg", "classification_fruit_more.jpg")
}
