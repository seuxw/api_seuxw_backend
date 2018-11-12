package main

import (
	"context"
	"seuxw/embrice/rdb/user"
	"seuxw/x/logger"
	"seuxw/x/web"
)

type server struct {
	*web.Server
	db  *user.Database
	log *logger.Logger
	ctx context.Context
}

// start 程序开始
func (svr *server) start() {
	svr.Server = web.NewServer(svr.log)
	/* 用户Handler */
	userRouter := svr.PathPrefix("/user").SubRouter()
	userRouter.HandleFunc("/create_user", svr.CreateUser).Methods("POST") // 创建用户

	svr.Serve("0.0.0.0:20000")
}

// stop 服务终止，断开数据库连接
func (svr *server) stop() {
	svr.db.Close()
}

// main 主程序入口
func main() {
	log := logger.NewStdLogger(true, true, true, true, true)
	s := &server{
		db:  user.NewDB(log, 10, 10),
		ctx: context.TODO(),
		log: log,
	}
	s.start()

	defer s.stop()
}
