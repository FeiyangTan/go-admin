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

type WechatUsers struct {
	service.Service
}

// GetPage 获取WechatUsers列表
func (e *WechatUsers) GetPage(c *dto.WechatUsersGetPageReq, p *actions.DataPermission, list *[]models.WechatUsers, count *int64) error {
	var err error
	var data models.WechatUsers

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("WechatUsersService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取WechatUsers对象
func (e *WechatUsers) Get(d *dto.WechatUsersGetReq, p *actions.DataPermission, model *models.WechatUsers) error {
	var data models.WechatUsers

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetWechatUsers error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建WechatUsers对象
func (e *WechatUsers) Insert(c *dto.WechatUsersInsertReq) error {
	var err error
	var data models.WechatUsers
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("WechatUsersService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改WechatUsers对象
func (e *WechatUsers) Update(c *dto.WechatUsersUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.WechatUsers{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	c.Generate(&data)

	db := e.Orm.Save(&data)
	if err = db.Error; err != nil {
		e.Log.Errorf("WechatUsersService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除WechatUsers
func (e *WechatUsers) Remove(d *dto.WechatUsersDeleteReq, p *actions.DataPermission) error {
	var data models.WechatUsers

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveWechatUsers error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}
