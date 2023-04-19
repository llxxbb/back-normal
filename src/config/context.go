package config

import (
	"cdel/demo/Normal/internal/dao"
	"cdel/demo/Normal/tool"
	"database/sql"

	"github.com/go-resty/resty/v2"
)

var CTX Context

type Context struct {
	Cfg           *DemoConfig
	DemoDB        *sql.DB
	GatewayClient *resty.Client
	App2Client    *resty.Client
	TmpDao        dao.TmpTableDaoI
	TmpCache      dao.TmpTableDaoI
}

func (c *Context) Init(cfg *DemoConfig) {
	c.Cfg = cfg
	// 数据库及 dao 初始化 --------------------------
	c.DemoDB = cfg.Mysql.DemoDBInit()
	c.TmpDao = dao.NewTmpDao(c.DemoDB)
	// 对数据进行缓存
	redisClient := cfg.Redis.GetRedisClient()                       // redis 初始化
	c.TmpCache = dao.GetCacheTmp(c.TmpDao, redisClient, &cfg.Redis) // cached dao

	// rpc 初始化 --------------------------
	c.GatewayClient = tool.RpcClient(cfg.Rpc.Timeout, cfg.Rpc.BUGateway)
	c.GatewayClient.SetHeader("HOST", "gateway.cdeledu.com")
	c.App2Client = tool.RpcClient(cfg.Rpc.Timeout, cfg.Rpc.BUApp)
}
