package dao

import (
	"back/demo/internal/entity"
	"context"
	"database/sql"
)

type TmpTableDaoI interface {
	SelectByName(ctx context.Context, name string) ([]entity.TmpTable, error)
	Delay() error
}

type tmpTableDao struct {
	sqlTmpSelect *sql.Stmt // tmp 数据表的预编译
	sqlDelay     *sql.Stmt // 用于超时测试
}

func (t *tmpTableDao) Delay() error {
	_, e := t.sqlDelay.Exec()
	if e != nil {
		return e
	}
	return nil
}

func (t *tmpTableDao) SelectByName(ctx context.Context, name string) ([]entity.TmpTable, error) {
	rows, e := t.sqlTmpSelect.QueryContext(ctx, name, 10)
	if e != nil {
		return nil, e
	}
	// 不要忘记关闭。
	defer rows.Close()

	var rtn []entity.TmpTable
	for rows.Next() {
		one := entity.TmpTable{}
		// 注意 null 类型
		var t1, t2 sql.NullString
		if e = rows.Scan(&one.Id, &one.Domain, &t1, &t2); e != nil {
			return nil, e
		}
		// 对 null 进行验证
		if t1.Valid {
			one.ResourcePath = t1.String
		}
		if t2.Valid {
			one.RealUrl = t2.String
		}
		rtn = append(rtn, one)
	}
	return rtn, nil
}

func NewTmpDao(db *sql.DB) TmpTableDaoI {
	dao := tmpTableDao{}
	// Tmp table 注意参数用 ？ 替代
	sql := "SELECT rm_id, domain, resourcePath, realUrl FROM routemap WHERE rm_id = ? LIMIt ?"
	var e error
	dao.sqlTmpSelect, e = db.Prepare(sql)
	if e != nil {
		panic(e)
	}

	// 超时测试
	sql = "select SLEEP(10)"
	dao.sqlDelay, e = db.Prepare(sql)
	if e != nil {
		panic(e)
	}
	return &dao
}
