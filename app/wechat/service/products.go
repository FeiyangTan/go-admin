// app/wechat/service/user.go
package service

import (
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"go-admin/app/wechat/models"
	"strings"
)

type WechatProductsService struct {
	api.Api
}

// NewWechatPhysiqueService 会在 Handler 里被 MakeService 调用
func NewWechatProductsService(e *api.Api) *WechatProductsService {
	return &WechatProductsService{*e}
}

// GetOrCreateUser 自动使用 e.Orm
func (s *WechatProductsService) GetProductInfo(mallProductID string) (map[string]interface{}, error) {
	db, _ := s.GetOrm() // 从 s.Api 已初始化的 Orm 拿到 *gorm.DB

	var product models.Product
	if err := db.Where("mall_product_id = ?", mallProductID).First(&product).Error; err != nil {
		return nil, err
	}

	// 组装 map
	item := map[string]interface{}{
		"id":            product.ID,
		"name":          product.ProductName,
		"price":         product.Price,
		"imageSrc":      product.ImageURL,
		"description":   product.Description,
		"ingredient":    product.Ingredients,
		"useMethod":     product.UsageMethod,
		"mallProductID": product.MallProductID,
	}

	return item, nil
}

func (s *WechatProductsService) GetProductsInfo(productIDs string) []map[string]interface{} {
	productIDs = strings.ReplaceAll(productIDs, "，", ",")
	ids := strings.Split(productIDs, ",")

	var result []map[string]interface{}

	for _, idStr := range ids {
		// 直接用 string 调用 GetProductsInfo
		product, err := s.GetProductInfo(idStr)
		if err != nil || product == nil {
			continue
		}
		// 放入结果数组
		result = append(result, product)
	}

	return result
}
