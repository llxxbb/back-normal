package tool

import (
	"go.uber.org/zap"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisConfig struct {
	Url          string
	User         string
	Password     string
	PoolSize     int
	PoolTimeout  time.Duration
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	DialTimeout  time.Duration
	MinIdleConns int
}

func (c *RedisConfig) AppendFieldMap(fMap map[string]string) {
	fMap["redis.url"] = "Redis.Url"
	fMap["redis.user"] = "Redis.User"
	fMap["redis.password"] = "Redis.Password"
	fMap["redis.poolSize"] = "Redis.PoolSize"
	fMap["redis.poolTimeout"] = "Redis.PoolTimeout"
	fMap["redis.readTimeout"] = "Redis.ReadTimeout"
	fMap["redis.writeTimeout"] = "Redis.WriteTimeout"
	fMap["redis.dialTimeout"] = "Redis.DialTimeout"
	fMap["redis.minIdleConns"] = "Redis.MinIdleConns"
}

func (c *RedisConfig) Print() {
	zap.L().Info("------------ redis ------------")
	zap.L().Info("-- ", zap.String("url", c.Url))
	zap.L().Info("-- ", zap.String("user", c.User))
	zap.L().Info("-- ", zap.Int("poolSize", c.PoolSize))
	zap.L().Info("-- ", zap.Duration("poolTimeout", c.PoolTimeout))
	zap.L().Info("-- ", zap.Duration("readTimeout", c.ReadTimeout))
	zap.L().Info("-- ", zap.Duration("writeTimeout", c.WriteTimeout))
	zap.L().Info("-- ", zap.Duration("dialTimeout", c.DialTimeout))
	zap.L().Info("-- ", zap.Int("minIdleConns", c.MinIdleConns))
}

func (c *RedisConfig) GetRedisClient() *redis.ClusterClient {

	clusterClient := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: strings.Split(c.Url, ","),
		NewClient: func(opt *redis.Options) *redis.Client {
			opt.Username = c.User
			opt.Password = c.Password
			opt.PoolSize = c.PoolSize
			opt.PoolTimeout = c.PoolTimeout
			opt.ReadTimeout = c.ReadTimeout
			opt.WriteTimeout = c.WriteTimeout
			opt.DialTimeout = c.DialTimeout
			opt.MinIdleConns = c.MinIdleConns
			return redis.NewClient(opt)
		},
		RouteByLatency: true,
	})

	// test redis
	//cSet := clusterClient.Set(context.Background(), "test-go", "hello", time.Second)
	//zap.S().Info("cached value set: ", cSet)
	//
	//rtn, _ := clusterClient.Get(context.Background(), "test-go").Result()
	//println("cached value is: ", rtn)

	return clusterClient
}
