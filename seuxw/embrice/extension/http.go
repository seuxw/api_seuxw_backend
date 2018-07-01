package extension

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net"
	"net/http"
	"seuxw/embrice/entity"
	"seuxw/x/logger"
	"strconv"
	"strings"
	"time"
)

// 获取用户的真实IP地址
func GetRealIP(r *http.Request) string {
	// todo: return x-forwarded-for for cdn
	if ip, _, e := net.SplitHostPort(r.RemoteAddr); e == nil {
		return ip
	}
	return r.RemoteAddr
}

// 发送HTTPGet请求
func HTTPGet(URL string, header map[string]string, timeout time.Duration) (error, int, []byte) {
	var (
		request  *http.Request
		resp     *http.Response
		client   http.Client
		respBody []byte
		err      error
	)

	request, err = http.NewRequest("GET", URL, nil)
	if err != nil {
		goto FAILED
	}

	if header != nil {
		for key := range header {
			request.Header.Set(key, header[key])
		}
	}

	client.Timeout = time.Duration(timeout * time.Second)

	resp, err = client.Do(request)
	if err != nil {
		goto FAILED
	}

	if resp == nil {
		goto FAILED
	}

	defer resp.Body.Close()

	respBody, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		goto FAILED
	}

	return nil, resp.StatusCode, respBody

FAILED:
	return err, 0, nil
}

// 发送HTTPPOST请求
func HTTPPost(URL string, header map[string]string, body []byte, timeout time.Duration) (error, int, []byte) {
	var (
		request    *http.Request
		resp       *http.Response
		client     http.Client
		bodyBuffer *bytes.Reader
		respBody   []byte
		err        error
	)

	bodyBuffer = bytes.NewReader(body)

	request, err = http.NewRequest("POST", URL, bodyBuffer)
	if err != nil {
		goto FAILED
	}

	if header != nil {
		for key := range header {
			request.Header.Set(key, header[key])
		}
	}

	client.Timeout = time.Duration(timeout * time.Second)

	resp, err = client.Do(request)
	if err != nil {
		goto FAILED
	}

	if resp == nil {
		goto FAILED
	}

	defer resp.Body.Close()

	respBody, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		goto FAILED
	}

	return nil, resp.StatusCode, respBody

FAILED:
	return err, 0, nil
}

// 发送简单的Get请求
func Get(url string, value interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(result, &value)
}

// 发送简单的POST请求
func Post(url string, params interface{}, value interface{}) error {

	data, err := json.Marshal(params)
	resp, err := http.Post(url, "application/json", strings.NewReader(string(data)))
	if err != nil {
		return err
	}

	result, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return err
	}
	if err = json.Unmarshal(result, value); err != nil {
		return err
	}
	return err
}

// 系统内部调用
func HttpMethod(url, method, param string) (*entity.Response, error) {
	var (
		err  error
		resp *http.Response
		req  *http.Request
		body []byte
		rsp  entity.Response
		log  *logger.Logger = logger.NewStdLogger(true, true, true, true, true)
	)
	defer func() {
		if resp != nil && resp.Body != nil {
			resp.Body.Close()
		}
	}()
	log.Trace("向接口【%s】传递的参数：%s", url, param)
	client := http.Client{}
	req, err = http.NewRequest(method, url, strings.NewReader(param))
	req.Header.Set("Content-type", "application/json")
	if resp, err = client.Do(req); err != nil {
		log.Trace("调用%s出错。\r\n%s", url, err)
		err = errors.New("系统调用错误。")
		goto FAILED
	}
	if resp.StatusCode == 200 {
		body, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Trace("违法的返回值。\r\n%s", err)
			err = errors.New("违法的返回值。")
			goto FAILED
		}
		log.Trace("接口【%s】返回值：%s", url, string(body))
		if err = json.Unmarshal(body, &rsp); err != nil {
			log.Trace("解析返回值失败。 \r\n%s", err)
			err = errors.New("解析返回值失败。")
			goto FAILED
		}
		return &rsp, nil
	}
FAILED:
	return nil, err

}

//获取用户ID
func FetchUserID(r *http.Request) int64 {
	if userID, err := strconv.ParseInt(r.Header.Get("Account-ID"), 10, 64); err == nil {
		return userID
	}
	return 0
}
