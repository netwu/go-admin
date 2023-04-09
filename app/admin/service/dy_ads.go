package service

import (
	"errors"

	"github.com/go-admin-team/go-admin-core/sdk/service"
	"gorm.io/gorm"

	"go-admin/app/admin/models"
	"go-admin/app/admin/service/dto"
	"go-admin/common/actions"
	cDto "go-admin/common/dto"
)

type DyAds struct {
	service.Service
}

// GetPage 获取DyAds列表
func (e *DyAds) GetPage(c *dto.DyAdsGetPageReq, p *actions.DataPermission, list *[]models.DyAds, count *int64) error {
	var err error
	var data models.DyAds

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("DyAdsService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取DyAds对象
func (e *DyAds) Get(d *dto.DyAdsGetReq, p *actions.DataPermission, model *models.DyAds) error {
	var data models.DyAds

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetDyAds error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建DyAds对象
func (e *DyAds) Insert(c *dto.DyAdsInsertReq) error {
	var err error
	var data models.DyAds
	c.Generate(&data)

	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("DyAdsService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改DyAds对象
func (e *DyAds) Update(c *dto.DyAdsUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.DyAds{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	c.Generate(&data)

	db := e.Orm.Save(&data)
	if err = db.Error; err != nil {
		e.Log.Errorf("DyAdsService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除DyAds
func (e *DyAds) Remove(d *dto.DyAdsDeleteReq, p *actions.DataPermission) error {
	var data models.DyAds

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveDyAds error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}

// Remove 删除DyAds
func (e *DyAds) GetUserProfile(createBy int, user *models.SysUser) error {
	err := e.Orm.Preload("Dept").First(user, createBy).Error
	if err != nil {
		return err
	}
	return nil
}
