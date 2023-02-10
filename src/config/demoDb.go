package config

import (
	"database/sql"

	"go.uber.org/zap"
)

func DemoDBInit(cfg *MysqlConfig) *sql.DB {
	// more default config for mysql
	cfg.Net = "tcp"
	cfg.AllowNativePasswords = true

	// connect
	db, err := sql.Open("mysql", cfg.FormatDSN())
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
