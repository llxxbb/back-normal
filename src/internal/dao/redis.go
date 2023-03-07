package dao

import (
	"cdel/demo/Normal/internal/entity"
	"cdel/demo/Normal/tool"
	"context"
	"github.com/go-redis/cache/v9"
	"github.com/redis/go-redis/v9"
)

type CachedTmp struct {
	TmpTableDaoI
	cacheDao *cache.Cache
	redis    *redis.ClusterClient
	cfg      *tool.RedisConfig
}

func GetCacheTmp(daoI TmpTableDaoI, redisClient *redis.ClusterClient, cfg *tool.RedisConfig) *CachedTmp {
	mycache := cache.New(&cache.Options{
		Redis: redisClient,
	})
	tmp := CachedTmp{daoI, mycache, redisClient, cfg}
	return &tmp
}

func (ct *CachedTmp) SelectByName(ctx context.Context, name string) ([]entity.TmpTable, error) {
	var rtn []entity.TmpTable
	key := ct.cfg.Prefix + name
	e := ct.cacheDao.Get(ctx, key, &rtn)
	if e != nil {
		if e.Error() != "cache: key is missing" {
			return ct.TmpTableDaoI.SelectByName(ctx, name)
		}
	}
	if len(rtn) > 0 {
		return rtn, nil
	}

	rtn, e = ct.TmpTableDaoI.SelectByName(ctx, name)
	if e != nil {
		return nil, e
	}
	_ = ct.cacheDao.Set(&cache.Item{
		TTL: ct.cfg.Ttl,
		Ctx: ctx,
		Key: key,
		//Key:   ct.prefix + name,
		Value: rtn,
	})
	return rtn, nil
}

//func (ct *CachedTmp) SelectByName(ctx context.Context, name string) ([]entity.TmpTable, error) {
//	item := &cache.Item{
//		Key: name,
//		Do: func(itm *cache.Item) (interface{}, error) {
//			return "hello", nil
//		},
//	}
//	err := ct.cacheDao.Once(item)
//	rtn, _ := item.Value.(string)
//	println(rtn)
//	return nil, err
//}

// Delay is not useful for cache demo, cache demo see SelectByName
func (ct *CachedTmp) Delay() error {
	return ct.TmpTableDaoI.Delay()
}

// test redis
//cSet := clusterClient.Set(context.Background(), "test-go", "hello", time.Second)
//zap.S().Info("cached value set: ", cSet)
//
//rtn, _ := clusterClient.Get(context.Background(), "test-go").Result()
//println("cached value is: ", rtn)
