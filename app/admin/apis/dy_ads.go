package apis

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	_ "github.com/go-admin-team/go-admin-core/sdk/pkg/response"

	"go-admin/app/admin/models"
	"go-admin/app/admin/service"
	"go-admin/app/admin/service/dto"
	"go-admin/common/actions"
)

type DyAds struct {
	api.Api
}

// GetPage 获取抖音计划管理列表
// @Summary 获取抖音计划管理列表
// @Description 获取抖音计划管理列表
// @Tags 抖音计划管理
// @Param account query string false "抖音号"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.DyAds}} "{"code": 200, "data": [...]}"
// @Router /api/v1/dy-ads [get]
// @Security Bearer
func (e DyAds) GetPage(c *gin.Context) {
	req := dto.DyAdsGetPageReq{}
	s := service.DyAds{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	p := actions.GetPermissionFromContext(c)
	list := make([]models.DyAds, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取抖音计划管理失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取抖音计划管理
// @Summary 获取抖音计划管理
// @Description 获取抖音计划管理
// @Tags 抖音计划管理
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.DyAds} "{"code": 200, "data": [...]}"
// @Router /api/v1/dy-ads/{id} [get]
// @Security Bearer
func (e DyAds) Get(c *gin.Context) {
	req := dto.DyAdsGetReq{}
	s := service.DyAds{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	var object models.DyAds

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取抖音计划管理失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(object, "查询成功")
}

// Insert 创建抖音计划管理
// @Summary 创建抖音计划管理
// @Description 创建抖音计划管理
// @Tags 抖音计划管理
// @Accept application/json
// @Product application/json
// @Param data body dto.DyAdsInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/dy-ads [post]
// @Security Bearer
func (e DyAds) Insert(c *gin.Context) {
	req := dto.DyAdsInsertReq{}
	s := service.DyAds{}

	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	// 设置创建人
	req.SetCreateBy(user.GetUserId(c))
	req.DeptId = user.GetDeptId(c)

	err = s.Insert(&req)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("创建抖音计划管理失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改抖音计划管理
// @Summary 修改抖音计划管理
// @Description 修改抖音计划管理
// @Tags 抖音计划管理
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.DyAdsUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/dy-ads/{id} [put]
// @Security Bearer
func (e DyAds) Update(c *gin.Context) {
	req := dto.DyAdsUpdateReq{}
	s := service.DyAds{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	req.SetUpdateBy(user.GetUserId(c))
	p := actions.GetPermissionFromContext(c)

	err = s.Update(&req, p)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("修改抖音计划管理失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "修改成功")
}

// Delete 删除抖音计划管理
// @Summary 删除抖音计划管理
// @Description 删除抖音计划管理
// @Tags 抖音计划管理
// @Param data body dto.DyAdsDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/dy-ads [delete]
// @Security Bearer
func (e DyAds) Delete(c *gin.Context) {
	s := service.DyAds{}
	req := dto.DyAdsDeleteReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	// req.SetUpdateBy(user.GetUserId(c))
	p := actions.GetPermissionFromContext(c)

	err = s.Remove(&req, p)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("删除抖音计划管理失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "删除成功")
}
