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

type WechatProducts struct {
	service.Service
}

// GetPage 获取WechatProducts列表
func (e *WechatProducts) GetPage(c *dto.WechatProductsGetPageReq, p *actions.DataPermission, list *[]models.WechatProducts, count *int64) error {
	var err error
	var data models.WechatProducts

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("WechatProductsService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取WechatProducts对象
func (e *WechatProducts) Get(d *dto.WechatProductsGetReq, p *actions.DataPermission, model *models.WechatProducts) error {
	var data models.WechatProducts

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetWechatProducts error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建WechatProducts对象
func (e *WechatProducts) Insert(c *dto.WechatProductsInsertReq) error {
	var err error
	var data models.WechatProducts
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("WechatProductsService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改WechatProducts对象
func (e *WechatProducts) Update(c *dto.WechatProductsUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.WechatProducts{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	c.Generate(&data)

	db := e.Orm.Save(&data)
	if err = db.Error; err != nil {
		e.Log.Errorf("WechatProductsService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除WechatProducts
func (e *WechatProducts) Remove(d *dto.WechatProductsDeleteReq, p *actions.DataPermission) error {
	var data models.WechatProducts

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveWechatProducts error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}
