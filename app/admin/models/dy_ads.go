package models

import (
	"go-admin/common/models"
)

type DyAds struct {
	models.Model

	Account        string `json:"account" gorm:"type:varchar(128);comment:抖音号"`
	Code           string `json:"code" gorm:"type:varchar(32);comment:授权码"`
	UserId         int64  `json:"userId" gorm:"type:int;comment:抖音id"`
	FollowCnt      string `json:"followCnt" gorm:"type:int;comment:关注数"`
	Status         string `json:"status" gorm:"type:tinyint;comment:任务状态"`
	CheckStatus    int64  `json:"checkStatus" gorm:"type:tinyint;comment:检查状态0检查抖音号，1检查直播状态"`
	Msg            string `json:"msg" gorm:"type:varchar(64);comment:提示信息"`
	DeptId         int    `json:"deptId" gorm:"type:int;comment:所属组织id"`
	ConversionRate string `json:"conversionRate" gorm:"type:double;comment:转化率"`
	ConversionCost int64  `json:"conversionCost" gorm:"type:double;comment:消耗"`
	models.ModelTime
	models.ControlBy
}

func (DyAds) TableName() string {
	return "dy_ads"
}

func (e *DyAds) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *DyAds) GetId() interface{} {
	return e.Id
}
