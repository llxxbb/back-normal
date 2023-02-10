package config

import (
	"github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
)

type MysqlConfig struct {
	mysql.Config
	MaxOpen int
	MaxIdle int
}

func (c *MysqlConfig) AppendFieldMap(fMap map[string]string) {
	fMap["mysql.user"] = "Mysql.User"
	fMap["mysql.password"] = "Mysql.Passwd"
	fMap["mysql.address"] = "Mysql.Addr"
	fMap["mysql.db"] = "Mysql.DBName"
	fMap["mysql.conns.timeout"] = "Mysql.Timeout"
	fMap["mysql.conns.readTimeout"] = "Mysql.ReadTimeout"
	fMap["mysql.conns.maxOpen"] = "Mysql.MaxOpen"
	fMap["mysql.conns.maxIdle"] = "Mysql.MaxIdle"
}

func (c *MysqlConfig) Print() {
	zap.L().Info("------------ mysql ------------")
	zap.L().Info("-- ", zap.String("addr", c.Addr))
	zap.L().Info("-- ", zap.String("dbName", c.DBName))
	zap.L().Info("-- ", zap.String("timeout", c.Timeout.String()))
	zap.L().Info("-- ", zap.String("readTimeout", c.ReadTimeout.String()))
	zap.L().Info("-- ", zap.Int("maxOpen", c.MaxOpen))
	zap.L().Info("-- ", zap.Int("MaxIdle", c.MaxIdle))
}
