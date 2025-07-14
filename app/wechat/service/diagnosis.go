// app/wechat/service/user.go
package service

import (
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
func (s *WechatDiagnosisService) GetDiagnosisList(openid string, diagnosisType string) (diagnoses []models.Diagnosis, err error) {
	db, _ := s.GetOrm() // 从 s.Api 已初始化的 Orm 拿到 *gorm.DB

	// 执行查询
	err = db.
		Model(&models.Diagnosis{}).
		Where("open_id = ? AND diagnosis_type = ?", openid, diagnosisType).
		Order("created_at DESC").
		Find(&diagnoses).Error
	if err != nil {
		return nil, err
	}

	return diagnoses, nil
}
