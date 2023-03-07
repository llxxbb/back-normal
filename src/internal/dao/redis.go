package dao

import (
	"cdel/demo/Normal/internal/entity"
	"context"
	"fmt"
	"github.com/go-redis/cache/v9"
	"go.uber.org/zap"

	"github.com/redis/go-redis/v9"
)

type CachedTmp struct {
	TmpTableDaoI
	cacheDao *cache.Cache
	appid    int
	prefix   string
}

func GetCacheTmp(daoI TmpTableDaoI, redisClient *redis.ClusterClient, appId int) *CachedTmp {
	mycache := cache.New(&cache.Options{
		Redis: redisClient,
	})
	tmp := CachedTmp{daoI, mycache, appId, fmt.Sprint(appId, ":")}
	zap.L().Info("CachedTmp", zap.String("redis prefix", tmp.prefix))
	return &tmp
}

func (ct *CachedTmp) SelectByName(ctx context.Context, name string) ([]entity.TmpTable, error) {
	var rtn []entity.TmpTable
	key := ct.prefix + name
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
