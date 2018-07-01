package extension

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"seuxw/embrice/entity"
	"seuxw/x/logger"
	"time"
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

// Api接口调用终点 - label END
func EndLabel(log *logger.Logger, err error, w http.ResponseWriter, Pagination *entity.Pagination, ResponseData ...interface{}) {
	var (
		data       []byte
		pagination []byte
	)
	// 如果请求结果有错误
	if err != nil {
		log.Error(fmt.Sprintf("%s", err))
		data = []byte(fmt.Sprintf(`{"code":1,"message":"%v"}`, err))
		goto END
	}

	// Marshal返回的data发生错误
	if data, err = json.Marshal(&ResponseData); err != nil {
		data = []byte(`{"code":1,"message":"json Marshal Err"}`)
		goto END
	}

	/* 如果Pagination为空的情况（一般不是列表的查询结果都是空）*/
	if Pagination == nil {
		// 如果返回的数据为空
		if string(data) == "[{}]" {
			data = []byte(fmt.Sprintf(`{"code":0,"message":"success.","data":{}}`))
			goto END
		}
		// 如果返回的data有数据
		data = []byte(fmt.Sprintf(`{"code":0,"message":"success.","data":%v}`, string(data)[1:len(string(data))-1]))
		goto END
	}

	/* 如果Pagination不为空的情况（一般data为数组）*/
	if pagination, err = json.Marshal(Pagination); err != nil {
		data = []byte(`{"code":1,"message":"json Marshal Err"}`)
		goto END
	}
	// 如果返回的数组为空
	if string(data) == "[null]" {
		data = []byte(fmt.Sprintf(`{"code":0,"message":"success.","data":[], "pagination":%v}`, string(pagination)))
		goto END
	}
	// 如果返回的是有数据的数组
	data = []byte(fmt.Sprintf(`{"code":0,"message":"success.","data":%v, "pagination":%v}`, string(string(data[1:len(data)-1])), string(pagination)))

END:
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write([]byte(data))
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
