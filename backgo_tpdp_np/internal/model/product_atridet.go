package model

type Product_attridet struct {
	ProductID        int `json:"product_id" gorm:"column:product_id"`
	AtributeDetID   int `json:"atributedet_id" gorm:"column:atributedet_id"`
}

func (Product_attridet) TableName() string { return "product_atridet" }
