// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameKcql = "kcql"

// Kcql mapped from table <kcql>
type Kcql struct {
	KCAAA  string  `gorm:"column:KCAAA" json:"KCAAA"`
	DWAACJ string  `gorm:"column:DWAACJ" json:"DWAACJ"`
	DWAADJ string  `gorm:"column:DWAADJ" json:"DWAADJ"`
	MDAAC  string  `gorm:"column:MDAAC" json:"MDAAC"`
	MDAREA float32 `gorm:"column:MDAREA" json:"MDAREA"`
	MDBA   string  `gorm:"column:MDBA" json:"MDBA"`
	MDCP   string  `gorm:"column:MDCP" json:"MDCP"`
	MDEG   string  `gorm:"column:MDEG" json:"MDEG"`
	MDBFNQ string  `gorm:"column:MDBFNQ" json:"MDBFNQ"`
	PKBBB  float64 `gorm:"column:PKBBB" json:"PKBBB"`
	Shuju     string  `gorm:"column:数据" json:"数据"`
}

// TableName Kcql's table name
func (*Kcql) TableName() string {
	return TableNameKcql
}
