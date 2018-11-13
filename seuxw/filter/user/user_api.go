package main

import (
	"encoding/json"
	"fmt"
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

	// 获取用户 QQ 信息
	if user.QQID != 0 {
		// todo: 通过已登录账号的 cookies 和目标用户的 QQ 账号获取其他用户的基础信息
		// 实现方案1：grpc 或者 NATS 延迟调用py脚本
		// 实现方案2：添加函数执行

		// 获取用户头像图片 http://q4.qlogo.cn/g?b=qq&nk={qq_id}&s=140

	}

	// 获取用户一卡通信息
	if user.CardID != 0 {
		// todo: 通过一卡通账号获取用户基础信息
		// 实现方案1：grpc 或者 NATS 延迟调用py脚本

		// 获取基础信息 http://xk.urp.seu.edu.cn/jw_service/service/stuCurriculum.action
		// 18 级以后课表需要统一身份认证

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
