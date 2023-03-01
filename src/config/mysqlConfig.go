package config

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"

	_ "github.com/pinpoint-apm/pinpoint-go-agent/plugin/mysql"
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

func (cfg *MysqlConfig) DemoDBInit() *sql.DB {
	// more default config for mysql
	cfg.Net = "tcp"
	cfg.AllowNativePasswords = true

	// connect
	db, err := sql.Open("mysql-pinpoint", cfg.FormatDSN())
	if err != nil {
		zap.S().Fatal(err)
		panic(err)
	}
	db.SetMaxOpenConns(cfg.MaxOpen)
	db.SetMaxIdleConns(cfg.MaxIdle)
	err = db.Ping()
	if err != nil {
		zap.S().Fatal(err)
		panic(err)
	}
	zap.S().Info("connected for : ", cfg.DBName)
	return db
}
