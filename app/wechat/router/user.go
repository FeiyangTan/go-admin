package router

import (
	"github.com/gin-gonic/gin"
	"go-admin/app/wechat/apis"
)

// 在 init 中追加到无需鉴权的列表（如登录、注册）
func init() {
	routerNoCheckRole = append(routerNoCheckRole, registerUserRouter)
}

func registerUserRouter(v1 *gin.RouterGroup) {
	api := apis.UserAPI{}
	// 这里设定你想要的二级路由前缀
	r := v1.Group("/wechat")
	{
		r.POST("/login", api.Login)       // POST /api/v1/wechat/login
		r.POST("/signup", api.Signup)     // POST /api/v1/wechat/signup
		r.GET("/profile", api.GetProfile) // GET  /api/v1/wechat/profile
		// ... 按需添加更多
	}
}
