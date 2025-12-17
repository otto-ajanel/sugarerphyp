package repository

import (
	"github.com/otto-ajanel/backgo_tpdp_np/internal/model"
	"gorm.io/gorm"
)

type AtributeDetailRepo struct{}

func NewAtributeDetailRepo() *AtributeDetailRepo { return &AtributeDetailRepo{} }

func (r *AtributeDetailRepo) GetAll(db *gorm.DB) ([]model.AtributeDetail, error) {

	var arr []model.AtributeDetail

	if err := db.
		Select(`
	atribute_detail.*,
	atribute.atribute_des
	`).
		Joins("LEFT JOIN atribute ON atribute.atribute_id = atribute_detail.atribute_id").
		Find(&arr).Error; err != nil {
		return nil, err
	}
	return arr, nil
}

func (r *AtributeDetailRepo) Create(tx *gorm.DB, a *model.AtributeDetail) error {
	return tx.Create(a).Error
}

// GetByAtributeID devuelve los detalles asociados a un atributo específico.
func (r *AtributeDetailRepo) GetByAtributeID(db *gorm.DB, atributeID int) ([]model.AtributeDetail, error) {
	var arr []model.AtributeDetail
	if err := db.Preload("Atribute").Where("atribute_id = ?", atributeID).Find(&arr).Error; err != nil {
		return nil, err
	}
	return arr, nil
}
