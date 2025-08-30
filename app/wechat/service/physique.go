// app/wechat/service/user.go
package service

import (
	"errors"
	"fmt"
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
		// 没找到记录，则把新的体质名字添加到数据库中
		if err == gorm.ErrRecordNotFound {
			p := &models.Physique{PhysiqueName: physiqueName}
			cre, err := s.CreatePhysique(p)
			fmt.Println("~~~~~~新的体质添加到数据库", cre)
			if err != nil {
				return nil, err
			}
			return nil, nil
		}
		// 其他错误
		return nil, err
	}

	// 如果找到的特性，没有数据，使用“默认特征”
	if physique.AcupunctureMethod == "" && physiqueName != "默认体质" {
		fmt.Println("~~~~~~默认体质")
		physiqueInfo, err := s.GetPhysiqueInfo("默认体质")
		if err != nil {
			return nil, errors.New("获取默认体质信息错误: " + err.Error())
		}
		return physiqueInfo, nil
	}
	return &physique, nil
}

// CreatePhysique 直接创建一条新纪录
func (s *WechatPhysiqueService) CreatePhysique(p *models.Physique) (*models.Physique, error) {
	db, _ := s.GetOrm()
	if err := db.Create(p).Error; err != nil {
		return nil, err
	}
	return p, nil
}
