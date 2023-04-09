package dto

import (
	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type DyAdsGetPageReq struct {
	dto.Pagination `search:"-"`
	Account        string `form:"account"  search:"type:contains;column:account;table:dy_ads" comment:"抖音号"`
	DyAdsOrder
}

type DyAdsOrder struct {
	Id             string `form:"idOrder"  search:"type:order;column:id;table:dy_ads"`
	Account        string `form:"accountOrder"  search:"type:order;column:account;table:dy_ads"`
	Code           string `form:"codeOrder"  search:"type:order;column:code;table:dy_ads"`
	UserId         string `form:"userIdOrder"  search:"type:order;column:user_id;table:dy_ads"`
	FollowCnt      string `form:"followCntOrder"  search:"type:order;column:follow_cnt;table:dy_ads"`
	Status         string `form:"statusOrder"  search:"type:order;column:status;table:dy_ads"`
	CheckStatus    string `form:"checkStatusOrder"  search:"type:order;column:check_status;table:dy_ads"`
	Msg            string `form:"msgOrder"  search:"type:order;column:msg;table:dy_ads"`
	DeptId         string `form:"deptIdOrder"  search:"type:order;column:dept_id;table:dy_ads"`
	ConversionRate string `form:"conversionRateOrder"  search:"type:order;column:conversion_rate;table:dy_ads"`
	ConversionCost string `form:"conversionCostOrder"  search:"type:order;column:conversion_cost;table:dy_ads"`
	CreatedAt      string `form:"createdAtOrder"  search:"type:order;column:created_at;table:dy_ads"`
	UpdatedAt      string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:dy_ads"`
	DeletedAt      string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:dy_ads"`
	CreateBy       string `form:"createByOrder"  search:"type:order;column:create_by;table:dy_ads"`
	UpdateBy       string `form:"updateByOrder"  search:"type:order;column:update_by;table:dy_ads"`
}

func (m *DyAdsGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type DyAdsInsertReq struct {
	Id             int    `json:"-" comment:"主键编码"` // 主键编码
	Account        string `json:"account" comment:"抖音号"`
	Code           string `json:"code" comment:"授权码"`
	UserId         int64  `json:"userId" comment:"抖音id"`
	FollowCnt      string `json:"followCnt" comment:"关注数"`
	Status         string `json:"status" comment:"任务状态"`
	CheckStatus    int64  `json:"checkStatus" comment:"检查状态0检查抖音号，1检查直播状态"`
	Msg            string `json:"msg" comment:"提示信息"`
	DeptId         int    `json:"deptId" comment:"所属组织id"`
	ConversionRate string `json:"conversionRate" comment:"转化率"`
	ConversionCost int64  `json:"conversionCost" comment:"消耗"`
	common.ControlBy
}

func (s *DyAdsInsertReq) Generate(model *models.DyAds) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Account = s.Account
	model.Code = s.Code
	model.UserId = s.UserId
	model.FollowCnt = s.FollowCnt
	model.Status = s.Status
	model.CheckStatus = s.CheckStatus
	model.Msg = s.Msg
	model.DeptId = s.DeptId
	model.ConversionRate = s.ConversionRate
	model.ConversionCost = s.ConversionCost
	model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *DyAdsInsertReq) GetId() interface{} {
	return s.Id
}

type DyAdsUpdateReq struct {
	Id             int    `uri:"id" comment:"主键编码"` // 主键编码
	Account        string `json:"account" comment:"抖音号"`
	Code           string `json:"code" comment:"授权码"`
	UserId         int64  `json:"userId" comment:"抖音id"`
	FollowCnt      string `json:"followCnt" comment:"关注数"`
	Status         string `json:"status" comment:"任务状态"`
	CheckStatus    int64  `json:"checkStatus" comment:"检查状态0检查抖音号，1检查直播状态"`
	Msg            string `json:"msg" comment:"提示信息"`
	DeptId         int    `json:"deptId" comment:"所属组织id"`
	ConversionRate string `json:"conversionRate" comment:"转化率"`
	ConversionCost int64  `json:"conversionCost" comment:"消耗"`
	common.ControlBy
}

func (s *DyAdsUpdateReq) Generate(model *models.DyAds) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Account = s.Account
	model.Code = s.Code
	model.UserId = s.UserId
	model.FollowCnt = s.FollowCnt
	model.Status = s.Status
	model.CheckStatus = s.CheckStatus
	model.Msg = s.Msg
	model.DeptId = s.DeptId
	model.ConversionRate = s.ConversionRate
	model.ConversionCost = s.ConversionCost
	model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *DyAdsUpdateReq) GetId() interface{} {
	return s.Id
}

// DyAdsGetReq 功能获取请求参数
type DyAdsGetReq struct {
	Id int `uri:"id"`
}

func (s *DyAdsGetReq) GetId() interface{} {
	return s.Id
}

// DyAdsDeleteReq 功能删除请求参数
type DyAdsDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *DyAdsDeleteReq) GetId() interface{} {
	return s.Ids
}
