package entity

import "database/sql"

type TmpTable struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	T1   string `json:"t1"`
	T2   string `json:"t2"`
}

func ToTmpTable(row *sql.Row) *TmpTable {

}
