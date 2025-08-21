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

type WechatPhysique struct {
	api.Api
}

// GetPage 获取体质表列表
// @Summary 获取体质表列表
// @Description 获取体质表列表
// @Tags 体质表
// @Param physiqueName query string false "体质名称"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.WechatPhysique}} "{"code": 200, "data": [...]}"
// @Router /api/v1/wechat-physique [get]
// @Security Bearer
func (e WechatPhysique) GetPage(c *gin.Context) {
	req := dto.WechatPhysiqueGetPageReq{}
	s := service.WechatPhysique{}
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
	list := make([]models.WechatPhysique, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取体质表失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取体质表
// @Summary 获取体质表
// @Description 获取体质表
// @Tags 体质表
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.WechatPhysique} "{"code": 200, "data": [...]}"
// @Router /api/v1/wechat-physique/{id} [get]
// @Security Bearer
func (e WechatPhysique) Get(c *gin.Context) {
	req := dto.WechatPhysiqueGetReq{}
	s := service.WechatPhysique{}
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
	var object models.WechatPhysique

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取体质表失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(object, "查询成功")
}

// Insert 创建体质表
// @Summary 创建体质表
// @Description 创建体质表
// @Tags 体质表
// @Accept application/json
// @Product application/json
// @Param data body dto.WechatPhysiqueInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/wechat-physique [post]
// @Security Bearer
func (e WechatPhysique) Insert(c *gin.Context) {
	req := dto.WechatPhysiqueInsertReq{}
	s := service.WechatPhysique{}
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
		e.Error(500, err, fmt.Sprintf("创建体质表失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改体质表
// @Summary 修改体质表
// @Description 修改体质表
// @Tags 体质表
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.WechatPhysiqueUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/wechat-physique/{id} [put]
// @Security Bearer
func (e WechatPhysique) Update(c *gin.Context) {
	req := dto.WechatPhysiqueUpdateReq{}
	s := service.WechatPhysique{}
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
		e.Error(500, err, fmt.Sprintf("修改体质表失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "修改成功")
}

// Delete 删除体质表
// @Summary 删除体质表
// @Description 删除体质表
// @Tags 体质表
// @Param data body dto.WechatPhysiqueDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/wechat-physique [delete]
// @Security Bearer
func (e WechatPhysique) Delete(c *gin.Context) {
	s := service.WechatPhysique{}
	req := dto.WechatPhysiqueDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除体质表失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "删除成功")
}
