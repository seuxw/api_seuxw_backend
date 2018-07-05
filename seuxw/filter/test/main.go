package main

import (
	"context"
	"seuxw/embrice/rdb/test"
	"seuxw/x/logger"
	"seuxw/x/web"
)

type server struct {
	*web.Server
	db  *test.Database
	log *logger.Logger
	ctx context.Context
}

// start 程序开始
func (self *server) start() {
	self.Server = web.NewServer(self.log)
	router := self.PathPrefix("/").SubRouter()

	/* -测试Handler- */
	// 测试接口 将会返回输入的日期对应的日出日落时间（需要对应时间）
	router.HandleFunc("/test", self.Test).Methods("GET", "POST")

	self.Serve("0.0.0.0:10000")
}

// stop 服务终止，断开数据库连接
func (self *server) stop() {
	self.db.Close()
}

// main 主程序入口
func main() {
	log := logger.NewStdLogger(true, true, true, true, true)
	s := &server{
		db:  test.NewDB(log, 10, 10),
		ctx: context.TODO(),
		log: log,
	}

	s.start()

	defer s.stop()
}
