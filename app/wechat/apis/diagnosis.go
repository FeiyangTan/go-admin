package apis

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"go-admin/app/wechat/client"
	"go-admin/app/wechat/service"
	"gorm.io/datatypes"
	"net/http"
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
	OpenID string `form:"open_id" binding:"required"`
}

type DiagnosisListReq struct {
	OpenID        string `form:"open_id" binding:"required"`
	DiagnosisType string `form:"diagnosis_type" binding:"required"`
	PageSize      int    `form:"page_size" binding:"required"`
	PageIndex     int    `form:"page_index" binding:"required"`
}

// 对应 thirdApiData
type ThirdAPIData struct {
	Scene   int    `json:"scene" binding:"required"`
	FFImage string `json:"ff_image"`
	TFImage string `json:"tf_image"`
	TBImage string `json:"tb_image"`
	Gender  string `json:"gender"`
}

// 整体请求体
type DiagnosisReq struct {
	ThirdAPIData  ThirdAPIData `json:"thirdApiData" binding:"required"`
	OpenID        string       `json:"open_id" binding:"required"`
	DiagnosisType string       `json:"diagnosis_type" binding:"required"`
}

// POST /api/v1/wechat/addDiagnosis
func (e *DiagnosisAPI) AddDiagnosis(c *gin.Context) {

	// 1. 初始化 Api 上下文和 ORM
	if err := e.MakeContext(c).MakeOrm().Errors; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "初始化失败: " + err.Error()})
		return
	}

	// 2. 获取响应数据
	var req AddDiagnosisReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "参数错误: " + err.Error()})
		return
	}
	// 2.1 获取响应数据中的检测结果
	var resMap map[string]interface{}
	if err := json.Unmarshal(req.DiagnosisResult, &resMap); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "result 解析失败: " + err.Error()})
		return
	}

	// 3 存入数据库：
	openID := req.OpenID
	diagnosisType := req.DiagnosisType

	fmt.Println("~~~诊疗类型：", diagnosisType)
	// 3.1 用户数据库，diagnosis数量+1
	s1 := service.NewWechatUserService(&e.Api)
	if err := s1.AddDiagnosisCount(openID, diagnosisType); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "添加诊疗数量错误: " + err.Error()})
		return
	}

	// 3.2 诊疗数据库，记录diagnosis
	s2 := service.NewWechatDiagnosisService(&e.Api)
	if err := s2.AddDiagnosis(openID, diagnosisType, resMap); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "添加诊疗记录错误: " + err.Error()})
		return
	}

	// 4. 返回结果
	c.JSON(http.StatusOK, gin.H{
		"state": "success",
	})
}

// POST /api/v1/wechat/newDiagnosis
func (e *DiagnosisAPI) NewDiagnosis(c *gin.Context) {

	// 1.1 初始化 Api 上下文和 ORM
	if err := e.MakeContext(c).MakeOrm().Errors; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "初始化失败: " + err.Error()})
		return
	}

	// 1.2 获取参数
	var req DiagnosisReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "请求参数错误: " + err.Error()})
		return
	}

	// 2 获得ai检测结果
	// 2.1 转回为 JSON，后续当作请求体使用
	forwardJSON, err := json.Marshal(req.ThirdAPIData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "JSON 序列化失败: " + err.Error()})
		return
	}

	// 2.2 发送请求
	data, statusCode, err := client.DetDiagnosisFromServer(forwardJSON)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "外部获取检测结果失败:  " + err.Error()})
		return
	}

	// 2.3 如果是舌诊，且包含舌后照片，获取不到舌后信息，尝试吧舌后信息删除
	if code, ok := data["code"].(float64); ok && int(code) == 20150 {
		if msg, ok := data["msg"].(string); ok && msg == "舌下络脉目标检测失败" {
			//删除请求中TBImage对应的内容，转回为 JSON，后续当作请求体使用
			req.ThirdAPIData.TBImage = ""
			forwardJSON, err := json.Marshal(req.ThirdAPIData)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"message": "JSON 序列化失败: " + err.Error()})
				return
			}
			//重新发起请求
			data, statusCode, err = client.DetDiagnosisFromServer(forwardJSON)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"message": "外部获取检测结果失败:  " + err.Error()})
				return
			}
		}
	}

	// 3 修改结果
	if code, ok := data["code"].(float64); ok && int(code) == 20000 {
		if dataMap, ok := data["data"].(map[string]interface{}); ok {

			// 3.1 获取第一个体质名称，以及修改后的physique_name参数
			firstPhysiqueName, newNameForm := service.GetPhysiqueName(dataMap)
			dataMap["physique_name"] = newNameForm
			//fmt.Println("physique string final:", firstPhysiqueName)

			// 3.2 从体质表中，获取对应的prodict_ids和治疗方法
			s1 := service.NewWechatPhysiqueService(&e.Api)
			physiqueInfo, err := s1.GetPhysiqueInfo(firstPhysiqueName)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"message": "获取体质对应信息错误: " + err.Error()})
				return
			}

			// 3.3 添加商品信息
			s2 := service.NewWechatProductsService(&e.Api)
			dataMap["goods"] = s2.GetProductsInfo(physiqueInfo.ProductIDs)

			// 3.4 添加自定义疗法
			advicesMap := service.GetAcupunctureMethod(dataMap, physiqueInfo)
			dataMap["advices"] = advicesMap
			data["data"] = dataMap

			// 3.5 添加随机获取问题
			s3 := service.NewWechatQuestionService(&e.Api)
			data["inquiry_questions"], _ = s3.GetRandomQuestion()

			// 4 返回结果
			c.JSON(statusCode, data)
			return
		}
	}
	// 或者直接返回
	c.JSON(statusCode, data)

}

// GET /api/v1/wechat/UserDiagnosisNum
func (e *DiagnosisAPI) UserDiagnosisNum(c *gin.Context) {

	// 1. 初始化 Api 上下文和 ORM
	if err := e.MakeContext(c).MakeOrm().Errors; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "初始化失败: " + err.Error()})
		return
	}

	// 2. 获取响应数据
	var req GetUserDiagnosisNumReq
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "参数错误: " + err.Error()})
		return
	}

	// 3. 数据库：
	// 查询用户表，诊疗数量
	s := service.NewWechatUserService(&e.Api)

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
func (e *DiagnosisAPI) DiagnosisList(c *gin.Context) {

	// 1. 初始化 Api 上下文和 ORM
	if err := e.MakeContext(c).MakeOrm().Errors; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "初始化失败: " + err.Error()})
		return
	}

	// 2. 获取响应数据
	var req DiagnosisListReq
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "参数错误: " + err.Error()})
		return
	}
	fmt.Println(req)
	// 3. 数据库：
	// 获取列表
	s := service.NewWechatDiagnosisService(&e.Api)

	diagnosisList, total, err := s.GetDiagnosisList(req.OpenID, req.DiagnosisType, req.PageSize, req.PageIndex)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "查询用户诊疗列表错误: " + err.Error()})
		return
	}

	// 4. 返回结果
	c.JSON(http.StatusOK, gin.H{
		"diagnosisList": diagnosisList,
		"total":         total,
	})
}
