package main

import (
	"context"
	py "github.com/sbinet/go-python"
	"seuxw/embrice/rdb/user"
	"seuxw/x/logger"
	"seuxw/x/web"
)

type server struct {
	*web.Server
	db  *user.Database
	log *logger.Logger
	ctx context.Context
	py  *py.PyObject
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
		py:  NewPyModule("/data/code/seuxw/api/api_seuxw_backend/py_seuxw/painter", "treehole"),
	}
	s.start()

	defer s.stop()
}

var (
	PyStr = py.PyString_FromString
	GoStr = py.PyString_AS_STRING
)

// NewPyModule
func NewPyModule(dir, name string) *py.PyObject {
	_ = py.Initialize()
	InsertBeforeSysPath("/usr/bin/python3")
	return ImportModule(dir, name)
}

// ImportModule
func ImportModule(dir, name string) *py.PyObject {
	sysModule := py.PyImport_ImportModule("sys") // import sys
	path := sysModule.GetAttrString("path")      // path = sys.path
	py.PyList_Insert(path, 0, PyStr(dir))        // path.insert(0, dir)
	return py.PyImport_ImportModule(name)        // return __import__(name)
}

// InsertBeforeSysPath
func InsertBeforeSysPath(p string) string {
	sysModule := py.PyImport_ImportModule("sys")
	path := sysModule.GetAttrString("path")
	py.PyList_Insert(path, 0, PyStr(p))
	return GoStr(path.Repr())
}
