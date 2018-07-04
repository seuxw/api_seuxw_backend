package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"seuxw/embrice/entity/test"
	"seuxw/embrice/entity"
	"seuxw/embrice/extension"
)

func (self *server) Test(w http.ResponseWriter, r *http.Request) {
	var (
		date 			string
		response 		entity.Response
		responseData    test.Test
		err     		error
		data 			[]byte
	)

	seuxwRequest := entity.GetSeuxwRequest(r)
	processName := "Test"
	params := make(map[string]string, 0)

	body, _ := extension.HandlerRequestLog(seuxwRequest, processName)

	// 获取date参数
	if r.Method == "GET"{
		date = r.FormValue("date")

	} else {
		err = json.Unmarshal(body, &params)
		date = params["date"]
		if err != nil {
			err = fmt.Errorf("Json 解析错误")
			goto END
		}
	}
	
	fmt.Println("DEBUG:", date)

	// 添加无值默认值 date = Today
	if len(date) == 0 {
		date = extension.CurrentDateInStr()
	}

	// 数据库操作
	responseData, err = self.db.Test(date)
	if err != nil {
		err = fmt.Errorf("数据库调用错误！ %s", err)
		goto END
	}
	response.Data = responseData

END:
	if err != nil {
		response.Code = 1
		response.Message = fmt.Sprintf("%s", err)
	}
	data, _ = json.Marshal(&response)

	extension.HandlerResponseLog(seuxwRequest, response, processName, true)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write([]byte(data))
}
