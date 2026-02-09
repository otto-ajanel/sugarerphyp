package model

import "encoding/json"

type AviableStore struct {
	StoreId int
	//StoreDes string
	CountAviable int
}

type ProductAviable struct {
	ProductId       int
	ProductSku      string
	ProductDes      string
	CategoryDes     string
	Price           float64
	TotalAvailable  int
	InventoryStatus string
	Locations       json.RawMessage `gorm:"column:locations" json:"locations"`
}
