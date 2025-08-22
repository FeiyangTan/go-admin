package apis

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"go-admin/app/wechat/models"
	"go-admin/app/wechat/service"
	"gorm.io/datatypes"
	"io"
	"net/http"
	"strings"

	"go-admin/app/wechat/config"
)

type DiagnosisAPI struct {
	api.Api
}

type AddDiagnosisReq struct {
	OpenID          string         `json:"open_id" binding:"required"`
	DiagnosisType   string         `json:"diagnosis_type" binding:"required"`
	DiagnosisResult datatypes.JSON `json:"diagnosis_result" binding:"required"`
}

type NewDiagnosisReq struct {
	Scene         int    `json:"scene" binding:"required"`
	FFImage       string `json:"ff_image"`
	TFImage       string `json:"tf_image"`
	TBImage       string `json:"tb_image"`
	Gender        string `json:"gender"`
	OpenID        string `json:"open_id"`
	DiagnosisType string `json:"diagnosis_type"`
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

// POST /api/v1/wechat/newDiagnosis
func (d *DiagnosisAPI) NewDiagnosis(c *gin.Context) {

	// 1. 初始化 Api 上下文和 ORM
	if err := d.MakeContext(c).MakeOrm().Errors; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "初始化失败: " + err.Error()})
		return
	}

	// 2 发送请求
	// 2.1 直接读 Body
	raw, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "读取请求失败: " + err.Error()})
		return
	}

	// 2.2 为“对外请求”做一个可编辑的副本：转 map、删除字段
	var forward map[string]interface{}
	if err := json.Unmarshal(raw, &forward); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "请求体不是有效 JSON: " + err.Error()})
		return
	}
	// 删除指定字段（如果不存在 delete 也不会报错）
	//fmt.Println(forward)
	requestScene := 0
	if scene, ok := forward["scene"].(float64); ok {
		requestScene = int(scene)
	}
	openID, _ := forward["open_id"].(string)
	//fmt.Println("openID:", openID)

	diagnosisType, _ := forward["diagnosis_type"].(string)
	//fmt.Println("diagnosisType:", diagnosisType)

	delete(forward, "open_id")
	delete(forward, "diagnosis_type")

	// 2.3 转回为 JSON，后续当作请求体使用
	forwardJSON, err := json.Marshal(forward)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "JSON 序列化失败: " + err.Error()})
		return
	}

	// 2.4 构造 HTTP 请求(ai检测)
	url := "https://ali-market-tongue-detect-v2.macrocura.com/diagnose/face-tongue/result/"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(forwardJSON))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "请求构造失败: " + err.Error()})
		return
	}

	// 设置 header
	AIAPPCODE := config.AIAPPCODE
	req.Header.Set("Authorization", AIAPPCODE)
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	// 2.5 发请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "请求失败: " + err.Error()})
		return
	}
	defer resp.Body.Close()

	// 3 处理结果（修改、存储、返回给前端）
	// 3.1 读取外部响应
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		c.String(http.StatusBadGateway, "读取外部响应失败: %v", err)
		return
	}

	// 若是 JSON
	ct := resp.Header.Get("Content-Type")

	if strings.Contains(strings.ToLower(ct), "application/json") && requestScene == 2 {
		fmt.Println("2222222222222")

		// 3.2 把结果变成map
		// 用 map[string]interface{} 方便动态修改
		var data map[string]interface{}
		if err := json.Unmarshal(bodyBytes, &data); err != nil {
			// JSON 解析失败就原样返回
			c.Data(resp.StatusCode, ct, bodyBytes)
			return
		}

		// 3.3 修改结果
		// 如果 data["data"] 是一个 map，可以直接修改
		if dataMap, ok := data["data"].(map[string]interface{}); ok {
			// 3.3.1 获取第一个体质名称
			physique := dataMap["physique_name"]
			//fmt.Println("~~~physique:", physique)

			var physiqueString string
			if s, ok := physique.(string); ok {
				physiqueString = s
			} else {
				physiqueString = "血瘀体质"
			}
			//fmt.Println("physique string:", physiqueString)

			var physiqueStringFinal string
			if idx := strings.Index(physiqueString, "、"); idx != -1 {
				physiqueStringFinal = physiqueString[:idx]
			} else {
				physiqueStringFinal = physiqueString
			}
			//fmt.Println("physique string final:", physiqueStringFinal)

			// 3.3.2 从体质表中，获取对应的prodict_ids和治疗方法
			s0 := service.NewWechatPhysiqueService(&d.Api)
			physiqueInfo, err := s0.GetPhysiqueInfo(physiqueStringFinal)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"message": "获取体质对应信息错误: " + err.Error()})
				return
			} else if physiqueInfo == nil || physiqueInfo.AcupunctureMethod == "" {
				// 未找到对应的体质信息，使用默认体质
				fmt.Println("~~~~~~默认体质")
				physiqueInfo, err = s0.GetPhysiqueInfo("默认体质")
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"message": "获取默认体质信息错误: " + err.Error()})
					return
				}
				// 把新的体质添加到数据库
				if physiqueInfo == nil {
					fmt.Println("~~~~~~新的体质添加到数据库")
					p := &models.Physique{PhysiqueName: physiqueStringFinal}
					cre, err := s0.CreatePhysique(p)
					fmt.Println("~~~~~~新的体质添加到数据库", cre)

					if err != nil {
						c.JSON(http.StatusInternalServerError, gin.H{"message": "添加新体制错误: " + err.Error()})
						return
					}
				}
			}
			// 提取 acupuncture_method 和 product_ids
			acupunctureMethod := physiqueInfo.AcupunctureMethod
			productIDs := physiqueInfo.ProductIDs
			//fmt.Println("acupunctureMethod:", acupunctureMethod)
			fmt.Println("productIDs:", productIDs)

			// 3.3.3 添加商品信息
			// 从商品表中获取商品信息
			productIDs = strings.ReplaceAll(productIDs, "，", ",")
			ids := strings.Split(productIDs, ",")
			fmt.Println("ids:", ids)

			var result []map[string]interface{}
			s := service.NewWechatProductsService(&d.Api)
			for _, idStr := range ids {
				//fmt.Println("~~~~~~~~~~~~~~")
				//fmt.Println("idStr: ", idStr)

				idStr = strings.TrimSpace(idStr)
				if idStr == "" {
					continue
				}

				// 直接用 string 调用 GetProductsInfo
				product, err := s.GetProductsInfo(idStr)
				if err != nil || product == nil {
					continue
				}

				// 组装 map
				item := map[string]interface{}{
					"id":            product.ID,
					"name":          product.ProductName,
					"price":         product.Price,
					"imageSrc":      product.ImageURL,
					"description":   product.Description,
					"ingredient":    product.Ingredients,
					"useMethod":     product.UsageMethod,
					"mallProductID": product.MallProductID,
				}

				// 放入结果数组
				result = append(result, item)
			}

			//jsonBytes, _ := json.Marshal(result)
			//jsonString := string(jsonBytes)
			//fmt.Println("jsonString: ", jsonString)
			dataMap["goods"] = result

			// 3.3.4
			// 添加自定义疗法
			if advicesMap, ok := dataMap["advices"].(map[string]interface{}); ok {
				if treatments, ok := advicesMap["treatment"].([]interface{}); ok {
					// 构造新的一组内容
					newItem := map[string]interface{}{
						"advice": acupunctureMethod,
						"title":  "三才河洛灸",
					}
					// 添加到数组
					//treatments = append(treatments, newItem)
					treatments = append([]interface{}{newItem}, treatments...)
					advicesMap["treatment"] = treatments
					dataMap["advices"] = advicesMap
					data["data"] = dataMap
				}
			}

		}
		// 换了3家，
		// 3.4 返回结果
		c.JSON(resp.StatusCode, data)

		// 3.5 存入数据库：
		// 用户数据库，diagnosis数量+1
		s := service.NewWechatUserService(&d.Api)
		if err := s.AddDiagnosisCount(openID, diagnosisType); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "添加诊疗数量错误: " + err.Error()})
			return
		}

		// 诊疗数据库，记录diagnosis
		s2 := service.NewWechatDiagnosisService(&d.Api)

		// 获取 data 中的 "data" 字段
		var diagnosisJSON datatypes.JSON
		if val, ok := data["data"]; ok {
			if b, err := json.Marshal(val); err == nil {
				diagnosisJSON = datatypes.JSON(b)
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"message": "诊断结果序列化失败: " + err.Error()})
				return
			}
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"message": "响应中缺少 data 字段"})
			return
		}

		fmt.Println(diagnosisJSON)

		if err := s2.AddDiagnosis(openID, diagnosisType, diagnosisJSON); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "添加诊疗记录错误: " + err.Error()})
			return
		}
		return
	}
	// 或者直接返回
	c.Data(resp.StatusCode, ct, bodyBytes)

}

// GET /api/v1/wechat/UserDiagnosisNum
func (d *DiagnosisAPI) UserDiagnosisNum(c *gin.Context) {

	// 1. 初始化 Api 上下文和 ORM
	if err := d.MakeContext(c).MakeOrm().Errors; err != nil {
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
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "参数错误: " + err.Error()})
		return
	}
	fmt.Println(req)
	// 3. 数据库：
	// 获取列表
	s := service.NewWechatDiagnosisService(&d.Api)

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
