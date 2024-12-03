package dao

import (
	"cdel/demo/Normal/config"
	"github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGormMysql(t *testing.T) {
	gormTest(t)
}

func gormTest(t *testing.T) {
	timeout := time.Duration(time.Duration.Milliseconds(300))
	dao, err := New(&config.MysqlConfig{
		Config: mysql.Config{
			User:        "dev_user",
			Passwd:      "dev_password",
			Addr:        "192.168.172.50:6008",
			DBName:      "doorman",
			Timeout:     timeout,
			ReadTimeout: timeout,
		},
		MaxOpen: 10,
		MaxIdle: 1,
	})
	assert.Nil(t, err)
	errNative := dao.Delay()
	assert.Nil(t, errNative)
}
