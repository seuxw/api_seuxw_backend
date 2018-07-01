package test

// Test接口传入参数
type Test struct {
	Date        string `json:"date" db:"date"`                   // 日期
	SunRiseTime string `json:"sun_rise_time" db:"sun_rise_time"` // 日出时间
	SunDownTime string `json:"sun_down_time" db:"sun_down_time"` // 日落时间
}
