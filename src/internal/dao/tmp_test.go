package dao

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestSelectByName(t *testing.T) {
	// 缺省使用
	db, mock, _ := sqlmock.New()
	// 如果完全匹配请使用下面的构建方式
	// db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	defer db.Close()

	// mock check ------------------
	columns := []string{"id", "name", "t1", "t2"}
	rows := sqlmock.NewRows(columns).FromCSVString(`
	1,"tom",,
	2,"tom","21",
	3,"tom",,"31"
	4,"tom","41","42"
	`)
	// 注意下面 ？的处理，
	byName := mock.ExpectPrepare("SELECT (.+) FROM tmp WHERE name = \\? LIMIt \\?")
	// 注意下面的 () 的处理
	mock.ExpectPrepare("select SLEEP(.+)")
	// 注意顺序，不能和上一行对调
	byName.ExpectQuery().WithArgs("tom", 10).WillReturnRows(rows)

	// business logic ------------------------
	dao := NewTmpDao(db)
	rtn, e := dao.SelectByName(context.Background(), "tom")
	if e != nil {
		panic(e)
	}
	assert.Equal(t, 4, len(rtn))
}
