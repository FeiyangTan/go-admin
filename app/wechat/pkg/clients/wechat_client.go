package clients

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Jscode2SessionResp struct {
	OpenID     string `json:"openid"`
	SessionKey string `json:"session_key"`
	ErrCode    int    `json:"errcode"`
	ErrMsg     string `json:"errmsg"`
}

type WechatClient interface {
	Jscode2Session(ctx context.Context, appid, secret, code string) (*Jscode2SessionResp, error)
}

type httpWechatClient struct {
	http *http.Client
}

func NewWechatClient() WechatClient {
	return &httpWechatClient{
		http: &http.Client{Timeout: 5 * time.Second},
	}
}

func (w *httpWechatClient) Jscode2Session(ctx context.Context, appid, secret, code string) (*Jscode2SessionResp, error) {
	url := fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code", appid, secret, code)
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	resp, err := w.http.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var out Jscode2SessionResp
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, err
	}
	if out.ErrCode != 0 {
		return nil, fmt.Errorf("wechat error: %d %s", out.ErrCode, out.ErrMsg)
	}
	return &out, nil
}
