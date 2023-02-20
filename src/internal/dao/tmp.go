package dao

import (
	"cdel/demo/Normal/internal/entity"
	"database/sql"
)

type tmpTableDao struct {
	sqlTmpSelect *sql.Stmt // tmp 数据表的预编译
	sqlDelay     *sql.Stmt // 用于超时测试
}

type TmpTableDaoI interface {
	SelectByName(name string) ([]entity.TmpTable, error)
	Delay() error
}

func (t *tmpTableDao) Delay() error {
	_, e := t.sqlDelay.Exec()
	if e != nil {
		return e
	}
	return nil
}

func (t *tmpTableDao) SelectByName(name string) ([]entity.TmpTable, error) {
	rows, e := t.sqlTmpSelect.Query(name, 10)
	if e != nil {
		return nil, e
	}
	defer rows.Close()

	var rtn []entity.TmpTable
	for rows.Next() {
		one := entity.TmpTable{}
		var t1, t2 sql.NullString
		if e = rows.Scan(&one.Id, &one.Name, &t1, &t2); e != nil {
			return nil, e
		}
		if t1.Valid {
			one.T1 = t1.String
		}
		if t2.Valid {
			one.T2 = t2.String
		}
		rtn = append(rtn, one)
	}
	return rtn, nil
}

func NewTmpDao(db *sql.DB) TmpTableDaoI {
	dao := tmpTableDao{}
	// Tmp table
	sql := "SELECT id, name, t1, t2 FROM tmp WHERE name = ? LIMIt ?"
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
