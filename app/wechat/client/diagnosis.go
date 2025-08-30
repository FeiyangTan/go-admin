package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"go-admin/app/wechat/config"
	"io"
	"net/http"
)

func DetDiagnosisFromServer(forwardJSON []byte) (map[string]interface{}, int, error) {

	// 2 构造 HTTP 请求(ai检测)
	url := "https://ali-market-tongue-detect-v2.macrocura.com/diagnose/face-tongue/result/"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(forwardJSON))
	if err != nil {
		return nil, 0, errors.New("请求构造失败: " + err.Error())
	}

	// 设置 header
	AIAPPCODE := config.AIAPPCODE
	req.Header.Set("Authorization", AIAPPCODE)
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	// 3 发请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, 0, errors.New("请求失败: " + err.Error())
	}

	defer resp.Body.Close()

	// 读成功响应（加上限，避免 OOM）
	bodyBytes, err := io.ReadAll(io.LimitReader(resp.Body, 8<<20)) // 8MB 上限
	if err != nil {
		return nil, resp.StatusCode, fmt.Errorf("读取响应失败: %w", err)
	}

	var data map[string]interface{}
	if err := json.Unmarshal(bodyBytes, &data); err != nil {
		// JSON 解析失败就原样返回
		return nil, resp.StatusCode, fmt.Errorf("JSON 解析失败: %w", err)
	}

	return data, resp.StatusCode, nil

	//return resp, nil
}
