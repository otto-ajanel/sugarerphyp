package model

type Product_attridet struct {
	ProductID     int `json:"product_id" gorm:"column:product_id"`
	AtributeDetID int `json:"atridet_id" gorm:"column:atridet_id"`
}

func (Product_attridet) TableName() string { return "product_atridet" }
