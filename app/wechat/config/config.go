// config/config.go
package config

import "os"

var (
	// 微信小程序的 AppID 和 AppSecret
	AppID     = os.Getenv("WEAPP_APPID")
	AppSecret = os.Getenv("WEAPP_APPSECRET")

	// 用于签发和校验 JWT 的密钥，建议至少 32 字节随机字符串
	JwtSecret = os.Getenv("Jwt_SECRET")

	// 新用户的默认昵称和头像
	DefaultNickName  = "艾小蕲用户"
	DefaultAvatarURL = "https://ai-qiai.oss-cn-guangzhou.aliyuncs.com/panda.png"

	AIAPPCODE = os.Getenv("AI_APPCODE")
)

// OSSConfig holds configuration for Alibaba Cloud OSS
type OSSConfig struct {
	AccessKeyID     string
	AccessKeySecret string
	Bucket          string
	Endpoint        string
	DirPrefix       string
	ExpireSeconds   int64
}

// OSS contains the OSS configuration values
var OSS = OSSConfig{
	AccessKeyID:     os.Getenv("OSS_ACCESS_KEY_ID"),
	AccessKeySecret: os.Getenv("OSS_ACCESS_KEY_SECRET"),
	Bucket:          "ai-qiai",
	Endpoint:        "oss-cn-guangzhou.aliyuncs.com",
	DirPrefix:       "wechat/", // 可选，前缀目录
	ExpireSeconds:   30,        // 签名过期秒数
}

// GetOSSConfig returns the OSS configuration
func GetOSSConfig() OSSConfig {
	return OSS
}
