package entity

// import "database/sql"

type TmpTable struct {
	Id           int    `json:"id" gorm:"primaryKey;column:rm_id"`
	Domain       string `json:"oomain"`
	ResourcePath string `json:"resourcePath"`
	RealUrl      string `json:"realUrl"`
}

// func ToTmpTable(row *sql.Row) *TmpTable {

// }
