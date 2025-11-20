package model

// Category representa la entidad category usada por el API.
type Category struct {
    ID  int    `json:"id" gorm:"column:id_category;primaryKey"`
    Des string `json:"des" gorm:"column:des_category"`
}

func (Category) TableName() string { return "categories" }
