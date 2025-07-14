package apis

import (
	"encoding/json"
	"fmt"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"go-admin/app/wechat/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"go-admin/app/wechat/config"
	"go-admin/app/wechat/util"
)

// WxSessionResp 用于接收微信 jscode2session 接口的返回
type WxSessionResp struct {
	OpenID     string `json:"openid"`
	SessionKey string `json:"session_key"`
	ErrCode    int    `json:"errcode"`
	ErrMsg     string `json:"errmsg"`
}

// LoginReq 是前端传来的登录请求参数
type LoginReq struct {
	Code          string `json:"code" binding:"required"`
	EncryptedData string `json:"encryptedData" binding:"required"`
	Iv            string `json:"iv" binding:"required"`
	RawData       string `json:"rawData"`
	Signature     string `json:"signature"`
}

type UserAPI struct {
	api.Api
}

// POST /api/v1/wechat/login
func (u *UserAPI) Login(c *gin.Context) {

	// 0. 初始化 Api 上下文和 ORM
	if err := u.MakeContext(c).MakeOrm().Errors; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "初始化失败: " + err.Error()})
		return
	}

	// 1. 用 code 换取 session_key 和 openid
	var req LoginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "参数错误: " + err.Error()})
		return
	}
	wxURL := fmt.Sprintf(
		"https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code",
		config.AppID, config.AppSecret, req.Code,
	)
	resp, err := http.Get(wxURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "调用微信接口失败: " + err.Error()})
		return
	}
	defer resp.Body.Close()

	var wxResp WxSessionResp
	if err := json.NewDecoder(resp.Body).Decode(&wxResp); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "解析微信返回数据失败: " + err.Error()})
		return
	}
	if wxResp.ErrCode != 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "微信接口错误: " + wxResp.ErrMsg})
		return
	}

	fmt.Println(wxResp)
	// 2. 解密用户信息，以获取昵称、头像等公开信息
	decrypted, err := util.DecryptWeChatData(wxResp.SessionKey, req.EncryptedData, req.Iv)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "解密用户数据失败: " + err.Error()})
		return
	}

	var profile struct {
		NickName  string `json:"nickName"`
		AvatarURL string `json:"avatarUrl"`
	}
	if err := json.Unmarshal(decrypted, &profile); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "解析用户数据失败: " + err.Error()})
		return
	}
	fmt.Println(profile)

	// 3. 数据库：获取或创建用户，使用 jscode2session 返回的 openid
	//db := c.MustGet("DB").(*gorm.DB)
	//user, err := models.GetOrCreateUser(db, wxResp.OpenID, profile.NickName, profile.AvatarURL)
	s := service.NewWechatUserService(&u.Api)
	fmt.Println(s)

	user, err := s.GetOrCreateUser(wxResp.OpenID, profile.NickName, profile.AvatarURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "数据库操作失败: " + err.Error()})
		return
	}

	fmt.Println(user)

	// 4. 生成 JWT，返回结果
	token, err := util.GenerateToken(user.OpenID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "生成 Token 失败: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
		"user": gin.H{
			"openId":    user.OpenID,
			"nickname":  user.NickName,
			"avatarUrl": user.AvatarURL,
		},
	})

}

// POST /api/v1/wechat/signup
func (u *UserAPI) Signup(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 2,
	})
}

// GET /api/v1/wechat/profile
func (u *UserAPI) GetProfile(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 3,
	})
}
