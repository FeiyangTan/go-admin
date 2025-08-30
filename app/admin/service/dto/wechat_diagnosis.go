package dto

import (
	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type WechatDiagnosisGetPageReq struct {
	dto.Pagination     `search:"-"`
	OpenId             string `form:"openId"  search:"type:exact;column:open_id;table:wechat_diagnosis" comment:"微信昵称"`
	DiagnosisType      string `form:"diagnosisType"  search:"type:exact;column:diagnosis_type;table:wechat_diagnosis" comment:"诊疗类型"`
	DiagnosisPhysique1 string `form:"diagnosisPhysique1"  search:"type:exact;column:diagnosis_physique1;table:wechat_diagnosis" comment:"体质1名称"`
	DiagnosisPhysique2 string `form:"diagnosisPhysique2"  search:"type:exact;column:diagnosis_physique2;table:wechat_diagnosis" comment:"体质2名称"`
	DiagnosisResult    string `form:"diagnosisResult"  search:"type:contains;column:diagnosis_result;table:wechat_diagnosis" comment:"诊疗结果 JSON 字符串"`
	WechatDiagnosisOrder
}

type WechatDiagnosisOrder struct {
	Id                 string `form:"idOrder"  search:"type:order;column:id;table:wechat_diagnosis"`
	OpenId             string `form:"openIdOrder"  search:"type:order;column:open_id;table:wechat_diagnosis"`
	DiagnosisType      string `form:"diagnosisTypeOrder"  search:"type:order;column:diagnosis_type;table:wechat_diagnosis"`
	DiagnosisPhysique1 string `form:"diagnosisPhysique1Order"  search:"type:order;column:diagnosis_physique1;table:wechat_diagnosis"`
	DiagnosisPhysique2 string `form:"diagnosisPhysique2Order"  search:"type:order;column:diagnosis_physique2;table:wechat_diagnosis"`
	DiagnosisResult    string `form:"diagnosisResultOrder"  search:"type:order;column:diagnosis_result;table:wechat_diagnosis"`
	CreatedAt          string `form:"createdAtOrder"  search:"type:order;column:created_at;table:wechat_diagnosis"`
	UpdatedAt          string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:wechat_diagnosis"`
	DeletedAt          string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:wechat_diagnosis"`
	CreateBy           string `form:"createByOrder"  search:"type:order;column:create_by;table:wechat_diagnosis"`
	UpdateBy           string `form:"updateByOrder"  search:"type:order;column:update_by;table:wechat_diagnosis"`
}

func (m *WechatDiagnosisGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type WechatDiagnosisInsertReq struct {
	Id                 int    `json:"-" comment:"自增主键"` // 自增主键
	OpenId             string `json:"openId" comment:"微信昵称"`
	DiagnosisType      string `json:"diagnosisType" comment:"诊疗类型"`
	DiagnosisPhysique1 string `json:"diagnosisPhysique1" comment:"体质1名称"`
	DiagnosisPhysique2 string `json:"diagnosisPhysique2" comment:"体质2名称"`
	DiagnosisResult    string `json:"diagnosisResult" comment:"诊疗结果 JSON 字符串"`
	common.ControlBy
}

func (s *WechatDiagnosisInsertReq) Generate(model *models.WechatDiagnosis) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.OpenId = s.OpenId
	model.DiagnosisType = s.DiagnosisType
	model.DiagnosisPhysique1 = s.DiagnosisPhysique1
	model.DiagnosisPhysique2 = s.DiagnosisPhysique2
	model.DiagnosisResult = s.DiagnosisResult
	model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *WechatDiagnosisInsertReq) GetId() interface{} {
	return s.Id
}

type WechatDiagnosisUpdateReq struct {
	Id                 int    `uri:"id" comment:"自增主键"` // 自增主键
	OpenId             string `json:"openId" comment:"微信昵称"`
	DiagnosisType      string `json:"diagnosisType" comment:"诊疗类型"`
	DiagnosisPhysique1 string `json:"diagnosisPhysique1" comment:"体质1名称"`
	DiagnosisPhysique2 string `json:"diagnosisPhysique2" comment:"体质2名称"`
	DiagnosisResult    string `json:"diagnosisResult" comment:"诊疗结果 JSON 字符串"`
	common.ControlBy
}

func (s *WechatDiagnosisUpdateReq) Generate(model *models.WechatDiagnosis) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.OpenId = s.OpenId
	model.DiagnosisType = s.DiagnosisType
	model.DiagnosisPhysique1 = s.DiagnosisPhysique1
	model.DiagnosisPhysique2 = s.DiagnosisPhysique2
	model.DiagnosisResult = s.DiagnosisResult
	model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *WechatDiagnosisUpdateReq) GetId() interface{} {
	return s.Id
}

// WechatDiagnosisGetReq 功能获取请求参数
type WechatDiagnosisGetReq struct {
	Id int `uri:"id"`
}

func (s *WechatDiagnosisGetReq) GetId() interface{} {
	return s.Id
}

// WechatDiagnosisDeleteReq 功能删除请求参数
type WechatDiagnosisDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *WechatDiagnosisDeleteReq) GetId() interface{} {
	return s.Ids
}
