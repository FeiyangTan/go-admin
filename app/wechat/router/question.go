package router

import (
	"github.com/gin-gonic/gin"
	"go-admin/app/wechat/apis"
)

// 在 init 中追加到无需鉴权的列表（如登录、注册）
func init() {
	routerNoCheckRole = append(routerNoCheckRole, registerQuestionRouter)
}

func registerQuestionRouter(v1 *gin.RouterGroup) {
	api := apis.QuestionAPI{}
	// 这里设定你想要的二级路由前缀
	r := v1.Group("/wechat")
	{
		r.GET("/question/get", api.Question_get) // POST /api/v1/wechat/question/get
	}
}
