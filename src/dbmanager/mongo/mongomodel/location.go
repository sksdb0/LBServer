package mongomodel

import (
	"time"
)

type LocationModel struct {
	View    *DailyLocationView
	Typemap map[string]string
}

func NewLocationModel(date time.Time) *LocationModel {
	locationmodel := LocationModel{
		View:    newDailyLocationView(date),
		Typemap: make(map[string]string),
	}
	locationmodel.Typemap["北京市"] = "北京"
	locationmodel.Typemap["天津市"] = "天津"
	locationmodel.Typemap["河北省"] = "河北"
	locationmodel.Typemap["山西省"] = "山西"
	locationmodel.Typemap["内蒙古自治区"] = "内蒙古"
	locationmodel.Typemap["辽宁省"] = "辽宁"
	locationmodel.Typemap["吉林省"] = "吉林"
	locationmodel.Typemap["黑龙江省"] = "黑龙江"
	locationmodel.Typemap["上海市"] = "上海"
	locationmodel.Typemap["江苏省"] = "江苏"
	locationmodel.Typemap["浙江省"] = "浙江"
	locationmodel.Typemap["安徽省"] = "安徽"
	locationmodel.Typemap["福建省"] = "福建"
	locationmodel.Typemap["江西省"] = "江西"
	locationmodel.Typemap["山东省"] = "山东"
	locationmodel.Typemap["河南省"] = "河南"
	locationmodel.Typemap["湖北省"] = "湖北"
	locationmodel.Typemap["湖南省"] = "湖南"
	locationmodel.Typemap["广东省"] = "广东"
	locationmodel.Typemap["广西壮族自治区"] = "广西"
	locationmodel.Typemap["海南省"] = "海南"
	locationmodel.Typemap["重庆市"] = "重庆"
	locationmodel.Typemap["四川省"] = "四川"
	locationmodel.Typemap["贵州省"] = "贵州"
	locationmodel.Typemap["云南省"] = "云南"
	locationmodel.Typemap["西藏自治区"] = "西藏"
	locationmodel.Typemap["陕西省"] = "陕西"
	locationmodel.Typemap["甘肃省"] = "甘肃"
	locationmodel.Typemap["青海省"] = "青海"
	locationmodel.Typemap["宁夏回族自治区"] = "宁夏"
	locationmodel.Typemap["新疆维吾尔自治区"] = "新疆"
	locationmodel.Typemap["台湾省"] = "台湾"
	locationmodel.Typemap["香港特别行政区"] = "香港"
	locationmodel.Typemap["澳门特别行政区"] = "澳门"

	return &locationmodel
}

func (this *LocationModel) AddChinaView(index string, value int) {
	this.View.ViewData = append(this.View.ViewData, item{Key: this.Typemap[index], Value: value})
}

func (this *LocationModel) AddGlobalView(index string, value int) {
	this.View.ViewData = append(this.View.ViewData, item{Key: index, Value: value})
}

type DailyLocationView struct {
	Date     time.Time
	ViewData []item
}

func newDailyLocationView(date time.Time) *DailyLocationView {
	return &DailyLocationView{
		Date: date,
	}
}
