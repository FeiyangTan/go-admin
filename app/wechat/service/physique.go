// app/wechat/service/user.go
package service

import (
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"go-admin/app/wechat/models"
	"gorm.io/gorm"
)

type WechatPhysiqueService struct {
	api.Api
}

// NewWechatPhysiqueService 会在 Handler 里被 MakeService 调用
func NewWechatPhysiqueService(e *api.Api) *WechatPhysiqueService {
	return &WechatPhysiqueService{*e}
}

// GetOrCreateUser 自动使用 e.Orm
func (s *WechatPhysiqueService) GetPhysiqueInfo(physiqueName string) (*models.Physique, error) {
	db, _ := s.GetOrm() // 从 s.Api 已初始化的 Orm 拿到 *gorm.DB
	var physique models.Physique
	if err := db.Where("physique_name = ?", physiqueName).First(&physique).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			// 没找到记录，返回 nil
			return nil, nil
		}
		// 其他错误
		return nil, err
	}

	return &physique, nil
}
