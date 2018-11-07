# MeizuWeather Apidoc

## BasicInfo

- url: http://aider.meizu.com/app/weather/listWeather
- method: `GET`
- params:
    - cityIds 城市 ID 列表
- sample:

```py
requests.get("http://aider.meizu.com/app/weather/listWeather?cityIds=101190104")
```

- returns:
```json
{
    "code":0,                       // 错误码
    "message":"",                   // 错误信息
    "redirect":"",                  // 重定向信息
    "value":[                       // Value
        {
            "alarms":[              // 预警信息 如：台风预警

            ],
            "city":"",              // 城市名称
            "cityid":0,             // 城市 ID
            "indexes":[{            // 生活指数信息
                "abbreviation":"",  // 简写
                "alias":"",         // 别名
                "content":"",       // 内容
                "level":"",         // 等级
                "name":""           // 名称
            }],
            "pm25":{                // 空气质量
                "advice":"",        // 
                "aqi":"",           // 
                "citycount":0,      //
                "cityrank":0,       //
                "co":"",            // 一氧化碳含量
                "color":"",         //
                "level":"",         // 空气质量等级
                "no2":"",           // 二氧化氮含量
                "o3":"",            // 臭氧含量
                "pm10":"",          // PM10
                "pm25":"",          // PM2.5
                "quality":"",       // 空气质量
                "so2":"",           // 二氧化硫含量
                "timestamp":"",     // （废弃）时间戳
                "updateTime":""     // 更新时间
            },
            "provinceName":"",      // 省份名称
            "realtime":{            // 实时
                "img":"",           //
                "sb":"",            //
                "sendibleTemp":"",  // 体感温度
                "temp":"",          // 温度
                "time":"",          // 更新时间
                "wD":"",            // 风向
                "wS":"",            // 风等级
                "weather":"",       // 天气
                "ziwaixian":""      // 紫外线
            },
            "weatherDetailsInfo":{                  // 天气详细信息
                "publishTime":"",                   // 发布时间
                "weather3HoursDetailsInfos":[{      // 短时天气预报
                    "endTime":"",                   // 结束时间
                    "highestTemperature":"",        // 最高温度
                    "img":"",                       //
                    "isRainFail":"",                // 是否下雨
                    "lowerestTemperature":"",       // 最低温度
                    "precipitation":"",             //
                    "startTime":"",                 // 开始时间
                    "wd":"",                        //
                    "ws":"",                        //
                    "weather":""                    // 天气
                }]
            },
            "weathers":[{                           // 七日天气预报
                "date":"",                          // 日期
                "img":"",                           //
                "sun_down_time":"",                 // 日落时间
                "sun_rise_time":"",                 // 日出时间
                "temp_day_c":"",                    // 白天温度 ℃
                "temp_day_f":"",                    // 白天温度 ℉
                "temp_night_c":"",                  // 晚上温度 ℃
                "temp_night_f":"",                  // 晚上温度 ℉
                "wd":"",                            // 
                "ws":"",                            //
                "weather":"",                       // 天气
                "week":""                           // 周几
            }]
        }
    ]
}
```