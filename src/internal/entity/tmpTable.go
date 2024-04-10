package entity

// import "database/sql"

type TmpTable struct {
	Id   int    `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"uniqueIndex:idx_name"`
	T1   string `json:"t1"`
	T2   string `json:"t2"`
}

// func ToTmpTable(row *sql.Row) *TmpTable {

// }
