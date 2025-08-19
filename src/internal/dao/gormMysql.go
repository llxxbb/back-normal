package dao

import (
	"back/demo/config"
	"back/demo/internal/entity"
	"context"
	"strings"

	"github.com/llxxbb/platform-common/def"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type GromTableDao struct {
	db *gorm.DB
}

// CustomNamingStrategy 自定义 NamingStrategy
type CustomNamingStrategy struct {
	schema.NamingStrategy
}

// ColumnName 实现 ColumnName 方法，将 Go 结构体字段名转换为数据库字段名
func (ns CustomNamingStrategy) ColumnName(_, column string) string {
	return strings.ToLower(column[:1]) + column[1:]
}

func New(cfg *config.MysqlConfig) (*GromTableDao, *def.CustomError) {
	cfg.Net = "tcp"
	cfg.AllowNativePasswords = true
	dsn := cfg.FormatDSN()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		// 结构与数据库字段仅首字母不同
		NamingStrategy: CustomNamingStrategy{},
		// 如果数据库字段与结构体字段完全一致时，使用以下配置
		//NamingStrategy: schema.NamingStrategy{
		//	NoLowerCase:   true, // 用于跳过小写转换, 设置为 true 则不会强行使用蛇形格式，保持定义的形式。
		//	SingularTable: true,
		//},
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
