package model

type Store struct {
	StoreID   int    `json:"store_id" gorm:"column:id_store;primaryKey;autoIncrement"`
	CompanyID int    `json:"company_id" gorm:"column:id_company"`
	StoreName string `json:"store_name" gorm:"column:store_name"`
	Address   string `json:"address" gorm:"column:address"`
	Phone     string `json:"phone" gorm:"column:phone"`
	Active    bool   `json:"active" gorm:"column:active"`
}

func (Store) TableName() string { return "stores" }
