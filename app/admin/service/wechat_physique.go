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

type WechatPhysique struct {
	service.Service
}

// GetPage 获取WechatPhysique列表
func (e *WechatPhysique) GetPage(c *dto.WechatPhysiqueGetPageReq, p *actions.DataPermission, list *[]models.WechatPhysique, count *int64) error {
	var err error
	var data models.WechatPhysique

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("WechatPhysiqueService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取WechatPhysique对象
func (e *WechatPhysique) Get(d *dto.WechatPhysiqueGetReq, p *actions.DataPermission, model *models.WechatPhysique) error {
	var data models.WechatPhysique

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetWechatPhysique error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建WechatPhysique对象
func (e *WechatPhysique) Insert(c *dto.WechatPhysiqueInsertReq) error {
	var err error
	var data models.WechatPhysique
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("WechatPhysiqueService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改WechatPhysique对象
func (e *WechatPhysique) Update(c *dto.WechatPhysiqueUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.WechatPhysique{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	c.Generate(&data)

	db := e.Orm.Save(&data)
	if err = db.Error; err != nil {
		e.Log.Errorf("WechatPhysiqueService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除WechatPhysique
func (e *WechatPhysique) Remove(d *dto.WechatPhysiqueDeleteReq, p *actions.DataPermission) error {
	var data models.WechatPhysique

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveWechatPhysique error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}
