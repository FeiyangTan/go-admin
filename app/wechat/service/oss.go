package service

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"time"

	"go-admin/app/wechat/config"
)

// PolicySignature 包含前端上传所需的签名字段
type PolicySignature struct {
	AccessID  string `json:"accessid"`
	Policy    string `json:"policy"`
	Signature string `json:"signature"`
	Host      string `json:"host"`
	Dir       string `json:"dir"`
	Expire    int64  `json:"expire"`
}

// GeneratePolicySignature 生成 OSS 表单直传的签名数据
func GeneratePolicySignature(file string) (*PolicySignature, error) {
	cfg := config.GetOSSConfig()

	// 构造上传前缀，按日期分目录
	dir := cfg.DirPrefix + file

	// 计算过期时间
	expireTime := time.Now().Add(time.Duration(cfg.ExpireSeconds) * time.Second).UTC()
	expireISO := expireTime.Format("2006-01-02T15:04:05Z")

	// 构造 policy JSON
	policyMap := map[string]interface{}{
		"expiration": expireISO,
		"conditions": []interface{}{[]interface{}{"starts-with", "$key", dir}},
	}
	policyBytes, err := json.Marshal(policyMap)
	if err != nil {
		return nil, err
	}

	// Base64 编码 policy
	policyBase64 := base64.StdEncoding.EncodeToString(policyBytes)

	// HMAC-SHA1 签名
	mac := hmac.New(sha1.New, []byte(cfg.AccessKeySecret))
	mac.Write([]byte(policyBase64))
	signature := base64.StdEncoding.EncodeToString(mac.Sum(nil))

	// 构造 host
	host := "https://" + cfg.Bucket + "." + cfg.Endpoint

	return &PolicySignature{
		AccessID:  cfg.AccessKeyID,
		Policy:    policyBase64,
		Signature: signature,
		Host:      host,
		Dir:       dir,
		Expire:    expireTime.Unix(),
	}, nil
}
