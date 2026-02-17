package model

// Atribute representa la tabla atribute
type AtributeDetail struct {
	AtributeDetailID          int    `json:"atribute_detail_id" gorm:"column:atridet_id;primaryKey;autoIncrement"`
	AtributeID                int    `json:"atribute_id" gorm:"column:atribute_id"`
	AtributeDetailDescription string `json:"atribute_detail_description" gorm:"column:atridet_description"`
	// Relación opcional al Atribute padre

	Atribute_des string `json:"atribute_des" gorm:"column:atribute_des"`
}

func (AtributeDetail) TableName() string { return "atribute_detail" }
