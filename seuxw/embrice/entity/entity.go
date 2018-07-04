package entity

import (
	"net/http"
)

// Response 请求返回
type Response struct {
	Code       int64       `json:"code"`
	Data       interface{} `json:"data,omitempty"`
	Message    string      `json:"message"`
	Pagination *Pagination `json:"pagination,omitempty"`
}

// Pagination 分页信息
type Pagination struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
	Total    int `json:"total"`
}

// SeuxwRequest Seuxw 项目请求基本信息结构
type SeuxwRequest struct {
	TraceID      string
	AccountID    int64
	Token        string
	EncryptToken string
	IP           string
	Action       string
	Info         string
	AccountName  string
	RealName     string
	RawRequest	 *http.Request
}

// GetRequestHeader 获取 http 请求头参数
func (l *SeuxwRequest) GetRequestHeader() map[string]string {
	header := make(map[string]string)
	header["Pragma"] = "no-cache"
	header["Accept-Encoding"] = "gzip, deflate, sdch"
	header["Accept-Language"] = "zh-CN,zh;q=0.8"
	header["Upgrade-Insecure-Requests"] = "1"
	header["User-Agent"] = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/50.0.2661.102 Safari/537.36"
	header["Accept"] = "text/javascript, text/html, application/xml, text/xml, */"
	header["Connection"] = "keep-alive"
	header["Content-Type"] = "application/x-www-form-urlencoded; charset=UTF-8"
	header["Cache-Control"] = "no-cache"
	header["X-Trace"] = l.TraceID
	header["Token"] = l.Token
	header["Encrypt-Token"] = l.EncryptToken
	return header
}

// GetSeuxwRequest 获取请求中头部参数
func GetSeuxwRequest(r *http.Request) *SeuxwRequest {
	// 内部调用账号ID传递方式
	// if accountID < 1 {
	// 	accountID = strconv.ParseInt(r.FormValue("account_id"))
	// }
	return &SeuxwRequest{
		TraceID:      r.Header.Get("X-Trace"),
		Token:        r.Header.Get("Token"),
		EncryptToken: r.Header.Get("Encrypt-Token"),
		AccountID:    0,
		IP:           r.Header.Get("X-Forwarded-For"),
		Action:       r.RequestURI,
		RealName:     r.Header.Get("Account-RealName"),
		RawRequest:	  r,
	}
}