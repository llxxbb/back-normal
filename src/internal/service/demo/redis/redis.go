package redis

import (
	"cdel/demo/Normal/internal/dao"
	"cdel/demo/Normal/internal/entity"
	"context"

	"github.com/redis/go-redis/v9"
)

type CachedTmp struct {
	dao.TmpTableDaoI
	cache *redis.ClusterClient
}

func (ct *CachedTmp) SelectByName(ctx context.Context, name string) ([]entity.TmpTable, error) {
	return ct.TmpTableDaoI.SelectByName(ctx, name)
}

func (ct *CachedTmp) Delay() error {
	return ct.TmpTableDaoI.Delay()
}
