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

type WechatDiagnosis struct {
	service.Service
}

// GetPage 获取WechatDiagnosis列表
func (e *WechatDiagnosis) GetPage(c *dto.WechatDiagnosisGetPageReq, p *actions.DataPermission, list *[]models.WechatDiagnosis, count *int64) error {
	var err error
	var data models.WechatDiagnosis

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("WechatDiagnosisService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取WechatDiagnosis对象
func (e *WechatDiagnosis) Get(d *dto.WechatDiagnosisGetReq, p *actions.DataPermission, model *models.WechatDiagnosis) error {
	var data models.WechatDiagnosis

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetWechatDiagnosis error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建WechatDiagnosis对象
func (e *WechatDiagnosis) Insert(c *dto.WechatDiagnosisInsertReq) error {
	var err error
	var data models.WechatDiagnosis
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("WechatDiagnosisService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改WechatDiagnosis对象
func (e *WechatDiagnosis) Update(c *dto.WechatDiagnosisUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.WechatDiagnosis{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	c.Generate(&data)

	db := e.Orm.Save(&data)
	if err = db.Error; err != nil {
		e.Log.Errorf("WechatDiagnosisService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除WechatDiagnosis
func (e *WechatDiagnosis) Remove(d *dto.WechatDiagnosisDeleteReq, p *actions.DataPermission) error {
	var data models.WechatDiagnosis

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveWechatDiagnosis error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}
