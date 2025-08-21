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

type WechatProducts struct {
	api.Api
}

// GetPage 获取商城商品表列表
// @Summary 获取商城商品表列表
// @Description 获取商城商品表列表
// @Tags 商城商品表
// @Param productName query string false "商品名称"
// @Param mallProductId query string false "商城商品编号"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.WechatProducts}} "{"code": 200, "data": [...]}"
// @Router /api/v1/wechat-products [get]
// @Security Bearer
func (e WechatProducts) GetPage(c *gin.Context) {
	req := dto.WechatProductsGetPageReq{}
	s := service.WechatProducts{}
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
	list := make([]models.WechatProducts, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取商城商品表失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取商城商品表
// @Summary 获取商城商品表
// @Description 获取商城商品表
// @Tags 商城商品表
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.WechatProducts} "{"code": 200, "data": [...]}"
// @Router /api/v1/wechat-products/{id} [get]
// @Security Bearer
func (e WechatProducts) Get(c *gin.Context) {
	req := dto.WechatProductsGetReq{}
	s := service.WechatProducts{}
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
	var object models.WechatProducts

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取商城商品表失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(object, "查询成功")
}

// Insert 创建商城商品表
// @Summary 创建商城商品表
// @Description 创建商城商品表
// @Tags 商城商品表
// @Accept application/json
// @Product application/json
// @Param data body dto.WechatProductsInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/wechat-products [post]
// @Security Bearer
func (e WechatProducts) Insert(c *gin.Context) {
	req := dto.WechatProductsInsertReq{}
	s := service.WechatProducts{}
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
		e.Error(500, err, fmt.Sprintf("创建商城商品表失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改商城商品表
// @Summary 修改商城商品表
// @Description 修改商城商品表
// @Tags 商城商品表
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.WechatProductsUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/wechat-products/{id} [put]
// @Security Bearer
func (e WechatProducts) Update(c *gin.Context) {
	req := dto.WechatProductsUpdateReq{}
	s := service.WechatProducts{}
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
		e.Error(500, err, fmt.Sprintf("修改商城商品表失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "修改成功")
}

// Delete 删除商城商品表
// @Summary 删除商城商品表
// @Description 删除商城商品表
// @Tags 商城商品表
// @Param data body dto.WechatProductsDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/wechat-products [delete]
// @Security Bearer
func (e WechatProducts) Delete(c *gin.Context) {
	s := service.WechatProducts{}
	req := dto.WechatProductsDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除商城商品表失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "删除成功")
}
