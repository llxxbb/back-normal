package demo

import (
	"back/demo/config"
	"back/demo/internal/dao"

	"github.com/go-resty/resty/v2"
)

var tmpDao dao.TmpTableDaoI
var tmpCache *dao.CachedTmp

var gatewayClient *resty.Client
var app2Client *resty.Client

func InitDemo(cfg *config.ProjectConfig) {
	// 数据库及 dao 初始化 --------------------------
	demoDB := cfg.Mysql.DemoDBInit()
	tmpDao = dao.NewTmpDao(demoDB)
	// 对数据进行缓存
	redisClient := cfg.Redis.GetRedisClient()                   // redis 初始化
	tmpCache = dao.GetCacheTmp(tmpDao, redisClient, &cfg.Redis) // cached dao

	// rpc 初始化 --------------------------
	gatewayClient := cfg.GateWay.NewClient()
	gatewayClient.SetHeader("HOST", "gateway.cdeledu.com")
	app2Client = cfg.App2.NewClient()
}
