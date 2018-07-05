package extension

import (
	"fmt"
	"encoding/json"
	"math"
	"seuxw/embrice/entity"
	"seuxw/x/logger"
	"time"
	"strings"
	"runtime"
	"io/ioutil"
)

const TIME_FORMAT = "2006-01-02 15:04:05"
const DATE_FORMAT = "2006-01-02"

// 获取当前时间字符串
func CurrentTimeInStr() string {
	return time.Now().Format(TIME_FORMAT)
}

// 获取当前日期字符串
func CurrentDateInStr() string {
	return time.Now().Format(DATE_FORMAT)
}

// 获取当前时间戳
func CurrentTimeStamp() int64 {
	return time.Now().Unix()
}

// 将时间字符串转换为时间戳
func Date2TimeStamp(date string) (timestamp int64) {
	t, err := time.ParseInLocation(TIME_FORMAT, "2017-07-25 08:54:14", time.Local)
	if err != nil {
		panic(err)
	}
	timestamp = t.Unix()
	return timestamp
}

// 将时间戳转换为时间字符串
func TimeStamp2Str(timestamp int64) (result string) {
	t := time.Unix(timestamp, 0)
	result = t.Format("2006-01-02 15:04:05")
	return result
}

// 将时间戳转换为时间字符串
func TS2Str(timestamp int64) (result string) {
	if timestamp > 0 {
		result = time.Unix(timestamp, 0).Format("2006-01-02 15:04:05")
	} else {
		result = "0000-00-00 00:00:00"
	}
	return result
}

//获取不带时、分、秒的时间；day是增加/减少的天数
func GetDateTime(inputTime time.Time, day int) int64 {
	return time.Date(inputTime.Year(), inputTime.Month(), inputTime.Day()+day, 0, 0, 0, 0, time.Local).Unix()
}

// 控制小数位
func Precision(f float64, prec int, round bool) float64 {
	pow10_n := math.Pow10(prec)
	if round {
		if f >= 0 {
			return math.Trunc((f+0.5/pow10_n)*pow10_n) / pow10_n
		}
		return math.Trunc((f-0.5/pow10_n)*pow10_n) / pow10_n
	}
	return math.Trunc((f)*pow10_n) / pow10_n
}

type FuncInfo struct {
	FuncName string
	FileName string
	Line     int
	Status   bool
}

// GetFuncInfo 获取当前函数信息
func GetFuncInfo() FuncInfo {
	p, file, line, ok := runtime.Caller(2)
	f := runtime.FuncForPC(p)
	return FuncInfo{
		FuncName: strings.Split(f.Name(), ".")[2],
		FileName: file,
		Line:     line,
		Status:   ok,
	}
}

// HandlerRequestLog 路由层请求添加Log日志
func HandlerRequestLog(SeuxwRequest *entity.SeuxwRequest, processName string) ([]byte, error) {
	log := logger.NewStdLogger(true, true, true, true, true)
	funcInfo := GetFuncInfo()
	method := SeuxwRequest.RawRequest.Method
	body, _ := ioutil.ReadAll(SeuxwRequest.RawRequest.Body)

	if method == "GET" {
		log.Trace("[%s:%d] %s流程开始～ => TraceID: %s, URI: %s, UsrID: %d, UsrNm: %s", funcInfo.FuncName, funcInfo.Line, processName,
			SeuxwRequest.TraceID, SeuxwRequest.Action, SeuxwRequest.AccountID, SeuxwRequest.RealName)
	} else if method == "POST" {
		
		log.Trace("[%s:%d] %s流程开始～ => TraceID: %s, URI: %s, Params: %s, UsrID: %d, UsrNm: %s", funcInfo.FuncName, funcInfo.Line, processName,
			SeuxwRequest.TraceID, SeuxwRequest.Action, string(body), SeuxwRequest.AccountID, SeuxwRequest.RealName)
	} else {
		return body, fmt.Errorf("系统内部日志生成错误！")
	}
	return body, nil
}

// HandlerResponseLog 路由层请求添加Log日志
func HandlerResponseLog(SeuxwRequest *entity.SeuxwRequest, response entity.Response, processName string, showData bool) []byte{
	var (
		result string
	)
	data, _ := json.Marshal(&response)
	log := logger.NewStdLogger(true, true, true, true, true)
	funcInfo := GetFuncInfo()
	if response.Code == 0 {
		result = "成功"
	} else {
		result = "失败"
	}
	if showData {
		log.Trace("[%s:%d] %s流程%s～ => TraceID: %s, Return: %s", funcInfo.FuncName, funcInfo.Line,
			processName, result, SeuxwRequest.TraceID, string(data))
	} else {
		log.Trace("[%s:%d] %s流程%s～ => TraceID: %s, Message:%s, Pagination: %+v", funcInfo.FuncName, funcInfo.Line,
			processName, result, SeuxwRequest.TraceID, response.Message, response.Pagination)
	}
	return data
}