package main

import (
	"context"
	"fmt"
	"net/http"
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

func (self *server) start() {
	self.Server = web.NewServer(self.log)
	router := self.PathPrefix("/").SubRouter()

	/* -测试Handler- */
	// 测试接口 将会返回输入的日期对应的日出日落时间（需要对应时间）
	router.HandleFunc("/test", self.TestGet).Methods("GET")
	router.HandleFunc("/test", self.TestPost).Methods("POST")

	self.Serve("0.0.0.0:10000")
}

func (self *server) stop() {
	self.db.Close()
}

// Trace详情打印函数
func (self *server) PrintTrace(r *http.Request, para string, head string) {
	header := r.Header.Get("X-Trace")
	env := fmt.Sprintf("ENV:{TraceID:%v,Parameters:[%s]}", header, para)
	self.log.Trace("%s: %s", head, env)
}

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
