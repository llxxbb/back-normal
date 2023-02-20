package config

import (
	"cdel/demo/Normal/internal/dao"
	"cdel/demo/Normal/tool"
	"database/sql"

	"github.com/go-resty/resty/v2"
)

var CTX Context

type Context struct {
	Cfg    *DemoConfig
	DemoDB *sql.DB
	TmpDao dao.TmpTableDaoI
	Client *resty.Client
}

func (c *Context) Init(cfg *DemoConfig) {
	c.Cfg = cfg
	c.DemoDB = DemoDBInit(&cfg.Mysql)
	// 初始化DAO
	c.TmpDao = dao.NewTmpDao(c.DemoDB)
	// 初始化 rest client
	c.Client = tool.NewRestClient(cfg.RestTimeout)
}
