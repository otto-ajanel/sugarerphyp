package model

// Atribute representa la tabla atribute
type Atribute struct {
	AtributeID       int    `json:"atribute_id" gorm:"column:atribute_id;primaryKey;autoIncrement"`
	AtributeDes      string `json:"atribute_des" gorm:"column:atribute_des"`
	AtributeTypedata string `json:"atribute_typedata" gorm:"column:atribute_typedata"`
}

func (Atribute) TableName() string { return "atribute" }
