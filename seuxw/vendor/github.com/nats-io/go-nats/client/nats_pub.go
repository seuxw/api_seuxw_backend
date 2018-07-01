// Copyright 2012-2016 Apcera Inc. All rights reserved.

package client

import (
	"seuxw/x/logger"
	"time"

	"github.com/nats-io/go-nats"
)

// 发送消息(接收返回消息)
// url：发送消息的服务地址
// subject：发送消息的主题
// content：发送消息的内容
// timeout：等待返回消息时间（多久不再等待）
// callback：回调函数（接收到回复消息后的处理逻辑函数）
func PubWithRply(url, subject, content string, timeout time.Duration, callback func([]byte)) {
	log := logger.NewStdLogger(true, true, true, true, true)
	nc, _ := nats.Connect(url)
	log.Trace("发送消息开始。\r\n消息主题：%s\r\n消息内容：%s", subject, content)
	defer nc.Close()
	msg, err := nc.Request(subject, []byte(content), timeout)
	if err != nil {
		log.Trace("发送消息失败：%s", err)
		return
	}
	log.Trace("接收到的回复内容：%s", msg.Data)
	callback(msg.Data)
	log.Trace("发送消息结束。")
}
