package apis

import (
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"go-admin/app/wechat/service"
	"gorm.io/datatypes"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DiagnosisAPI struct {
	api.Api
}

type AddDiagnosisReq struct {
	OpenID          string         `json:"open_id" binding:"required"`
	DiagnosisType   string         `json:"diagnosis_type" binding:"required"`
	DiagnosisResult datatypes.JSON `json:"diagnosis_result" binding:"required"`
}

type GetUserDiagnosisNumReq struct {
	OpenID string `json:"open_id" binding:"required"`
}

type DiagnosisListReq struct {
	OpenID        string `json:"open_id" binding:"required"`
	DiagnosisType string `json:"diagnosis_type" binding:"required"`
}

// POST /api/v1/wechat/addDiagnosis
func (d *DiagnosisAPI) AddDiagnosis(c *gin.Context) {

	// 1. 初始化 Api 上下文和 ORM
	if err := d.MakeContext(c).MakeOrm().Errors; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "初始化失败: " + err.Error()})
		return
	}

	// 2. 获取响应数据
	var req AddDiagnosisReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "参数错误: " + err.Error()})
		return
	}

	// 3. 数据库：
	// 用户数据库，diagnosis数量+1
	s := service.NewWechatUserService(&d.Api)
	if err := s.AddDiagnosisCount(req.OpenID, req.DiagnosisType); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "添加诊疗数量错误: " + err.Error()})
		return
	}

	// 诊疗数据库，记录diagnosis
	s2 := service.NewWechatDiagnosisService(&d.Api)
	if err := s2.AddDiagnosis(req.OpenID, req.DiagnosisType, req.DiagnosisResult); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "添加诊疗记录错误: " + err.Error()})
		return
	}

	// 4. 返回结果
	c.JSON(http.StatusOK, gin.H{
		"state": "success",
	})
}

// POST /api/v1/wechat/UserDiagnosisNum
func (d *DiagnosisAPI) UserDiagnosisNum(c *gin.Context) {

	// 1. 初始化 Api 上下文和 ORM
	if err := d.MakeContext(c).MakeOrm().Errors; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "初始化失败: " + err.Error()})
		return
	}

	// 2. 获取响应数据
	var req GetUserDiagnosisNumReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "参数错误: " + err.Error()})
		return
	}

	// 3. 数据库：
	// 查询用户表，诊疗数量
	s := service.NewWechatUserService(&d.Api)

	faceCount, tongueCount, err := s.GetUserDiagnosisNum(req.OpenID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "查询用户诊疗数量错误: " + err.Error()})
		return
	}

	// 4. 返回结果
	c.JSON(http.StatusOK, gin.H{
		"faceCount":   faceCount,
		"tongueCount": tongueCount,
	})
}

// POST /api/v1/wechat/DiagnosisList
func (d *DiagnosisAPI) DiagnosisList(c *gin.Context) {

	// 1. 初始化 Api 上下文和 ORM
	if err := d.MakeContext(c).MakeOrm().Errors; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "初始化失败: " + err.Error()})
		return
	}

	// 2. 获取响应数据
	var req DiagnosisListReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "参数错误: " + err.Error()})
		return
	}

	// 3. 数据库：
	// 获取列表
	s := service.NewWechatDiagnosisService(&d.Api)

	diagnosisList, err := s.GetDiagnosisList(req.OpenID, req.DiagnosisType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "查询用户诊疗列表错误: " + err.Error()})
		return
	}

	// 4. 返回结果
	c.JSON(http.StatusOK, gin.H{
		"diagnosisList": diagnosisList,
	})
}
