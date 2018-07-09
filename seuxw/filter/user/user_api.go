package main

import (
	"encoding/json"
	"fmt"
	py "github.com/sbinet/go-python"
	"net/http"
	"seuxw/embrice/entity"
	"seuxw/embrice/entity/user"
	"seuxw/embrice/extension"
)

// CreateUser 创建用户 Handler
// 目前 *仅接受* 群内签到注册申请用户和一卡通账号用户，一般群聊用户请不要调用此接口
func (svr *server) CreateUser(w http.ResponseWriter, r *http.Request) {
	var (
		response entity.Response
		err      error
		data     []byte
		user     user.User
	)

	seuxwRequest := entity.GetSeuxwRequest(r)
	processName := "CreateUser"

	body, _ := extension.HandlerRequestLog(seuxwRequest, processName)

	pyExec := svr.py.GetAttrString("exec")
	fmt.Println(pyExec)
	res := pyExec.Call(py.Py_None, py.Py_None)
	fmt.Println(res)

	// 参数分析
	err = json.Unmarshal(body, &user)

	if err != nil {
		err = fmt.Errorf("Json 解析错误")
		goto END
	}

	if user.QQID == 0 && user.CardID == 0 {
		err = fmt.Errorf("缺少必要参数 qq_id 或 card_id，请检查！")
		goto END
	}

	// 获取用户头像图片 http://q4.qlogo.cn/g?b=qq&nk={qq_id}&s=140
	if user.QQID != 0 {

	}

	// 数据库操作
	err = svr.db.CreateUserDB(user)
	if err != nil {
		err = fmt.Errorf("数据库调用错误！ %s", err)
		goto END
	}

END:
	if err != nil {
		response.Code = 1
		response.Message = fmt.Sprintf("%s", err)
	}

	data = extension.HandlerResponseLog(seuxwRequest, response, processName, true)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write([]byte(data))
}
