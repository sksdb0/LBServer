package mongo

import (
	"logger"

	"gopkg.in/mgo.v2/bson"
)

func (this *MongoManager) IsExist(db string, coll string, query bson.M) bool {
	c := this.GetDB(db).C(coll)
	count, err := c.Find(query).Count()
	if err != nil || count == 0 {
		logger.LOGLINE(err)
		return false
	}
	return true
}

func (this *MongoManager) Count(db string, coll string, query bson.M) (int, error) {
	c := this.GetDB(db).C(coll)
	count, err := c.Find(query).Count()
	if err != nil {
		return -1, err
	}
	return count, nil
}

func (this *MongoManager) DistinctCount(db string, coll string, query bson.M, distinct string) (int, error) {
	c := this.GetDB(db).C(coll)
	var tmp []int
	err := c.Find(query).Distinct(distinct, &tmp)
	if err != nil {
		logger.PRINTLINE(err)
		return -1, err
	}

	return len(tmp), nil
}

func (this *MongoManager) IsCollExist(db string, coll string) bool {
	c := this.GetDB(db).C(coll)
	n, _ := c.Count()

	if n > 0 {
		return true
	}

	return false
}

func (this *MongoManager) Max(db string, coll string, query bson.M, sort string, i interface{}) bool {
	c := this.GetDB(db).C(coll)
	count, err := c.Find(query).Count()
	if err != nil {
		logger.LOGLINE(err)
		return false
	}
	if count == 0 {
		logger.PRINTLINE(count)
		return false
	}

	err = c.Find(query).Sort(sort).Limit(1).All(i)
	if err != nil {
		logger.LOGLINE(err)
		return false
	}
	return true
}

func (this *MongoManager) Find(db string, coll string, query bson.M, field bson.M, i interface{}) bool {
	c := this.GetDB(db).C(coll)
	_, err := c.Find(query).Count()
	if err != nil {
		logger.LOGLINE(err)
		return false
	}

	err = c.Find(query).Select(field).One(i)
	if err != nil {
		logger.LOGLINE(err)
		return false
	}
	return true
}

func (this *MongoManager) FindAll(db string, coll string, query bson.M, sort string, fields bson.M, i interface{}) bool {
	c := this.GetDB(db).C(coll)
	_, err := c.Find(query).Count()
	if err != nil {
		logger.LOGLINE(err)
		return false
	}

	if sort == "" {
		err = c.Find(query).Select(fields).All(i)
	} else {
		err = c.Find(query).Sort(sort).Select(fields).All(i)
	}
	if err != nil {
		logger.LOGLINE(err)
		return false
	}
	return true
}

func (this *MongoManager) FullFind(db string, coll string, query bson.M, sort string, fields bson.M, limit int, i interface{}) bool {
	c := this.GetDB(db).C(coll)
	_, err := c.Find(query).Count()
	if err != nil {
		logger.LOGLINE(err)
		return false
	}

	err = c.Find(query).Sort(sort).Limit(limit).All(i)
	if err != nil {
		logger.LOGLINE(err)
		return false
	}
	return true
}

func (this *MongoManager) Update(db string, coll string, query bson.M, i interface{}) bool {
	c := this.GetDB(db).C(coll)
	err := c.Update(query, i)
	if err != nil {
		logger.LOGLINE(err)
		return false
	}
	return true
}

func (this *MongoManager) Insert(db string, coll string, i interface{}) bool {
	c := this.GetDB(db).C(coll)
	err := c.Insert(i)
	if err != nil {
		logger.LOGLINE(err)
		return false
	}
	return true
}

func (this *MongoManager) Remove(db string, coll string, i interface{}) {
	c := this.GetDB(db).C(coll)
	err := c.Remove(i)
	if err != nil {
		panic(err)
	}
}
