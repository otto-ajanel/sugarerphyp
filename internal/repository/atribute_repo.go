package repository

import (
	"sugarerpgo/internal/model"

	"gorm.io/gorm"
)

type AtributeRepo struct{}

func NewAtributeRepo() *AtributeRepo { return &AtributeRepo{} }

func (r *AtributeRepo) GetAll(db *gorm.DB) ([]model.Atribute, error) {
	var arr []model.Atribute
	if err := db.Find(&arr).Error; err != nil {
		return nil, err
	}
	return arr, nil
}

func (r *AtributeRepo) Create(tx *gorm.DB, a *model.Atribute) error {
	return tx.Create(a).Error
}
