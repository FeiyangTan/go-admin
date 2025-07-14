package apis

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go-admin/app/wechat/service"
)

// OSSOptionsResponse 定义返回给前端的数据结构
type OSSOptionsResponse struct {
	AccessID  string `json:"accessid"`
	Policy    string `json:"policy"`
	Signature string `json:"signature"`
	Host      string `json:"host"`
	Dir       string `json:"dir"`
	Expire    int64  `json:"expire"`
}

// GetOSSOptions 小程序调用此接口获取上传签名
func GetOSSOptions(c *gin.Context) {
	sig, err := service.GeneratePolicySignature()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// 转换为前端定义的结构体（可选）
	resp := OSSOptionsResponse{
		AccessID:  sig.AccessID,
		Policy:    sig.Policy,
		Signature: sig.Signature,
		Host:      sig.Host,
		Dir:       sig.Dir,
		Expire:    sig.Expire,
	}
	c.JSON(http.StatusOK, resp)
}
