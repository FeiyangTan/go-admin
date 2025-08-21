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

type WechatUsers struct {
	api.Api
}

// GetPage 获取用户表列表
// @Summary 获取用户表列表
// @Description 获取用户表列表
// @Tags 用户表
// @Param openId query string false "微信 OpenID"
// @Param nickName query string false "昵称"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.WechatUsers}} "{"code": 200, "data": [...]}"
// @Router /api/v1/wechat-users [get]
// @Security Bearer
func (e WechatUsers) GetPage(c *gin.Context) {
	req := dto.WechatUsersGetPageReq{}
	s := service.WechatUsers{}
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
	list := make([]models.WechatUsers, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取用户表失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取用户表
// @Summary 获取用户表
// @Description 获取用户表
// @Tags 用户表
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.WechatUsers} "{"code": 200, "data": [...]}"
// @Router /api/v1/wechat-users/{id} [get]
// @Security Bearer
func (e WechatUsers) Get(c *gin.Context) {
	req := dto.WechatUsersGetReq{}
	s := service.WechatUsers{}
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
	var object models.WechatUsers

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取用户表失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(object, "查询成功")
}

// Insert 创建用户表
// @Summary 创建用户表
// @Description 创建用户表
// @Tags 用户表
// @Accept application/json
// @Product application/json
// @Param data body dto.WechatUsersInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/wechat-users [post]
// @Security Bearer
func (e WechatUsers) Insert(c *gin.Context) {
	req := dto.WechatUsersInsertReq{}
	s := service.WechatUsers{}
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
		e.Error(500, err, fmt.Sprintf("创建用户表失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改用户表
// @Summary 修改用户表
// @Description 修改用户表
// @Tags 用户表
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.WechatUsersUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/wechat-users/{id} [put]
// @Security Bearer
func (e WechatUsers) Update(c *gin.Context) {
	req := dto.WechatUsersUpdateReq{}
	s := service.WechatUsers{}
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
		e.Error(500, err, fmt.Sprintf("修改用户表失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "修改成功")
}

// Delete 删除用户表
// @Summary 删除用户表
// @Description 删除用户表
// @Tags 用户表
// @Param data body dto.WechatUsersDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/wechat-users [delete]
// @Security Bearer
func (e WechatUsers) Delete(c *gin.Context) {
	s := service.WechatUsers{}
	req := dto.WechatUsersDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除用户表失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "删除成功")
}
