// app/wechat/service/user.go
package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"go-admin/app/wechat/models"
	"gorm.io/datatypes"
	"strings"
)

type WechatDiagnosisService struct {
	api.Api
}

// NewWechatDiagnosisService 会在 Handler 里被 MakeService 调用
func NewWechatDiagnosisService(e *api.Api) *WechatDiagnosisService {
	return &WechatDiagnosisService{*e}
}

// 新增Diagnosis记录
func (s *WechatDiagnosisService) AddDiagnosis(openid string, diagnosisType string, diagnosisPhysique string, dataMap map[string]interface{}) error {
	db, _ := s.GetOrm() // 从 s.Api 已初始化的 Orm 拿到 *gorm.DB

	// 获取 data 中的 "data" 字段
	var diagnosisJSON datatypes.JSON
	if b, err := json.Marshal(dataMap); err == nil {
		diagnosisJSON = datatypes.JSON(b)
	} else {
		return errors.New("诊断结果序列化失败: " + err.Error())
	}

	// 获取diagnosisPhysique中的信息
	var majorDiagnosis string
	var minorDiagnosis string
	majorDiagnosis, minorDiagnosis, ok := strings.Cut(diagnosisPhysique, "、")
	if !ok {
		majorDiagnosis = diagnosisPhysique
		minorDiagnosis = ""
	}

	// 新增Diagnosis记录
	record := models.Diagnosis{
		OpenID:             openid,
		DiagnosisResult:    diagnosisJSON,
		DiagnosisType:      diagnosisType,
		DiagnosisPhysique1: majorDiagnosis,
		DiagnosisPhysique2: minorDiagnosis,
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

// 获取第一个体质名称
func GetPhysiqueName(dataMap map[string]interface{}) (string, string) {
	//fmt.Println("~~~physique:", physique)
	// 1. 获取体质名称
	var physiqueString string
	var newPhysiqueString string
	if s, ok := dataMap["physique_name"].(string); ok {
		physiqueString = s
		newPhysiqueString = strings.ReplaceAll(s, "体质", "")
	} else {
		physiqueString = "默认体质"
	}
	//fmt.Println("physique string:", physiqueString)
	// 2. 获取第一个体质名称
	var firstPhysique string
	if idx := strings.Index(physiqueString, "、"); idx != -1 {
		firstPhysique = physiqueString[:idx]
	} else {
		firstPhysique = physiqueString
	}

	return firstPhysique, newPhysiqueString
}

// 添加自定义疗法
func GetAcupunctureMethod(dataMap map[string]interface{}, acupunctureMethod *models.Physique) map[string]interface{} {
	if advicesMap, ok := dataMap["advices"].(map[string]interface{}); ok {
		if treatments, ok := advicesMap["treatment"].([]interface{}); ok {

			// 构造新的一组内容
			newItems := []interface{}{
				map[string]interface{}{"advice": acupunctureMethod.AcupunctureMethod,
					"title": "三才河洛灸组穴"},
				map[string]interface{}{"advice": acupunctureMethod.EighteenMethod,
					"title": "药食同源亿草十八宝调理"},
				map[string]interface{}{"advice": acupunctureMethod.WellnessMethod,
					"title": "日常养生建议"},
			}

			// 添加到数组
			treatments = append(newItems, treatments...)

			advicesMap["treatment"] = treatments

		}
		return advicesMap
	}
	return nil
}

// 修改Diagnosis的node
func (s *WechatDiagnosisService) EditDiagnosisNote(id int, note string) error {
	db, _ := s.GetOrm() // 从 s.Api 已初始化的 Orm 拿到 *gorm.DB

	// 更新记录
	result := db.Model(&models.Diagnosis{}).
		Where("id = ?", id).
		Update("diagnosis_note", note)

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("no rows updated / record not found")
	}

	return nil
}
