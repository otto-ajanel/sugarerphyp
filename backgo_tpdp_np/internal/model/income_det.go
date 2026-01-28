package model

type IncomeDet struct {
	IncomeDetID    int     `json:"incomedet_id" gorm:"column:incomedet_id;primaryKey;autoIncrement" `
	ProductID      int     `gorm:"column:product_id" json:"product_id"`
	IncomeDetCount float64 `gorm:"column:incomedet_count" json:"incomedet_count"`
	IncomeDetVal   float64 `gorm:"column:incomedet_val" json:"incomedet_val"`
	IncomeID       int     `gorm:"column:income_id" json:"income_id"`
}

func (IncomeDet) TableName() string {
	return "income_det"
}
