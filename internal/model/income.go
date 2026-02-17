package model

type Income struct {
	IncomeID     int    `gorm:"column:income_id;primaryKey;autoIncrement" json:"income_id"`
	IncomeDoc    string `gorm:"column:income_doc" json:"income_doc"`
	IncomeDateIn string `gorm:"column:income_dateing" json:"income_dateing"`
	UserID       int    `gorm:"column:user_id" json:"user_id"`
	SupplierID   int    `gorm:"column:supplier_id" json:"supplier_id"`
	StoreID      int    `gorm:"column:store_id" json:"store_id"`
}

func (Income) TableName() string {
	return "incomes"
}

type ResultIncome struct {
	ID            uint   `gorm:"column:income_id"`
	IncomeDoc     string `gorm:"column:income_doc"`
	IncomeDateIng string `gorm:"column:income_dateing"`
	Username      string `gorm:"column:username"`
	StoreName     string `gorm:column:store_name`
}
