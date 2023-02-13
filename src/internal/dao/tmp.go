package dao

import (
	"cdel/demo/Normal/internal/entity"
	"database/sql"
)

type TmpTableDao struct {
	sqlTmpSelect *sql.Stmt // tmp 数据表的预编译
	sqlDelay     *sql.Stmt // 用于超时测试
}

func (t *TmpTableDao) Delay() error {
	_, e := t.sqlDelay.Exec()
	if e != nil {
		return e
	}
	return nil
}

func (t *TmpTableDao) SelectByName(name string) ([]entity.TmpTable, error) {
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

func (t *TmpTableDao) Prepare(db *sql.DB) {
	// Tmp table
	sql := "SELECT id, name, t1, t2 FROM tmp WHERE name = ? LIMIt ?"
	var e error
	t.sqlTmpSelect, e = db.Prepare(sql)
	if e != nil {
		panic(e)
	}

	// 超时测试
	sql = "select SLEEP(10)"
	t.sqlDelay, e = db.Prepare(sql)
	if e != nil {
		panic(e)
	}
}
