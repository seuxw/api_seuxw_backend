package entity

// Database 数据库连接设置
type Database struct {
	DBHost string `ini:"db_host"`	// Host
	DBPort string `ini:"db_port"`	// 端口
	DBUser string `ini:"db_user"`	// 用户名
	DBPwd  string `ini:"db_pwd"`	// 密码
	DBName string `ini:"db_name"`	// 数据库名
}