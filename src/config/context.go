package config

import "database/sql"

var CTX Context

type Context struct {
	C      *Config
	DemoDB *sql.DB
}

func (c *Context) Init(cfg *Config) {
	c.C = cfg
	c.DemoDB = DemoDBInit(&cfg.Mysql)
}
