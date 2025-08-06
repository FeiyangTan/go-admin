// app/wechat/service/user.go
package service

import (
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"go-admin/app/wechat/models"
	"gorm.io/gorm"
)

type WechatUserService struct {
	api.Api
}

// NewWechatUserService 会在 Handler 里被 MakeService 调用
func NewWechatUserService(e *api.Api) *WechatUserService {
	return &WechatUserService{*e}
}

// GetOrCreateUser 自动使用 e.Orm
func (s *WechatUserService) GetOrCreateUser(openid, nickname, avatar string) (*models.User, error) {
	db, _ := s.GetOrm() // 从 s.Api 已初始化的 Orm 拿到 *gorm.DB
	var user models.User
	if err := db.Where("open_id = ?", openid).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			user = models.User{OpenID: openid, NickName: nickname, AvatarURL: avatar}
			if err := db.Create(&user).Error; err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}
	return &user, nil
}

// Diagnosis数量+1
func (s *WechatUserService) AddDiagnosisCount(openid string, diagnosisType string) error {
	db, _ := s.GetOrm() // 从 s.Api 已初始化的 Orm 拿到 *gorm.DB

	switch diagnosisType {
	case "face":
		if err := db.
			Model(&models.User{}).
			Where("open_id = ?", openid).
			UpdateColumn("face_count", gorm.Expr("face_count + ?", 1)).
			Error; err != nil {
			return err
		}
	default:
		if err := db.
			Model(&models.User{}).
			Where("open_id = ?", openid).
			UpdateColumn("tongue_count", gorm.Expr("tongue_count + ?", 1)).
			Error; err != nil {
			return err
		}
	}

	return nil
}

// 查询用户Diagnosis数量
func (s *WechatUserService) GetUserDiagnosisNum(openid string) (faceCount uint16, tongueCount uint16, err error) {
	db, _ := s.GetOrm() // 从 s.Api 已初始化的 Orm 拿到 *gorm.DB

	var user models.User
	// 只查询 face_count 和 tongue_count 两列
	if err = db.
		Model(&models.User{}).
		Select("face_count", "tongue_count").
		Where("open_id = ?", openid).
		First(&user).Error; err != nil {
		return 0, 0, err
	}

	// 从查询结果中取出计数
	faceCount = user.FaceCount
	tongueCount = user.TongueCount
	return faceCount, tongueCount, nil
}

// 修改用户名称和头像
func (s *WechatUserService) SetUserInfo(openid, nickname, avatarUrl string) error {
	db, _ := s.GetOrm() // 从 s.Api 已初始化的 Orm 拿到 *gorm.DB
	// 修改用户名称和头像
	if err := db.
		Model(&models.User{}).
		Where("open_id = ?", openid).
		Updates(map[string]interface{}{
			"nick_name":  nickname,
			"avatar_url": avatarUrl,
		}).
		Error; err != nil {
		return err
	}
	return nil
}
