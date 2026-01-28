package model

type Supplier struct {
	SupplierID  int    `gorm:"column:id_supplier;primaryKey;autoIncrement" json:"id_supplier"`
	CompanyName string `gorm:"column:company_name" json:"company_name"`
	Address     string `gorm:"column:address" json:"address"`
	Phone       string `gorm:"column:phone" json:"phone"`
	Email       string `gorm:"column:email" json:"email"`
	Active      bool   `gorm:"column:active" json:"active"`
}

func (Supplier) TableName() string {
	return "suppliers"
}
