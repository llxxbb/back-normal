package config

import (
	"github.com/go-sql-driver/mysql"
)

type MysqlConfig struct {
	mysql.Config
	MaxOpen int
	MaxIdle int
}

func AppendMysqlConfigMap(fMap map[string]string) {
	fMap["mysql.user"] = "Mysql.User"
	fMap["mysql.password"] = "Mysql.Passwd"
	fMap["mysql.address"] = "Mysql.Addr"
	fMap["mysql.db"] = "Mysql.DBName"
	fMap["mysql.conns.timeout"] = "Mysql.Timeout"
	fMap["mysql.conns.readTimeout"] = "Mysql.ReadTimeout"
	fMap["mysql.conns.maxOpen"] = "Mysql.MaxOpen"
	fMap["mysql.conns.maxIdle"] = "Mysql.MaxIdle"
}
