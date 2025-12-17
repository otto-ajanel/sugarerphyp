package model

type Product struct {
	ProductID       int    `json:"product_id" gorm:"column:product_id;primaryKey;autoIncrement"`
	ProductSKU      string `json:"product_sku" gorm:"column:product_sku"`
	ProductDes      string `json:"product_des" gorm:"column:product_des"`
	CategoryID      int    `json:"category_id" gorm:"column:category_id"`
	TypeProductID   int    `json:"typeproduct_id" gorm:"column:typeproduct_id"`
	BillingPolicyID int    `json:"billingpolicy_id" gorm:"column:billingpolicy_id"`
	StatusID        int    `json:"status_id" gorm:"column:status_id"`
}

func (Product) TableName() string { return "products" }

type ProductsPrice struct {
	ID          int     `json:"id" gorm:"column:prodprec_id;primaryKey;autoIncrement"`
	Price       float64 `json:"prodprec_price" gorm:"column:prodprec_price"`
	TaxPrice    float64 `json:"prodprec_taxprice" gorm:"column:prodprec_taxprice"`
	ProductID   int     `json:"product_id" gorm:"column:product_id"`
	Purchase    float64 `json:"prodprec_purchase" gorm:"column:prodprec_purchase"`
	PurchaseTax float64 `json:"producprec_purchasetax" gorm:"column:producprec_purchasetax"`
}

func (ProductsPrice) TableName() string { return "products_price" }

// ModuleAviability representa module availability (moduleaviability)
type ModuleAviability struct {
	ID        int `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	ModuleID  int `json:"module_id" gorm:"column:module_id"`
	ProductID int `json:"product_id" gorm:"column:product_id"`
}

func (ModuleAviability) TableName() string { return "moduleaviability" }
