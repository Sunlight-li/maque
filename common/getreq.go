package common

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

type DingdingResp struct {
	Errcode int64  `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

// 发送告警信息
func Send_info(a, baseUrl, keyword, title, info_type string) int {
	// log.Println(a)
	req := map[string]interface{}{
		"at": map[string]interface{}{
			"atMobiles": map[string]interface{}{},
			"atUserIds": map[string]interface{}{},
			"isAtAll":   true,
		},
		"title": title,
		"text": map[string]interface{}{
			"content": keyword + info_type + ": \n  系统名称：" + title + "\n " + a,
		},
		"msgtype": "text",
	}
	resp, err := dingding(req, baseUrl)
	if err != nil {
		log.Println(err)
	}
	body, err := ToJsonBuff(resp)
	if err != nil {
		log.Println("json转换失败", err)
	}
	log.Printf("输出返回结果：%s", body)
	if body.String() == "null" {
		return -1
	}
	return 1
}
func dingding(req map[string]interface{}, baseUrl string) (*DingdingResp, error) {
	url, err := url.Parse(baseUrl)
	if err != nil {
		return nil, err
	}
	body, err := ToJsonBuff(req)
	if err != nil {
		return nil, err
	}
	log.Println("hson请求体：", body)
	r, err := http.NewRequest(http.MethodPost, url.String(), body)
	if err != nil {
		return nil, err
	}
	r.Header.Add("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("响应失败，状态码为 %s", resp.Status)
	}
	var gcresp DingdingResp
	if err := ScanJson(resp, &gcresp); err != nil {
		return nil, err
	}
	return &gcresp, nil
}

// ToJsonBuff 转成 json 格式的 buff，作为 http body
func ToJsonBuff(v any) (*bytes.Buffer, error) {
	var b bytes.Buffer
	e := json.NewEncoder(&b)
	// 设置禁止html转义
	e.SetEscapeHTML(false)
	err := e.Encode(v)
	return &b, err
}

// ScanJson 按json格式解析resp的payload（Body）
func ScanJson(resp *http.Response, v interface{}) error {
	return json.NewDecoder(resp.Body).Decode(v)
}
