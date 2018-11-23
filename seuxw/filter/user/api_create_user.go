package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"seuxw/embrice/constant/constuser"
	"seuxw/embrice/entity"
	"seuxw/embrice/entity/user"
	"seuxw/embrice/extension"
	"strconv"
)

// CreateUser 创建用户 Handler
// 目前 *仅接受* 群内签到注册申请用户和一卡通账号用户，一般群聊用户请不要调用此接口
func (svr *server) CreateUser(w http.ResponseWriter, r *http.Request) {
	var (
		response entity.Response
		err      error
		data     []byte
	)

	seuxwRequest := entity.GetSeuxwRequest(r)
	processName := "CreateUser"
	userObj := new(user.User)
	cardObj := new(user.Card)

	body, _ := extension.HandlerRequestLog(seuxwRequest, processName)

	// 参数分析
	err = json.Unmarshal(body, &userObj)

	if err != nil {
		err = fmt.Errorf("Json 解析错误")
		goto END
	}

	if userObj.QQID == 0 && userObj.CardID == 0 {
		err = fmt.Errorf("缺少必要参数 qq_id 或 card_id，请检查！")
		goto END
	}

	// 获取用户 QQ 信息
	if userObj.QQID != 0 {
		// todo: 通过已登录账号的 cookies 和目标用户的 QQ 账号获取其他用户的基础信息
		// 实现方案1：grpc 或者 NATS 延迟调用py脚本
		// 实现方案2：添加函数执行

		// 获取用户头像图片 http://q4.qlogo.cn/g?b=qq&nk={qq_id}&s=140

	}

	// 获取用户一卡通信息
	if userObj.CardID != 0 {
		// todo: 通过一卡通账号获取用户基础信息
		// 实现方案1：grpc 或者 NATS 延迟调用py脚本

		// 获取基础信息 http://xk.urp.seu.edu.cn/jw_service/service/stuCurriculum.action
		// 18 级以后课表需要统一身份认证

		if err = getInfoFromCard(userObj, cardObj); err != nil {
			goto END
		}

	}

	userObj.UserUUID = extension.NewUUIDString()

	// 数据库操作
	err = svr.db.CreateUserDB(userObj)
	if err != nil {
		err = fmt.Errorf("数据库调用错误！ %s", err)
		goto END
	}

	if userObj.CardID != 0 {
		err = svr.db.CreateCardDB(cardObj)
		if err != nil {
			err = fmt.Errorf("数据库调用错误！ %s", err)
			goto END
		}
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

// getInfoFromCard 从一卡通中获取信息
func getInfoFromCard(user *user.User, card *user.Card) error {

	var (
		err   error
		grade int
	)

	if user.CardID == 0 {
		return fmt.Errorf("user 信息不存在 CardID")
	}

	if isInvalidCardID(user.CardID) {
		return fmt.Errorf("非法值 cardID:%d", user.CardID)
	}

	cardIDStr := strconv.Itoa(int(user.CardID))
	card.CardID = user.CardID
	if grade, err = strconv.Atoi(string([]byte(cardIDStr)[3:5])); err != nil {
		return err
	}
	card.Grade = int64(grade)
	card.Identity = constuser.DictCardType[string([]byte(cardIDStr)[0:3])]

	return nil
}

// isInvalidCardID 是否是非法一卡通号
func isInvalidCardID(cardID int64) bool {

	matched := false
	var err error
	cardIDStr := strconv.Itoa(int(cardID))

	if len(cardIDStr) != 9 {
		return true
	}

	if matched, err = regexp.MatchString(`(110|213|220|230)\d{6}`, cardIDStr); err != nil {
		return true
	}

	return !matched

}

// isValidCardID 是否是合法一卡通号
func isValidCardID(cardID int64) bool {
	return !isInvalidCardID(cardID)
}
