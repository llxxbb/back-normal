package config

import (
	"cdel/demo/Normal/internal/dao"
	"database/sql"
)

var CTX Context

type Context struct {
	C      *DemoConfig
	DemoDB *sql.DB
	TmpDao *dao.TmpTableDao
}

func (c *Context) Init(cfg *DemoConfig) {
	c.C = cfg
	c.DemoDB = DemoDBInit(&cfg.Mysql)
	// 初始化DAO
	c.TmpDao = &dao.TmpTableDao{}
	c.TmpDao.Prepare(c.DemoDB)
}
