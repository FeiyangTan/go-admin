package dto

// Login 接口前端传来的登录请求参数
type LoginReq struct {
	Code string `json:"code" binding:"required"`
}

// WxSessionResp 用于接收微信 jscode2session 接口的返回
type LoginResp struct {
	Token string `json:"token"`
	User  struct {
		OpenID    string `json:"openId"`
		NickName  string `json:"nickname"`
		AvatarURL string `json:"avatarUrl"`
	} `json:"user"`
}

// setUserInfo 接口前端传来的登录请求参数
type SetUserInfoReq struct {
	OpenID    string `json:"open_id"  binding:"required"`
	NickName  string `json:"nick_name" binding:"required"`
	AvatarURL string `json:"avatar_url" binding:"required"`
}
