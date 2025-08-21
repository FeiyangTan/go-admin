// app/wechat/service/user.go
package service

import (
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"go-admin/app/wechat/models"
)

type WechatProductsService struct {
	api.Api
}

// NewWechatPhysiqueService 会在 Handler 里被 MakeService 调用
func NewWechatProductsService(e *api.Api) *WechatProductsService {
	return &WechatProductsService{*e}
}

// GetOrCreateUser 自动使用 e.Orm
func (s *WechatProductsService) GetProductsInfo(mallProductID string) (*models.Product, error) {
	db, _ := s.GetOrm() // 从 s.Api 已初始化的 Orm 拿到 *gorm.DB

	var product models.Product
	if err := db.Where("mall_product_id = ?", mallProductID).First(&product).Error; err != nil {
		return nil, err
	}
	return &product, nil
}
