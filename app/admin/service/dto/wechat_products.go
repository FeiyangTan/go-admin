package dto

import (
	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type WechatProductsGetPageReq struct {
	dto.Pagination `search:"-"`
	ProductName    string `form:"productName"  search:"type:exact;column:product_name;table:wechat_products" comment:"商品名称"`
	MallProductId  string `form:"mallProductId"  search:"type:exact;column:mall_product_id;table:wechat_products" comment:"商城商品编号"`
	WechatProductsOrder
}

type WechatProductsOrder struct {
	Id            string `form:"idOrder"  search:"type:order;column:id;table:wechat_products"`
	ProductName   string `form:"productNameOrder"  search:"type:order;column:product_name;table:wechat_products"`
	ImageUrl      string `form:"imageUrlOrder"  search:"type:order;column:image_url;table:wechat_products"`
	Description   string `form:"descriptionOrder"  search:"type:order;column:description;table:wechat_products"`
	Ingredients   string `form:"ingredientsOrder"  search:"type:order;column:ingredients;table:wechat_products"`
	UsageMethod   string `form:"usageMethodOrder"  search:"type:order;column:usage_method;table:wechat_products"`
	Price         string `form:"priceOrder"  search:"type:order;column:price;table:wechat_products"`
	MallProductId string `form:"mallProductIdOrder"  search:"type:order;column:mall_product_id;table:wechat_products"`
	CreatedAt     string `form:"createdAtOrder"  search:"type:order;column:created_at;table:wechat_products"`
	UpdatedAt     string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:wechat_products"`
	DeletedAt     string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:wechat_products"`
}

func (m *WechatProductsGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type WechatProductsInsertReq struct {
	Id            int    `json:"-" comment:"自增主键"` // 自增主键
	ProductName   string `json:"productName" comment:"商品名称"`
	ImageUrl      string `json:"imageUrl" comment:"商品图片 URL"`
	Description   string `json:"description" comment:"产品介绍"`
	Ingredients   string `json:"ingredients" comment:"产品成分"`
	UsageMethod   string `json:"usageMethod" comment:"使用方法"`
	Price         string `json:"price" comment:"商品价格（单位：元）"`
	MallProductId string `json:"mallProductId" comment:"商城商品编号"`
	common.ControlBy
}

func (s *WechatProductsInsertReq) Generate(model *models.WechatProducts) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.ProductName = s.ProductName
	model.ImageUrl = s.ImageUrl
	model.Description = s.Description
	model.Ingredients = s.Ingredients
	model.UsageMethod = s.UsageMethod
	model.Price = s.Price
	model.MallProductId = s.MallProductId
}

func (s *WechatProductsInsertReq) GetId() interface{} {
	return s.Id
}

type WechatProductsUpdateReq struct {
	Id            int    `uri:"id" comment:"自增主键"` // 自增主键
	ProductName   string `json:"productName" comment:"商品名称"`
	ImageUrl      string `json:"imageUrl" comment:"商品图片 URL"`
	Description   string `json:"description" comment:"产品介绍"`
	Ingredients   string `json:"ingredients" comment:"产品成分"`
	UsageMethod   string `json:"usageMethod" comment:"使用方法"`
	Price         string `json:"price" comment:"商品价格（单位：元）"`
	MallProductId string `json:"mallProductId" comment:"商城商品编号"`
	common.ControlBy
}

func (s *WechatProductsUpdateReq) Generate(model *models.WechatProducts) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.ProductName = s.ProductName
	model.ImageUrl = s.ImageUrl
	model.Description = s.Description
	model.Ingredients = s.Ingredients
	model.UsageMethod = s.UsageMethod
	model.Price = s.Price
	model.MallProductId = s.MallProductId
}

func (s *WechatProductsUpdateReq) GetId() interface{} {
	return s.Id
}

// WechatProductsGetReq 功能获取请求参数
type WechatProductsGetReq struct {
	Id int `uri:"id"`
}

func (s *WechatProductsGetReq) GetId() interface{} {
	return s.Id
}

// WechatProductsDeleteReq 功能删除请求参数
type WechatProductsDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *WechatProductsDeleteReq) GetId() interface{} {
	return s.Ids
}
