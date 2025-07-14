package router

import (
	"github.com/gin-gonic/gin"
	"go-admin/app/wechat/apis"
)

// 在 init 中追加到无需鉴权的列表（如 OSS 上传签名接口）
func init() {
	routerNoCheckRole = append(routerNoCheckRole, registerOSSRouter)
}

// registerOSSRouter 将 OSS 签名接口注册到无需鉴权的路由组
func registerOSSRouter(v1 *gin.RouterGroup) {
	// 二级路由前缀，例如 /api/v1/oss
	r := v1.Group("/oss")
	{
		r.GET("/signature", apis.GetOSSOptions) // GET /api/v1/oss/signature
	}
}
