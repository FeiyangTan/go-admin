package apis

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"go-admin/app/wechat/models"
	"go-admin/app/wechat/service"
	"net/http"
)

type QuestionAPI struct {
	api.Api
}

// 前端请求体：inquiry_questions 为 {name, value} 的数组
type inquiryItem struct {
	Name  string `json:"name"  binding:"required"` // 必填
	Value string `json:"value" binding:"required"` // 必填（允许为空就去掉 required）
}

type questionAddReq struct {
	InquiryQuestions []inquiryItem `json:"inquiry_questions" binding:"required,min=1"`
}

// GET /api/v1/wechat/question/get
func (q *QuestionAPI) Question_get(c *gin.Context) {

	// 0. 初始化 Api 上下文和 ORM
	if err := q.MakeContext(c).MakeOrm().Errors; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "初始化失败: " + err.Error()})
		return
	}

	// 1. 数据库：随机获取问题
	s := service.NewWechatQuestionService(&q.Api)
	questions, err := s.GetRandomQuestion()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "获取问题，数据库操作失败: " + err.Error()})
		return
	}
	fmt.Println(questions)

	c.JSON(http.StatusOK, gin.H{})
}

// POST /api/v1/wechat/question/add
func (q *QuestionAPI) Question_add(c *gin.Context) {

	// 0. 初始化 Api 上下文和 ORM
	if err := q.MakeContext(c).MakeOrm().Errors; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "初始化失败: " + err.Error()})
		return
	}

	// 1) 解析并校验请求体
	var req questionAddReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "参数错误",
			"error":   err.Error(),
		})
		return
	}

	// 2) 直接映射到模型切片
	items := make([]models.Question, 0, len(req.InquiryQuestions))
	for _, it := range req.InquiryQuestions {
		items = append(items, models.Question{
			Name:  it.Name,
			Value: it.Value,
		})
	}

	// 3) 调用 Service：基于唯一索引（name）的 UPSERT 忽略冲突
	s := service.NewWechatQuestionService(&q.Api)
	inserted, err := s.AddInquiryQuestionsUpsert(items)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "数据库操作失败: " + err.Error()})
		return
	}
	fmt.Println("~~~~前端提交的问题总数：", len(req.InquiryQuestions))

	fmt.Println("~~~~成功添加问题数量：", inserted)

	// 4) 返回统计
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}
