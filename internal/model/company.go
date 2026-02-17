package model

type Company struct {
	CompanyID   int    `json:"id_company" gorm:"column:id_company;primaryKey;autoIncrement"`
	CompanyName string `json:"company_name" gorm:"column:company_name"`
	Address     string `json:"address" gorm:"column:address"`
	Email       string `json:"email" gorm:"column:email"`
}

func (Company) TableName() string { return "companies" }
