// app/wechat/service/user.go
package service

import (
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"go-admin/app/wechat/models"
	"gorm.io/gorm"
)

type WechatUserService1 struct {
	api.Api
}

// NewWechatUserService 会在 Handler 里被 MakeService 调用
func NewWechatUserService(e *api.Api) *WechatUserService1 {
	return &WechatUserService1{*e}
}

// GetOrCreateUser 自动使用 e.Orm
func (s *WechatUserService1) GetOrCreateUser(openid, nickname, avatar string) (*models.User, error) {
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
func (s *WechatUserService1) AddDiagnosisCount(openid string, diagnosisType string) error {
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
func (s *WechatUserService1) GetUserDiagnosisNum(openid string) (faceCount uint16, tongueCount uint16, err error) {
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

//修改用户名称和头像

func (s *WechatUserService1) SetUserInfo(openid, nickname, avatarUrl string) error {
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

//type WechatUserService struct {
//	sdkservice.Service
//}
//
//type WechatUser struct {
//	sdkservice.Service
//	wc clients.WechatClient
//}
//
//func NewWechatUser(svc *sdkservice.Service) *WechatUser {
//	r := &WechatUser{wc: clients.NewWechatClient()}
//	if svc != nil {
//		r.Service = *svc
//	}
//	return r
//}
//
//// 登录：微信 code -> openid -> GetOrCreate -> 生成 JWT -> 返回
//func (s *WechatUser) Login(ctx context.Context, in *dto.LoginReq) (*dto.LoginResp, error) {
//	wx, err := s.wc.Jscode2Session(ctx, config.AppID, config.AppSecret, in.Code)
//	if err != nil {
//		return nil, err
//	}
//
//	// ✅ 复用当前上下文：直接字面量构造，不再调用旧的 NewWechatUserService(...)
//	us := &WechatUserService{Service: s.Service}
//
//	user, err := us.GetOrCreateUser(wx.OpenID, config.DefaultNickName, config.DefaultAvatarURL)
//	if err != nil {
//		return nil, err
//	}
//
//	token, err := util.GenerateToken(user.OpenID)
//	if err != nil {
//		return nil, err
//	}
//
//	out := &dto.LoginResp{Token: token}
//	out.User.OpenID = user.OpenID
//	out.User.NickName = user.NickName
//	out.User.AvatarURL = user.AvatarURL
//	return out, nil
//}
//
//// GetOrCreateUser 根据 open_id 查找用户，不存在则按给定昵称与头像创建
//func (s *WechatUserService) GetOrCreateUser(openid, nickname, avatarURL string) (models.User, error) {
//	var user models.User
//	db := s.Orm
//
//	// open_id 建议在 models.User 上有唯一索引
//	err := db.Where("open_id = ?", openid).First(&user).Error
//	if errors.Is(err, gorm.ErrRecordNotFound) {
//		user = models.User{
//			OpenID:    openid,
//			NickName:  nickname,
//			AvatarURL: avatarURL,
//		}
//		if err := db.Create(&user).Error; err != nil {
//			return models.User{}, err
//		}
//		return user, nil
//	}
//	if err != nil {
//		return models.User{}, err
//	}
//	return user, nil
//}
//
//// SetUserInfo 更新昵称与头像
//func (s *WechatUserService) SetUserInfo(openid, nickname, avatarURL string) error {
//	db := s.Orm
//	return db.
//		Model(&models.User{}).
//		Where("open_id = ?", openid).
//		Updates(map[string]any{
//			"nick_name":  nickname,
//			"avatar_url": avatarURL,
//		}).Error
//}
