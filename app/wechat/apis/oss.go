package apis

import (
	"fmt"
	"net/http"
	"time"

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

// GetOSSImageSignature 小程序调用此接口获取照片上传签名
func GetOSSImageSignature(c *gin.Context) {
	file := time.Now().Format("2006-01-02") + "/"
	sig, err := service.GeneratePolicySignature(file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(sig)
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

// GetOSSImageSignature 小程序调用此接口获取头像上传签名
func GetOSSAvatarSignature(c *gin.Context) {
	sig, err := service.GeneratePolicySignature("avatar/")
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
