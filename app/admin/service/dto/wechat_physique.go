package dto

import (
	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type WechatPhysiqueGetPageReq struct {
	dto.Pagination `search:"-"`
	PhysiqueName   string `form:"physiqueName"  search:"type:exact;column:physique_name;table:wechat_physique" comment:"体质名称"`
	WechatPhysiqueOrder
}

type WechatPhysiqueOrder struct {
	Id                string `form:"idOrder"  search:"type:order;column:id;table:wechat_physique"`
	PhysiqueName      string `form:"physiqueNameOrder"  search:"type:order;column:physique_name;table:wechat_physique"`
	AcupunctureMethod string `form:"acupunctureMethodOrder"  search:"type:order;column:acupuncture_method;table:wechat_physique"`
	ProductIds        string `form:"productIdsOrder"  search:"type:order;column:product_ids;table:wechat_physique"`
	CreatedAt         string `form:"createdAtOrder"  search:"type:order;column:created_at;table:wechat_physique"`
	UpdatedAt         string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:wechat_physique"`
	DeletedAt         string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:wechat_physique"`
}

func (m *WechatPhysiqueGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type WechatPhysiqueInsertReq struct {
	Id                int    `json:"-" comment:"自增主键"` // 自增主键
	PhysiqueName      string `json:"physiqueName" comment:"体质名称"`
	AcupunctureMethod string `json:"acupunctureMethod" comment:"针灸法"`
	ProductIds        string `json:"productIds" comment:"商品ID数组（以逗号分隔的字符串）"`
	common.ControlBy
}

func (s *WechatPhysiqueInsertReq) Generate(model *models.WechatPhysique) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.PhysiqueName = s.PhysiqueName
	model.AcupunctureMethod = s.AcupunctureMethod
	model.ProductIds = s.ProductIds
}

func (s *WechatPhysiqueInsertReq) GetId() interface{} {
	return s.Id
}

type WechatPhysiqueUpdateReq struct {
	Id                int    `uri:"id" comment:"自增主键"` // 自增主键
	PhysiqueName      string `json:"physiqueName" comment:"体质名称"`
	AcupunctureMethod string `json:"acupunctureMethod" comment:"针灸法"`
	ProductIds        string `json:"productIds" comment:"商品ID数组（以逗号分隔的字符串）"`
	common.ControlBy
}

func (s *WechatPhysiqueUpdateReq) Generate(model *models.WechatPhysique) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.PhysiqueName = s.PhysiqueName
	model.AcupunctureMethod = s.AcupunctureMethod
	model.ProductIds = s.ProductIds
}

func (s *WechatPhysiqueUpdateReq) GetId() interface{} {
	return s.Id
}

// WechatPhysiqueGetReq 功能获取请求参数
type WechatPhysiqueGetReq struct {
	Id int `uri:"id"`
}

func (s *WechatPhysiqueGetReq) GetId() interface{} {
	return s.Id
}

// WechatPhysiqueDeleteReq 功能删除请求参数
type WechatPhysiqueDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *WechatPhysiqueDeleteReq) GetId() interface{} {
	return s.Ids
}
