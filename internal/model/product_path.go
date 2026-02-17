package model

type Product_path struct {
	Product_pathID int    `json:"productpath_id" gorm:"column:productpath_id;primaryKey;autoIncrement"`
	ProductID      int    `json:"product_id" gorm:"column:product_id"`
	ProductPath    string `json:"path" gorm:"column:productpath_path"`
}

func (Product_path) TableName() string { return "product_path" }
