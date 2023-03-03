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
		Addrs:          strings.Split(c.Url, ","),
		Username:       c.User,
		Password:       c.Password,
		PoolSize:       c.PoolSize,
		PoolTimeout:    c.PoolTimeout,
		ReadTimeout:    c.ReadTimeout,
		WriteTimeout:   c.WriteTimeout,
		DialTimeout:    c.DialTimeout,
		MinIdleConns:   c.MinIdleConns,
		RouteByLatency: true,
	})

	return clusterClient
}
