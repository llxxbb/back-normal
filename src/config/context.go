package config

import (
	"cdel/demo/Normal/internal/dao"
	"database/sql"
	"github.com/redis/go-redis/v9"

	"github.com/go-resty/resty/v2"
)

var CTX Context

type Context struct {
	Cfg           *DemoConfig
	DemoDB        *sql.DB
	TmpDao        dao.TmpTableDaoI
	GatewayClient *resty.Client
	App2Client    *resty.Client
	Redis         *redis.ClusterClient
}

func (c *Context) Init(cfg *DemoConfig) {
	c.Cfg = cfg
	// 数据库及 dao 初始化 --------------------------
	c.DemoDB = cfg.Mysql.DemoDBInit()
	c.TmpDao = dao.NewTmpDao(c.DemoDB)

	// rpc 初始化 --------------------------
	c.GatewayClient = cfg.Rpc.NewGateWayClient()
	c.App2Client = cfg.Rpc.NewApp2Client()

	// redis 初始化 --------------------------
	c.Redis = cfg.Redis.GetRedisClient()
}
