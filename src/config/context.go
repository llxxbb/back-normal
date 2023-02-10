package config

import "database/sql"

var CTX Context

type Context struct {
	C      *DemoConfig
	DemoDB *sql.DB
}

func (c *Context) Init(cfg *DemoConfig) {
	c.C = cfg
	c.DemoDB = DemoDBInit(&cfg.Mysql)
}
