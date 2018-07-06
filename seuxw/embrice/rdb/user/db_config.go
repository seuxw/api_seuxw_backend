package user

import (
	"database/sql"
	"fmt"
	"seuxw/embrice/config"
	"seuxw/x/logger"
	_ "seuxw/x/mysql"
	"seuxw/x/sqlx"
)


type Database struct {
	*sqlx.DB
	log *logger.Logger
}


func NewDB(log *logger.Logger, maxConns, maxIdles int) *Database {
	config, err := config.ReadDBConfig()
	if err != nil {
		log.Fatal("Connected to database failed. Please check the config file. Err:%s", err)
	}
	db := &Database{
		log: log,
		DB: sqlx.MustConnect("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		config.DBUser, config.DBPwd, config.DBHost, config.DBPort, config.DBName)), //连接数据库
	}
	db.SetMaxOpenConns(maxConns)
	db.SetMaxIdleConns(maxIdles)
	log.Trace("Connect to database [%s] succeed.", config.DBName)
	return db
}

// 判断Update执行语句后影响的函数并输出错误
func (self *Database) JudgeAffect(Result sql.Result) error {
	var (
		err       error
		AffectRow int64
	)
	if AffectRow, err = Result.RowsAffected(); err != nil {
		err = fmt.Errorf("获取更新数据库执行状态失败！ err:%s", err)
		goto END
	}

	if AffectRow == 0 {
		err = fmt.Errorf("没有数据被修改！")
		goto END
	}
END:
	return err
}
