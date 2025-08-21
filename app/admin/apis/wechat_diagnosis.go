package apis

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	_ "github.com/go-admin-team/go-admin-core/sdk/pkg/response"

	"go-admin/app/admin/models"
	"go-admin/app/admin/service"
	"go-admin/app/admin/service/dto"
	"go-admin/common/actions"
)

type WechatDiagnosis struct {
	api.Api
}

// GetPage 获取诊疗结果列表
// @Summary 获取诊疗结果列表
// @Description 获取诊疗结果列表
// @Tags 诊疗结果
// @Param openId query string false "微信 OpenID"
// @Param diagnosisType query string false "诊疗类型"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.WechatDiagnosis}} "{"code": 200, "data": [...]}"
// @Router /api/v1/wechat-diagnosis [get]
// @Security Bearer
func (e WechatDiagnosis) GetPage(c *gin.Context) {
	req := dto.WechatDiagnosisGetPageReq{}
	s := service.WechatDiagnosis{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	p := actions.GetPermissionFromContext(c)
	list := make([]models.WechatDiagnosis, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取诊疗结果失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取诊疗结果
// @Summary 获取诊疗结果
// @Description 获取诊疗结果
// @Tags 诊疗结果
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.WechatDiagnosis} "{"code": 200, "data": [...]}"
// @Router /api/v1/wechat-diagnosis/{id} [get]
// @Security Bearer
func (e WechatDiagnosis) Get(c *gin.Context) {
	req := dto.WechatDiagnosisGetReq{}
	s := service.WechatDiagnosis{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	var object models.WechatDiagnosis

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取诊疗结果失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(object, "查询成功")
}

// Insert 创建诊疗结果
// @Summary 创建诊疗结果
// @Description 创建诊疗结果
// @Tags 诊疗结果
// @Accept application/json
// @Product application/json
// @Param data body dto.WechatDiagnosisInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/wechat-diagnosis [post]
// @Security Bearer
func (e WechatDiagnosis) Insert(c *gin.Context) {
	req := dto.WechatDiagnosisInsertReq{}
	s := service.WechatDiagnosis{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	// 设置创建人
	req.SetCreateBy(user.GetUserId(c))

	err = s.Insert(&req)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("创建诊疗结果失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改诊疗结果
// @Summary 修改诊疗结果
// @Description 修改诊疗结果
// @Tags 诊疗结果
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.WechatDiagnosisUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/wechat-diagnosis/{id} [put]
// @Security Bearer
func (e WechatDiagnosis) Update(c *gin.Context) {
	req := dto.WechatDiagnosisUpdateReq{}
	s := service.WechatDiagnosis{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	req.SetUpdateBy(user.GetUserId(c))
	p := actions.GetPermissionFromContext(c)

	err = s.Update(&req, p)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("修改诊疗结果失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "修改成功")
}

// Delete 删除诊疗结果
// @Summary 删除诊疗结果
// @Description 删除诊疗结果
// @Tags 诊疗结果
// @Param data body dto.WechatDiagnosisDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/wechat-diagnosis [delete]
// @Security Bearer
func (e WechatDiagnosis) Delete(c *gin.Context) {
	s := service.WechatDiagnosis{}
	req := dto.WechatDiagnosisDeleteReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	// req.SetUpdateBy(user.GetUserId(c))
	p := actions.GetPermissionFromContext(c)

	err = s.Remove(&req, p)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("删除诊疗结果失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "删除成功")
}
