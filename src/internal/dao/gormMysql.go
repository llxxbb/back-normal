package dao

import (
	"cdel/demo/Normal/config"
	"cdel/demo/Normal/internal/entity"
	"context"
	"gitlab.cdel.local/platform/go/platform-common/def"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type GromTableDao struct {
	db *gorm.DB
}

func New(cfg *config.MysqlConfig) (*GromTableDao, *def.CustomError) {
	cfg.Net = "tcp"
	cfg.AllowNativePasswords = true
	dsn := cfg.FormatDSN()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			NoLowerCase:   true, // 用于跳过小写转换, 设置为 true 则不会强行使用蛇形格式，保持定义的形式。
			SingularTable: true,
		},
		PrepareStmt: true,
	})
	if err != nil {
		return nil, def.NewCustomError(def.ET_ENV, def.ENV_C, err.Error(), nil)
	}
	baseDb, err := db.DB()
	if err != nil {
		return nil, def.NewCustomError(def.ET_ENV, def.ENV_C, err.Error(), nil)
	}
	baseDb.SetMaxOpenConns(cfg.MaxOpen)
	baseDb.SetMaxIdleConns(cfg.MaxIdle)
	rtn := GromTableDao{
		db,
	}
	return &rtn, nil
}

// SelectByName 一个查询示例，其他请参考：https://gorm.io/zh_CN/docs/index.html
func (gt *GromTableDao) SelectByName(_ context.Context, name string) ([]entity.TmpTable, error) {
	var dest []entity.TmpTable
	gt.db.Where("name = ?", name).Find(&dest)
	return dest, nil
}

func (gt *GromTableDao) Delay() error {
	gt.db.Exec("select sleep(1) from dual")
	return nil
}
