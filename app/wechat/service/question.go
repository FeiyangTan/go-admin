// app/wechat/service/question.go
package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"go-admin/app/wechat/models"
	"gorm.io/gorm/clause"
	"math/rand"
	"net/http"
	"strings"
)

type WechatQuestionService struct {
	api.Api
}

// NewWechatQuestionService 会在 Handler 里被 MakeService 调用
func NewWechatQuestionService(e *api.Api) *WechatQuestionService {
	return &WechatQuestionService{*e}
}

// AddQuestion 向数据库新增一条 Question
func (s *WechatQuestionService) AddQuestion(q *models.Question) error {
	db, _ := s.GetOrm()

	// 使用 gorm 的 Create 方法
	if err := db.Create(q).Error; err != nil {
		return err
	}
	return nil
}

// AddInquiryQuestionsUpsert 批量新增 questions；遇到 name 冲突（唯一索引）就跳过
func (s *WechatQuestionService) AddInquiryQuestionsUpsert(items []models.Question) (int64, error) {
	if len(items) == 0 {
		return 0, errors.New("empty payload")
	}

	db, _ := s.GetOrm()

	tx := db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "name"}}, // 与唯一索引列一致
		DoNothing: true,                            // 冲突时不做任何事（忽略）
	}).Create(&items)

	return tx.RowsAffected, tx.Error
}

func (s *WechatQuestionService) AddMutipleQuestions(bodyBytes []byte, c *gin.Context) {
	fmt.Println("~~~~添加问题集合")
	// 3.2 先用 map[string]interface{} 解析
	var data map[string]interface{}
	if err := json.Unmarshal(bodyBytes, &data); err != nil {
		// JSON 解析失败就原样返回
		c.JSON(http.StatusInternalServerError, gin.H{"message": "添加问题集合，解析1错误: " + err.Error()})
		return
	}

	// 3.3 从 data.inquiry_questions 提取并写入数据库（基于唯一索引 name，冲突忽略）
	// 需要：import "strings" 以及已引入 models / service 包
	if dataMap, ok := data["data"].(map[string]interface{}); ok {
		iqRaw, exists := dataMap["inquiry_questions"]
		if exists {
			if list, ok := iqRaw.([]interface{}); ok {

				items := make([]models.Question, 0, len(list))
				for _, v := range list {
					obj, ok := v.(map[string]interface{})
					if !ok {
						continue
					}
					name, _ := obj["name"].(string)
					value, _ := obj["value"].(string)
					if strings.TrimSpace(name) == "" {
						continue
					}
					items = append(items, models.Question{
						Name:  name,
						Value: value,
					})
				}

				// 3) 调用 Service：基于唯一索引（name）的 UPSERT 忽略冲突
				_, err := s.AddInquiryQuestionsUpsert(items)
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"message": "添加问题集合,数据库操作失败: " + err.Error()})
					return
				}
				//fmt.Println("~~~~前端提交的问题总数：", len(items))
				//
				//fmt.Println("~~~~成功添加问题数量：", inserted)
			}

		}
	}
}

// GetRandomQuestions 从数据库随机获取 10～15 个 Questio
func (s *WechatQuestionService) GetRandomQuestion() ([]interface{}, error) {
	db, _ := s.GetOrm() // 从 s.Api 已初始化的 Orm 拿到 *gorm.DB
	// 生成一个随机数，范围在 [10, 15]
	count := 10 + rand.Intn(6)

	var questions []models.Question
	// 随机从数据库中获取数据
	err := db.Order("RAND()").Limit(count).Find(&questions).Error
	if err != nil {
		return nil, err
	}

	var questionsArray []interface{}
	for _, question := range questions {
		questionMap := map[string]interface{}{
			"name":  question.Name,
			"value": question.Value,
		}
		questionsArray = append(questionsArray, questionMap)
	}

	return questionsArray, nil
}
