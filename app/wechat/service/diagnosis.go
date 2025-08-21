// app/wechat/service/user.go
package service

import (
	"fmt"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"go-admin/app/wechat/models"
	"gorm.io/datatypes"
)

type WechatDiagnosisService struct {
	api.Api
}

// NewWechatDiagnosisService 会在 Handler 里被 MakeService 调用
func NewWechatDiagnosisService(e *api.Api) *WechatDiagnosisService {
	return &WechatDiagnosisService{*e}
}

// 新增Diagnosis记录
func (s *WechatDiagnosisService) AddDiagnosis(openid string, diagnosisType string, diagnosisResult datatypes.JSON) error {
	db, _ := s.GetOrm() // 从 s.Api 已初始化的 Orm 拿到 *gorm.DB

	// 新增Diagnosis记录
	record := models.Diagnosis{
		OpenID:          openid,
		DiagnosisResult: diagnosisResult,
		DiagnosisType:   diagnosisType,
	}

	// 执行插入
	if err := db.Create(&record).Error; err != nil {
		return err
	}

	return nil

}

// GetDiagnosisList 查询符合条件的所有诊疗记录列表
func (s *WechatDiagnosisService) GetDiagnosisList(openid string, diagnosisType string, pageSize int, pageIndex int) (diagnoses []models.Diagnosis, total int64, err error) {

	db, _ := s.GetOrm() // 从 s.Api 已初始化的 Orm 拿到 *gorm.DB
	fmt.Println(pageSize, pageIndex)

	if pageSize <= 0 {
		pageSize = 10
	}
	if pageIndex <= 0 {
		pageIndex = 1
	}
	offset := (pageIndex - 1) * pageSize

	fmt.Println(pageSize, pageIndex)
	// 1. 查询总记录数
	if err = db.Model(&models.Diagnosis{}).
		Where("open_id = ? AND diagnosis_type = ?", openid, diagnosisType).
		Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 2. 查询分页数据
	if err = db.Model(&models.Diagnosis{}).
		Where("open_id = ? AND diagnosis_type = ?", openid, diagnosisType).
		Order("created_at DESC").
		Limit(pageSize).
		Offset(offset).
		Find(&diagnoses).Error; err != nil {
		return nil, 0, err
	}

	return diagnoses, total, nil
}
