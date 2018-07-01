// Copyright 2012-2016 Apcera Inc. All rights reserved.

package client

import (
	"runtime"
	"seuxw/x/logger"

	"github.com/nats-io/go-nats"
)

// 监听消息主题
// url：监听消息的服务地址
// subject：监听消息的主题
// reply：是否回复。true：回复，false：不回复。
// callback：监听到消息之后的处理函数。函数的返回值作为回复时的内容
func Sub(url, subject string, reply bool, callback func([]byte) []byte) {
	log := logger.NewStdLogger(true, true, true, true, true)
	nc, _ := nats.Connect(url)
	nc.Subscribe(subject, func(msg *nats.Msg) {
		log.Trace("接收的消息内容：%s\r\n", msg.Data)
		jsonStr := callback(msg.Data)
		if reply {
			log.Trace("回复【%s】的内容为：%s", subject, jsonStr)
			nc.Publish(msg.Reply, jsonStr)
		}
	})

	nc.Flush()
	runtime.Goexit()
}

// 监听消息主题(使用队列)
// url：监听消息的服务地址
// subject：监听消息的主题
// reply：是否回复。true：回复，false：不回复。
// callback：监听到消息之后的处理函数。函数的返回值作为回复时的内容
func SubWithQueen(url, subject string, reply bool, callback func([]byte) []byte) {
	log := logger.NewStdLogger(true, true, true, true, true)
	nc, _ := nats.Connect(url)
	nc.QueueSubscribe(subject, subject+"-queen", func(msg *nats.Msg) {
		log.Trace("接收的消息内容：%s\r\n", msg.Data)
		jsonStr := callback(msg.Data)
		if reply {
			log.Trace("回复【%s】的内容为：%s", subject, jsonStr)
			nc.Publish(msg.Reply, jsonStr)
		}
	})

	nc.Flush()
	runtime.Goexit()
}
