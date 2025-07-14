package router

import (
	"github.com/gin-gonic/gin"
	"go-admin/app/wechat/apis"
)

// 在 init 中追加到无需鉴权的列表（如登录、注册）
func init() {
	routerNoCheckRole = append(routerNoCheckRole, registerDiagnosisRouter)
}

func registerDiagnosisRouter(v1 *gin.RouterGroup) {
	api := apis.DiagnosisAPI{}
	// 这里设定你想要的二级路由前缀
	r := v1.Group("/wechat")
	{
		r.POST("/addDiagnosis", api.AddDiagnosis)         // POST /api/v1/wechat/addDiagnosis
		r.POST("/userDiagnosisNum", api.UserDiagnosisNum) // POST /api/v1/wechat/UserDiagnosisNum
		r.POST("/diagnosisList", api.DiagnosisList)       // POST /api/v1/wechat/UserDiagnosisNum
	}
}
